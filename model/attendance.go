package model

import "time"

type Attendance struct {
	ID        uint32    `json:"id" gorm:"primaryKey;notNull"`
	UserID    uint32    `json:"user_id" gorm:"notNull;foreignKey:UserID"`
	User      User      `json:"user"`
	EventID   uint32    `json:"event_id" gorm:"notNull;foreignKey:EventID"`
	Event     Event     `json:"event"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (Attendance) TableName() string {
	return "attendance"
}
