package controllers

import (
	"go-web/cmd/app/models"
	"go-web/common/database"
	"go-web/common/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := middleware.GetSession(c)
		c.HTML(http.StatusOK, "user_create.html", gin.H{
			"title": "Create Data User",
			"user":  user,
		})
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		user := middleware.GetSession(c)
		if err := c.ShouldBind(&input); err != nil {
			c.HTML(http.StatusBadRequest, "user_create.html", gin.H{
				"message": err.Error(),
				"user":    user,
			})
			return
		}

		database.DB.Create(&input)
		c.Redirect(http.StatusCreated, "/user/")
	}
}
