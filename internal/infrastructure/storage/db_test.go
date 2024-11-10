package storage

import (
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/repositories/abstract"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type dbTest struct {
	suite.Suite
}

const TestDBDSN = "host=localhost user=authmcrsrv_user password=authmcrsrv_pass dbname=authmcrsrv_db_test sslmode=disable"

func initDB() abstract.QueryMaker {
	return NewDB(TestDBDSN)
}

func (suite *dbTest) TestDbExists() {
	_db := initDB()
	err := _db.Ping()
	assert.NoError(suite.T(), err)
}

func (suite *dbTest) TestDbTableExists() {
	_db := initDB()
	rawSQLString := "select * from users limit 1;"
	_, err := _db.Exec(rawSQLString)
	assert.NoError(suite.T(), err)

	rawSQLString = "select * from orders limit 1;"
	_, err = _db.Exec(rawSQLString)
	assert.NoError(suite.T(), err)
}

func TestDb(t *testing.T) {
	suite.Run(t, new(dbTest))
}
