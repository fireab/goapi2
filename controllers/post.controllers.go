package controllers

import (
	"net/http"

	"github.com/fireab/goapi2/initializers"
	"github.com/fireab/goapi2/models"
	"github.com/gin-gonic/gin"
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
	initializers.DB.Create(&post) // pass pointer of data to Create

	// return
	c.JSON(http.StatusOK, gin.H{
		"message": post,
	})
}

func PostsFind(c *gin.Context) {

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
