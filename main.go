package main

import (
	"github.com/fireab/goapi2/initializers"
	"github.com/fireab/goapi2/routers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}
func main() {
	r := routers.InitRoutes()
	r.Run()
}
