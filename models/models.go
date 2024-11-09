package models

import "time"

type User struct {
	ID                   uint   `json:"id,omitempty"`
	Firstname            string `json:"firstname,omitempty"`
	Lastname             string `json:"lastname,omitempty"`
	Username             string `json:"username,omitempty"`
	Password             string `json:"-"`
	FriendInvitationCode string `json:"friend_invitation_code,omitempty"`
}

type FriendLink struct {
	ID       uint `json:"id,omitempty"`
	User     User `json:"user,omitempty"`
	UserID   uint `json:"user_id,omitempty"`
	Friend   User `json:"friend,omitempty"`
	FriendID uint `json:"friend_id,omitempty"`
}

type Event struct {
	ID          uint      `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	EventStart  time.Time `json:"event_start,omitempty"`
	EventEnd    time.Time `json:"event_end,omitempty"`
}

type EventParticipant struct {
	ID      uint  `json:"id,omitempty"`
	User    User  `json:"user,omitempty"`
	UserID  uint  `json:"user_id,omitempty"`
	Event   Event `json:"event,omitempty"`
	EventID uint  `json:"event_id,omitempty"`
}

type EventAsset struct {
	ID      uint   `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	User    User   `json:"user,omitempty"`
	UserID  uint   `json:"user_id,omitempty"`
	Event   Event  `json:"event,omitempty"`
	EventID uint   `json:"event_id,omitempty"`
	Type    string `json:"type,omitempty"`
	Private bool   `json:"private,omitempty"`
	AssetID string `json:"asset_id,omitempty"`
}

type Session struct {
	ID     uint
	User   User
	UserID uint
	Token  string
}
