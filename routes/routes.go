package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup) {
    SetupAuthRoutes(router)
    SetupFriendsRoutes(router)
    SetupEventsRoutes(router)
    SetUpImagesRoutes(router)
}
