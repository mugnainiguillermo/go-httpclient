package gohttp

import "net/http"

func (c *httpClient) do(method string, url string, customHeaders http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic("unable to create request")
	}

	request.Header = c.getRequestHeaders(customHeaders)

	return client.Do(request)
}

func (c *httpClient) getRequestHeaders(customHeaders http.Header) http.Header {
	headers := make(http.Header)

	for header, value := range c.CommonHeaders {
		if len(value) > 0 {
			headers.Set(header, value[0])
		}
	}

	for header, value := range customHeaders {
		if len(value) > 0 {
			headers.Set(header, value[0])
		}
	}
}
