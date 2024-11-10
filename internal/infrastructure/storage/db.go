package storage

import (
	"database/sql"
	"fmt"

	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/storage/migrations"
	"github.com/gennadyterekhov/auth-microservice/internal/logger"
	"github.com/gennadyterekhov/auth-microservice/internal/repositories"
	"github.com/gennadyterekhov/auth-microservice/internal/repositories/abstract"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB(dsn string) abstract.QueryMaker {
	fmt.Println("opening database connection")

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Debugln("could not connect to db ", err.Error())
		panic(err)
	}

	fmt.Println("running migrations")
	err = migrations.RunMigrationsOnConnection(conn)
	if err != nil {
		logger.Debugln("could not run migrations ", err.Error())
		panic(err)
	}

	return conn
}

// NewRepo exists because this pkg can depend on repo, but repo cannot depend on this pkg
func NewRepo(dsn string) *repositories.Repository {
	return repositories.New(NewDB(dsn))
}
