package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yigittopm/test/internal/users/dtos"
	"github.com/yigittopm/test/internal/users/usecase"
	"github.com/yigittopm/test/pkg/utils/response"
)

type Handler interface {
	GetAllUsers(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUserByID(c *fiber.Ctx) error
	DeleteUserByID(c *fiber.Ctx) error
}

type handler struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) Handler {
	return &handler{uc: uc}
}

// CreateUser godoc
// @Summary returns the HTTP headers
// @Description use this to inspect the headers set by the portal and received by the service
// @Produce json
// @Router /v1/users [post]
func (h *handler) CreateUser(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
		payload     dtos.CreateUserRequest
	)
	defer cancel()

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Parse error.")
	}

	if err := payload.Validate(); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Validate error.")
	}

	userID, err := h.uc.Create(ctx, payload)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "User create error.")
	}

	return response.SuccessResponse(c, http.StatusOK, userID)
}

func (h *handler) GetAllUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), time.Duration(10*time.Second))
	defer cancel()

	users, err := h.uc.GetAll(ctx)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Error fetch users.")
	}

	return response.SuccessResponse(c, http.StatusOK, users)
}

func (h *handler) UpdateUserByID(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
		payload     dtos.UpdateUserRequest
	)
	defer cancel()

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Parse error.")
	}

	if err := payload.Validate(); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Validate error.")
	}

	userID, err := h.uc.Update(ctx, payload)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "User update error.")
	}

	return response.SuccessResponse(c, http.StatusOK, userID)
}

func (h *handler) DeleteUserByID(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
		payload     dtos.DeleteUserByIdRequest
	)
	defer cancel()

	payload.ID = c.Query("userID")
	if payload.ID == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, "User id not be null.")
	}

	userId, err := h.uc.Delete(ctx, payload)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Error delete user.")
	}

	return response.SuccessResponse(c, http.StatusOK, userId)
}
