package serializers

import (
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/serializer"
)

func New() *serializer.Pack {
	return &serializer.Pack{
		Categories: NewCategory(),
		Users:      NewBase(),
		Orders:     NewOrder(),
		Register:   NewBase(),
		Login:      NewBase(),
	}
}
