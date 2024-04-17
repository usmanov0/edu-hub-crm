package entity

import "time"

type User struct {
	Id               int
	FirstName        string
	LastName         string
	Role             string
	Email            string
	Password         string
	PhoneNumber      string
	Address          string
	RegistrationDate time.Time
	UpdatedAt        time.Time
}

type NewUser struct {
	FirstName   string
	LastName    string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
}

// Role (Moderator, Admin, Student, Mentor)
