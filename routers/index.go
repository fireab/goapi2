package routers

import (
	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all grouped routes
func InitRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	PostRoutes(api)
	UserRoutes(api)

	return r
}
