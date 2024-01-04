package repository

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/yigittopm/test/internal/users/entities"
)

type Repository interface {
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

	rows, err := stmt.Exec(user.Username, user.Email, user.Password, user.UserType, user.IsActive, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		return "", err
	}

	id, err := rows.RowsAffected()
	if err != nil {
		return "", err
	}

	res := strconv.FormatInt(id, 10)
	return res, nil
}
