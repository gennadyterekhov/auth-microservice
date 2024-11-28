package requests

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Register struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
