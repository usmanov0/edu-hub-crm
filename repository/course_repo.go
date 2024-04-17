package repository

import (
	"context"
	"edu-sphere-crm/entity"
	"edu-sphere-crm/pkg/customErrors"
	"github.com/jackc/pgx"
	"time"
)

type CourseRepository interface {
	Save(ctx context.Context, course *entity.Course) error
	GetById(ctx context.Context, courseId int) (entity.Course, error)
	GetAll(ctx context.Context) ([]entity.Course, error)
	Delete(ctx context.Context, courseId int) error
}

type courseRepository struct {
	db *pgx.Conn
}

func NewCourseRepo(db *pgx.Conn) CourseRepository {
	return &courseRepository{db: db}
}

func (c *courseRepository) Save(ctx context.Context, course *entity.Course) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
			INSERT INTO courses(course_name, description, user_id, capacity, updated_at, start_date, end_date)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING id
			`

	err := c.db.QueryRow(
		query,
		course.CourseName,
		course.Description,
		course.UserId,
		course.Capacity,
		course.UpdatedAt,
		course.StartDate,
		course.EndDate,
	).Scan(&course.Id)

	if err != nil {
		return err
	}

	return nil
}

func (c *courseRepository) GetById(ctx context.Context, courseId int) (entity.Course, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
			SELECT c.id, 
			       c.name, 
			       c.description, 
			       c.user_id, 
			       c.capacity, 
			       c.updated_at, 
			       c.start_date, 
			       c.end_date
			FROM courses c
			WHERE c.id = $1`

	var course entity.Course
	err := c.db.QueryRow(
		query,
		courseId,
	).Scan(
		&course.Id,
		&course.CourseName,
		&course.Description,
		&course.UserId,
		&course.Capacity,
		&course.UpdatedAt,
		&course.StartDate,
		&course.EndDate,
	)

	if err != nil {
		return entity.Course{}, err
	}

	return course, nil
}

func (c *courseRepository) GetAll(ctx context.Context) ([]entity.Course, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
			SELECT c.id, 
			       c.name, 
			       c.description, 
			       c.user_id, 
			       c.capacity, 
			       c.updated_at, 
			       c.start_date, 
			       c.end_date
			FROM courses c
			WHERE c.id = $1`

	row, err := c.db.Query(query)
	if err != nil {
		return []entity.Course{}, err
	}
	var courses []entity.Course
	for row.Next() {
		var course entity.Course
		err = row.Scan(
			&course.Id,
			&course.CourseName,
			&course.Description,
			&course.UserId,
			&course.Capacity,
			&course.UpdatedAt,
			&course.StartDate,
			&course.EndDate,
		)
		if err != nil {
			return []entity.Course{}, customErrors.ErrIdScanFailed
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (c *courseRepository) Delete(ctx context.Context, courseId int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `DELETE FROM courses WHERE id = $1`
	_, err := c.db.Exec(query, courseId)
	if err != nil {
		return err
	}
	return nil
}
