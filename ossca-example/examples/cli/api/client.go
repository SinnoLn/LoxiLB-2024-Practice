package api

import (
	"net/http"
	"time"
)

const (
	loxiProvider    = "company_name"
	loxiApiVersion  = "v1"
	AccountResource = "account/all"
)

type LoxiClient struct {
	restClient RESTClient
}

func NewLoxiClient(o *RESTOptions) *LoxiClient {
	return &LoxiClient{
		restClient: RESTClient{
			Options: *o,
			Client: &http.Client{
				Timeout: time.Second * time.Duration(o.Timeout),
			},
		},
	}
}

func (l *LoxiClient) Account() *Account {
	return &Account{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   loxiProvider,
				apiVersion: loxiApiVersion,
				resource:   AccountResource,
			},
		},
	}
}
