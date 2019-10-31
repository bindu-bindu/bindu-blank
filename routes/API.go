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
	// bindu generate routes User --group v1 --auth true
	// bindu generate scaffold User --group v1 --auth true
	// bindu generate controller User --group v1 --auth true
	r.Use(middlewares.Cors())
	authMiddleware, err := jwt.New(middlewares.GinJwtMiddlewareHandler())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	api := r.Group("/api")
	api.POST("/login", authMiddleware.LoginHandler)
	api.POST("/register", controllers.Registration)
	v1 := api.Group("/v1")

	v1.Use(authMiddleware.MiddlewareFunc())
	{

	}
	return r
}
