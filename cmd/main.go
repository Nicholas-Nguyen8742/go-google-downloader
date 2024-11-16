package main

import (
	"go-google-downloader/api/routes"
	"go-google-downloader/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	bootstrap.App()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}