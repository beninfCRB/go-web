package controllers

import (
	"go-web/cmd/app/global"
	"go-web/common/utils"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(global.Userkey)
		if user != nil {
			c.Redirect(http.StatusMovedPermanently, "/dashboard")
			// return
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
		session := sessions.Default(c)
		user := session.Get(global.Userkey)
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

		session.Set(global.Userkey, username)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func LogoutGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(global.Userkey)
		log.Println("logging out user:", user)
		if user == nil {
			log.Println("Invalid session token")
			return
		}

		session.Delete(global.Userkey)
		session.Options(sessions.Options{MaxAge: -1})
		if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func IndexGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(global.Userkey)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Halaman Index",
			"content": "Anda harus melakukan login akun terlebih dahulu ..",
			"user":    user,
		})
	}
}

func DashboardGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(global.Userkey)
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"title":   "Halaman Dashboard",
			"content": "This is a dashboard",
			"user":    user,
		})
	}
}
