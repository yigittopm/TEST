package users

import (
	"github.com/gofiber/fiber/v2"
	usersHandler "github.com/yigittopm/test/internal/users/delivery/http/v1"
	userEntities "github.com/yigittopm/test/internal/users/entities"
	usersRepository "github.com/yigittopm/test/internal/users/repository"
	usersUsecase "github.com/yigittopm/test/internal/users/usecase"
	"gorm.io/gorm"
)

func Setup(router fiber.Router, db *gorm.DB) {
	// Migration
	db.AutoMigrate(&userEntities.User{})

	// Dependency Injection
	repo := usersRepository.New(db)
	usecase := usersUsecase.New(repo)
	handler := usersHandler.New(usecase)

	// Routes
	route := router.Group("/auth")

	route.Post("/register", handler.Register)
	route.Post("/login", handler.Login)
}
