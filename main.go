package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/photoline-club/backend/config"
	"github.com/photoline-club/backend/database"
	"github.com/photoline-club/backend/middleware"
	"github.com/photoline-club/backend/routes"
)

func main() {
    db := database.InitialiseDB(config.GetDBConfig())

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "online")
	})

	router := r.Group("/api")
    router.Use(middleware.InjectDB(db))
    router.Use(middleware.CORSMiddleware())
	routes.SetupRoutes(router)

	r.Run("0.0.0.0:8080")
}
