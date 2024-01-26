package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/yigittopm/test/database"

	"github.com/yigittopm/test/internal/users"
)

func NewApp() {
	// Initial Database
	db, err := database.Start()
	if err != nil {
		log.Fatal(err)
	}

	// New Fiber App
	app := fiber.New()

	// Cors Middleware
	app.Use(cors.New())

	// Swagger implementation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Logger
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} ${latency} ${method} ${path}\n",
	}))

	// Handler Version
	v1 := app.Group("/api/v1")

	// Users
	users.Setup(v1, db)

	// Listening http port on :8080
	log.Fatal(app.Listen(":8080"))
}
