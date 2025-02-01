package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type PaymentStatus string

const (
	PaymentPending PaymentStatus = "pending"
	PaymentPaid    PaymentStatus = "paid"
	PaymentFailed  PaymentStatus = "failed"
)

type Payment struct {
	PaymentID     uuid.UUID     `json:"payment_id"`
	BookingID     uuid.UUID     `json:"booking_id"`
	Amount        float64       `json:"amount"`
	PaymentMethod string        `json:"payment_method"`
	Status        PaymentStatus `json:"status"`
	CreatedAt     time.Time     `json:"created_at"`
}

func NewPayment(bookingID uuid.UUID, amount float64, PaymentMethod string) (*Payment, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	return &Payment{
		PaymentID:     uuid.New(),
		BookingID:     bookingID,
		Amount:        amount,
		PaymentMethod: PaymentMethod,
		Status:        PaymentPending,
		CreatedAt:     time.Now(),
	}, nil
}

func (p *Payment) MarkAsPaid() error {
	if p.Status != PaymentPending {
		return errors.New("only pending payments can be marked as paid")
	}
	p.Status = PaymentPaid
	return nil
}

func (p *Payment) MarkAsFailed() error {
	if p.Status != PaymentPending {
		return errors.New("only pending payments can be marked as failed")
	}
	p.Status = PaymentFailed
	return nil
}
