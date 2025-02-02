package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// AuditLog представляет запись аудиторского лога.
type AuditLog struct {
	LogID     uuid.UUID `json:"log_id"`
	UserID    uuid.UUID `json:"user_id"`    // Кто выполнил действие; может быть пустым, если действие системное
	Action    string    `json:"action"`     // Описание действия (например, "user_created")
	Details   string    `json:"details"`    // Дополнительные сведения об изменениях
	CreatedAt time.Time `json:"created_at"` // Время создания записи
}

// NewAuditLog – фабричный метод для создания новой записи аудита.
func NewAuditLog(userID uuid.UUID, action, details string) (*AuditLog, error) {
	if action == "" {
		return nil, errors.New("action cannot be empty")
	}
	return &AuditLog{
		LogID:     uuid.New(),
		UserID:    userID,
		Action:    action,
		Details:   details,
		CreatedAt: time.Now(),
	}, nil
}
