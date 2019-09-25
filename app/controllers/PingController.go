package controllers

import (
	"fmt"
	"time"

	"github.com/bindu-bindu/bindu-blank/db"
	"github.com/gin-gonic/gin"
)

//
// @Summary Ping a test api
// @Description get string
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /ping [get]
func Ping(c *gin.Context) {

	pingErr := db.DB.DB().Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
		c.AbortWithError(400, pingErr)
	} else {
		fmt.Println("Connected")
	}
	time.Sleep(time.Duration(3) * time.Second)

	c.JSON(200, "pong")
}
