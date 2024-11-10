package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/photoline-club/backend/database"
	"github.com/photoline-club/backend/middleware"
	"github.com/photoline-club/backend/models"
	"gorm.io/gorm"
)

func GetInvitationToken(ctx *gin.Context) {
	user := middleware.GetUser(ctx)
	ctx.JSON(http.StatusOK, gin.H{"code": user.FriendInvitationCode})
}

func ListFriends(ctx *gin.Context) {
	db := middleware.GetDB(ctx)
	user := middleware.GetUser(ctx)
	var friends []models.FriendLink
	db.Model(&models.FriendLink{}).
		Where(&models.FriendLink{UserID: user.ID}).
		Preload("Friend").
		Find(&friends)
	ctx.JSON(http.StatusOK, gin.H{"data": friends})
}

func AddFriend(ctx *gin.Context) {
	db := middleware.GetDB(ctx)
	user := middleware.GetUser(ctx)
	var body models.AddFriendRequest
	if ctx.BindJSON(&body) != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
		return
	}
	var friend models.User
	if errors.Is(db.Model(&models.User{}).
		Where(&models.User{
			FriendInvitationCode: body.Token,
		}).
		First(&friend).Error, gorm.ErrRecordNotFound) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{})
		return
	}

	if database.UsersAreFriends(db, user.ID, friend.ID) || friend.ID == user.ID {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{})
		return
	}
	db.Create(&models.FriendLink{
		UserID:   user.ID,
		FriendID: friend.ID,
	})
	db.Create(&models.FriendLink{
		UserID:   friend.ID,
		FriendID: user.ID,
	})
	ctx.JSON(http.StatusCreated, gin.H{})

}

func SetupFriendsRoutes(router *gin.RouterGroup) {
	router.GET("/friends", middleware.Authenticate(), ListFriends)
	router.POST("/friends", middleware.Authenticate(), AddFriend)
	router.GET("/friendcode", middleware.Authenticate(), GetInvitationToken)
}
