package repository

import (
	"log"

	"github.com/linothomas14/hadir-in-api/model"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Insert(attendance model.Attendance) (model.Attendance, error)
}

type attendanceConnection struct {
	connection *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceConnection{
		connection: db,
	}
}

func (db *attendanceConnection) Insert(attendance model.Attendance) (model.Attendance, error) {
	err := db.connection.Preload("User").Preload("Event").Save(&attendance).Find(&attendance).Error

	if err != nil {
		return model.Attendance{}, err
	}
	log.Println("user : ", attendance.Event)
	return attendance, err
}
