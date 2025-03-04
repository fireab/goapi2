package main

import (
	"log"

	"github.com/fireab/goapi2/initializers"
	"github.com/fireab/goapi2/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{}, &models.Post{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database Migrated")
	}

}
