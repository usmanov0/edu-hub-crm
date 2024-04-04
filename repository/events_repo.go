package repository

import (
	"edu-sphere-crm/entity"
	"github.com/jackc/pgx"
)

type EventRepository interface {
	Save(event *entity.Event) error
	GetEventById(eventId int) (entity.Event, error)
	Update(eventId int) error
	Delete(eventId int) error
}

type eventRepository struct {
	db *pgx.Conn
}

func NewEventsRepo(db *pgx.Conn) EventRepository {
	return &eventRepository{db: db}
}

func (e *eventRepository) Save(event *entity.Event) error {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) GetEventById(eventId int) (entity.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) Update(eventId int) error {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) Delete(eventId int) error {
	//TODO implement me
	panic("implement me")
}
