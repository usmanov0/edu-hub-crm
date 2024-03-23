package entity

import "time"

type Enrollments struct {
	Id             int
	StudentId      int
	CourseId       int
	EnrollmentDate time.Time
	Status         string
}

// Status can be ('Completed', 'Enrolled', 'In Progress', 'Withdrawn', 'Failed', 'On Hold', 'Cancelled')
