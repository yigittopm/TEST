package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yigittopm/wl-auth/internal/users/dtos"
	"github.com/yigittopm/wl-auth/internal/users/usecase"
	"github.com/yigittopm/wl-auth/pkg/utils/response"
)

type Handler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	Profile(c *fiber.Ctx) error
}

type handler struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) Handler {
	return &handler{uc: uc}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the given details
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body dtos.RegisterRequest true "User details for registration"
// @Success 200 {object} dtos.RegisterResponse "Successfully registered user"
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

	userID, status, err := h.uc.Register(ctx, payload)
	if err != nil {
		return response.ErrorResponse(c, status, fmt.Sprintf("Register error: %v", err.Error()))
	}

	return response.SuccessResponse(c, status, userID)
}

// Login godoc
// @Summary Log in a user
// @Description Log in a user with the given credentials
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body dtos.LoginRequest true "User credentials for login"
// @Success 200 {object} dtos.LoginResponse "Successfully logged in user"
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

	user, status, err := h.uc.Login(ctx, payload)
	if err != nil {
		return response.ErrorResponse(c, status, fmt.Sprintf("Login error: %v", err.Error()))
	}

	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   "Bearer " + user.AccessToken,
		Expires: time.Now().Add(time.Hour * 24),
	})

	c.Cookie(&fiber.Cookie{
		Name:    "role",
		Value:   user.Roles[0].Name,
		Expires: time.Now().Add(time.Hour * 24),
	})

	return response.SuccessResponse(c, status, user)
}

// Logout godoc
// @Summary Log out a user
// @Description Log out the authenticated user
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {string} string "Successfully logged out"
// @Failure 400 {object} ErrorResponse "Failed to log out user"
// @Router /v1/auth/logout [post]
func (h *handler) Logout(c *fiber.Ctx) error {
	c.Locals("userId", nil)
	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	return response.SuccessResponse(c, http.StatusOK, "Successfully logged out")
}

// Profile godoc
// @Summary Get user profile
// @Description Get the profile of the authenticated user
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} dtos.ProfileResponse "Successfully retrieved user profile"
// @Failure 400 {object} ErrorResponse "Failed to retrieve user profile"
// @Router /v1/auth/profile [get]
func (h *handler) Profile(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
	)
	defer cancel()

	userId := c.Locals("userId").(uint)

	user, status, err := h.uc.Profile(ctx, dtos.ProfileRequest{ID: userId})
	if err != nil {
		return response.ErrorResponse(c, status, fmt.Sprintf("Profile error: %v", err.Error()))
	}

	return response.SuccessResponse(c, status, user)
}
