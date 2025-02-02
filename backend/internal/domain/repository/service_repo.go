package repository

import (
	"context"

	"GoCRM/internal/domain/entity"

	"github.com/google/uuid"
)

type ServiceRepository interface {
	Create(ctx context.Context, s *entity.Service) error

	GetByID(ctx context.Context, id uuid.UUID) (*entity.Service, error)

	List(ctx context.Context) ([]*entity.Service, error)

	Update(ctx context.Context, s *entity.Service) (*entity.Service, error)

	Delete(ctx context.Context, id uuid.UUID) error
}
