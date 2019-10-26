package controllers

import (
	"github.com/bindu-bindu/bindu-blank/app/models"
	"github.com/bindu-bindu/bindu-blank/db"
	"github.com/gin-gonic/gin"
)

// Registration to a new user
func Registration(c *gin.Context) {
	var db = db.DB
	var user models.User
	c.BindJSON(&user)

	db.Create(&user)
	c.JSON(200, user)
}
