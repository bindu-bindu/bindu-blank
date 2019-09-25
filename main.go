package main

import (
	"github.com/bindu-bindu/bindu-blank/bindu"
	"github.com/bindu-bindu/bindu-blank/db"
	"github.com/bindu-bindu/bindu-blank/routes"
	"github.com/gin-gonic/gin"
)

var err error

func main() {

	bindu.Init()
	db.Con()
	r := gin.Default()
	r = routes.API(r)

	defer db.DB.Close()
	r.Run()
}
