package repository

import (
	"GoCRM/internal/domain/entity"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, u *entity.User) error
	// GetUser(ctx context.Context, u *entity.User) (*entity.User, error)
	Update(ctx context.Context, u *entity.User) (*entity.User, error)
	Delete(ctx context.Context, u *entity.User) error

	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	// GetByTgID(ctx context.Context, tgID *entity.User) (*entity.User, error)
}
