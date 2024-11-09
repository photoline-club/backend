package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/photoline-club/backend/config"
	"github.com/photoline-club/backend/database"
	"github.com/photoline-club/backend/middleware"
	"github.com/photoline-club/backend/routes"
)

// need to changet these, the images, have who posted it the private etc. we send all this to the front 
type todo struct{ 
	ID		string 	`json:"id"`
		Item		string `json:"item"`
		Completed 		bool `json:"completed"`
}

var todos = []todo{
	{ID:"1", Item: "", Completed:false},
}

func getImage(context *gin.Context){ // the context conatins the infor for the incoming http request
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
    db := database.InitialiseDB(config.GetDBConfig())

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "online")
	})

	router := r.Group("/api")
    router.Use(middleware.InjectDB(db))
	routes.SetupRoutes(router)

	r.Run("0.0.0.0:8080")
	
	r.GET("/Images", getImage) // call the get imaghes
	
}

