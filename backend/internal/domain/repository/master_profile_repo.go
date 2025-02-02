package repository

import (
	"context"

	"GoCRM/internal/domain/entity"

	"github.com/google/uuid"
)

type MasterProfileRepository interface {
	Create(ctx context.Context, mp *entity.MasterProfile) error
	GetByID(ctx context.Context, masterID uuid.UUID) (*entity.MasterProfile, error)
	Update(ctx context.Context, mp *entity.MasterProfile) (*entity.MasterProfile, error)
	Delete(ctx context.Context, masterID uuid.UUID) error
	List(ctx context.Context) ([]*entity.MasterProfile, error)
}
