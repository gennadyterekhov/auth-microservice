package suites

import (
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
)

type WithFactory struct {
	WithDb
	factory interfaces.Interface
}

var _ interfaces.WithFactory = &WithFactory{}

// SetupSuite - when overriding, don't forget to call InitBaseSuite
func (s *WithFactory) SetupSuite() {
	inits.InitFactorySuite(s)
}

func (s *WithFactory) SetFactory(fact interfaces.Interface) {
	s.factory = fact
}

func (s *WithFactory) GetFactory() interfaces.Interface {
	return s.factory
}
