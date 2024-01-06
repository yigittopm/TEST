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
	GetAllUsers() fiber.Handler
	CreateUser() fiber.Handler
	UpdateUserByID() fiber.Handler
	DeleteUserByID() fiber.Handler
}

type handler struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) Handler {
	return &handler{uc: uc}
}

func (h *handler) CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
			payload     dtos.CreateUserRequest
		)
		defer cancel()

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(http.StatusBadRequest).JSON(response.NewResponseError("Parse error.", err.Error()))
		}

		if err := payload.Validate(); err != nil {
			return c.Status(http.StatusBadRequest).JSON(response.NewResponseError("Validate error.", err.Error()))
		}

		userID, err := h.uc.Create(ctx, payload)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(response.NewResponseError("Create error.", err.Error()))
		}

		return c.Status(http.StatusOK).JSON(response.NewResponseSuccess("Successfully created.", userID))
	}
}

func (h *handler) GetAllUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), time.Duration(30*time.Second))
		defer cancel()

		users, err := h.uc.GetAll(ctx)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(response.NewResponseError("Error fetch users.", err.Error()))
		}

		return c.Status(http.StatusOK).JSON(response.NewResponseSuccess("Successfully get all users;", users))
	}
}

func (h *handler) UpdateUserByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON(response.NewResponseSuccess("", nil))
	}
}

func (h *handler) DeleteUserByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
			payload     dtos.DeleteUserByIdRequest
		)
		defer cancel()

		payload.ID = c.Query("userID")
		if payload.ID == "" {
			return c.Status(http.StatusBadRequest).JSON(response.NewResponseError("User id not be null.", ""))
		}

		userId, err := h.uc.Delete(ctx, payload)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(response.NewResponseError("Error delete user.", err.Error()))
		}

		return c.Status(http.StatusOK).JSON(response.NewResponseSuccess("Deleted user.", userId))
	}
}
