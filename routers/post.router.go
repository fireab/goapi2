package routers

import (
	"github.com/fireab/goapi2/controllers"
	"github.com/gin-gonic/gin"
)

func PostRoutes(rg *gin.RouterGroup) {
	post := rg.Group("/posts")

	post.GET("/", controllers.PostsFind)
	post.POST("/", controllers.PostsCreate)
	post.GET("/:id", controllers.FindOne)

}
