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
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		}

		// Token'ı "Bearer token" formatından ayıkla
		tokenString := authHeader[7:]

		// Token'ı doğrula
		_, err := jwt.Verify(tokenString)
		if err != nil {
			return response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		}

		// Token doğrulandı, bir sonraki handler'ı çağır
		return c.Next()
	}
}
