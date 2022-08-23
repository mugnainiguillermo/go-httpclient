package gohttp

import (
	"testing"
)

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
