package controllers

import (
	"net/http"

	"github.com/fireab/goapi2/initializers"
	"github.com/fireab/goapi2/models"
	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	// get the body
	var body struct {
		FullName string
		Email    string
		Password string
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// check if email already exists
	var currentUser models.User
	initializers.DB.Where("email = ?", body.Email).First(&currentUser)
	if currentUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exists",
		})
		return
	}

	// create the user
	user := models.User{FullName: body.FullName, Email: body.Email, Password: body.Password}
	initializers.DB.Create(&user) // pass pointer of data to Create

	c.JSON(http.StatusOK, gin.H{
		"response": user,
	})
}
