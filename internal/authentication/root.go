package authentication

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/yigittopm/wl-auth/internal/authentication/delivery/http/v1"
	privilegesEntity "github.com/yigittopm/wl-auth/internal/authentication/entities/privilege"
	roleEntity "github.com/yigittopm/wl-auth/internal/authentication/entities/role"
	"github.com/yigittopm/wl-auth/internal/authentication/repository"
	"github.com/yigittopm/wl-auth/internal/authentication/usecase"
	"github.com/yigittopm/wl-auth/pkg/middleware"
	"gorm.io/gorm"
)

func Setup(router fiber.Router, db *gorm.DB) {

	// Migration
	db.AutoMigrate(&roleEntity.Role{})
	db.AutoMigrate(&privilegesEntity.Privilege{})

	// Dependency Injection for Role
	roleRepo := repository.NewRoleRepository(db)
	roleUsecase := usecase.NewRoleUsecase(roleRepo)
	roleHandler := handler.NewRoleHandler(roleUsecase)

	// Dependency Injection for Privileges
	privilegeRepo := repository.NewPrivilegeRepository(db)
	privilegeUsecase := usecase.NewPrivilegeUsecase(privilegeRepo)
	privilegeHandler := handler.NewPrivilegeHandler(privilegeUsecase)

	// Routes
	route := router.Group("/auth")

	// Roles Routes
	route.Post("/role", middleware.AuthRequired(), roleHandler.CreateRole)

	// Privilege Routes
	route.Post("/privilege", middleware.AuthRequired(), privilegeHandler.CreatePrivilege)

}
