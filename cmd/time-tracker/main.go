package main

import (
	"fmt"
	"os"
	"time-tracker-app/internal/config"
	db "time-tracker-app/internal/db/sqlc"
	"time-tracker-app/internal/repository"
	"time-tracker-app/internal/service"
	"time-tracker-app/internal/transport/http/handler"
	"time-tracker-app/internal/transport/http/router"
	"time-tracker-app/pkg/database"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	newConfig, err := config.NewConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create config: %v\n", err)
		os.Exit(1)
	}

	err = database.RunMigrations(newConfig.DBSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to execute migrations: %v\n", err)
		os.Exit(1)
	}

	newPool, err := database.NewPGXPool(newConfig.DBSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer newPool.Close()

	queries := db.New(newPool)

	newRepository := repository.NewRepository(queries)
	newService := service.NewService(newRepository)
	newHandler := handler.NewHandler(newService)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	newRouter := router.NewRouter(newHandler)
	newRouter.SetRoutes(app)

	app.Listen(":3000")
}
