package suites

import (
	"fmt"

	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
	factoriesInterface "github.com/gennadyterekhov/auth-microservice/internal/interfaces/factories"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
)

type WithFactory struct {
	WithDb
	factory factoriesInterface.Interface
}

var _ interfaces.WithFactory = &WithFactory{}

// SetupSuite - when overriding, don't forget to call InitBaseSuite
func (s *WithFactory) SetupSuite() {
	fmt.Println("(s *WithFactory) WithFactory() ")
	fmt.Println()

	inits.InitFactorySuite(s)
}

func (s *WithFactory) SetFactory(fact factoriesInterface.Interface) {
	s.factory = fact
}

func (s *WithFactory) GetFactory() factoriesInterface.Interface {
	return s.factory
}
