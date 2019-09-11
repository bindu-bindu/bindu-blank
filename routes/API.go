package routes

import (
	"github.com/gin-gonic/gin"
)

// API REST ROUTES
func API(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	return r
}
