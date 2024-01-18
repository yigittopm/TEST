package users

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	usersHandler "github.com/yigittopm/test/internal/users/delivery/http/v1"
	usersRepository "github.com/yigittopm/test/internal/users/repository"
	usersUsecase "github.com/yigittopm/test/internal/users/usecase"
)

func Setup(router fiber.Router, db *sql.DB) {
	UserMigrate(db)

	repo := usersRepository.New(db)
	service := usersUsecase.New(repo)
	handler := usersHandler.New(service)

	router.Get("/users", handler.GetAllUsers)
	router.Post("/users", handler.CreateUser)
	router.Patch("/users", handler.UpdateUserByID)
	router.Delete("/users", handler.DeleteUserByID)
}

func UserMigrate(DB *sql.DB) {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id serial primary key,
			username varchar(255) not null,
			email varchar(255) not null,
			password varchar(255) not null,
			user_type varchar(255),
			is_active boolean,
			created_at TIMESTAMP not null default NOW(),
			created_by varchar(255) not null,
			updated_at TIMESTAMP not null default NOW(),
			updated_by varchar(255) not null
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success created users table")
}
