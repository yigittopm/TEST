package response

import (
	"github.com/gofiber/fiber/v2"
)

func NewResponseError(message string, err string) *fiber.Map {
	return &fiber.Map{
		"status":  false,
		"message": message,
		"error":   err,
	}
}

func NewResponseSuccess(message string, data any) *fiber.Map {
	return &fiber.Map{
		"status":  true,
		"message": message,
		"data":    data,
	}
}
