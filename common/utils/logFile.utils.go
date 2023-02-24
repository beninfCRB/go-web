package utils

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func LogFile() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}
