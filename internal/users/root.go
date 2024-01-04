package users

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	usersHandler "github.com/yigittopm/test/internal/users/delivery/http/v1"
	usersRepository "github.com/yigittopm/test/internal/users/repository"
	usersUsecase "github.com/yigittopm/test/internal/users/usecase"
)

func Setup(router fiber.Router, db *sql.DB) {
	repo := usersRepository.New(db)
	service := usersUsecase.New(repo)
	handler := usersHandler.New(service)

	router.Post("/users", handler.CreateUser())
}
