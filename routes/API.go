package routes

import (
	"github.com/bindu-bindu/bindu-blank/app/controllers"
	"github.com/bindu-bindu/bindu/World/app/middlewares"
	"github.com/gin-gonic/gin"
)

// API REST ROUTES
func API(r *gin.Engine) *gin.Engine {
	r.Use(middlewares.Cors())
	r.GET("/ping", controllers.Ping)

	return r
}
