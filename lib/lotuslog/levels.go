package lotuslog

import (
	logging "github.com/ipfs/go-log/v2"
)

func SetupLogLevels() {
	logging.SetLogLevel("*", "fatal")
	//if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
	//	logging.SetLogLevel("*", "fatal")
	//	_ = logging.SetLogLevel("*", "INFO")
	//	_ = logging.SetLogLevel("dht", "ERROR")
	//	_ = logging.SetLogLevel("swarm2", "WARN")
	//	_ = logging.SetLogLevel("bitswap", "WARN")
	//	//_ = logging.SetLogLevel("pubsub", "WARN")
	//	_ = logging.SetLogLevel("connmgr", "WARN")
	//	_ = logging.SetLogLevel("advmgr", "DEBUG")
	//	_ = logging.SetLogLevel("stores", "DEBUG")
	//	_ = logging.SetLogLevel("nat", "INFO")
	//}
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
