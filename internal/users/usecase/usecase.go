package usecase

import (
	"context"

	"github.com/yigittopm/test/internal/users/dtos"
	"github.com/yigittopm/test/internal/users/entities"
	"github.com/yigittopm/test/internal/users/repository"
)

type Usecase interface {
	GetAll(ctx context.Context) ([]dtos.GetAllUsersResponse, error)
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
	return uc.repo.SaveNewUser(ctx, entities.New(payload))
}

func (uc *usecase) GetAll(ctx context.Context) ([]dtos.GetAllUsersResponse, error) {
	return uc.repo.GetAllUsers(ctx)
}
