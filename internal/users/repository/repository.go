package repository

import (
	"context"

	"github.com/yigittopm/wl-auth/internal/users/dtos"
	"github.com/yigittopm/wl-auth/internal/users/entities"
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
	result := repo.db.Create(&user)

	return dtos.RegisterResponse{
		ID: user.ID,
	}, result.Error
}

func (repo *repository) Login(ctx context.Context, payload dtos.LoginRequest) (uint, error) {
	user := entities.User{}
	result := repo.db.First(&user, "username = ? AND password = ?", payload.Username, payload.Password)
	return user.ID, result.Error
}

func (repo *repository) Profile(ctx context.Context, payload dtos.ProfileRequest) (entities.User, error) {
	var user entities.User
	result := repo.db.Find(&user, "id = ?", payload.ID)
	return user, result.Error
}
