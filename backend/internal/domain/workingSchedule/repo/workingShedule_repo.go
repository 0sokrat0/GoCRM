package repo

import (
	"GoCRM/internal/domain/workingSchedule/entity"
	"time"

	"github.com/google/uuid"
)

type ScheduleRepository interface {
	Create(schedule *entity.WorkingSchedule) error
	GetByMasterIDAndDay(masterID uuid.UUID, day int) ([]*entity.WorkingSchedule, error)
	GetByMasterID(masterID uuid.UUID) ([]*entity.WorkingSchedule, error)
	IsSlotAvailable(masterID uuid.UUID, bookingTime time.Time) bool
	Update(schedule *entity.WorkingSchedule) error
	Delete(scheduleID uuid.UUID) error
}
