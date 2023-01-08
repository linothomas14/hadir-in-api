package param

import "time"

type CreateEvent struct {
	Title        string    `json:"time"`
	Date         time.Time `json:"date"`
	ExpiredToken time.Time `json:"expired_token"`
}
