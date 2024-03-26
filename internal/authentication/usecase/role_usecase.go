package usecase

import (
	"context"
	dtos "github.com/yigittopm/wl-auth/internal/authentication/dtos/role"
	entities "github.com/yigittopm/wl-auth/internal/authentication/entities/role"
	"github.com/yigittopm/wl-auth/internal/authentication/repository"
	"net/http"
)

type RoleUsecase interface {
	GetAllRoles(context.Context) ([]entities.Role, int, error)
	GetRoleByID(context.Context, dtos.GetRoleByIdRequest) (entities.Role, int, error)
	CreateRole(context.Context, dtos.CreateRoleRequest) (dtos.CreateRoleResponse, int, error)
	UpdateRole(context.Context, dtos.UpdateRoleRequest) (entities.Role, int, error)
	DeleteRole(context.Context, dtos.DeleteRoleRequest) (int, error)
}

type roleUsecase struct {
	repo repository.RoleRepository
}

func NewRoleUsecase(repo repository.RoleRepository) RoleUsecase {
	return &roleUsecase{
		repo: repo,
	}
}

func (uc *roleUsecase) GetAllRoles(ctx context.Context) ([]entities.Role, int, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *roleUsecase) GetRoleByID(ctx context.Context, request dtos.GetRoleByIdRequest) (entities.Role, int, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *roleUsecase) CreateRole(ctx context.Context, request dtos.CreateRoleRequest) (dtos.CreateRoleResponse, int, error) {
	role := entities.New(request)
	response, err := uc.repo.CreateRole(ctx, role)
	if err != nil {
		return dtos.CreateRoleResponse{}, http.StatusBadRequest, err
	}

	return response, http.StatusOK, nil
}

func (uc *roleUsecase) UpdateRole(ctx context.Context, request dtos.UpdateRoleRequest) (entities.Role, int, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *roleUsecase) DeleteRole(ctx context.Context, request dtos.DeleteRoleRequest) (int, error) {
	//TODO implement me
	panic("implement me")
}
