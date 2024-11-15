package config

import (
	"os"
	"path"

	"github.com/gennadyterekhov/auth-microservice/internal/logger"
	"github.com/gennadyterekhov/auth-microservice/internal/project"
	"github.com/joho/godotenv"
)

const (
	defaultAddr  = "localhost:8080"
	defaultDbUrl = "host=psql port=5432 user=authmcrsrv_user password=authmcrsrv_pass dbname=authmcrsrv_db sslmode=disable"
)

type Config struct {
	Addr  string
	DBDsn string
}

func New() *Config {
	return getConfig()
}

func getConfig() *Config {
	pr, err := project.GetProjectRoot()
	if err != nil {
		logger.Errorln("could not find project root", err.Error())
	}

	err = godotenv.Load(path.Join(pr, ".env"))
	if err != nil {
		logger.Errorln("could not load env file", err.Error())
	}

	return &Config{
		Addr:  getStringFromEnvOrFallback("RUN_ADDRESS", defaultAddr),
		DBDsn: getStringFromEnvOrFallback("DATABASE_URI", defaultDbUrl),
	}
}

func getStringFromEnvOrFallback(envKey string, fallback string) string {
	fromEnv, ok := os.LookupEnv(envKey)
	if ok {
		return fromEnv
	}

	return fallback
}
