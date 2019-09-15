package routes

import (
	"github.com/bindu-bindu/bindu-blank/app/controllers"

	"github.com/gin-gonic/gin"
)

// API REST ROUTES
func API(r *gin.Engine) *gin.Engine {
	r.GET("/ping", controllers.Ping)

	return r
}
