package routes

import (
	"github.com/bindu-bindu/bindu-blank/app/controllers"
	"github.com/bindu-bindu/bindu-blank/app/middlewares"
	"github.com/gin-gonic/gin"
)

// API REST ROUTES
func API(r *gin.Engine) *gin.Engine {
	r.Use(middlewares.Cors())
	r.POST("/login", middlewares.GinJwtMiddlewareHandler().LoginHandler)
	r.POST("/register", controllers.Registration)

	return r
}
