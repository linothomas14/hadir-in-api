package model

type Attendance struct {
}

func (Attendance) TableName() string {
	return "attendance"
}
