package middleware

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var Secret = []byte(os.Getenv("SECRET_KEY"))

const Userkey = "user"

func SetSession() gin.HandlerFunc {
	store := cookie.NewStore(Secret)
	age, _ := strconv.Atoi(os.Getenv("MAX_AGE"))
	if age > 0 {
		store.Options(sessions.Options{MaxAge: age * 3600, Path: "/", Secure: true, HttpOnly: true, SameSite: http.SameSiteStrictMode})
	}
	return sessions.Sessions(os.Getenv("SESSION"), store)
}

func GetSession(c *gin.Context) interface{} {
	session := sessions.Default(c)
	user := session.Get(Userkey)
	return user
}

func SaveSession(c *gin.Context, user string) interface{} {
	session := sessions.Default(c)
	session.Set(Userkey, user)
	s := session.Save()
	return s
}

func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(Userkey)
	log.Println("logging out user:", user)
	if user == nil {
		log.Println("Invalid session token")
		return
	}

	session.Delete(Userkey)
	session.Save()
}
