package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func SetupApp() *gin.Engine {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Welcome to Go Backend Engine built with Gin!",
		})
	})

	return app
}
