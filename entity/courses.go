package entity

import "time"

type Course struct {
	Id           int
	CourseName   string
	Description  string
	InstructorId int
	Capacity     int
	StartDate    time.Time
	EndDate      time.Time
}
