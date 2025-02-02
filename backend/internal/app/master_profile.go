package app

import (
	"context"
	"errors"

	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"

	"github.com/google/uuid"
)

// MasterProfileService инкапсулирует бизнес-логику для агрегата MasterProfile.
type MasterProfileService struct {
	masterRepo repository.MasterProfileRepository
}

// NewMasterProfileService создает новый экземпляр MasterProfileService.
func NewMasterProfileService(repo repository.MasterProfileRepository) *MasterProfileService {
	return &MasterProfileService{
		masterRepo: repo,
	}
}

// CreateMasterProfile создает новый профиль мастера.
func (s *MasterProfileService) CreateMasterProfile(ctx context.Context, mp *entity.MasterProfile) error {
	if mp == nil {
		return errors.New("master profile cannot be nil")
	}
	// Если MasterID не задан, генерируем его.
	if mp.MasterID == uuid.Nil {
		mp.MasterID = uuid.New()
	}
	// Дополнительные проверки и бизнес-логика по расписанию можно добавить здесь.
	return s.masterRepo.Create(ctx, mp)
}

// GetMasterProfile возвращает профиль мастера по его идентификатору.
func (s *MasterProfileService) GetMasterProfile(ctx context.Context, id uuid.UUID) (*entity.MasterProfile, error) {
	mp, err := s.masterRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if mp == nil {
		return nil, errors.New("master profile not found")
	}
	return mp, nil
}

// UpdateMasterProfile обновляет существующий профиль мастера.
func (s *MasterProfileService) UpdateMasterProfile(ctx context.Context, mp *entity.MasterProfile) (*entity.MasterProfile, error) {
	if mp == nil {
		return nil, errors.New("master profile cannot be nil")
	}
	return s.masterRepo.Update(ctx, mp)
}

// DeleteMasterProfile удаляет профиль мастера по его идентификатору.
func (s *MasterProfileService) DeleteMasterProfile(ctx context.Context, id uuid.UUID) error {
	return s.masterRepo.Delete(ctx, id)
}
