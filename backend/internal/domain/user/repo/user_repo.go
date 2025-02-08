package repo

import (
	"GoCRM/internal/domain/user/entity"
	"context"
	"errors"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("record not found")

type UserRepository interface {
	Create(ctx context.Context, u *entity.User) error
	Update(ctx context.Context, u *entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetByTelegramID(ctx context.Context, tgID int64) (*entity.User, error)
	GetByPhone(ctx context.Context, phone string) (*entity.User, error)
}
