package repository

import (
	"context"
	"database/sql"

	"github.com/yigittopm/test/internal/users/dtos"
	"github.com/yigittopm/test/internal/users/entities"
)

type Repository interface {
	GetAllUsers(context.Context) ([]dtos.GetAllUsersResponse, error)
	SaveNewUser(context.Context, entities.User) (string, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) SaveNewUser(ctx context.Context, user entities.User) (string, error) {
	stmt, err := repo.db.Prepare(`
	INSERT INTO 
	users  (username, email, password, user_type, is_active, created_by, updated_by) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`)

	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var id string
	_ = stmt.QueryRow(user.Username, user.Email, user.Password, user.UserType, user.IsActive, user.CreatedBy, user.UpdatedBy).Scan(&id)

	return id, nil
}

func (repo *repository) GetAllUsers(ctx context.Context) ([]dtos.GetAllUsersResponse, error) {
	rows, err := repo.db.Query("SELECT username, email, user_type FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []dtos.GetAllUsersResponse
	for rows.Next() {
		var user dtos.GetAllUsersResponse
		if err := rows.Scan(
			&user.Username,
			&user.Email,
			&user.UserType,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, err
}
