package repository

import (
	"context"

	"GoCRM/internal/domain/entity"

	"github.com/google/uuid"
)

// AuditLogRepository определяет операции для записи и получения аудиторских логов.
type AuditLogRepository interface {
	// Create записывает новое событие в журнал аудита.
	Create(ctx context.Context, a *entity.AuditLog) error

	// GetByID возвращает аудиторскую запись по её идентификатору.
	GetByID(ctx context.Context, id uuid.UUID) (*entity.AuditLog, error)

	// List возвращает список аудиторских записей с возможными фильтрами (например, по userID или действию).
	List(ctx context.Context, filter map[string]interface{}) ([]*entity.AuditLog, error)

	// Delete удаляет запись аудита (если необходимо, обычно логи хранятся на протяжении определенного времени).
	Delete(ctx context.Context, id uuid.UUID) error
}
