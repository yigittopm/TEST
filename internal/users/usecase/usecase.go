package usecase

import (
	"context"

	"github.com/yigittopm/test/internal/users/dtos"
	"github.com/yigittopm/test/internal/users/entities"
	"github.com/yigittopm/test/internal/users/repository"
)

type Usecase interface {
	Register(ctx context.Context, payload dtos.RegisterRequest) (dtos.RegisterResponse, error)
	Login(ctx context.Context, payload dtos.LoginRequest) (dtos.LoginResponse, error)
}

type usecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Register(ctx context.Context, payload dtos.RegisterRequest) (dtos.RegisterResponse, error) {
	user := entities.New(payload)

	return uc.repo.Register(ctx, user)
}

func (uc *usecase) Login(ctx context.Context, payload dtos.LoginRequest) (dtos.LoginResponse, error) {
	return uc.repo.Login(ctx, payload)
}
