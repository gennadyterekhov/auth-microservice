package requests

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Register struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// YandexCloudRequest https://yandex.cloud/ru/docs/functions/concepts/function-invoke#request
type YandexCloudRequest struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`

	HTTPMethod      string            `json:"httpMethod"`
	Path            string            `json:"path"`
	Headers         map[string]string `json:"headers"`
	QueryString     map[string]string `json:"queryStringParameters"`
	RequestContext  map[string]string `json:"requestContext"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
}
