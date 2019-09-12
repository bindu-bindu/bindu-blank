package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/bindu-bindu/bindu-blank/routes"
	"github.com/bindu-bindu/bindu-blank/models"
)

// API REST ROUTES
func API(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	return r
}
