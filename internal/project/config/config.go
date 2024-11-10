package config

import (
	"flag"
	"os"
	"path"

	"github.com/gennadyterekhov/auth-microservice/internal/logger"
	"github.com/gennadyterekhov/auth-microservice/internal/project"
	"github.com/joho/godotenv"
)

type Config struct {
	Addr  string
	DBDsn string
}

func New() *Config {
	return getConfig()
}

func getConfig() *Config {
	var addressFlag *string
	var DBDsnFlag *string

	if flag.Lookup("a") == nil {
		addressFlag = flag.String(
			"a",
			"localhost:8080",
			"[address] Net address host:port without protocol",
		)
	}
	if flag.Lookup("d") == nil {
		DBDsnFlag = flag.String(
			"d",
			"",
			"[db dsn] format: `host=%s user=%s password=%s dbname=%s sslmode=%s`",
		)
	}

	flag.Parse()
	flags := Config{}

	if addressFlag != nil {
		flags.Addr = *addressFlag
	}
	if DBDsnFlag != nil {
		flags.DBDsn = *DBDsnFlag
	}

	overwriteWithEnv(&flags)

	return &flags
}

func overwriteWithEnv(flags *Config) {
	pr, err := project.GetProjectRoot()
	if err != nil {
		logger.Errorln("could not find project root", err.Error())
	}

	err = godotenv.Load(path.Join(pr, ".env"))
	if err != nil {
		logger.Errorln("could not load env file", err.Error())
	}

	flags.Addr = getAddress(flags.Addr)
	flags.DBDsn = getDBDsn(flags.DBDsn)
}

func getAddress(current string) string {
	return getStringFromEnvOrFallback("RUN_ADDRESS", current)
}

func getDBDsn(current string) string {
	return getStringFromEnvOrFallback("DATABASE_URI", current)
}

func getStringFromEnvOrFallback(envKey string, fallback string) string {
	fromEnv, ok := os.LookupEnv(envKey)
	if ok {
		return fromEnv
	}

	return fallback
}
