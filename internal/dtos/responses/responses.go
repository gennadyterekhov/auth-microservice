package responses

type Login struct {
	Token string `json:"token"`
}

type Register struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}
