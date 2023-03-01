package main

import (
	"fmt"
	"go-web/cmd/app/models"
	"go-web/common/database"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	//load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("\nerror Loading .env File")
	}
}

func main() {
	//load funchandler database
	db, _ := database.ConnectDB()

	//run migration
	var mg = db.AutoMigrate(&models.User{})

	//check migration
	if mg != nil {
		fmt.Println("\n\n Migration Unsuccessfully !!")
	} else {
		fmt.Println("\n Migration complete")
	}
}
