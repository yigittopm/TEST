package repository

import (
	"context"
	dtos "github.com/yigittopm/wl-auth/internal/authentication/dtos/role"
	"github.com/yigittopm/wl-auth/internal/authentication/entities/role"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetAllRoles(context.Context) (role.Role, error)
	GetRoleById(context.Context, dtos.GetRoleByIdRequest) (role.Role, error)
	CreateRole(context.Context, role.Role) (dtos.CreateRoleResponse, error)
	UpdateRole(context.Context, dtos.UpdateRoleRequest) (role.Role, error)
	DeleteRole(context.Context, dtos.DeleteRoleRequest) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (repo *roleRepository) GetAllRoles(ctx context.Context) (role.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *roleRepository) GetRoleById(ctx context.Context, payload dtos.GetRoleByIdRequest) (role.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *roleRepository) CreateRole(ctx context.Context, role role.Role) (dtos.CreateRoleResponse, error) {
	result := repo.db.Create(&role)
	return dtos.CreateRoleResponse{
		ID:         role.ID,
		Name:       role.Name,
		Privileges: role.Privileges,
	}, result.Error
}

func (repo *roleRepository) UpdateRole(ctx context.Context, payload dtos.UpdateRoleRequest) (role.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *roleRepository) DeleteRole(ctx context.Context, payload dtos.DeleteRoleRequest) error {
	//TODO implement me
	panic("implement me")
}
