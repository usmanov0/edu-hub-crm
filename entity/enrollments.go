package entity

import "time"

type Enrollments struct {
	Id             int
	StudentId      int
	CourseId       int
	EnrollmentDate time.Time
	Status         EnrollmentStatus
}

type EnrollmentStatus string

const (
	Completed  EnrollmentStatus = "Completed"
	Enrolled   EnrollmentStatus = "Enrolled"
	InProgress EnrollmentStatus = "In Progress"
	Withdrawn  EnrollmentStatus = "Withdrawn"
	Failed     EnrollmentStatus = "Failed"
	OnHold     EnrollmentStatus = "On Hold"
	Cancelled  EnrollmentStatus = "Cancelled"
)

// Status can be ('Completed', 'Enrolled', 'In Progress', 'Withdrawn', 'Failed', 'On Hold', 'Cancelled')
