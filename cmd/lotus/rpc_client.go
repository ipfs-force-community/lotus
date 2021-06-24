package main

import (
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/venus-auth/auth"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type JWTClient struct {
	cli *resty.Client
}

func NewJWTClient(url string) *JWTClient {
	client := resty.New().
		SetHostURL(url).
		SetRetryCount(0).
		SetTimeout(time.Second).
		SetHeader("Accept", "application/json")

	return &JWTClient{
		cli: client,
	}
}

// Verify: post method for Verify token
// @spanId: local service unique Id
// @serviceName: e.g. venus
// @preHost: the IP of the request server
// @host: local service IP
// @token: jwt token gen from this service
func (c *JWTClient) Verify(spanId, serviceName, preHost, host, token string) (*auth.VerifyResponse, error) {
	response, err := c.cli.R().SetHeader("X-Forwarded-For", host).
		SetHeader("X-Real-Ip", host).
		SetHeader("spanId", spanId).
		SetHeader("preHost", preHost).
		SetHeader("svcName", serviceName).
		SetHeader("Origin", host).
		SetFormData(map[string]string{
			"token": token,
		}).Post("/verify")
	if err != nil {
		return nil, err
	}
	switch response.StatusCode() {
	case http.StatusOK:
		var res = new(auth.VerifyResponse)
		response.Body()
		err = json.Unmarshal(response.Body(), res)
		return res, err
	default:
		response.Result()
		return nil, fmt.Errorf("response code is : %d, msg:%s", response.StatusCode(), response.Body())
	}
}
