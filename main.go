package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(ctx *gin.Context) {
      ctx.String(http.StatusOK, "online")
    })
    r.Run("0.0.0.0:8080")
  }
