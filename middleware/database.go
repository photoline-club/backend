package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const CTX_DB_KEY = "CTX_DB_KEY"

func InjectDB(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(CTX_DB_KEY, db)
	}
}

func GetDB(ctx *gin.Context) *gorm.DB {
	db, exists := ctx.Get(CTX_DB_KEY)
	if !exists {
		return nil
	}
	return db.(*gorm.DB)
}
