package routers

import (
	"github.com/fireab/goapi2/controllers"
	"github.com/fireab/goapi2/middlewares"
	"github.com/gin-gonic/gin"
)

func PostRoutes(rg *gin.RouterGroup) {
	post := rg.Group("/posts")
	post.Use(middlewares.AuthMiddleware())

	post.GET("/", controllers.PostsFind)
	post.POST("/", controllers.PostsCreate)
	post.POST("/upload", controllers.UploadFile)
	post.POST("/upload/multiple", controllers.UploadFiles)
	// post.POST("/upload/multiple/anyfile", controllers.UploadAnyFile)

	post.GET("/:id", controllers.FindOne)

}
