package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/apistruct"
	"github.com/filecoin-project/venus-auth/core"
	manet "github.com/multiformats/go-multiaddr/net"
	"go.opencensus.io/tag"
	"golang.org/x/xerrors"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-multiaddr"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-jsonrpc/auth"

	"github.com/filecoin-project/lotus/metrics"
	"github.com/filecoin-project/lotus/node"
	"github.com/filecoin-project/lotus/node/impl"
	auth2 "github.com/filecoin-project/venus-auth/auth"
)

var log = logging.Logger("main")

type Handler2 struct {
	Verify func(spanId, serviceName, preHost, host, token string) (*auth2.VerifyResponse, error)
	Next   http.HandlerFunc
}

func MacAddr() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("net interfaces" + err.Error())
	}
	mac := ""
	for _, netInterface := range interfaces {
		mac = netInterface.HardwareAddr.String()
		if len(mac) == 0 {
			continue
		}
		break
	}
	return mac
}

func (h *Handler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token := r.Header.Get("Authorization")
	// if other nodes on the same PC, the permission check will passes directly
	if strings.Split(r.RemoteAddr, ":")[0] == "127.0.0.1" {
		ctx = auth.WithPerm(ctx, []auth.Permission{"read", "write", "sign", "admin"})
	} else {
		if token == "" {
			token = r.FormValue("token")
			if token != "" {
				token = "Bearer " + token
			}
		}
		if token != "" {
			if !strings.HasPrefix(token, "Bearer ") {
				log.Warn("missing Bearer prefix in venusauth header")
				w.WriteHeader(401)
				return
			}
			fmt.Println(token)
			token = strings.TrimPrefix(token, "Bearer ")
			res, err := h.Verify(MacAddr(), "lotus", r.RemoteAddr, r.Host, token)
			if err != nil {
				log.Warnf("JWT Verification failed (originating from %s): %s", r.RemoteAddr, err)
				w.WriteHeader(401)
				return
			}
			perms := core.AdaptOldStrategy(res.Perm)
			perms2 := make([]auth.Permission, 0)
			for _, v := range perms {
				perms2 = append(perms2, auth.Permission(v))
			}
			ctx = auth.WithPerm(ctx, perms2)
		}
	}
	h.Next(w, r.WithContext(ctx))
}

func serveRPC(a api.FullNode, authEndpoint string, stop node.StopFunc, addr multiaddr.Multiaddr, shutdownCh <-chan struct{}, maxRequestSize int64) error {
	serverOptions := make([]jsonrpc.ServerOption, 0)
	if maxRequestSize != 0 { // config set
		serverOptions = append(serverOptions, jsonrpc.WithMaxRequestSize(maxRequestSize))
	}
	serveRpc := func(path string, hnd interface{}) {
		rpcServer := jsonrpc.NewServer(serverOptions...)
		rpcServer.Register("Filecoin", hnd)

		if authEndpoint != "" {
			cli := NewJWTClient(authEndpoint)
			ah := &Handler2{
				Verify: cli.Verify,
				Next:   rpcServer.ServeHTTP,
			}
			http.Handle(path, ah)
			fmt.Println("✅ venus auth")
		} else {
			ah := &auth.Handler{
				Verify: a.AuthVerify,
				Next:   rpcServer.ServeHTTP,
			}
			http.Handle(path, ah)
		}

	}

	//pma := api.PermissionedFullAPI(metrics.MetricedFullAPI(a))
	//serveRpc("/rpc/v1", pma)
	//serveRpc("/rpc/v0", &v0api.WrapperV1Full{FullNode: pma})
	serveRpc("/rpc/v0", apistruct.PermissionedFullAPI(metrics.MetricedFullAPI(a)))

	importAH := &auth.Handler{
		Verify: a.AuthVerify,
		Next:   handleImport(a.(*impl.FullNodeAPI)),
	}

	http.Handle("/rest/v0/import", importAH)

	http.Handle("/debug/metrics", metrics.Exporter())
	http.Handle("/debug/pprof-set/block", handleFractionOpt("BlockProfileRate", runtime.SetBlockProfileRate))
	http.Handle("/debug/pprof-set/mutex", handleFractionOpt("MutexProfileFraction",
		func(x int) { runtime.SetMutexProfileFraction(x) },
	))

	lst, err := manet.Listen(addr)
	if err != nil {
		return xerrors.Errorf("could not listen: %w", err)
	}

	srv := &http.Server{
		Handler: http.DefaultServeMux,
		BaseContext: func(listener net.Listener) context.Context {
			ctx, _ := tag.New(context.Background(), tag.Upsert(metrics.APIInterface, "lotus-daemon"))
			return ctx
		},
	}

	sigCh := make(chan os.Signal, 2)
	shutdownDone := make(chan struct{})
	go func() {
		select {
		case sig := <-sigCh:
			log.Warnw("received shutdown", "signal", sig)
		case <-shutdownCh:
			log.Warn("received shutdown")
		}

		log.Warn("Shutting down...")
		if err := srv.Shutdown(context.TODO()); err != nil {
			log.Errorf("shutting down RPC server failed: %s", err)
		}
		if err := stop(context.TODO()); err != nil {
			log.Errorf("graceful shutting down failed: %s", err)
		}
		log.Warn("Graceful shutdown successful")
		_ = log.Sync() //nolint:errcheck
		close(shutdownDone)
	}()
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	err = srv.Serve(manet.NetListener(lst))
	if err == http.ErrServerClosed {
		<-shutdownDone
		return nil
	}
	return err
}

func handleImport(a *impl.FullNodeAPI) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			w.WriteHeader(404)
			return
		}
		if !auth.HasPerm(r.Context(), nil, apistruct.PermWrite) {
			w.WriteHeader(401)
			_ = json.NewEncoder(w).Encode(struct{ Error string }{"unauthorized: missing write permission"})
			return
		}

		c, err := a.ClientImportLocal(r.Context(), r.Body)
		if err != nil {
			w.WriteHeader(500)
			_ = json.NewEncoder(w).Encode(struct{ Error string }{err.Error()})
			return
		}
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(struct{ Cid cid.Cid }{c})
		if err != nil {
			log.Errorf("/rest/v0/import: Writing response failed: %+v", err)
			return
		}
	}
}
