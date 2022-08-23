package examples

import (
	"github.com/mugnainiguillermo/go-httpclient/gohttp"
	"github.com/mugnainiguillermo/go-httpclient/gomime"
	"net/http"
	"time"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.HttpClient {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	client := gohttp.Builder().
		SetHeaders(headers).
		SetMaxIdleConnections(5).
		SetConnectionTimeout(1 * time.Second).
		SetResponseTimeout(1 * time.Second).
		SetUserAgent("Guille-Computer").
		Build()

	return client
}
