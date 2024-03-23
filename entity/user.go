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

type UserRole struct {
	Id       int
	RoleName string
}

// Role (Moderator, Admin, Student, Mentor)
