package serializers

import (
	"encoding/json"
)

func Serialize(data interface{}) ([]byte, error) {
	serialized, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}

	return serialized, nil
}
