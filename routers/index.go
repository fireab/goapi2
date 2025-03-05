package routers

import (
	"github.com/fireab/goapi2/middlewares"
	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all grouped routes
func InitRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		uploads := r.Group("/uploads")
		uploads.Use(middlewares.AuthMiddleware())
		uploads.Static("/", "./uploads")
	}

	PostRoutes(api)
	UserRoutes(api)

	return r
}
