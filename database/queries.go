package database

import (
	"errors"

	"github.com/photoline-club/backend/models"
	"gorm.io/gorm"
)

func UsernameExists(db *gorm.DB, username string) bool {
	var user models.User
	return !errors.Is(db.Model(&models.User{}).Where("username = ?", username).First(&user).Error,
		gorm.ErrRecordNotFound)
}

func UsersAreFriends(db *gorm.DB, userA uint, userB uint) bool {
	var obj models.FriendLink
	return !errors.Is(db.Model(&models.FriendLink{}).
		Where(&models.FriendLink{UserID: userA, FriendID: userB}).
		First(&obj).Error,
		gorm.ErrRecordNotFound)
}

