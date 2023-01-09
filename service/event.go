package service

import (
	"time"

	"github.com/linothomas14/hadir-in-api/helper"
	"github.com/linothomas14/hadir-in-api/helper/param"
	"github.com/linothomas14/hadir-in-api/model"
	"github.com/linothomas14/hadir-in-api/repository"
)

type EventService interface {
	CreateEvent(event param.CreateEvent) (model.Event, error)
}

type eventService struct {
	eventRepository repository.EventRepository
}

func NewEventService(eventRep repository.EventRepository) EventService {
	return &eventService{
		eventRepository: eventRep,
	}
}

func (service *eventService) CreateEvent(event param.CreateEvent) (model.Event, error) {
	var eventModel model.Event

	token := helper.TokenGenerator()

	formatTime := "2006-01-02 15:04:05 MST"

	event.Date = event.Date + " WIB"
	event.ExpiredToken = event.ExpiredToken + " WIB"

	date, err := time.Parse(formatTime, event.Date)

	if err != nil {
		return model.Event{}, err
	}

	expired_token, err := time.Parse(formatTime, event.ExpiredToken)

	if err != nil {
		return model.Event{}, err
	}

	eventModel.Title = event.Title
	eventModel.Date = date
	eventModel.Token = token
	eventModel.ExpiredToken = expired_token

	eventModel, err = service.eventRepository.CreateEvent(eventModel)

	if err != nil {
		return model.Event{}, err
	}

	return eventModel, err
}
