package factories

import (
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/responses"
	models2 "github.com/gennadyterekhov/auth-microservice/internal/models"
)

type Interface interface {
	NewUser(name string) *models2.User
	RegisterForTest(login string, password string) *responses.Register
	RegisterRandForTest() *responses.Register
}
