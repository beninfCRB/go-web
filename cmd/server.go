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
	r := gin.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("\nerror Loading .env File")
	}

	database.ConnectDB()
	r.Use(middleware.SetSession())
	r.Use(static.Serve("/", static.LocalFile("assets", false)))
	r.LoadHTMLGlob("cmd/app/views/**/*")

	r.Use(gin.Recovery(), utils.CustomLogger())
	utils.LogFile()

	middleware.NotFound(r)
	public := r.Group("/")
	routes.PublicRoutes(public)

	private := r.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("\nport not setup in file .env")
		port = "3000"
	}

	r.Run(":" + port)
}
