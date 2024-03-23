package entity

import "time"

type Event struct {
	Id          int
	EventName   string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Location    string
}
