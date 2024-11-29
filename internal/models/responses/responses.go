package responses

import "net/http"

type Login struct {
	Token string `json:"token"`
}

type Register struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}

type YandexCloudResponse struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
	Headers    http.Header `json:"headers"`
}
