package tests

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/storage/migrations"
	"github.com/gennadyterekhov/auth-microservice/internal/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func CreatePostgresContainerAndRunMigrations(ctx context.Context) (testcontainers.Container, string, error) {
	cont, dbName, err := createPostgresContainer(ctx)
	err = runMigrations(dbName)
	if err != nil {
		panic(err)
	}
	return cont, dbName, err
}

func createPostgresContainer(ctx context.Context) (testcontainers.Container, string, error) {
	randint := rand.Intn(100)
	dbName := fmt.Sprintf("authmcrsrv_db_test_%d", randint)
	req := testcontainers.ContainerRequest{
		Image:        "postgres:17-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_USER":     "user",
			"POSTGRES_DB":       dbName,
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"), // Wait until the port is open
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true, // Start the container immediately
	})
	if err != nil {
		return nil, "", fmt.Errorf("failed to start container: %w", err)
	}

	dbURL, err := getDbUrl(ctx, container, dbName)
	if err != nil {
		return nil, "", err
	}
	return container, dbURL, nil
}

func getDbUrl(ctx context.Context, container testcontainers.Container, dbName string) (string, error) {
	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return "", fmt.Errorf("failed to get mapped port: %w", err)
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get host IP: %w", err)
	}

	dbURL := fmt.Sprintf("postgres://user:password@%s:%s/%s?sslmode=disable", hostIP, mappedPort.Port(), dbName)
	return dbURL, nil
}

func runMigrations(dbURL string) error {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return fmt.Errorf("failed to open db connection: %w", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Debugln("error when closing", err.Error())
		}
	}(db)

	err = migrations.RunMigrationsOnConnection(db)
	if err != nil {
		return fmt.Errorf("failed to run migrations: " + err.Error())
	}

	return nil
}
