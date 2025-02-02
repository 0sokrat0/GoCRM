package repository

import (
	"context"

	"GoCRM/internal/domain/entity"

	"github.com/google/uuid"
)

// ScheduleRepository определяет операции для управления рабочим расписанием мастеров.
type ScheduleRepository interface {
	// Create добавляет новое рабочее расписание для мастера.
	Create(ctx context.Context, ws *entity.WorkingSchedule) error

	// GetByID возвращает запись расписания по её идентификатору.
	GetByID(ctx context.Context, id uuid.UUID) (*entity.WorkingSchedule, error)

	// ListByMaster возвращает все записи расписания для конкретного мастера.
	ListByMaster(ctx context.Context, masterID uuid.UUID) ([]*entity.WorkingSchedule, error)

	// Update обновляет запись расписания.
	Update(ctx context.Context, ws *entity.WorkingSchedule) (*entity.WorkingSchedule, error)

	// Delete удаляет запись расписания по её идентификатору.
	Delete(ctx context.Context, id uuid.UUID) error
}
