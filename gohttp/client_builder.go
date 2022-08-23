package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	client            *http.Client
	headers           http.Header
	maxIdleConnection int
	connectionTimeout time.Duration
	responseTimeout   time.Duration
	disableTimeouts   bool
	userAgent         string
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetMaxIdleConnections(max int) ClientBuilder
	SetConnectionTimeout(ms time.Duration) ClientBuilder
	SetResponseTimeout(ms time.Duration) ClientBuilder
	SetDisableTimeouts(disable bool) ClientBuilder
	SetHttpClient(c *http.Client) ClientBuilder
	SetUserAgent(userAgent string) ClientBuilder

	Build() HttpClient
}

func Builder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() HttpClient {
	client := httpClient{
		builder: c,
	}

	return &client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(i int) ClientBuilder {
	c.maxIdleConnection = i
	return c
}

func (c *clientBuilder) SetConnectionTimeout(duration time.Duration) ClientBuilder {
	c.connectionTimeout = duration
	return c
}

func (c *clientBuilder) SetResponseTimeout(duration time.Duration) ClientBuilder {
	c.responseTimeout = duration
	return c
}

func (c *clientBuilder) SetDisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}

func (c *clientBuilder) SetUserAgent(userAgent string) ClientBuilder {
	c.userAgent = userAgent
	return c
}
