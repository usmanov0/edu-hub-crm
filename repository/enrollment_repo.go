package repository

import (
	"edu-sphere-crm/entity"
	"github.com/jackc/pgx"
)

type EnrollmentRepository interface {
	Save(enrollment *entity.Enrollments) error
	GetById(enrollmentId int) (*entity.Enrollments, error)
	Update(enrollmentId int) error
	Delete(enrollmentId int) error
}

type enrollmentRepository struct {
	db *pgx.Conn
}

func NewEnrollmentRepo(db *pgx.Conn) EnrollmentRepository {
	return &enrollmentRepository{db: db}
}

func (e *enrollmentRepository) Save(enrollment *entity.Enrollments) error {
	//TODO implement me
	panic("implement me")
}

func (e *enrollmentRepository) GetById(enrollmentId int) (*entity.Enrollments, error) {
	//TODO implement me
	panic("implement me")
}

func (e *enrollmentRepository) Update(enrollmentId int) error {
	//TODO implement me
	panic("implement me")
}

func (e *enrollmentRepository) Delete(enrollmentId int) error {
	//TODO implement me
	panic("implement me")
}
