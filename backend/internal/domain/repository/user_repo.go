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
	GetByTelegramID(ctx context.Context, tgID int64) (*entity.User, error)

	//List(ctx context.Context, filter map[string]interface{}) ([]*entity.User, error)
}
