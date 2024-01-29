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
// @Tags User
// @Description create product.
// @Summary create product
// @Accept json
// @Produce json
// @Param request body model.ProductCreateOrUpdateModel true "Request Body"
// @Success 200 {object} dtos.CreateUserResponse
// @Security JWT
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

// GetAllUsers godoc
// @Tags User
// @Description create product.
// @Summary create product
// @Accept json
// @Produce json
// @Param request body model.ProductCreateOrUpdateModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/users [get]
func (h *handler) GetAllUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), time.Duration(10*time.Second))
	defer cancel()

	users, err := h.uc.GetAll(ctx)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Error fetch users.")
	}

	return response.SuccessResponse(c, http.StatusOK, users)
}

// UpdateUserByID godoc
// @Tags User
// @Description create product.
// @Summary create product
// @Accept json
// @Produce json
// @Param request body model.ProductCreateOrUpdateModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/users [put]
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

// DeleteUserByID godoc
// @Tags User
// @Description create product.
// @Summary create product
// @Accept json
// @Produce json
// @Param request body model.ProductCreateOrUpdateModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/users [delete]
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
