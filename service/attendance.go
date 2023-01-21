package service

import (
	"github.com/linothomas14/hadir-in-api/helper/response"
	"github.com/linothomas14/hadir-in-api/model"
	"github.com/linothomas14/hadir-in-api/repository"
)

type AttendanceService interface {
	Attend(attendance model.Attendance, token_event string) (response.PresentResponse, error)
}

type attendanceService struct {
	attendanceRepository repository.AttendanceRepository
	userRepository       repository.UserRepository
	eventRepository      repository.EventRepository
}

func NewAttendanceService(attendanceRep repository.AttendanceRepository, userRep repository.UserRepository, eventRep repository.EventRepository) AttendanceService {
	return &attendanceService{
		attendanceRepository: attendanceRep,
		userRepository:       userRep,
		eventRepository:      eventRep,
	}
}

func (service *attendanceService) Attend(attendance model.Attendance, token_event string) (response.PresentResponse, error) {
	event, err := service.eventRepository.GetEventByToken(token_event)

	if err != nil {
		return response.PresentResponse{}, err
	}

	attendance.EventID = event.ID

	attendance, err = service.attendanceRepository.Insert(attendance)

	user, err := service.userRepository.GetUser(int(attendance.UserID))
	if err != nil {
		return response.PresentResponse{}, err
	}

	res := ParsePresentResponse(attendance, user)

	if err != nil {
		return response.PresentResponse{}, err
	}

	return res, err
}

func ParsePresentResponse(m model.Attendance, u model.User) response.PresentResponse {

	var res response.PresentResponse

	res.ID = m.ID
	res.User.ID = m.User.ID
	res.User.Name = m.User.Name
	res.User.Email = m.User.Email
	res.Event.ID = m.Event.ID
	res.Event.Title = m.Event.Title
	res.Event.Token = m.Event.Token
	res.Event.Date = m.Event.Date
	res.Event.ExpiredToken = m.Event.ExpiredToken
	res.Event.UserID = m.Event.UserID
	res.Event.User.ID = u.ID
	res.Event.User.Name = u.Name
	res.Event.User.Email = u.Email

	return res
}
