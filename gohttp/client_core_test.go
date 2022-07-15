package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	client := httpClient{}

	commonHeaders := make(http.Header)
	commonHeaders.Set("User-Agent", "cool-http-client")
	commonHeaders.Set("Content-Type", "application/json")

	client.SetCommonHeaders(commonHeaders)

	customHeaders := make(http.Header)
	customHeaders.Set("X-Request-ID", "abc-123")

	headers := client.getRequestHeaders(customHeaders)

	if len(headers) != 3 {
		t.Error("we expected 3 headers")
	}

	if headers.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid user agent header")
	}

	if headers.Get("Content-Type") != "application/json" {
		t.Error("invalid content type header")
	}

	if headers.Get("X-Request-ID") != "abc-123" {
		t.Error("invalid request id header")
	}
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}

	t.Run("WithNilBody", func(t *testing.T) {
		requestBody, err := client.getRequestBody("", nil)

		if err != nil {
			t.Error("no error expected when passing nil body")
		}

		if requestBody != nil {
			t.Error("no body expected when passing nil body")
		}
	})

	t.Run("WithJson", func(t *testing.T) {
		body := []string{"one", "two"}
		requestBody, err := client.getRequestBody("application/json", body)

		if err != nil {
			t.Error("no error expected when marshalling slice as json")
		}

		if string(requestBody) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})

	t.Run("WithXml", func(t *testing.T) {
		body := []string{"two", "three"}

		requestBody, err := client.getRequestBody("application/xml", body)

		if err != nil {
			t.Error("no error expected when marshalling slice as xml")
		}

		if string(requestBody) != `<string>two</string><string>three</string>` {
			t.Error("invalid xml body obtained")
		}
	})

	t.Run("WithJsonByDefault", func(t *testing.T) {
		body := []string{"three", "four"}

		requestBody, err := client.getRequestBody("", body)

		if err != nil {
			t.Error("no error expected when marshalling slice as json by default")
		}

		if string(requestBody) != `["three","four"]` {
			t.Error("invalid json body obtained")
		}
	})
}
