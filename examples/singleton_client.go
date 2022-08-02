package examples

import (
	"github.com/mugnainiguillermo/go-httpclient/gohttp"
	"time"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.HttpClient {
	client := gohttp.Builder().
		SetMaxIdleConnections(5).
		SetConnectionTimeout(1 * time.Second).
		SetResponseTimeout(1 * time.Second).
		Build()

	return client
}
