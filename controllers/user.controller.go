package controllers

import (
	"fmt"
	"net/http"

	"github.com/fireab/goapi2/initializers"
	"github.com/fireab/goapi2/models"
	"github.com/fireab/goapi2/utils"
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

	// hash password
	hashedPassword := utils.HashPassword(body.Password)

	// create the user
	user := models.User{FullName: body.FullName, Email: body.Email, Password: hashedPassword}
	initializers.DB.Create(&user) // pass pointer of data to Create

	c.JSON(http.StatusOK, gin.H{
		"response": user,
	})
}

func LoginUser(c *gin.Context) {
	var LoginBody struct {
		Email    string
		Password string
	}
	// get email and password
	c.BindJSON(&LoginBody)
	fmt.Println(LoginBody)
	var user models.User
	initializers.DB.Where("email = ?", LoginBody.Email).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found"})
		return
	}

	comparePassword := utils.CheckPasswordHash(LoginBody.Password, user.Password)
	if !comparePassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"user":  user,
			"token": token,
		})
	}
}
