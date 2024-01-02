package v1

import (
	"github.com/labstack/echo"
	"github.com/yigittopm/test/config"
)

func UserRoute(version *echo.Group, h Handler, cfg config.Config) {
	users := version.Group("users")

	users.POST("/create", h.CreateUser)
	users.PATCH("/", h.UpdateUser)
}
