package response

import "time"

type EventRes struct {
	ID           uint32       `json:"id" `
	Title        string       `json:"title" `
	Date         time.Time    `json:"date" `
	Token        string       `json:"token"`
	ExpiredToken time.Time    `json:"expired_token"`
	UserID       uint32       `json:"user_id"` // who created the event
	User         UserResponse `json:"user"`
	CreatedAt    time.Time    `json:"-"`
	UpdatedAt    time.Time    `json:"-"`
}
