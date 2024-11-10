package routes

import (
	"net/http"
	"github.com/photoline-club/backend/models"
	"github.com/photoline-club/backend/middleware"
	"github.com/photoline-club/backend/auth"
	"github.com/gin-gonic/gin"
	"strings"
	"strconv"
)

type ImageReq struct{
	Title string `json:"user"` // OPtioal 
	Private bool   `json:"private"`
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

func UpdateImage(context *gin.Context, filename string){
	// we need: Title, Event, Type, Private --> title optional
	var imageIn ImageReq
	if context.BindJSON(&imageIn) != nil{
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Image could not be altered"})
		return 
	}

	context.IndentedJSON(http.StatusOK, imageIn)
	}

func AddImages(context *gin.Context){
		// Multipart form
		db := middleware.GetDB(context)
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]
		id_string := context.Param("id")
		id, err := strconv.ParseUint(id_string, 10, 64)
		private, _ := context.GetQuery("private")
	if err != nil{
		context.IndentedJSON(http.StatusBadRequest, gin.H{})
		return 
	}
		user := middleware.GetUser(context)
		

	for _, file := range files {
			//log.Println(file.Filename)
			parts := strings.Split(file.Filename, ".")
			file_id := auth.GenerateUID(32)
			filename := "images/" + file_id + parts[len(parts)-1]// 32 for security
			// Upload the file to specific dst.
			asset := models.EventAsset{
			UserID: user.ID,
			EventID: uint(id), 
			AssetID: file_id, 
			Type: parts[len(parts)-1],
			Private: (private == "true"), 
		}
			db.Save(&asset)
			
			context.SaveUploadedFile(file, filename)
		}
}

func SetUpImagesRoutes(router *gin.RouterGroup){
	router.GET("/events/:id", middleware.Authenticate(), GetImages)
	router.POST("/events/:id", middleware.Authenticate(), AddImages)
}
