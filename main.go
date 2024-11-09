package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/photoline-club/backend/config"
	"github.com/photoline-club/backend/database"
)

func main() {
	database.InitialiseDB(config.GetDBConfig())

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "online")
	})
	r.Run("0.0.0.0:8080")
}
