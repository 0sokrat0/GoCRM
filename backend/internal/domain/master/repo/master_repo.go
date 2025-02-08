package repo

import (
	"GoCRM/internal/domain/master/entity"

	"github.com/google/uuid"
)

type MasterRepository interface {
	Create(master *entity.Master) error
	GetByID(masterID uuid.UUID) (*entity.Master, error)
	GetByTelegramID(telegramID int64) (*entity.Master, error)
	GetAll() ([]*entity.Master, error)
	Update(master *entity.Master) error
	Delete(masterID uuid.UUID) error
}
