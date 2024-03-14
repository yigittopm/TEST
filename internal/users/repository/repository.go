package repository

import (
	"context"

	"github.com/yigittopm/wl-auth/internal/users/dtos"
	"github.com/yigittopm/wl-auth/internal/users/entities"
	"gorm.io/gorm"
)

type Repository interface {
	Register(context.Context, entities.User) (dtos.RegisterResponse, error)
	Login(context.Context, dtos.LoginRequest) (dtos.LoginResponse, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) Register(ctx context.Context, user entities.User) (dtos.RegisterResponse, error) {
	result := repo.db.Create(&user)

	return dtos.RegisterResponse{
		ID: user.ID,
	}, result.Error
}

func (repo *repository) Login(ctx context.Context, payload dtos.LoginRequest) (dtos.LoginResponse, error) {
	user := entities.User{}
	result := repo.db.First(&user, "username = ? AND password = ?", payload.Username, payload.Password)

	return dtos.LoginResponse{
		ID:       user.ID,
		IsActive: user.IsActive,
	}, result.Error
}
