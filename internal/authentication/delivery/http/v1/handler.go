package v1

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Refresh(c *fiber.Ctx) error
}
