package repository

import (
	"context"

	"GoCRM/internal/domain/entity"

	"github.com/google/uuid"
)

type PaymentRepository interface {
	// Create создает новый платеж.
	Create(ctx context.Context, p *entity.Payment) error

	// GetByID возвращает платеж по его уникальному идентификатору.
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Payment, error)

	// List возвращает список платежей с возможными фильтрами.
	List(ctx context.Context, filter map[string]interface{}) ([]*entity.Payment, error)

	// Update обновляет данные платежа.
	Update(ctx context.Context, p *entity.Payment) (*entity.Payment, error)

	// Delete удаляет платеж по его уникальному идентификатору.
	Delete(ctx context.Context, id uuid.UUID) error
}
