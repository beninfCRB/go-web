package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "notFound.html", gin.H{
			"title": "Page Not Found",
		})
	})
}
