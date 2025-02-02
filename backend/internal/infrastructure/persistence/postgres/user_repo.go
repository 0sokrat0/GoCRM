// repository/postgres/user_repo.go
package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"

	"github.com/google/uuid"
)

type PGUserRepo struct {
	db *sql.DB
}

func NewPGUserRepo(db *sql.DB) repository.UserRepository {
	return &PGUserRepo{db: db}
}

func (r *PGUserRepo) Create(ctx context.Context, u *entity.User) error {
	query := `
		INSERT INTO users (
			id, role, telegram_id, username, 
			first_name, last_name, lang_code, 
			phone, session_hash, created_at, login_date
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.db.ExecContext(ctx, query,
		u.ID,
		u.Role,
		u.TelegramID,
		u.Username,
		u.FirstName,
		u.LastName,
		u.LanguageCode,
		u.Phone,
		u.SessionHash,
		u.CreatedAt,
		u.LoginDate,
	)
	return err
}

func (r *PGUserRepo) Update(ctx context.Context, u *entity.User) (*entity.User, error) {
	query := `
		UPDATE users SET
			username = $2,
			first_name = $3,
			last_name = $4,
			lang_code = $5,
			phone = $6,
			session_hash = $7,
			login_date = $8
		WHERE id = $1
		RETURNING *
	`

	updated := &entity.User{}
	err := r.db.QueryRowContext(ctx, query,
		u.ID,
		u.Username,
		u.FirstName,
		u.LastName,
		u.LanguageCode,
		u.Phone,
		u.SessionHash,
		u.LoginDate,
	).Scan(
		&updated.ID,
		&updated.Role,
		&updated.TelegramID,
		&updated.Username,
		&updated.FirstName,
		&updated.LastName,
		&updated.LanguageCode,
		&updated.Phone,
		&updated.SessionHash,
		&updated.CreatedAt,
		&updated.LoginDate,
	)

	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return updated, nil
}

func (r *PGUserRepo) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	query := `
		SELECT * FROM users 
		WHERE id = $1
	`

	u := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&u.ID,
		&u.Role,
		&u.TelegramID,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.LanguageCode,
		&u.Phone,
		&u.SessionHash,
		&u.CreatedAt,
		&u.LoginDate,
	)

	if err != nil {
		return nil, fmt.Errorf("error getting user by ID: %w", err)
	}

	return u, nil
}

func (r *PGUserRepo) GetByTelegramID(ctx context.Context, tgID int64) (*entity.User, error) {
	query := `
		SELECT * FROM users 
		WHERE telegram_id = $1
	`

	u := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, tgID).Scan(
		&u.ID,
		&u.Role,
		&u.TelegramID,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.LanguageCode,
		&u.Phone,
		&u.SessionHash,
		&u.CreatedAt,
		&u.LoginDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting user by Telegram ID: %w", err)
	}

	return u, nil
}
