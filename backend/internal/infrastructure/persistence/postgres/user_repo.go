package postgres

import (
	"context"
	"database/sql"

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
		INSERT INTO users (user_id, name, email, role, phone, telegram_id, password, created_at, login_date)
		VALUES ($1, $2, $3, COALESCE(NULLIF($4, ''), 'client'), $5, $6, $7, $8, $9)
	`

	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	_, err := r.db.ExecContext(ctx, query,
		u.ID, u.Name, u.Email, u.Role, u.Phone, u.TelegramID, u.Password, u.CreatedAt, u.LoginDate,
	)
	return err
}

func (r *PGUserRepo) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	query := `
		SELECT user_id, name, email, role, phone, telegram_id, password, created_at, login_date
		FROM users
		WHERE user_id = $1
	`
	var u entity.User
	if err := r.db.QueryRowContext(ctx, query, id).Scan(
		&u.ID, &u.Name, &u.Email, &u.Role, &u.Phone, &u.TelegramID, &u.Password, &u.CreatedAt, &u.LoginDate,
	); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *PGUserRepo) Update(ctx context.Context, u *entity.User) (*entity.User, error) {
	query := `
		UPDATE users
		SET name = $2, email = $3, phone = $4, telegram_id = $5, role = $6, password = $7, login_date = $8
		WHERE user_id = $1
		RETURNING user_id, name, email, phone, telegram_id, role, password, created_at, login_date
	`
	var updated entity.User
	if err := r.db.QueryRowContext(ctx, query,
		u.ID, u.Name, u.Email, u.Phone, u.TelegramID, u.Role, u.Password, u.LoginDate,
	).Scan(&updated.ID, &updated.Name, &updated.Email, &updated.Phone, &updated.TelegramID, &updated.Role, &updated.Password, &updated.CreatedAt, &updated.LoginDate); err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *PGUserRepo) Delete(ctx context.Context, u *entity.User) error {
	query := `
		DELETE FROM users
		WHERE user_id = $1
	`
	_, err := r.db.ExecContext(ctx, query, u.ID)
	return err
}

func (r *PGUserRepo) GetByTelegramID(ctx context.Context, tgID int64) (*entity.User, error) {
	query := `
        SELECT user_id, name, email, role, phone, telegram_id, password, created_at, login_date
        FROM users
        WHERE telegram_id = $1
    `
	var u entity.User
	if err := r.db.QueryRowContext(ctx, query, tgID).Scan(
		&u.ID, &u.Name, &u.Email, &u.Role, &u.Phone, &u.TelegramID, &u.Password, &u.CreatedAt, &u.LoginDate,
	); err != nil {
		return nil, err
	}
	return &u, nil
}
