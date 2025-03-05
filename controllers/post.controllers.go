package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fireab/goapi2/initializers"
	"github.com/fireab/goapi2/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostsCreate(c *gin.Context) {
	// get the body
	var body struct {
		Title string
		Body  string
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// create the post
	post := models.Post{Title: body.Title, Body: body.Body}
	initializers.DB.Create(&post)

	// return
	c.JSON(http.StatusOK, gin.H{
		"message": post,
	})
}

func PostsFind(c *gin.Context) {
	val, _ := c.Get("user_id")
	fmt.Println(val)
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}

func FindOne(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	// check if post is found
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UploadFile(c *gin.Context) {
	// single file
	file, err1 := c.FormFile("file")

	if err1 != nil {
		fmt.Println("error", err1)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read file",
		})
		return
	}

	fmt.Println("file information", file.Filename, file.Size, file.Header)
	fileExtention := strings.Split(file.Filename, ".")[len(strings.Split(file.Filename, "."))-1]
	updatedFileName := uuid.New().String() + "." + fileExtention

	err := c.SaveUploadedFile(file, "./uploads/"+updatedFileName)
	if err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to upload file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"file": updatedFileName,
	})
}
