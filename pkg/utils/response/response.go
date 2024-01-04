package response

import (
	"github.com/gofiber/fiber/v2"
)

func NewResponse(statusCode int, message string, data interface{}) *fiber.Map {
	return &fiber.Map{
		"status":  statusCode,
		"data":    data,
		"message": message,
	}
}

func NewResponseError(statusCode int, messageStatus string, details string) *fiber.Map {
	return &fiber.Map{
		"status":  statusCode,
		"data":    "",
		"message": messageStatus,
		"error":   details,
	}
}
