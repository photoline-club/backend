package models

import "time"

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type AddFriendRequest struct {
	Token string `binding:"required" json:"token,omitempty"`
}

type CreateEventRequest struct {
	Title       string    `json:"title,omitempty" binding:"required"`
	Description string    `json:"description,omitempty" binding:"required"`
	Users       []uint    `json:"users,omitempty" binding:"required"`
	StartDate   time.Time `json:"start_date,omitempty" binding:"required"`
	EndDate     time.Time `json:"end_date,omitempty" binding:"required"`
}
