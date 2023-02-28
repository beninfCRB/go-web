package routes

import (
	"go-web/cmd/app/controllers"
	"go-web/common/middleware"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/login", controllers.LoginGet())
	g.POST("/login", controllers.LoginPost())
	g.GET("/", controllers.IndexGet())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.Use(middleware.AuthRequired)
	g.GET("/dashboard", controllers.DashboardGet())
	g.GET("/logout", controllers.LogoutGet())

}
