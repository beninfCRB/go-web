package utils

import (
	"go-web/cmd/app/global"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Session(c *gin.Engine) {
	store := cookie.NewStore(global.Secret)
	age, _ := strconv.Atoi(os.Getenv("MAX_AGE"))
	if age > 0 {
		store.Options(sessions.Options{MaxAge: age, Path: "/", Secure: true})
	}
	c.Use(sessions.Sessions(os.Getenv("SESSION"), store))
}
