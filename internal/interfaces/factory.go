package interfaces

import (
	"github.com/gennadyterekhov/auth-microservice/internal/models/responses"
)

type Interface interface {
	RegisterForTest(login string, password string) *responses.Register
}
