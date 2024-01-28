package client

import (
	"net/http"
	"net/url"
)

const baseURL = "https://www.thecocktaildb.com/api/json/v1/1/"

type Parameters = map[string]string

func NewClient() http.Client {
	return http.Client{}
}

func NewRequest(endpoint string, params Parameters) (*http.Request, error) {
	uri, errUri := url.JoinPath(baseURL, endpoint)
	if errUri != nil {
		return nil, errUri
	}
	request, errReq := http.NewRequest("GET", uri, nil)
	if errReq != nil {
		return nil, errReq
	}

	queries := request.URL.Query()
	for k, v := range params {
		queries.Add(k, v)
	}
	return request, nil
}
