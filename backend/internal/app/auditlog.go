package app

import (
	"context"
	"errors"

	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"

	"github.com/google/uuid"
)

// AuditLogService инкапсулирует бизнес-логику для аудиторских записей.
type AuditLogService struct {
	auditRepo repository.AuditLogRepository
}

// NewAuditLogService создаёт новый экземпляр AuditLogService.
func NewAuditLogService(repo repository.AuditLogRepository) *AuditLogService {
	return &AuditLogService{
		auditRepo: repo,
	}
}

// CreateAuditLog создаёт новую запись аудита.
func (als *AuditLogService) CreateAuditLog(ctx context.Context, a *entity.AuditLog) error {
	if a == nil {
		return errors.New("audit log cannot be nil")
	}
	return als.auditRepo.Create(ctx, a)
}

// GetAuditLog возвращает аудиторскую запись по ID.
func (als *AuditLogService) GetAuditLog(ctx context.Context, id uuid.UUID) (*entity.AuditLog, error) {
	a, err := als.auditRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, errors.New("audit log not found")
	}
	return a, nil
}

// ListAuditLogs возвращает список аудиторских записей с возможными фильтрами.
func (als *AuditLogService) ListAuditLogs(ctx context.Context, filter map[string]interface{}) ([]*entity.AuditLog, error) {
	return als.auditRepo.List(ctx, filter)
}

// DeleteAuditLog удаляет аудиторскую запись по ID.
func (als *AuditLogService) DeleteAuditLog(ctx context.Context, id uuid.UUID) error {
	return als.auditRepo.Delete(ctx, id)
}
