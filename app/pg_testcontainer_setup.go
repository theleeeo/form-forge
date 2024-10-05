package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	pgUser = "user"
	pgPass = "password"
)

type TestDB struct {
	container testcontainers.Container

	schema string

	Pool *pgxpool.Pool
}

func (t *TestDB) Shutdown() error {
	t.Pool.Close()

	if err := t.container.Terminate(context.Background()); err != nil {
		return fmt.Errorf("failed to terminate DB container: %w", err)
	}

	return nil
}

func (t *TestDB) Reset() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rows, err := t.Pool.Query(ctx, `
        SELECT tablename
        FROM pg_tables
        WHERE schemaname = 'public'
    `)
	if err != nil {
		return fmt.Errorf("failed to query table names: %w", err)
	}
	defer rows.Close()

	var tableNames []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return fmt.Errorf("failed to scan table name: %w", err)
		}
		tableNames = append(tableNames, tableName)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error occurred during rows iteration: %w", err)
	}

	for _, tableName := range tableNames {
		if _, err := t.Pool.Exec(ctx, fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName)); err != nil {
			return fmt.Errorf("failed to drop table %s: %w", tableName, err)
		}
	}

	if _, err := t.Pool.Exec(ctx, t.schema); err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	return err
}

func SetupTestPostgresql(dbName string, schema string) (testdb *TestDB, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	container, err := createContainer(ctx, dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to create container: %w", err)
	}
	defer func() {
		if err != nil {
			if terminateErr := container.Terminate(ctx); terminateErr != nil {
				log.Println("Failed to terminate the container:", terminateErr)
			}
		}

	}()

	host, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get container host: %w", err)
	}

	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return nil, fmt.Errorf("failed to get container port: %w", err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", pgUser, pgPass, host, port.Port(), dbName)
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}

	return &TestDB{
		container: container,
		schema:    schema,
		Pool:      pool,
	}, nil
}

func createContainer(ctx context.Context, dbName string) (testcontainers.Container, error) {
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:16",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_DB":       dbName,
				"POSTGRES_USER":     pgUser,
				"POSTGRES_PASSWORD": pgPass,
			},
			WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(30 * time.Second),
		},
		Started: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start PostgreSQL container: %w", err)
	}

	return container, nil
}
