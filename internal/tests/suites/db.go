package suites

import (
	"context"

	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
)

type WithDb struct {
	suite.Suite
	dbContainer testcontainers.Container
	repository  interfaces.RepositoryInterface
}

var _ interfaces.WithDb = &WithDb{}

func (s *WithDb) SetupTest() {
	if s.GetRepository() != nil {
		s.GetRepository().Clear()
	}
}

// SetupSuite - when overriding, don't forget to call InitBaseSuite
func (s *WithDb) SetupSuite() {
	inits.InitDbSuite(s)
}

func (s *WithDb) TearDownSuite() {
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

func (s *WithDb) SetRepository(repo interfaces.RepositoryInterface) {
	s.repository = repo
}

func (s *WithDb) GetRepository() interfaces.RepositoryInterface {
	return s.repository
}
