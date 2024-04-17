package entity

import "time"

type Course struct {
	Id          int
	CourseName  string
	Description string
	UserId      int
	Capacity    int
	UpdatedAt   time.Time
	StartDate   time.Time
	EndDate     time.Time
}
