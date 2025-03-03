package main

import (
	"github.com/fireab/goapi2/controllers"
	"github.com/fireab/goapi2/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}
func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/post", controllers.PostsFind)

	r.POST("/post", controllers.PostsCreate)
	r.GET(("/post/:id"), controllers.FindOne)
	r.Run()
}
