package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yigittopm/test/internal/users/dtos"
	"github.com/yigittopm/test/internal/users/usecase"
)

func newResponseError(err string) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err,
	}
}

func newResponseSuccess(data any) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

type Handler interface {
	CreateUser() fiber.Handler
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
			return c.Status(http.StatusBadRequest).JSON(newResponseError(err.Error()))
		}

		if err := payload.Validate(); err != nil {
			return c.Status(http.StatusBadRequest).JSON(newResponseError(err.Error()))
		}

		userID, err := h.uc.Create(ctx, payload)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(newResponseError(err.Error()))
		}

		return c.Status(http.StatusOK).JSON(newResponseSuccess(userID))
	}

}
