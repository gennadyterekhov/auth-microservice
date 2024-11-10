package migrations

import (
	"database/sql"
	"path"

	"github.com/gennadyterekhov/auth-microservice/internal/project"
	"github.com/pkg/errors"
	"github.com/pressly/goose"
)

func RunMigrationsOnConnection(db *sql.DB) error {
	dir, err := getMigrationsDir()
	if err != nil {
		return err
	}
	return goose.Up(db, dir)
}

func getMigrationsDir() (string, error) {
	pr, err := project.GetProjectRoot()
	if err != nil {
		return "", errors.Wrap(err, "error getting project root")
	}

	return path.Join(pr, "migrations"), nil
}
