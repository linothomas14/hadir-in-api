package model

import "time"

type User struct {
	ID        uint32    `json:"id" gorm:"primaryKey;notNull"`
	Email     string    `json:"email" gorm:"unique;notNull" validate:"email"`
	Name      string    `json:"name" gorm:"notNull"`
	Password  string    `json:"password" gorm:"notNull" validate:"min=6"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (User) TableName() string {
	return "user"
}
