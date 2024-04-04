package repository

import (
	"edu-sphere-crm/entity"
	"github.com/jackc/pgx"
)

type EventAttendRepository interface {
	Save(event *entity.Event) error
}

type eventAttendRepo struct {
	db *pgx.Conn
}

func NewEventAttendRepo(db *pgx.Conn) EventAttendRepository {
	return &eventAttendRepo{db: db}
}

func (e *eventAttendRepo) Save(event *entity.Event) error {
	//TODO implement me
	panic("implement me")
}
