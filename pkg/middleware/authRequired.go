package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yigittopm/wl-auth/pkg/jwt"
	"github.com/yigittopm/wl-auth/pkg/utils/response"
)

func RoleRequired(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtCookie := c.Cookies("jwt")
		if jwtCookie == "" {
			return response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		}

		//tokenString := jwtCookie[7:]
		tokenString := jwtCookie

		userId, err := jwt.Verify(tokenString)
		if err != nil {
			return response.ErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %v", err.Error()))
		}

		roleCookie := c.Cookies("role")
		if roleCookie == "" || roleCookie != role {
			return response.ErrorResponse(c, http.StatusForbidden, "Access denied")
		}

		c.Locals("userId", userId)
		c.Locals("role", role)

		return c.Next()
	}
}

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtCookie := c.Cookies("jwt")
		if jwtCookie == "" {
			return response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		}

		//tokenString := jwtCookie[7:]
		tokenString := jwtCookie

		userId, err := jwt.Verify(tokenString)
		if err != nil {
			return response.ErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %v", err.Error()))
		}

		c.Locals("userId", userId)

		return c.Next()
	}
}
