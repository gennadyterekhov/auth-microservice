package serializers

import (
	"encoding/json"

	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/serializer"
)

type Base struct{}

var _ serializer.Serializer = New()

func New() *Base {
	return &Base{}
}

func (s *Base) Serialize(data interface{}) ([]byte, error) {
	serialized, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}

	return serialized, nil
}
