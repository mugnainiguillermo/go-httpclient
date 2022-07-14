package main

import (
	"fmt"
	"github.com/mugnainiguillermo/go-httpclient/gohttp"
	"io"
	"net/http"
)

var (
	githubClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Accept", "application/json")

	client.SetCommonHeaders(commonHeaders)

	return client
}

func main() {
	getUrl(false)
	getUrl(true)
	getUrl(true)
	getUrl(false)
}

func getUrl(flag bool) {
	customHeaders := make(http.Header)

	if flag == true {
		customHeaders.Set("Accept", "application/xml")
	}

	response, err := githubClient.Get("https://api.github.com", customHeaders)

	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)

	bytes, _ := io.ReadAll(response.Body)

	fmt.Println(string(bytes))
}
