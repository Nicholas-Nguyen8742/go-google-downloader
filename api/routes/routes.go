package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message":"OK"})
	})
}