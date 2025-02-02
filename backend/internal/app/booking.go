package app

import (
	"context"
	"errors"
	"time"

	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"

	"github.com/google/uuid"
)

type BookingService struct {
	bookingRepo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) *BookingService {
	return &BookingService{
		bookingRepo: repo,
	}
}

// CreateBooking создает новое бронирование.
// Перед созданием бронирования можно добавить дополнительные проверки, например,
// проверку на доступность мастера в заданное время.
func (bs *BookingService) CreateBooking(ctx context.Context, b *entity.Booking) error {
	if b == nil {
		return errors.New("booking cannot be nil")
	}
	if b.BookingTime.Before(time.Now()) {
		return errors.New("booking time cannot be in the past")
	}
	// Здесь можно добавить вызов метода проверки доступности мастера,
	// например: if !masterProfile.IsAvailable(b.BookingTime, service.Duration) { ... }
	return bs.bookingRepo.Create(ctx, b)
}

// GetBooking возвращает бронирование по его идентификатору.
func (bs *BookingService) GetBooking(ctx context.Context, id uuid.UUID) (*entity.Booking, error) {
	b, err := bs.bookingRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, errors.New("booking not found")
	}
	return b, nil
}

// UpdateBooking обновляет данные бронирования.
// Здесь можно проверить, например, что обновление не нарушает инварианты (например, нельзя изменить время, если бронирование уже подтверждено).
func (bs *BookingService) UpdateBooking(ctx context.Context, b *entity.Booking) (*entity.Booking, error) {
	if b == nil {
		return nil, errors.New("booking cannot be nil")
	}
	if b.Status == entity.StatusConfirmed {
		return nil, errors.New("cannot update a confirmed booking")
	}
	return bs.bookingRepo.Update(ctx, b)
}

// CancelBooking изменяет статус бронирования на "canceled".
// Можно добавить дополнительную логику, например, уведомление клиента или возврат средств.
func (bs *BookingService) CancelBooking(ctx context.Context, id uuid.UUID) error {
	// Получаем бронирование
	b, err := bs.bookingRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if b == nil {
		return errors.New("booking not found")
	}
	// Если бронирование уже отменено, возвращаем ошибку
	if b.Status == entity.StatusCanceled {
		return errors.New("booking is already canceled")
	}
	// Меняем статус
	b.Status = entity.StatusCanceled
	_, err = bs.bookingRepo.Update(ctx, b)
	return err
}

// RescheduleBooking изменяет время бронирования.
// Можно добавить проверку, что новое время не в прошлом и соответствует рабочему расписанию мастера.
func (bs *BookingService) RescheduleBooking(ctx context.Context, id uuid.UUID, newTime time.Time) (*entity.Booking, error) {
	if newTime.Before(time.Now()) {
		return nil, errors.New("new booking time cannot be in the past")
	}
	b, err := bs.bookingRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, errors.New("booking not found")
	}
	// Если бронирование уже подтверждено, возможно, перенести его нельзя или требуется другая логика.
	if b.Status == entity.StatusConfirmed {
		return nil, errors.New("cannot reschedule a confirmed booking")
	}
	b.BookingTime = newTime
	return bs.bookingRepo.Update(ctx, b)
}
