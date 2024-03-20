package repository

import (
	"context"
	"errors"

	"github.com/yigittopm/wl-auth/internal/users/dtos"
	"github.com/yigittopm/wl-auth/internal/users/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	Register(context.Context, entities.User) (dtos.RegisterResponse, error)
	Login(context.Context, dtos.LoginRequest) (uint, error)
	Profile(context.Context, dtos.ProfileRequest) (entities.User, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) Register(ctx context.Context, user entities.User) (dtos.RegisterResponse, error) {
	user.Password, _ = hashPassword(user.Password)
	result := repo.db.Create(&user)

	return dtos.RegisterResponse{
		ID: user.ID,
	}, result.Error
}

func (repo *repository) Login(ctx context.Context, payload dtos.LoginRequest) (uint, error) {
	user := entities.User{}
	result := repo.db.Find(&user, "username = ?", payload.Username)
	if ok := comparePassword(user.Password, payload.Password); !ok {
		return 0, errors.New("username or password is incorrect")
	}

	return user.ID, result.Error
}

func (repo *repository) Profile(ctx context.Context, payload dtos.ProfileRequest) (entities.User, error) {
	var user entities.User
	result := repo.db.Find(&user, "id = ?", payload.ID)
	return user, result.Error
}

func (repo *repository) IsExist(ctx context.Context, email string) (bool, error) {
	var user entities.User
	result := repo.db.First(&user, "email = ?", email)
	return result.RowsAffected > 0, result.Error
}

func hashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

func comparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
