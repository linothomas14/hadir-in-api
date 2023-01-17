package model

import (
	"time"
)

type Event struct {
	ID           uint32    `json:"id" gorm:"primaryKey;notNull"`
	Title        string    `json:"title" gorm:"notNull"`
	Date         time.Time `json:"date" gorm:"notNull"`
	Token        string    `json:"token" gorm:"notNull"`
	ExpiredToken time.Time `json:"expired_token" gorm:"notNull"`
	UserID       uint32    `json:"user_id" gorm:"notNull;foreignKey:UserID"` // who created the event
	User         User      `json:"user"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

func (Event) TableName() string {
	return "event"
}

// type myTime time.Time

// var _ json.Unmarshaler = &myTime{}

// func (mt *myTime) UnmarshalJSON(bs []byte) error {
// 	var s string
// 	err := json.Unmarshal(bs, &s)
// 	if err != nil {
// 		return err
// 	}
// 	t, err := time.ParseInLocation("2006-01-02 07:03", s, time.UTC)
// 	if err != nil {
// 		return err
// 	}
// 	*mt = myTime(t)
// 	return nil
// }
