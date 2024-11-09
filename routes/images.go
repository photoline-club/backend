package routes

import (
	"net/http"
	"github.com/photoline-club/backend/models"
	"github.com/photoline-club/backend/middleware"
	"github.com/gin-gonic/gin"
)

type ImageReq struct{
	Title string `json:"user"` // OPtioal 
	User    models.User   `json:"user", binding:"required"`
	EventID   uint  `json:"event, binding:"required"`
	Type    string `json:"type,binding:"required"`
	Private bool   `json:"private,binding:"required"`
}// TODO: why is above warngin 

// need to changet these, the images, have who posted it the private etc. we send all this to the front 


func GetImages(context *gin.Context){ // the context conatins the infor for the incoming http request
	db := middleware.GetDB(context)
	EventID := context.Param("id")
	
	user := middleware.GetUser(context)

	var images [] models.EventAsset
	db.Model(&models.EventAsset{}).Where("event_id = ? AND (NOT private OR (user_ID = ?))", EventID, user.ID).Find(&images)

	context.IndentedJSON(http.StatusOK, images)
	
}

func AddImage(context *gin.Context){
	// we need: Title, Event, Type, Private --> title optional
	var imageIn ImageReq
	if context.BindJSON(&imageIn) != nil{
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Image could not be added"})
		return 
	}

	context.IndentedJSON(http.StatusOK, imageIn)


		// Multipart form
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]

	for _, file := range files {
			log.Println(file.Filename)
			var filename = auth.GenerateUID()
			// Upload the file to specific dst.
			context.SaveUploadedFile(file, )
		}
		context.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	}



func SetUpImagesRoutes(router *gin.RouterGroup){
	router.GET("/events/:id", middleware.Authenticate(), GetImages)
	router.POST("/events/:id", middleware.Authenticate(), AddImage)
}
