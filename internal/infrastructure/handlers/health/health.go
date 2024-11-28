package health

import (
	"net/http"
)

func Health(res http.ResponseWriter, req *http.Request) {
	_, err := res.Write([]byte(`{"status":"ok", "code":200}`))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}
