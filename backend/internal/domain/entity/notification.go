package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// NotificationType задаёт допустимые типы уведомлений.
type NotificationType string

const (
	NotificationTypeEmail    NotificationType = "email"
	NotificationTypeSMS      NotificationType = "sms"
	NotificationTypeTelegram NotificationType = "telegram"
)

// NotificationStatus задаёт возможные статусы уведомления.
type NotificationStatus string

const (
	NotificationStatusPending NotificationStatus = "pending"
	NotificationStatusSent    NotificationStatus = "sent"
	NotificationStatusFailed  NotificationStatus = "failed"
)

// Notification представляет агрегат уведомления.
type Notification struct {
	NotificationID uuid.UUID          `json:"notification_id"`
	UserID         uuid.UUID          `json:"user_id"` // Получатель уведомления
	Type           NotificationType   `json:"type"`
	Message        string             `json:"message"`
	Status         NotificationStatus `json:"status"`
	CreatedAt      time.Time          `json:"created_at"`
}

// NewNotification – фабричный метод для создания нового уведомления с базовой валидацией.
func NewNotification(userID uuid.UUID, nType NotificationType, message string) (*Notification, error) {
	if message == "" {
		return nil, errors.New("notification message cannot be empty")
	}
	if userID == uuid.Nil {
		return nil, errors.New("userID is required for notification")
	}
	// Можно добавить дополнительные проверки: допустимость типа уведомления и т.д.
	return &Notification{
		NotificationID: uuid.New(),
		UserID:         userID,
		Type:           nType,
		Message:        message,
		Status:         NotificationStatusPending,
		CreatedAt:      time.Now(),
	}, nil
}
