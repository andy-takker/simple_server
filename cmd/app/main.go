package main

import (
	"context"
	"log"
	"time"

	config "github.com/andy-takker/simple_server/internal/adapters/database"
	user_repository "github.com/andy-takker/simple_server/internal/adapters/database/repositories/user"
	services "github.com/andy-takker/simple_server/internal/domain/services"
	handlers "github.com/andy-takker/simple_server/internal/presentors/rest/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dsn, err := config.LoadDatabaseCredentials()
	if err != nil {
		log.Fatalf("Unable to load database credentials: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	user_repo := user_repository.NewRepository(dbpool)
	user_service := services.NewService(user_repo)
	user_handler := handlers.NewUserHandler(user_service)
	healthcheck_handler := handlers.NewHealthcheckHandler()

	r := gin.Default()
	user_handler.RegisterRoutes(r)
	healthcheck_handler.RegisterRoutes(r)

	r.Run(":8000")
}
