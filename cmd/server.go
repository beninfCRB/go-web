package main

import (
	"fmt"
	"go-web/cmd/app/routes"
	"go-web/common/database"
	"go-web/common/middleware"
	"go-web/common/utils"
	"log"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//gin new router
	r := gin.New()

	//load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("\nerror Loading .env File")
	}

	//load funchandler database
	database.ConnectDB()

	//load session
	r.Use(middleware.SetSession())

	//load static file
	r.Use(static.Serve("/", static.LocalFile("assets", false)))

	//load html file
	r.LoadHTMLGlob("cmd/app/views/**/*")

	//load recovery & logger
	r.Use(gin.Recovery(), utils.CustomLogger())

	//create file logger
	utils.LogFile()

	//route region
	middleware.NotFound(r)
	public := r.Group("/")
	routes.PublicRoutes(public)

	private := r.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	//listen server
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("\nport not setup in file .env")
		port = "3000"
	}

	r.Run(":" + port)
}
