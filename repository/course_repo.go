package repository

import (
	"edu-sphere-crm/entity"
	"github.com/jackc/pgx"
)

type CourseRepository interface {
	Save(course *entity.Course) error
	GetById(courseId int) (entity.Course, error)
	Update(courseId int) error
	Delete(courseId int) error
}

type courseRepository struct {
	db *pgx.Conn
}

func NewCourseRepo(db *pgx.Conn) CourseRepository {
	return &courseRepository{db: db}
}

func (c *courseRepository) Save(course *entity.Course) error {
	//TODO implement me
	panic("implement me")
}

func (c *courseRepository) GetById(courseId int) (entity.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c *courseRepository) Update(courseId int) error {
	//TODO implement me
	panic("implement me")
}

func (c *courseRepository) Delete(courseId int) error {
	//TODO implement me
	panic("implement me")
}
