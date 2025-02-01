package postgres

import (
	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"
	"context"
	"database/sql"

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
	INSERT INTO users(id, name, email, role, created_at)
	VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(ctx, query, u.ID, u.Name, u.Email, u.Role, u.CreatedAt)

	return err
}

func (r *PGUserRepo) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {

	query := `
		SELECT id, name, email, role, created_at
		FROM users
		WHERE id = $1
	`

	var u entity.User

	if err := r.db.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.CreatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *PGUserRepo) Update(ctx context.Context, u *entity.User) (*entity.User, error) {
	query := `
	UPDATE users 
	set name = $2, email = $3, phone = $4 ,role = $5
	WHERE id = $1
	RETURNING id, name, email, phone ,role, created_at
	`

	var updated entity.User

	err := r.db.QueryRowContext(ctx, query, u.ID, u.Name, u.Email, u.Phone, u.Role).Scan(&updated.ID, &updated.Name, &updated.Email, &updated.Phone, &updated.Role, &updated.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *PGUserRepo) Delete(ctx context.Context, u *entity.User) error {
	query := `
	DELETE FROM users
	WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, u.ID)
	return err
}
