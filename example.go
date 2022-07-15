package main

import (
	"fmt"
	"github.com/mugnainiguillermo/go-httpclient/gohttp"
	"io"
	"net/http"
	"time"
)

var (
	githubClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	commonHeaders := make(http.Header)
	commonHeaders.Set("Accept", "application/json")

	client := gohttp.Builder().
		SetHeaders(commonHeaders).
		SetMaxIdleConnections(5).
		SetConnectionTimeout(1 * time.Second).
		SetResponseTimeout(1 * time.Second).
		Build()

	return client
}

func main() {
	getUrl()
}

func getUrl() {
	customHeaders := make(http.Header)

	//customHeaders.Set("Accept", "application/xml")

	response, err := githubClient.Get("https://api.github.com", customHeaders)

	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)

	bytes, _ := io.ReadAll(response.Body)

	fmt.Println(string(bytes))
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createUser(user User) {
	response, err := githubClient.Post("https://api.github.com", nil, user)

	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)

	bytes, _ := io.ReadAll(response.Body)

	fmt.Println(string(bytes))
}
