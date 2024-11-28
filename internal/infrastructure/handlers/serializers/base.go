package serializers

import (
	"encoding/json"
	"net/http"
)

func serialize(data interface{}) ([]byte, error) {
	serialized, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}

	return serialized, nil
}

func WriteToWriter(res http.ResponseWriter, data interface{}) error {
	resBody, err := serialize(data)
	if err != nil {
		return err
	}
	_, err = res.Write(resBody)

	return err
}
