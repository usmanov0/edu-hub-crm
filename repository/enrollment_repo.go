package repository

import (
	"context"
	"edu-sphere-crm/entity"
	"github.com/jackc/pgx"
)

type EnrollmentRepository interface {
	Save(ctx context.Context, enrollment *entity.Enrollments) error
	GetAllEnrollments(ctx context.Context) ([]entity.Enrollments, error)
	GetById(ctx context.Context, enrollmentId int) (*entity.Enrollments, error)
	Delete(ctx context.Context, enrollmentId int) error
}

type enrollmentRepository struct {
	db *pgx.Conn
}

func NewEnrollmentRepo(db *pgx.Conn) EnrollmentRepository {
	return &enrollmentRepository{db: db}
}

func (e *enrollmentRepository) Save(ctx context.Context, enrollment *entity.Enrollments) error {
	//TODO implement me
	panic("implement me")
}

func (e *enrollmentRepository) GetAllEnrollments(ctx context.Context) ([]entity.Enrollments, error) {
	//TODO implement me
	panic("implement me")
}

func (e *enrollmentRepository) GetById(ctx context.Context, enrollmentId int) (*entity.Enrollments, error) {
	//TODO implement me
	panic("implement me")
}

func (e *enrollmentRepository) Delete(ctx context.Context, enrollmentId int) error {
	//TODO implement me
	panic("implement me")
}
