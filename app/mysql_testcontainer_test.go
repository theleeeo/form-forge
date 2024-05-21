package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// GetTestMariadb returns a connection to a test MariaDB database running as a testcontainer.
// The caller is responsible for closing the connection and stopping the container using the returned stopFunc.
func GetTestMariadb(dbName string) (db *sql.DB, stopFunc func() error, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create the MySQL container
	container, err := createContainer(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to create MySQL container: %w", err)
	}

	// Get the host and port of the MySQL container
	host, err := container.Host(ctx)
	if err != nil {
		if terminateErr := container.Terminate(ctx); terminateErr != nil {
			log.Println("Failed to terminate the DB container:", terminateErr)
		}
		return nil, nil, fmt.Errorf("Failed to get MySQL container host: %w", err)
	}
	port, err := container.MappedPort(ctx, "3306")
	if err != nil {
		if terminateErr := container.Terminate(ctx); terminateErr != nil {
			log.Println("Failed to terminate the DB container:", terminateErr)
		}
		return nil, nil, fmt.Errorf("Failed to get MySQL container port: %w", err)
	}

	// Setup the test database
	if err := setupDatabase(ctx, host, port.Port()); err != nil {
		if terminateErr := container.Terminate(ctx); terminateErr != nil {
			log.Println("Failed to terminate the DB container:", terminateErr)
		}
		return nil, nil, fmt.Errorf("Failed to setup MySQL database: %w", err)
	}

	// Connect to the test database
	dsn := fmt.Sprintf("root:password@tcp(%s:%s)/%s", host, port.Port(), dbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		if terminateErr := container.Terminate(ctx); terminateErr != nil {
			log.Println("Failed to terminate the DB container:", terminateErr)
		}
		return nil, nil, fmt.Errorf("Failed to connect to MySQL: %w", err)
	}

	return db, func() error {
		var closeErr error
		if err := db.Close(); err != nil {
			closeErr = fmt.Errorf("Failed to close MySQL connection: %w", err)
		}

		if err := container.Terminate(ctx); err != nil {
			closeErr = errors.Join(closeErr, fmt.Errorf("Failed to terminate MySQL container: %w", err))
		}

		return closeErr
	}, nil
}

func createContainer(ctx context.Context) (testcontainers.Container, error) {
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mariadb:11",
			ExposedPorts: []string{"3306/tcp"},
			Env: map[string]string{
				"MYSQL_ROOT_PASSWORD": "password",
			},
			// WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
			WaitingFor: wait.ForListeningPort("3306/tcp").WithStartupTimeout(2 * time.Minute),
		},
		Started: true,
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to start MySQL container: %w", err)
	}

	return container, nil
}

func setupDatabase(ctx context.Context, host, port string) error {
	dsn := fmt.Sprintf("root:password@tcp(%s:%s)/", host, port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("Failed to connect to MySQL: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("Failed to ping MySQL: %w", err)
	}

	// Create the test database
	_, err = db.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE %s", dbName))
	if err != nil {
		return fmt.Errorf("Failed to create test database: %w", err)
	}

	_, err = db.ExecContext(ctx, fmt.Sprintf("USE %s", dbName))
	if err != nil {
		return fmt.Errorf("Failed to create test database: %w", err)
	}

	if err := setupTables(ctx, db); err != nil {
		return fmt.Errorf("Failed to setup tables: %w", err)
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("Failed to close MySQL connection: %w", err)
	}

	return nil
}

func setupTables(ctx context.Context, db *sql.DB) error {
	entries, err := os.ReadDir(filepath.Join("..", "migrations"))
	if err != nil {
		return fmt.Errorf("Failed to read migrations directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}

		migration, err := os.ReadFile(filepath.Join("..", "migrations", entry.Name()))
		if err != nil {
			return fmt.Errorf("Failed to read migration file: %w", err)
		}

		if _, err := db.ExecContext(ctx, string(migration)); err != nil {
			return fmt.Errorf("Failed to execute migration: %w", err)
		}
	}

	return nil
}
