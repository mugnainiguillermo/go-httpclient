package main

import (
	"fmt"
	"github.com/mugnainiguillermo/go-httpclient/gohttp"
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

type Endpoints struct {
	currentUserUrl                   string `json:"current_user_url"`
	currentUserAuthorizationsHtmlUrl string `json:"current_user_authorizations_html_url"`
	authorizationsUrl                string `json:"authorizations_url"`
}

func getUrl() {
	customHeaders := make(http.Header)

	response, err := githubClient.Get("https://api.github.com", customHeaders)

	if err != nil {
		panic(err)
	}
	fmt.Println(response.Status())

	var endpoints Endpoints

	if err := response.UnmarshallJson(&endpoints); err != nil {
		panic(err)
	}

	fmt.Println(endpoints.currentUserUrl)
	fmt.Println(endpoints.currentUserAuthorizationsHtmlUrl)
	fmt.Println(endpoints.authorizationsUrl)
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
	fmt.Println(response.StatusCode())
	fmt.Println(response.String())
}
