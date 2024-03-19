package usecase

import (
	"context"
	"time"

	"github.com/yigittopm/wl-auth/internal/users/dtos"
	"github.com/yigittopm/wl-auth/internal/users/entities"
	"github.com/yigittopm/wl-auth/internal/users/repository"
	"github.com/yigittopm/wl-auth/pkg/jwt"
)

type Usecase interface {
	Register(ctx context.Context, payload dtos.RegisterRequest) (dtos.RegisterResponse, error)
	Login(ctx context.Context, payload dtos.LoginRequest) (dtos.LoginResponse, error)
	Profile(ctx context.Context, payload dtos.ProfileRequest) (entities.User, error)
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
	//TODO: Check IsExist user

	userId, err := uc.repo.Login(ctx, payload)
	if err != nil {
		return dtos.LoginResponse{}, err
	}

	accessToken, err := jwt.Sign(userId, time.Hour*8)
	if err != nil {
		return dtos.LoginResponse{}, err
	}

	return dtos.LoginResponse{
		ID:          userId,
		AccessToken: accessToken,
	}, err
}

func (uc *usecase) Profile(ctx context.Context, payload dtos.ProfileRequest) (entities.User, error) {
	user, err := uc.repo.Profile(ctx, payload)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
