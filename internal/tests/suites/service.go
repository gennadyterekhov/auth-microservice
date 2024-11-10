package suites

import (
	"fmt"

	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
)

type WithService struct {
	WithFactory
	Service any
}

var _ interfaces.WithService = &WithService{}

func (s *WithService) SetupTest() {
	fmt.Println("(s *WithService) SetupTest()")
	fmt.Println()

	s.GetRepository().Clear()
}

// SetupSuite must be overloaded to pass real service to init
func (s *WithService) SetupSuite() {
	fmt.Println("(s *WithService) SetupSuite() in base ")
	fmt.Println()

	inits.InitServiceSuite(s, nil)
}

func (s *WithService) SetService(srv any) {
	s.Service = srv
}

func (s *WithService) GetService() any {
	return s.Service
}
