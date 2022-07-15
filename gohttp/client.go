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
	Get(url string, customHeaders http.Header) (*http.Response, error)
	Post(url string, customHeaders http.Header, body interface{}) (*http.Response, error)
	Put(url string, customHeaders http.Header, body interface{}) (*http.Response, error)
	Patch(url string, customHeaders http.Header, body interface{}) (*http.Response, error)
	Delete(url string, customHeaders http.Header) (*http.Response, error)
}

func (c *httpClient) Get(url string, customHeaders http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, customHeaders, nil)
}

func (c *httpClient) Post(url string, customHeaders http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, customHeaders, body)
}

func (c *httpClient) Put(url string, customHeaders http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, customHeaders, body)
}

func (c *httpClient) Patch(url string, customHeaders http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, customHeaders, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
