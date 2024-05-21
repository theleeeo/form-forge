package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/theleeeo/form-forge/runner"
	"gopkg.in/yaml.v3"
)

var (
	dbUser     string
	dbPassword string
	dbName     = "formforge"
	dbAddr     = "localhost:3307"
)

func executeSQLFile(db *sql.DB, filePath string) error {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading SQL file (%s): %v", filePath, err)
	}

	request := string(fileContent)
	_, err = db.Exec(request)
	if err != nil {
		return fmt.Errorf("error executing SQL file (%s): %v", filePath, err)
	}

	return nil
}

func load_env_vars() {
	if v := os.Getenv("DB_USER"); v != "" {
		dbUser = v
	} else {
		log.Fatal("DB_USER environment variable is not set")
	}

	if v := os.Getenv("DB_PASSWORD"); v != "" {
		dbPassword = v
	} else {
		log.Println("DB_PASSWORD environment variable is not set, using empty password")
	}

	if v := os.Getenv("DB_NAME"); v != "" {
		dbName = v
	} else {
		log.Printf("DB_NAME environment variable is not set, using default value (%s)\n", dbAddr)
	}

	if v := os.Getenv("DB_ADDR"); v != "" {
		dbAddr = v
	} else {
		log.Printf("DB_ADDR environment variable is not set, using default value (%s)\n", dbAddr)
	}
}

func loadConfig() (*runner.Config, error) {
	content, err := os.ReadFile("./.cfg.yml")
	if err != nil {
		return nil, err
	}

	var config runner.Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Println("error loading config file, moving on with env-vars:", err)
		load_env_vars()
	} else {
		dbUser = cfg.RepoCfg.User
		dbPassword = cfg.RepoCfg.Password
		dbName = cfg.RepoCfg.Database
		dbAddr = cfg.RepoCfg.Address
	}

	// Build the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		dbUser, dbPassword, dbAddr, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Could not ping the database: %v", err)
	}

	dir, err := os.ReadDir("migrations")
	if err != nil {
		log.Fatalf("Could not read migrations directory: %v", err)
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) != ".sql" {
			continue
		}

		if err := executeSQLFile(db, filepath.Join("migrations", file.Name())); err != nil {
			log.Fatalf("Failed to execute migration file (%s): %v", file.Name(), err)
		}
		log.Printf("Successfully executed migration file: %s", file.Name())
	}
}
