package repositories

import (
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces/repositories"
	"github.com/gennadyterekhov/auth-microservice/internal/logger"
	"github.com/gennadyterekhov/auth-microservice/internal/repositories/abstract"
)

// Repository is a single point of access to the database. it's not divided between entities - one repo for all
type Repository struct {
	DB                 abstract.QueryMaker
	abstractRepository *abstract.Repository
}

func New(db abstract.QueryMaker) *Repository {
	return &Repository{
		DB:                 db,
		abstractRepository: abstract.New(db),
	}
}

var (
	_ repositories.RepositoryInterface = NewMock()
	_ repositories.RepositoryInterface = New(nil)
)

// Clear is used only in tests
func (repo *Repository) Clear() {
	queries := []string{
		"delete from users;",
	}

	for _, v := range queries {
		_, err := repo.DB.Exec(v)
		if err != nil {
			logger.Errorln(err.Error())
		}
	}
}
