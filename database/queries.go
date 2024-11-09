package database

import (
	"errors"

	"github.com/photoline-club/backend/models"
	"gorm.io/gorm"
)

func UsernameExists(db *gorm.DB, username string) bool {
	var user models.User
	if errors.Is(db.Model(&models.User{}).Where("username = ?", username).First(&user).Error,
		gorm.ErrRecordNotFound) {
		return false
	}
	return true
}


