package storage

import (
	"database/sql"

	"github.com/gennadyterekhov/auth-microservice/internal/logger"
	"github.com/gennadyterekhov/auth-microservice/internal/repositories"
	"github.com/gennadyterekhov/auth-microservice/internal/repositories/abstract"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB(dsn string) abstract.QueryMaker {
	logger.Debugln("opening database connection with dsn ", dsn)

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Debugln("could not connect to db using dsn: " + dsn + " " + err.Error())
		panic(err)
	}

	return conn
}

// NewRepo exists because this pkg can depend on repo, but repo cannot depend on this pkg
func NewRepo(dsn string) *repositories.Repository {
	logger.Debugln("opening database connection with dsn ", dsn)
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Debugln("could not connect to db using dsn: " + dsn + " " + err.Error())
		panic(err)
	}

	return repositories.New(conn)
}
