package routers

import (
	"github.com/fireab/goapi2/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/users")
	user.POST("/", controllers.UserCreate)
}
