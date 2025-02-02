package repository

import (
	"context"

	"GoCRM/internal/domain/entity"

	"github.com/google/uuid"
)

// BookingRepository определяет операции для работы с бронированиями.
type BookingRepository interface {
	// Create создает новое бронирование.
	Create(ctx context.Context, b *entity.Booking) error

	// GetByID возвращает бронирование по его уникальному идентификатору.
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Booking, error)

	// List возвращает список бронирований с возможными фильтрами.
	List(ctx context.Context, filter map[string]interface{}) ([]*entity.Booking, error)

	// Update обновляет данные бронирования и возвращает обновленный объект.
	Update(ctx context.Context, b *entity.Booking) (*entity.Booking, error)

	// Delete удаляет бронирование по его уникальному идентификатору.
	Delete(ctx context.Context, id uuid.UUID) error
}
