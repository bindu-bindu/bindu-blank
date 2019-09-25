package main

import (
	"github.com/bindu-bindu/bindu-blank/bindu"
	"github.com/bindu-bindu/bindu-blank/db"
	"github.com/bindu-bindu/bindu-blank/routes"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "./docs"
)

var err error

func main() {
	// http://localhost:8080/swagger/index.html
	// @title Bindu Swagger API
	// @version 1.0
	// @description This is a bindu server bindu-blank server.
	// @termsOfService http://swagger.io/terms/

	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email support@swagger.io

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	// @host localhost:8080
	// @BasePath /

	bindu.Init()
	db.Con()
	r := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r = routes.API(r)

	defer db.DB.Close()
	r.Run()
}
