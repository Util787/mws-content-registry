package httpclients

import (
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	defaultTimeOut    = 30 * time.Second
	defaultRetryCount = 3
)

func NewRestyClient() *resty.Client {
	rclient := resty.New()

	rclient.
		SetTimeout(defaultTimeOut).
		SetRetryCount(defaultRetryCount)

	return rclient
}
