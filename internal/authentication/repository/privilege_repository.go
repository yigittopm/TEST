package repository

import (
	"context"
	dtos "github.com/yigittopm/wl-auth/internal/authentication/dtos/privilege"
	"github.com/yigittopm/wl-auth/internal/authentication/entities/privilege"
	"github.com/yigittopm/wl-auth/internal/authentication/entities/role"
	"gorm.io/gorm"
)

type PrivilegeRepository interface {
	GetAllPrivileges(context.Context) (role.Role, error)
	GetPrivilegeById(context.Context, dtos.GetPrivilegeByIdRequest) (role.Role, error)
	CreatePrivilege(context.Context, privilege.Privilege) (dtos.CreatePrivilegeResponse, error)
	UpdatePrivilege(context.Context, dtos.UpdatePrivilegeRequest) (role.Role, error)
	DeletePrivilege(context.Context, dtos.UpdatePrivilegeRequest) error
}

type privilegeRepository struct {
	db *gorm.DB
}

func NewPrivilegeRepository(db *gorm.DB) PrivilegeRepository {
	return &privilegeRepository{db: db}
}

func (repo *privilegeRepository) GetAllPrivileges(ctx context.Context) (role.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *privilegeRepository) GetPrivilegeById(ctx context.Context, request dtos.GetPrivilegeByIdRequest) (role.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *privilegeRepository) CreatePrivilege(ctx context.Context, privilege privilege.Privilege) (dtos.CreatePrivilegeResponse, error) {
	result := repo.db.Create(&privilege)
	return dtos.CreatePrivilegeResponse{
		ID:          privilege.ID,
		Description: privilege.Description,
		Key:         privilege.Key,
	}, result.Error
}

func (repo *privilegeRepository) UpdatePrivilege(ctx context.Context, request dtos.UpdatePrivilegeRequest) (role.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *privilegeRepository) DeletePrivilege(ctx context.Context, request dtos.UpdatePrivilegeRequest) error {
	//TODO implement me
	panic("implement me")
}
