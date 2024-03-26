package usecase

import (
	"context"
	dtos "github.com/yigittopm/wl-auth/internal/authentication/dtos/privilege"
	entities "github.com/yigittopm/wl-auth/internal/authentication/entities/privilege"
	"github.com/yigittopm/wl-auth/internal/authentication/repository"
	"net/http"
)

type PrivilegeUsecase interface {
	GetAllPrivileges(context.Context) ([]entities.Privilege, int, error)
	GetPrivilegeByID(context.Context, dtos.GetPrivilegeByIdRequest) (dtos.GetPrivilegeByIdResponse, int, error)
	CreatePrivilege(context.Context, dtos.CreatePrivilegeRequest) (dtos.CreatePrivilegeResponse, int, error)
	UpdatePrivilege(context.Context, dtos.UpdatePrivilegeRequest) (dtos.UpdatePrivilegeResponse, int, error)
	DeletePrivilege(context.Context, dtos.DeletePrivilegeRequest) (int, error)
}

type privilegeUsecase struct {
	repo repository.PrivilegeRepository
}

func NewPrivilegeUsecase(repo repository.PrivilegeRepository) PrivilegeUsecase {
	return &privilegeUsecase{
		repo: repo,
	}
}

func (u *privilegeUsecase) GetAllPrivileges(ctx context.Context) ([]entities.Privilege, int, error) {
	//TODO implement me
	panic("implement me")
}

func (u *privilegeUsecase) GetPrivilegeByID(ctx context.Context, request dtos.GetPrivilegeByIdRequest) (dtos.GetPrivilegeByIdResponse, int, error) {
	//TODO implement me
	panic("implement me")
}

func (u *privilegeUsecase) CreatePrivilege(ctx context.Context, request dtos.CreatePrivilegeRequest) (dtos.CreatePrivilegeResponse, int, error) {
	privilege := entities.New(request)
	response, err := u.repo.CreatePrivilege(ctx, privilege)
	if err != nil {
		return dtos.CreatePrivilegeResponse{}, http.StatusBadRequest, err
	}

	return response, http.StatusOK, nil
}

func (u *privilegeUsecase) UpdatePrivilege(ctx context.Context, request dtos.UpdatePrivilegeRequest) (dtos.UpdatePrivilegeResponse, int, error) {
	//TODO implement me
	panic("implement me")
}

func (u *privilegeUsecase) DeletePrivilege(ctx context.Context, request dtos.DeletePrivilegeRequest) (int, error) {
	//TODO implement me
	panic("implement me")
}
