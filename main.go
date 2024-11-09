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
type image struct{ 
	ID		string 	`json:"id"`
		Item		string `json:"item"`
		Completed 		bool `json:"completed"`
}

var images = []image{
	{ID:"1", Item: "", Completed:false},
}

func GetImage(context *gin.Context){ // the context conatins the infor for the incoming http request
	context.IndentedJSON(http.StatusOK, images)
}

func AddImage(context *gin.Context){
	var NewImage image 

	if err := context.BindJSON(&NewImage); err != nil{
		return
	}
	
	images = append(images, NewImage)

	context.IndentedJSON(http.StatusCreated, NewImage)



}

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

	r.GET("/Images", GetImage) // call the get imaghes
	r.POST("/Images", GetImage)
	r.Run("0.0.0.0:8080")
	

	
}

