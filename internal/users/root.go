package users

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	usersHandler "github.com/yigittopm/test/internal/users/delivery/http/v1"
	usersRepository "github.com/yigittopm/test/internal/users/repository"
	usersUsecase "github.com/yigittopm/test/internal/users/usecase"
)

func Setup(router fiber.Router, db *sql.DB) {
	// Dependency Injection
	repo := usersRepository.New(db)
	usecase := usersUsecase.New(repo)
	handler := usersHandler.New(usecase)

	// Routes
	router.Get("/users", handler.GetAllUsers)
	router.Post("/users", handler.CreateUser)
	router.Patch("/users", handler.UpdateUserByID)
	router.Delete("/users", handler.DeleteUserByID)
}
