package app

import (
	"context"
	"errors"

	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"

	"github.com/google/uuid"
)

// NotificationService инкапсулирует бизнес-логику для уведомлений.
type NotificationService struct {
	notifRepo repository.NotificationRepository
}

// NewNotificationService создаёт новый экземпляр NotificationService.
func NewNotificationService(repo repository.NotificationRepository) *NotificationService {
	return &NotificationService{
		notifRepo: repo,
	}
}

// CreateNotification сохраняет новое уведомление.
func (ns *NotificationService) CreateNotification(ctx context.Context, n *entity.Notification) error {
	if n == nil {
		return errors.New("notification cannot be nil")
	}
	// Дополнительные проверки, например, валидность типа уведомления, можно добавить здесь.
	return ns.notifRepo.Create(ctx, n)
}

// GetNotification возвращает уведомление по ID.
func (ns *NotificationService) GetNotification(ctx context.Context, id uuid.UUID) (*entity.Notification, error) {
	n, err := ns.notifRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if n == nil {
		return nil, errors.New("notification not found")
	}
	return n, nil
}

// UpdateNotification обновляет уведомление.
func (ns *NotificationService) UpdateNotification(ctx context.Context, n *entity.Notification) (*entity.Notification, error) {
	if n == nil {
		return nil, errors.New("notification cannot be nil")
	}
	return ns.notifRepo.Update(ctx, n)
}

// DeleteNotification удаляет уведомление по ID.
func (ns *NotificationService) DeleteNotification(ctx context.Context, id uuid.UUID) error {
	return ns.notifRepo.Delete(ctx, id)
}

// ListNotifications возвращает уведомления с опциональными фильтрами.
func (ns *NotificationService) ListNotifications(ctx context.Context, filter map[string]interface{}) ([]*entity.Notification, error) {
	return ns.notifRepo.List(ctx, filter)
}
