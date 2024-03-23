package entity

import "time"

type User struct {
	Id               int
	RoleId           int
	FirstName        string
	LastName         string
	Email            string
	PhoneNumber      string
	Address          string
	RegistrationDate time.Time
}

// Role (Moderator, Admin, Student, Mentor)
