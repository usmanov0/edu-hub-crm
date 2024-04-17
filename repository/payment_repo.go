package repository

import (
	"context"
	"edu-sphere-crm/entity"
	"github.com/jackc/pgx"
)

type PaymentRepo interface {
	Save(ctx context.Context, payment *entity.Payment) error
}

type paymentRepository struct {
	db *pgx.Conn
}

func NewPaymentRepo(db *pgx.Conn) PaymentRepo {
	return &paymentRepository{db: db}
}

func (p *paymentRepository) Save(ctx context.Context, payment *entity.Payment) error {
	//TODO implement me
	panic("implement me")
}
