package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	roleDto "github.com/yigittopm/wl-auth/internal/authentication/dtos/role"
	"github.com/yigittopm/wl-auth/internal/authentication/usecase"
	"github.com/yigittopm/wl-auth/pkg/utils/response"
)

type RoleHandler interface {
	GetAllRoles(c *fiber.Ctx) error
	GetRoleById(c *fiber.Ctx) error
	CreateRole(c *fiber.Ctx) error
	DeleteRoleById(c *fiber.Ctx) error
	UpdateRoleById(c *fiber.Ctx) error
}

type roleHandler struct {
	uc usecase.RoleUsecase
}

func NewRoleHandler(uc usecase.RoleUsecase) RoleHandler {
	return &roleHandler{uc: uc}
}

func (h *roleHandler) GetAllRoles(c *fiber.Ctx) error {
	return nil
}

func (h *roleHandler) GetRoleById(c *fiber.Ctx) error {
	return nil
}

func (h *roleHandler) CreateRole(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
		payload     roleDto.CreateRoleRequest
	)
	defer cancel()

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Parse error: %v", err.Error()))
	}

	if err := payload.Validate(); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validate error: %v", err.Error()))
	}

	role, status, err := h.uc.CreateRole(ctx, payload)
	if err != nil {
		if err != nil {
			return response.ErrorResponse(c, status, fmt.Sprintf("Create Role error: %v", err.Error()))
		}
	}

	return response.SuccessResponse(c, status, role)
}

func (h *roleHandler) DeleteRoleById(c *fiber.Ctx) error {
	return nil
}

func (h *roleHandler) UpdateRoleById(c *fiber.Ctx) error {
	return nil
}
