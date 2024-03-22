package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/yigittopm/wl-auth/internal/users/dtos"
	"github.com/yigittopm/wl-auth/internal/users/entities"
	"github.com/yigittopm/wl-auth/internal/users/repository"
	"github.com/yigittopm/wl-auth/pkg/jwt"
)

type Usecase interface {
	Register(ctx context.Context, payload dtos.RegisterRequest) (dtos.RegisterResponse, int, error)
	Login(ctx context.Context, payload dtos.LoginRequest) (dtos.LoginResponse, int, error)
	Profile(ctx context.Context, payload dtos.ProfileRequest) (entities.User, int, error)
}

type usecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Register(ctx context.Context, payload dtos.RegisterRequest) (dtos.RegisterResponse, int, error) {
	user := entities.New(payload)
	response, err := uc.repo.Register(ctx, user)
	if err != nil {
		return dtos.RegisterResponse{}, http.StatusBadRequest, err
	}

	return response, http.StatusOK, nil
}

func (uc *usecase) Login(ctx context.Context, payload dtos.LoginRequest) (dtos.LoginResponse, int, error) {
	userId, err := uc.repo.Login(ctx, payload)
	if err != nil {
		return dtos.LoginResponse{}, http.StatusBadRequest, err
	}

	accessToken, err := jwt.Sign(userId, time.Hour*8)
	if err != nil {
		return dtos.LoginResponse{}, http.StatusUnauthorized, err
	}

	return dtos.LoginResponse{
		ID:          userId,
		AccessToken: accessToken,
	}, http.StatusOK, nil
}

func (uc *usecase) Profile(ctx context.Context, payload dtos.ProfileRequest) (entities.User, int, error) {
	user, err := uc.repo.Profile(ctx, payload)
	if err != nil {
		return entities.User{}, http.StatusBadRequest, err
	}

	return user, http.StatusOK, nil
}
