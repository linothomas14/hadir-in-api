package service

import (
	"github.com/linothomas14/hadir-in-api/model"
	"github.com/linothomas14/hadir-in-api/repository"
)

type AttendanceService interface {
	Attend(attendance model.Attendance, token_event string) (model.Attendance, error)
}

type attendanceService struct {
	attendanceRepository repository.AttendanceRepository
	eventRepository      repository.EventRepository
}

func NewAttendanceService(attendanceRep repository.AttendanceRepository, eventRep repository.EventRepository) AttendanceService {
	return &attendanceService{
		attendanceRepository: attendanceRep,
		eventRepository:      eventRep,
	}
}

func (service *attendanceService) Attend(attendance model.Attendance, token_event string) (model.Attendance, error) {
	event, err := service.eventRepository.GetEventByToken(token_event)

	if err != nil {
		return model.Attendance{}, err
	}

	attendance.EventID = event.ID

	attendance, err = service.attendanceRepository.Insert(attendance)

	if err != nil {
		return model.Attendance{}, err
	}
	return attendance, err
}
