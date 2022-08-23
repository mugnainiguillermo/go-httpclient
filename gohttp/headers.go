package gohttp

import (
	"github.com/mugnainiguillermo/go-httpclient/gomime"
	"net/http"
)

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *httpClient) getRequestHeaders(customHeaders http.Header) http.Header {
	headers := make(http.Header)

	for header, value := range c.builder.headers {
		if len(value) > 0 {
			headers.Set(header, value[0])
		}
	}

	for header, value := range customHeaders {
		if len(value) > 0 {
			headers.Set(header, value[0])
		}
	}

	if c.builder.userAgent != "" {
		if headers.Get(gomime.HeaderUserAgent) != "" {
			return headers
		}
		headers.Set(gomime.HeaderUserAgent, c.builder.userAgent)
	}
	return headers

	return headers
}
