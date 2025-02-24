package main

import (
	"log"

	config "github.com/andy-takker/simple_server/internal/adapters/database"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dsn, err := config.LoadDatabaseCredentials()
	if err != nil {
		log.Fatalf("Unable to load database credentials: %v\n", err)
	}

	m, err := migrate.New("file://migrations", dsn+"?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to create migrations: %v\n", err)
	}
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("Database already up to date")
			return
		}
		log.Fatalf("Unable to run migrations: %v\n", err)
	}
}
