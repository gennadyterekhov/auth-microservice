package config

import (
	"fmt"
	"os"
	"path"

	"github.com/gennadyterekhov/auth-microservice/internal/logger"
	"github.com/gennadyterekhov/auth-microservice/internal/project"
	"github.com/joho/godotenv"
)

const (
	defaultAddr  = "0.0.0.0:8080"
	defaultDbUrl = "host=localhost port=5432 user=authmcrsrv_user password=authmcrsrv_pass dbname=authmcrsrv_db sslmode=disable"
)

type Config struct {
	Addr    string
	DBDsn   string
	IsHttps bool
}

func New() (*Config, error) {
	conf := getConfig()

	if conf.Addr == "" || conf.DBDsn == "" {
		return nil, fmt.Errorf("some required values are empty")
	}
	return conf, nil
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
		Addr:    getStringFromEnvOrFallback("RUN_ADDRESS", defaultAddr),
		DBDsn:   getStringFromEnvOrFallback("DATABASE_URI", defaultDbUrl),
		IsHttps: getBoolFromEnvOrFallback("IS_HTTPS", false),
	}
}

func getBoolFromEnvOrFallback(envKey string, fallback bool) bool {
	fromEnv, ok := os.LookupEnv(envKey)
	if ok {
		if fromEnv == "true" || fromEnv == "True" || fromEnv == "TRUE" || fromEnv == "1" {
			return true
		}
		return false
	}

	return fallback
}

func getStringFromEnvOrFallback(envKey string, fallback string) string {
	fromEnv, ok := os.LookupEnv(envKey)
	if ok {
		return fromEnv
	}

	return fallback
}
