package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/photoline-club/backend/models"
	"gorm.io/gorm"
)

const CTX_USER_KEY = "CTX_USER_KEY"
const CTX_SESSION_KEY = "CTX_SESSION_KEY"

func GetUser(ctx *gin.Context) models.User {
	user, _ := ctx.Get(CTX_USER_KEY)
	return user.(models.User)
}

func GetSession(ctx *gin.Context) models.Session {
	user, _ := ctx.Get(CTX_SESSION_KEY)
	return user.(models.Session)
}

func Authenticate(preload ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		token, valid := strings.CutPrefix(header, "Bearer ")
		if !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		db := GetDB(ctx)
		var session models.Session
		if errors.Is(db.Model(&models.Session{}).Where(&models.Session{Token: token}).First(&session).Error,
			gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		var user models.User
		qry := db.Model(&models.User{}).Where(&models.User{
			ID: session.UserID,
		})
		for _, x := range preload {
			qry = qry.Preload(x)
		}
		qry.First(&user)
		ctx.Set(CTX_USER_KEY, user)
		ctx.Set(CTX_SESSION_KEY, session)
		ctx.Next()
	}
}
