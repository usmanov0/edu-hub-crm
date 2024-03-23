package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type Payment struct {
	Id            int
	StudentId     int
	Amount        decimal.Decimal
	PaymentDate   time.Time
	PaymentMethod string
}
