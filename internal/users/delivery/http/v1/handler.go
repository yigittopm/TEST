package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yigittopm/test/internal/users/dtos"
	"github.com/yigittopm/test/internal/users/usecase"
	"github.com/yigittopm/test/pkg/utils/response"
)

type Handler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type handler struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) Handler {
	return &handler{uc: uc}
}

// Register godoc
// @Tags User
// @Description register user.
// @Summary register user
// @Accept json
// @Produce json
// @Param request body dtos.RegisterRequest true "Request Body"
// @Success 200 {string} string "User ID"
// @Router /v1/auth/register [post]
func (h *handler) Register(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
		payload     dtos.RegisterRequest
	)
	defer cancel()

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Parse error: %v", err.Error()))
	}

	if err := payload.Validate(); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validate error: %v", err.Error()))
	}

	userID, err := h.uc.Register(ctx, payload)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Register error: %v", err.Error()))
	}

	return response.SuccessResponse(c, http.StatusOK, userID)
}

// Login godoc
// @Tags User
// @Description create product.
// @Summary create product
// @Accept json
// @Produce json
// @Param request body model.ProductCreateOrUpdateModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/auth/login [post]
func (h *handler) Login(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
		payload     dtos.LoginRequest
	)
	defer cancel()

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Parse error: %v", err.Error()))
	}

	if err := payload.Validate(); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validation error: %v", err.Error()))
	}

	user, err := h.uc.Login(ctx, payload)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Login error: %v", err.Error()))
	}

	return response.SuccessResponse(c, http.StatusOK, user)
}
