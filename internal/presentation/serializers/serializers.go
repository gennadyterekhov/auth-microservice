package serializers

import (
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/serializer"
)

func New() *serializer.Pack {
	return &serializer.Pack{
		Users:      NewBase(),
		Register:   NewBase(),
		Login:      NewBase(),
	}
}
