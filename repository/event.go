package repository

import (
	"github.com/linothomas14/hadir-in-api/model"
	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(event model.Event) (model.Event, error)
}

type eventConnection struct {
	connection *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventConnection{
		connection: db,
	}
}

func (db *eventConnection) CreateEvent(event model.Event) (model.Event, error) {
	err := db.connection.Save(&event).Error
	if err != nil {
		return model.Event{}, err
	}
	return event, err
}
