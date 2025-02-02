// repository/user_repository.go
package repository

import (
	"GoCRM/internal/domain/entity"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, u *entity.User) error
	Update(ctx context.Context, u *entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetByTelegramID(ctx context.Context, tgID int64) (*entity.User, error)
}
