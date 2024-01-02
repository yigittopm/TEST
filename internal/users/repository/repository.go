package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yigittopm/test/internal/users/entities"
)

type Repository interface {
	SaveNewUser(context.Context, entities.User) (string, error)
	UpdateUserById(context.Context, entities.UpdateUser)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) SaveNewUser(ctx context.Context, user entities.User) (string, error) {
	rows, err := repo.db.Query("INSERT INTO users (username, email, password) values ($1, $2, $3)",
		user.Username, user.Email, user.Password)
	if err != nil {
		return "", err
	}
	fmt.Println(rows)
	return "", nil
}

func (repo *repository) UpdateUserById(ctx context.Context, user entities.UpdateUser) {

}
