package gohttp

import (
	"net/http"
	"sync"
)

type httpClient struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

type HttpClient interface {
	Get(url string, customHeaders ...http.Header) (*Response, error)
	Post(url string, body interface{}, customHeaders ...http.Header) (*Response, error)
	Put(url string, body interface{}, customHeaders ...http.Header) (*Response, error)
	Patch(url string, body interface{}, customHeaders ...http.Header) (*Response, error)
	Delete(url string, customHeaders ...http.Header) (*Response, error)
	Options(url string, customHeaders ...http.Header) (*Response, error)
}

func (c *httpClient) Get(url string, customHeaders ...http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, getHeaders(customHeaders...), nil)
}

func (c *httpClient) Post(url string, body interface{}, customHeaders ...http.Header) (*Response, error) {
	return c.do(http.MethodPost, url, getHeaders(customHeaders...), body)
}

func (c *httpClient) Put(url string, body interface{}, customHeaders ...http.Header) (*Response, error) {
	return c.do(http.MethodPut, url, getHeaders(customHeaders...), body)
}

func (c *httpClient) Patch(url string, body interface{}, customHeaders ...http.Header) (*Response, error) {
	return c.do(http.MethodPatch, url, getHeaders(customHeaders...), body)
}

func (c *httpClient) Delete(url string, customHeaders ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, getHeaders(customHeaders...), nil)
}

func (c *httpClient) Options(url string, customHeaders ...http.Header) (*Response, error) {
	return c.do(http.MethodOptions, url, getHeaders(customHeaders...), nil)
}
