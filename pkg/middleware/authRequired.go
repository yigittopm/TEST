package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yigittopm/wl-auth/pkg/jwt"
	"github.com/yigittopm/wl-auth/pkg/utils/response"
)

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Authorization header'dan token'ı al
		jwtCookie := c.Cookies("jwt")
		if jwtCookie == "" {
			return response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		}

		// Token'ı "Bearer token" formatından ayıkla
		//tokenString := jwtCookie[7:]
		tokenString := jwtCookie

		// Token'ı doğrula
		userId, err := jwt.Verify(tokenString)
		if err != nil {
			return response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		}

		c.Locals("userId", userId)

		// Token doğrulandı, bir sonraki handler'ı çağır
		return c.Next()
	}
}
