package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/yigittopm/test/database"
	"github.com/yigittopm/test/internal/users"
)

func NewApp() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	db, err := database.Start()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} ${latency} ${method} ${path}\n",
	}))
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture psql!"))
	})

	v1 := app.Group("/api/v1")
	users.Setup(v1, db)

	log.Fatal(app.Listen(":8080"))
}
