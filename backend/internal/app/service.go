package app

import (
	"context"
	"errors"
	"time"

	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"

	"github.com/google/uuid"
)

// ServiceService инкапсулирует бизнес-логику для агрегата Service.
type ServiceService struct {
	serviceRepo repository.ServiceRepository
}

// NewServiceService создаёт новый экземпляр ServiceService.
func NewServiceService(repo repository.ServiceRepository) *ServiceService {
	return &ServiceService{
		serviceRepo: repo,
	}
}

// CreateService создаёт новую услугу с базовой валидацией.
func (ss *ServiceService) CreateService(ctx context.Context, s *entity.Service) error {
	if s == nil {
		return errors.New("service cannot be nil")
	}
	if s.Name == "" {
		return errors.New("service name is required")
	}
	if s.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if s.Duration <= 0 {
		return errors.New("duration must be greater than zero")
	}
	s.CreatedAt = time.Now()
	return ss.serviceRepo.Create(ctx, s)
}

// GetService возвращает услугу по ID.
func (ss *ServiceService) GetService(ctx context.Context, id uuid.UUID) (*entity.Service, error) {
	s, err := ss.serviceRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, errors.New("service not found")
	}
	return s, nil
}

// UpdateService обновляет данные услуги.
func (ss *ServiceService) UpdateService(ctx context.Context, s *entity.Service) (*entity.Service, error) {
	if s == nil {
		return nil, errors.New("service cannot be nil")
	}
	return ss.serviceRepo.Update(ctx, s)
}

// DeleteService удаляет услугу по ID.
func (ss *ServiceService) DeleteService(ctx context.Context, id uuid.UUID) error {
	return ss.serviceRepo.Delete(ctx, id)
}

// ListServices возвращает список всех услуг.
func (ss *ServiceService) ListServices(ctx context.Context) ([]*entity.Service, error) {
	return ss.serviceRepo.List(ctx)
}
