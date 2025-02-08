package repo

import (
	"GoCRM/internal/domain/booking/entity"
	"time"

	"github.com/google/uuid"
)

type BookingRepository interface {
	Create(booking *entity.Booking) error
	GetByID(bookingID uuid.UUID) (*entity.Booking, error)
	GetByUserID(userID uuid.UUID) ([]*entity.Booking, error)
	GetByMasterID(masterID uuid.UUID, from, to time.Time) ([]*entity.Booking, error)
	GetByDate(date time.Time) ([]*entity.Booking, error)
	UpdateStatus(bookingID uuid.UUID, status entity.Status) error
	Delete(bookingID uuid.UUID) error
}
