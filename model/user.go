package model

import "time"

type User struct {
	ID        uint32    `json:"id" gorm:"primaryKey;notNull"`
	Email     string    `json:"email" gorm:"unique;notNull"`
	Name      string    `json:"name" gorm:"notNull"`
	Password  string    `json:"password" gorm:"notNull"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (User) TableName() string {
	return "user"
}
