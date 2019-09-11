package main

import (
	"github.com/bindu-bindu/bindu-blank/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r = routes.API(r)
	r.Run()
}
