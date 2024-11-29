package inits

import (
	"context"

	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/storage"
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
	"github.com/gennadyterekhov/auth-microservice/internal/tests"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

// InitDbSuite is needed so we can initialize any kind of different suites with 1 piece of code.
// it's different from SetupSuite, because SetupSuite can be overridden to inject suite-specific dependencies later
func InitDbSuite[T interfaces.WithDb](genericSuite T) {
	ctx := context.Background()
	var cont testcontainers.Container
	var dbDsn string
	var err error

	if genericSuite.GetDBContainer() == nil {
		cont, dbDsn, err = tests.CreatePostgresContainerAndRunMigrations(ctx)
		require.NoError(genericSuite.T(), err)
		genericSuite.SetDBContainer(cont)
	}

	if genericSuite.GetRepository() != nil {
		genericSuite.GetRepository().Clear()
	}

	if genericSuite.GetRepository() == nil {
		repo, err := storage.NewRepo(dbDsn)
		require.NoError(genericSuite.T(), err)
		repo.Clear()
		genericSuite.SetRepository(repo)
	}
}
