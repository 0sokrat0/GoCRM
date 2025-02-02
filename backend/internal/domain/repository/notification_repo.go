package repository

import (
	"context"

	"GoCRM/internal/domain/entity"

	"github.com/google/uuid"
)

// NotificationRepository определяет операции для работы с уведомлениями.
type NotificationRepository interface {
	// Create создает новое уведомление.
	Create(ctx context.Context, n *entity.Notification) error

	// GetByID возвращает уведомление по его идентификатору.
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Notification, error)

	// List возвращает уведомления с возможными фильтрами (например, по userID).
	List(ctx context.Context, filter map[string]interface{}) ([]*entity.Notification, error)

	// Update обновляет уведомление (например, для изменения статуса на "read").
	Update(ctx context.Context, n *entity.Notification) (*entity.Notification, error)

	// Delete удаляет уведомление по его идентификатору.
	Delete(ctx context.Context, id uuid.UUID) error
}
