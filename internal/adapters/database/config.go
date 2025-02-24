package config

import (
	"os"

	"github.com/pkg/errors"
)

func LoadDatabaseCredentials() (string, error) {
	database_dsn := os.Getenv("APP_DATABASE_DSN")

	database_user := os.Getenv("APP_DATABASE_USER")
	if len(database_user) == 0 {
		return "", errors.New("APP_DATABASE_USER not set")
	}

	database_password := os.Getenv("APP_DATABASE_PASSWORD")
	if len(database_password) == 0 {
		return "", errors.New("APP_DATABASE_PASSWORD not set")
	}

	database_host := os.Getenv("APP_DATABASE_HOST")
	if len(database_host) == 0 {
		return "", errors.New("APP_DATABASE_HOST not set")
	}

	database_port := os.Getenv("APP_DATABASE_PORT")
	if len(database_port) == 0 {
		return "", errors.New("APP_DATABASE_PORT not set")
	}

	database_name := os.Getenv("APP_DATABASE_NAME")
	if len(database_name) == 0 {
		return "", errors.New("APP_DATABASE_NAME not set")
	}

	database_dsn = "postgres://" + database_user + ":" + database_password + "@" + database_host + ":" + database_port + "/" + database_name
	return database_dsn, nil
}
