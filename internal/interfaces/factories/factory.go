package factories

import (
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/responses"
)

type Interface interface {
	RegisterForTest(login string, password string) *responses.Register
}
