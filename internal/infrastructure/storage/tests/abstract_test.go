package tests

import (
	"context"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/storage"
	"github.com/gennadyterekhov/auth-microservice/internal/models"
	"github.com/gennadyterekhov/auth-microservice/internal/repositories/abstract"
	"github.com/gennadyterekhov/auth-microservice/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetOneBy(t *testing.T) {
	ctx := context.Background()

	cont, dbDsn, err := tests.CreatePostgresContainerAndRunMigrations(ctx)
	assert.NoError(t, err)

	repo := abstract.New(storage.NewDB(dbDsn))
	repo.Clear()

	_, err = repo.SelectOneBy(ctx, "id", 1, "name", "nm")
	assert.Equal(t, abstract.ErrorExpectedAtLeastOneResult, err.Error())
	if err != nil {
		return
	}

	_, err = repo.DB.ExecContext(ctx, "insert into categories (name) values('nm')")
	assert.NoError(t, err)

	cat, err := repo.SelectOneBy(ctx, "id", 1, "name", "nm")
	assert.NoError(t, err)
	assert.NotNil(t, cat)

	assert.NoError(t, cont.Terminate(ctx))
}

func TestGetOneByGeneric(t *testing.T) {
	ctx := context.Background()

	cont, dbDsn, err := tests.CreatePostgresContainerAndRunMigrations(ctx)
	assert.NoError(t, err)

	repo := abstract.New(storage.NewDB(dbDsn))
	repo.Clear()

	_, err = abstract.GenericSelectOneBy[*models.Category](ctx, repo, "id", 1, "name", "nm")
	assert.Equal(t, abstract.ErrorExpectedAtLeastOneResult, err.Error())
	if err != nil {
		return
	}

	_, err = repo.DB.ExecContext(ctx, "insert into categories (name) values('nm')")
	assert.NoError(t, err)

	cat, err := abstract.GenericSelectOneBy[*models.Category](ctx, repo, "id", 1, "name", "nm")
	assert.NoError(t, err)
	assert.NotNil(t, cat)

	_, err = repo.DB.ExecContext(ctx, "insert into users (login, password) values('a', 'b')")
	assert.NoError(t, err)
	user, err := abstract.GenericSelectOneBy[*models.User](ctx, repo, "login", "a")
	assert.NoError(t, err)
	assert.NotNil(t, user)

	assert.NoError(t, cont.Terminate(ctx))
}

func TestFindBy(t *testing.T) {
	t.Skipf("fix ")

	ctx := context.Background()

	cont, dbDsn, err := tests.CreatePostgresContainerAndRunMigrations(ctx)
	assert.NoError(t, err)

	repo := abstract.New(storage.NewDB(dbDsn))
	repo.Clear()

	cats, err := repo.FindBy(ctx, "id", 1, "name", "nm")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(cats))

	_, err = repo.DB.ExecContext(ctx, "insert into categories (name) values('nm1')")
	assert.NoError(t, err)
	_, err = repo.DB.ExecContext(ctx, "insert into categories (name) values('nm2')")
	assert.NoError(t, err)

	cats, err = repo.FindBy(ctx, "notAColumn", 1)
	assert.Error(t, err)
	assert.Equal(t, 0, len(cats))

	cats, err = repo.FindBy(ctx, "name", "nm1")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(cats))

	assert.NoError(t, cont.Terminate(ctx))
}

func TestFindByGeneric(t *testing.T) {
	ctx := context.Background()

	cont, dbDsn, err := tests.CreatePostgresContainerAndRunMigrations(ctx)
	assert.NoError(t, err)

	repo := abstract.New(storage.NewDB(dbDsn))
	repo.Clear()

	cats, err := abstract.FindByGeneric[*models.Category](ctx, repo, "id", 1, "name", "nm")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(cats))

	_, err = repo.DB.ExecContext(ctx, "insert into categories (name) values('nm1')")
	assert.NoError(t, err)
	_, err = repo.DB.ExecContext(ctx, "insert into categories (name) values('nm2')")
	assert.NoError(t, err)

	cats, err = abstract.FindByGeneric[*models.Category](ctx, repo, "notAColumn", 1)
	assert.Error(t, err)
	assert.Equal(t, 0, len(cats))

	cats, err = abstract.FindByGeneric[*models.Category](ctx, repo, "name", "nm1")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(cats))

	assert.NoError(t, cont.Terminate(ctx))
}
