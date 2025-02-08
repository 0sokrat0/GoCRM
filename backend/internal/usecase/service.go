package usecase

import (
	"GoCRM/internal/domain/service/entity"
	"GoCRM/internal/domain/service/repo"
	"context"
	"errors"

	"github.com/google/uuid"
)

type Service struct {
	repo repo.ServiceRepository
}

func NewService(repo repo.ServiceRepository) *Service {
	return &Service{repo: repo}
}

func (uc *Service) CreateService(ctx context.Context, name, description string, price float64, duration int) (*entity.Service, error) {
	if name == "" {
		return nil, errors.New("service name is required")
	}
	if price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}
	if duration <= 0 {
		return nil, errors.New("duration must be greater than zero")
	}

	service := entity.NewService(name, description, price, duration)

	err := uc.repo.Create(ctx, service)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (uc *Service) UpdateService(ctx context.Context, id uuid.UUID, name, description string, price float64, duration int) (*entity.Service, error) {
	service, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if service == nil {
		return nil, errors.New("service not found")
	}

	if name == "" {
		return nil, errors.New("service name cannot be empty")
	}
	if price <= 0 {
		return nil, errors.New("new price must be greater than zero")
	}
	if duration <= 0 {
		return nil, errors.New("new duration must be greater than zero")
	}

	// Обновляем поля
	service.Name = name
	service.Description = description
	service.Price = price
	service.Duration = duration

	updatedService, err := uc.repo.Update(ctx, service)
	if err != nil {
		return nil, err
	}
	return updatedService, nil
}

func (uc *Service) GetByID(ctx context.Context, id uuid.UUID) (*entity.Service, error) {
	service, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if service == nil {
		return nil, errors.New("service not found")
	}
	return service, nil
}

func (uc *Service) DeleteByID(ctx context.Context, id uuid.UUID) error {
	service, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if service == nil {
		return errors.New("service not found")
	}

	return uc.repo.Delete(ctx, id)
}

func (uc *Service) ListServices(ctx context.Context) ([]*entity.Service, error) {
	return uc.repo.List(ctx)
}
