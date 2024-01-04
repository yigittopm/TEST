package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yigittopm/test/internal/users/dtos"
	"github.com/yigittopm/test/internal/users/entities"
	"github.com/yigittopm/test/internal/users/repository"
)

type Usecase interface {
	Create(ctx context.Context, payload dtos.CreateUserRequest) (userID string, err error)
}

type usecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Create(ctx context.Context, payload dtos.CreateUserRequest) (userID string, err error) {
	userID, err = uc.repo.SaveNewUser(ctx, entities.New(payload))
	if err != nil {
		return userID, err
	}
	fmt.Println(userID, http.StatusOK, err)
	return userID, err
}
