package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/photoline-club/backend/database"
	"github.com/photoline-club/backend/middleware"
	"github.com/photoline-club/backend/models"
)

func ListEvents(ctx *gin.Context) {
	db := middleware.GetDB(ctx)
	user := middleware.GetUser(ctx)
	idstr, _ := ctx.GetQuery("user_id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	var events []models.Event
	if err != nil {
		events = database.VisibleEventsForUser(db, user.ID)
	} else {
		if !database.UsersAreFriends(db, user.ID, uint(id)) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
			return
		}
		events = database.GetMutualEvents(db, user.ID, uint(id))
	}
	ctx.JSON(http.StatusOK, gin.H{"events": events})
}

func CreateEvent(ctx *gin.Context) {
	var body models.CreateEventRequest
	if ctx.BindJSON(&body) != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
		return
	}
	db := middleware.GetDB(ctx)
	user := middleware.GetUser(ctx)
	for _, id := range body.Users {
		if !database.UsersAreFriends(db, user.ID, id) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
			return
		}
	}
	evt := models.Event{
		Title:       body.Title,
		Description: body.Description,
		EventStart:  body.StartDate,
		EventEnd:    body.StartDate,
	}
	db.Save(&evt)
	for _, id := range append(body.Users, user.ID) {
		db.Save(&models.EventParticipant{
			UserID:  id,
			EventID: evt.ID,
		})
	}
	ctx.JSON(http.StatusCreated, gin.H{"event": evt})
}

func SetupEventsRoutes(router *gin.RouterGroup) {
	router.GET("/events", middleware.Authenticate(), ListEvents)
	router.POST("/events", middleware.Authenticate(), CreateEvent)
}
