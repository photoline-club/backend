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

func VisibleEventsForUser(db *gorm.DB, uid uint) []models.Event {
	var res []models.Event
	db.Model(&models.Event{}).
		Joins("INNER JOIN event_participants ON event_participants.event_id = events.id").
		Where("event_participants.user_id = ?", uid).
		Order("events.event_start").
        Find(&res)
	return res
}

func GetMutualEvents(db *gorm.DB, uid uint, friendID uint) []models.Event {
	userEvents := VisibleEventsForUser(db, uid)
	for _, e := range userEvents {
		var evt models.EventParticipant
		if !errors.Is(db.Model(&models.EventParticipant{}).
			Where(&models.EventParticipant{
				UserID:  friendID,
				EventID: e.ID,
			}).
			First(&evt).Error, gorm.ErrRecordNotFound) {
			userEvents = append(userEvents, e)
		}
	}
	return userEvents
}
