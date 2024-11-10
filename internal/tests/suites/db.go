package suites

import (
	"context"
	"fmt"

	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces/repositories"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

type WithDb struct {
	Abstract
	dbContainer testcontainers.Container
	repository  repositories.RepositoryInterface
}

var _ interfaces.WithDb = &WithDb{}

func (s *WithDb) SetupTest() {
	fmt.Println("(s *WithDb) SetupTest()")
	fmt.Println()

	if s.GetRepository() != nil {
		s.GetRepository().Clear()
	}
}

// SetupSuite - when overriding, don't forget to call InitBaseSuite
func (s *WithDb) SetupSuite() {
	fmt.Println("(s *WithDb) SetupSuite() in base ")
	fmt.Println()

	inits.InitDbSuite(s)
}

func (s *WithDb) TearDownSuite() {
	fmt.Println("(s *WithDb) TearDownSuite() in base ")
	fmt.Println()
	if s.GetRepository() != nil {
		s.GetRepository().Clear()
	}
	if s.dbContainer != nil {
		assert.NoError(s.T(), s.dbContainer.Terminate(context.Background()))
		s.SetDBContainer(nil)
	}
}

func (s *WithDb) SetDBContainer(cont testcontainers.Container) {
	s.dbContainer = cont
}

func (s *WithDb) GetDBContainer() testcontainers.Container {
	return s.dbContainer
}

func (s *WithDb) SetRepository(repo repositories.RepositoryInterface) {
	s.repository = repo
}

func (s *WithDb) GetRepository() repositories.RepositoryInterface {
	return s.repository
}
