package service

import (
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

	eventModel.Title = event.Title
	eventModel.Date = event.Date
	eventModel.Token = token
	eventModel.ExpiredToken = event.ExpiredToken

	eventModel, err := service.eventRepository.CreateEvent(eventModel)

	if err != nil {
		return model.Event{}, err
	}

	return eventModel, err
}
