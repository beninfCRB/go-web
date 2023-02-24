package main

import (
	"fmt"
	"go-web/cmd/app/routes"
	"go-web/common/database"
	"go-web/common/middleware"
	"go-web/common/utils"
	"log"
	"net/http"
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
	r.Use(static.Serve("/", static.LocalFile("assets", false)))
	r.LoadHTMLGlob("cmd/app/views/**/*")

	r.Use(gin.Recovery(), utils.CustomLogger())
	utils.Session(r)
	utils.LogFile()

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "notFound.html", gin.H{
			"title": "Page Not Found",
		})
	})

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
