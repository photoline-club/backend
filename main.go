package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/photoline-club/backend/config"
	"github.com/photoline-club/backend/database"
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
	database.InitialiseDB(config.GetDBConfig())

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "online")
	})
	r.Run("0.0.0.0:8080")
	
	r.GET("/Images", getImage) // call the get imaghes
	
}

