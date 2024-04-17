package repository

import (
	"context"
	"edu-sphere-crm/entity"
	"github.com/jackc/pgx"
)

type EventRepository interface {
	Save(ctx context.Context, event *entity.Event) error
	GetAllEvents(ctx context.Context) ([]entity.Event, error)
	GetEventById(ctx context.Context, eventId int) (entity.Event, error)
	GetEventByName(ctx context.Context, eventName string) ([]entity.Event, error)
	Delete(ctx context.Context, eventId int) error
}

type eventRepository struct {
	db *pgx.Conn
}

func NewEventsRepo(db *pgx.Conn) EventRepository {
	return &eventRepository{db: db}
}

func (e *eventRepository) Save(ctx context.Context, event *entity.Event) error {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) GetAllEvents(ctx context.Context) ([]entity.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) GetEventById(ctx context.Context, eventId int) (entity.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) GetEventByName(ctx context.Context, eventName string) ([]entity.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (e *eventRepository) Delete(ctx context.Context, eventId int) error {
	//TODO implement me
	panic("implement me")
}
