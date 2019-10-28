package routes

import (
	"log"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/bindu-bindu/bindu-blank/app/controllers"
	"github.com/bindu-bindu/bindu-blank/app/middlewares"
	"github.com/gin-gonic/gin"
)

// API REST ROUTES
func API(r *gin.Engine) *gin.Engine {
	// adapter := gormadapter.NewAdapterByDB(db.DB)
	// ok, err := enforce(val.(string), obj, act, adapter)

	r.Use(middlewares.Cors())
	authMiddleware, err := jwt.New(middlewares.GinJwtMiddlewareHandler())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", controllers.Registration)
	r.Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("/admin", controllers.Ping)
	}
	return r
}
