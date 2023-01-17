package service

import (
	"time"

	"github.com/linothomas14/hadir-in-api/helper"
	"github.com/linothomas14/hadir-in-api/helper/param"
	"github.com/linothomas14/hadir-in-api/helper/response"
	"github.com/linothomas14/hadir-in-api/model"
	"github.com/linothomas14/hadir-in-api/repository"
)

type EventService interface {
	CreateEvent(event param.CreateEvent, userID uint32) (response.EventRes, error)
}

type eventService struct {
	eventRepository repository.EventRepository
}

func NewEventService(eventRep repository.EventRepository) EventService {
	return &eventService{
		eventRepository: eventRep,
	}
}

func (service *eventService) CreateEvent(event param.CreateEvent, userID uint32) (response.EventRes, error) {
	var eventModel model.Event

	token := helper.TokenGenerator()

	formatTime := "2006-01-02 15:04:05 MST"

	event.Date = event.Date + " WIB"
	event.ExpiredToken = event.ExpiredToken + " WIB"

	date, err := time.Parse(formatTime, event.Date)

	if err != nil {
		return response.EventRes{}, err
	}

	expired_token, err := time.Parse(formatTime, event.ExpiredToken)

	if err != nil {
		return response.EventRes{}, err
	}

	eventModel.UserID = userID
	eventModel.Title = event.Title
	eventModel.Date = date
	eventModel.Token = token
	eventModel.ExpiredToken = expired_token

	eventModel, err = service.eventRepository.CreateEvent(eventModel)

	if err != nil {
		return response.EventRes{}, err
	}

	res := ParseEventResponse(eventModel)

	return res, err
}

func ParseEventResponse(m model.Event) response.EventRes {

	var res response.EventRes

	res.ID = m.ID
	res.Title = m.Title
	res.Date = m.Date
	res.Token = m.Token
	res.ExpiredToken = m.ExpiredToken
	res.UserID = m.UserID
	res.User.ID = m.User.ID
	res.User.Name = m.User.Name
	res.User.Email = m.User.Email

	return res
}
