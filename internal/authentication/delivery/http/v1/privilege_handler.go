package v1

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	privilegeDto "github.com/yigittopm/wl-auth/internal/authentication/dtos/privilege"
	"github.com/yigittopm/wl-auth/internal/authentication/usecase"
	"github.com/yigittopm/wl-auth/pkg/utils/response"
	"net/http"
	"time"
)

type PrivilegeHandler interface {
	GetAllPrivileges(c *fiber.Ctx) error
	GetPrivilegeById(c *fiber.Ctx) error
	CreatePrivilege(c *fiber.Ctx) error
	DeletePrivilegeById(c *fiber.Ctx) error
	UpdatePrivilegeById(c *fiber.Ctx) error
}

type privilegeHandler struct {
	uc usecase.PrivilegeUsecase
}

func NewPrivilegeHandler(uc usecase.PrivilegeUsecase) PrivilegeHandler {
	return &privilegeHandler{uc: uc}
}

func (h *privilegeHandler) GetAllPrivileges(c *fiber.Ctx) error {
	return nil
}

func (h *privilegeHandler) GetPrivilegeById(c *fiber.Ctx) error {
	return nil
}

func (h *privilegeHandler) CreatePrivilege(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
		payload     privilegeDto.CreatePrivilegeRequest
	)
	defer cancel()

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Parse error: %v", err.Error()))
	}

	if err := payload.Validate(); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validate error: %v", err.Error()))
	}

	privilege, status, err := h.uc.CreatePrivilege(ctx, payload)
	if err != nil {
		if err != nil {
			return response.ErrorResponse(c, status, fmt.Sprintf("Create Role error: %v", err.Error()))
		}
	}

	return response.SuccessResponse(c, status, privilege)
}

func (h *privilegeHandler) DeletePrivilegeById(c *fiber.Ctx) error {
	return nil
}

func (h *privilegeHandler) UpdatePrivilegeById(c *fiber.Ctx) error {
	return nil
}
