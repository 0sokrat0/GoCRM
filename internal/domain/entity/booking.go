package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusConfirmed Status = "confirmed"
	StatusCanceled  Status = "canceled"
)

type Booking struct {
	BookingID uuid.UUID `json:"booking_id"`

	UserID    uuid.UUID `json:"user_id"`
	MasterID  uuid.UUID `json:"master_id"`
	ServiceID uuid.UUID `json:"service_id"`

	Status Status `json:"status"`

	BookingTime time.Time `json:"booking_time"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewBooking(userID, masterID, serviceID uuid.UUID, bookingTime time.Time) (*Booking, error) {
	if bookingTime.Before(time.Now()) {
		return nil, errors.New("booking time cannot be in the past")
	}
	return &Booking{
		BookingID:   uuid.New(),
		UserID:      userID,
		MasterID:    masterID,
		ServiceID:   serviceID,
		Status:      StatusPending,
		BookingTime: bookingTime,
		CreatedAt:   time.Now(),
	}, nil
}

func (b *Booking) Confirm() error {
	if b.Status == StatusCanceled {
		return errors.New("booking is already canceled")
	}
	b.Status = StatusCanceled
	return nil
}

func (b *Booking) Reschedule(newTime time.Time) error {
	if newTime.Before(time.Now()) {
		return errors.New("new booking time cannot be in the past")
	}
	b.BookingTime = newTime
	return nil
}
