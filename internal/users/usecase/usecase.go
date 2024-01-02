package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yigittopm/test/config"
	"github.com/yigittopm/test/internal/users/dtos"
	"github.com/yigittopm/test/internal/users/entities"
	"github.com/yigittopm/test/internal/users/repository"
)

type Usecase interface {
	Create(ctx context.Context, payload dtos.CreateUserRequest) (userID string, httpCode int, err error)
	Update(ctx context.Context, payload dtos.UpdateUserRequest) error
}

type usecase struct {
	repo repository.Repository
	cfg  config.Config
}

func New(repo repository.Repository, cfg config.Config) Usecase {
	return &usecase{
		repo: repo,
		cfg:  cfg,
	}
}

func (uc *usecase) Create(ctx context.Context, payload dtos.CreateUserRequest) (userID string, httpCode int, err error) {
	userID, err = uc.repo.SaveNewUser(ctx, entities.New(payload, uc.cfg))
	if err != nil {
		return userID, http.StatusInternalServerError, err
	}
	fmt.Println(userID, http.StatusOK, err)
	return userID, http.StatusOK, err
}

func (uc *usecase) Update(ctx context.Context, payload dtos.UpdateUserRequest) error {
	return nil
}
