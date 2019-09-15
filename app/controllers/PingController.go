package controllers

import (
	"fmt"
	"time"

	"github.com/bindu-bindu/bindu-blank/db"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {

	for {
		pingErr := db.DB.DB().Ping()
		if pingErr != nil {
			fmt.Println(pingErr)
		} else {
			fmt.Println("Connected")
		}
		time.Sleep(time.Duration(3) * time.Second)
	}

	c.JSON(200, "pong")
}
