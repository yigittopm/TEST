package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yigittopm/wl-auth/pkg/jwt"
	"github.com/yigittopm/wl-auth/pkg/utils/response"
)

func AuthRequired(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtCookie := c.Cookies("jwt")
		if jwtCookie == "" {
			return response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		}

		//bearerString := jwtCookie[:6]
		//if bearerString != "Bearer " {
		//	return response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Bearer token is missing")
		//}

		tokenString := jwtCookie[7:]
		userId, err := jwt.Verify(tokenString)
		if err != nil {
			return response.ErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %v", err.Error()))
		}

		roleCookie := c.Cookies("role")
		if roleCookie == "" || roleCookie != role {
			return response.ErrorResponse(c, http.StatusForbidden, "Access denied")
		}

		c.Locals("role", roleCookie)
		c.Locals("userId", userId)

		return c.Next()
	}
}
