package main

import (
	"github.com/fireab/goapi2/initializers"
	"github.com/fireab/goapi2/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := routers.InitRoutes()
	r.Use(gin.Logger())
	r.Run()
}
