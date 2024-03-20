package app

import (
	"log"

	"github.com/yigittopm/wl-auth/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
	"github.com/yigittopm/wl-auth/internal/users"
)

func NewApp() {
	// Initial Database
	db, err := database.Start()
	if err != nil {
		log.Fatal(err)
	}

	// New Fiber App
	app := fiber.New()

	// Cache Middleware
	app.Use(cache.New())

	// Cors Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Swagger implementation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Logger
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} ${latency} ${method} ${path}\n",
	}))

	// Metrics
	app.Get("/metrics", monitor.New(monitor.Config{Title: "WeLedger Auth Monitor"}))

	// Handler Version
	version := app.Group("/api/v1") // V1

	// Users
	users.Setup(version, db)

	// Listening http port on :8080
	log.Fatal(app.Listen(":8080"))
}
