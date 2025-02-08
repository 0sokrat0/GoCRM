package repositories

import (
	"context"
	"errors"
	"fmt"

	"GoCRM/internal/domain/user/entity"
	"GoCRM/internal/domain/user/repo"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PGUserRepo struct {
	db *gorm.DB
}

// ✅ Конструктор репозитория
func NewPGUserRepo(db *gorm.DB) repo.UserRepository {
	return &PGUserRepo{db: db}
}

// 🔹 Создание пользователя
func (r *PGUserRepo) Create(ctx context.Context, u *entity.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

// 🔹 Обновление пользователя
func (r *PGUserRepo) Update(ctx context.Context, u *entity.User) (*entity.User, error) {
	err := r.db.WithContext(ctx).Model(&entity.User{}).
		Where("id = ?", u.ID).
		Updates(map[string]interface{}{
			"username":      u.Username,
			"first_name":    u.FirstName,
			"last_name":     u.LastName,
			"language_code": u.LanguageCode, // ✅ Исправленное название колонки
			"phone":         u.Phone,
			"session_hash":  u.SessionHash,
			"login_date":    u.LoginDate,
			"updated_at":    gorm.Expr("NOW()"), // ✅ Явное обновление `updated_at`
		}).Error

	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	// 🔹 Получаем обновленные данные
	return r.GetByID(ctx, u.ID)
}

// 🔹 Получение пользователя по ID
func (r *PGUserRepo) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var u entity.User
	err := r.db.WithContext(ctx).First(&u, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // ✅ Исправлено: теперь возвращает `nil, nil`
	}
	if err != nil {
		return nil, fmt.Errorf("error getting user by ID: %w", err)
	}

	return &u, nil
}

// 🔹 Получение пользователя по Telegram ID
func (r *PGUserRepo) GetByTelegramID(ctx context.Context, tgID int64) (*entity.User, error) {
	var u entity.User
	err := r.db.WithContext(ctx).Where("telegram_id = ?", tgID).First(&u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // ✅ Исправлено: теперь возвращает `nil, nil`
	}
	if err != nil {
		return nil, fmt.Errorf("error getting user by Telegram ID: %w", err)
	}

	return &u, nil
}

func (r *PGUserRepo) GetByPhone(ctx context.Context, phone string) (*entity.User, error) {
	var u entity.User
	err := r.db.WithContext(ctx).Where("phone = ?", phone).First(&u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error getting user by phone: %w", err)
	}

	return &u, nil
}
