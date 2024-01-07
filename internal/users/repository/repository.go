package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/yigittopm/test/internal/users/dtos"
	"github.com/yigittopm/test/internal/users/entities"
)

type Repository interface {
	GetAllUsers(context.Context) ([]dtos.GetAllUsersResponse, error)
	SaveNewUser(context.Context, entities.User) (string, error)
	UpdateUserById(context.Context, dtos.UpdateUserRequest) (string, error)
	DeleteUserById(context.Context, string) (string, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) SaveNewUser(ctx context.Context, user entities.User) (string, error) {
	if ok := repo.userIsExists(ctx, user.Email); !ok {
		return "", errors.New("User already exists")
	}

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
	_ = stmt.QueryRowContext(ctx, user.Username, user.Email, user.Password, user.UserType, user.IsActive, user.CreatedBy, user.UpdatedBy).Scan(&id)

	return id, nil
}

func (repo *repository) GetAllUsers(ctx context.Context) ([]dtos.GetAllUsersResponse, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id,username, email, user_type FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []dtos.GetAllUsersResponse
	for rows.Next() {
		var user dtos.GetAllUsersResponse
		if err := rows.Scan(
			&user.ID,
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

func (repo *repository) UpdateUserById(ctx context.Context, payload dtos.UpdateUserRequest) (string, error) {
	stmt, err := repo.db.Prepare(`
	UPDATE users
	SET
	    username=$1,
	    email=$2,
	    password=$3
	WHERE 
	    id=$4
	`)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, payload.Username, payload.Email, payload.Password, payload.ID)
	if err != nil {
		return "", err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return "", errors.New("User id not found.")
	}

	return payload.ID, nil
}

func (repo *repository) DeleteUserById(ctx context.Context, userID string) (string, error) {
	stmt, err := repo.db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID)
	if err != nil {
		return "", err
	}

	return userID, err
}

func (repo *repository) userIsExists(ctx context.Context, email string) bool {
	stmt, err := repo.db.Prepare("SELECT id FROM users WHERE email=$1")
	if err != nil {
		return true
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, email)

	var userID int
	_ = row.Scan(&userID)

	if userID == 0 {
		return true
	}

	return false
}
