package config

import (
	"os"

	"github.com/pkg/errors"
)

func LoadDatabaseCredentials() (string, error) {
	database_dsn := os.Getenv("APP_DATABASE_DSN")

	if len(database_dsn) == 0 {
		return "", errors.New("APP_DATABASE_DSN not set")
	}

	return database_dsn, nil
}
