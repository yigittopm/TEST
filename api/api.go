// Package api
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	return router
}

