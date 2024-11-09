package routes

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/photoline-club/backend/auth"
	"github.com/photoline-club/backend/database"
	"github.com/photoline-club/backend/middleware"
	"github.com/photoline-club/backend/models"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	var body models.RegisterRequest
	if ctx.BindJSON(&body) != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	db := middleware.GetDB(ctx)
	if database.UsernameExists(db, body.Username) {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}
	user := models.User{
		Firstname:            body.FirstName,
		Lastname:             body.LastName,
		Username:             body.Username,
		Password:             auth.HashPassword(body.Password),
		FriendInvitationCode: auth.GenerateUID(6),
	}
	db.Create(&user)
	ctx.JSON(http.StatusCreated, user)

}

func Login(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	parts := strings.Split(header, ":")
	if len(parts) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	db := middleware.GetDB(ctx)
	var user models.User
	if errors.Is(db.Model(&models.User{}).Where(&models.User{Username: parts[0]}).First(&user).Error,
		gorm.ErrRecordNotFound) {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
		return
	}
	if !auth.PasswordValid(parts[1], user.Password) {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
		return
	}
	token := auth.GenerateToken()
	session := models.Session{
		UserID: user.ID,
		Token:  token,
	}
	db.Save(&session)
	ctx.JSON(http.StatusCreated, gin.H{"success": true, "token": token})
}

func Logout(ctx *gin.Context) {
	session := middleware.GetSession(ctx)
	db := middleware.GetDB(ctx)
	db.Delete(&session)
	ctx.JSON(http.StatusOK, gin.H{})
}

func SetupAuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", Login)
	router.POST("/logout", middleware.Authenticate(), Logout)
	router.POST("/register", Register)
}
