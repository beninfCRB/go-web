package routes

import (
	"go-web/cmd/app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	user := r.Group("/user")
	user.GET("/", controllers.ShowUser())
	user.POST("/create", controllers.CreateUser())
}
