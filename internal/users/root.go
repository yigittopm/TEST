package users

import (
	"github.com/gofiber/fiber/v2"
	usersHandler "github.com/yigittopm/wl-auth/internal/users/delivery/http/v1"
	userEntities "github.com/yigittopm/wl-auth/internal/users/entities"
	usersRepository "github.com/yigittopm/wl-auth/internal/users/repository"
	usersUsecase "github.com/yigittopm/wl-auth/internal/users/usecase"
	"github.com/yigittopm/wl-auth/pkg/constant"
	"github.com/yigittopm/wl-auth/pkg/middleware"
	"gorm.io/gorm"
)

func Setup(router fiber.Router, db *gorm.DB) {

	// Migration
	db.AutoMigrate(&userEntities.User{})
	db.AutoMigrate(&userEntities.UserDetail{})

	// Dependency Injection
	repo := usersRepository.New(db)
	usecase := usersUsecase.New(repo)
	handler := usersHandler.New(usecase)

	// Routes
	route := router.Group("/auth")

	route.Post("/register", handler.Register)
	route.Post("/login", handler.Login)
	route.Get("/logout", middleware.AuthRequired(constant.DEFAULT_ROLE), handler.Logout)
	route.Get("/profile", middleware.AuthRequired(constant.ROOT_ROLE), handler.Profile)
}
