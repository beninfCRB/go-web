package controllers

import (
	"go-web/common/middleware"
	"go-web/common/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := middleware.GetSession(c)
		if user != nil {
			c.Redirect(http.StatusMovedPermanently, "/dashboard")
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title":   "Halaman Login",
			"content": "",
			"user":    user,
		})
	}
}

func LoginPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := middleware.GetSession(c)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Please logout first"})
			return
		}

		username := c.PostForm("username")
		password := c.PostForm("password")

		if utils.EmptyUserPass(username, password) {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
			return
		}

		if !utils.CheckUserPass(username, password) {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect username or password"})
			return
		}

		session := middleware.SaveSession(c, username)
		if session != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func LogoutGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware.ClearSession(c)
		c.Redirect(http.StatusFound, "/")
	}
}

func IndexGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := middleware.GetSession(c)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Halaman Index",
			"content": "Anda harus melakukan login akun terlebih dahulu ..",
			"user":    user,
		})
	}
}

func DashboardGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := middleware.GetSession(c)
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"title":   "Halaman Dashboard",
			"content": "This is a dashboard",
			"user":    user,
		})
	}
}
