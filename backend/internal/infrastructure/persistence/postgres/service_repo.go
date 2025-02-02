package postgres

import (
	"context"
	"database/sql"
	"errors"

	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"

	"github.com/google/uuid"
)

type PGServiceRepo struct {
	db *sql.DB
}

func NewPGServiceRepo(db *sql.DB) repository.ServiceRepository {
	return &PGServiceRepo{db: db}
}

func (r *PGServiceRepo) Create(ctx context.Context, s *entity.Service) error {
	query := `
		INSERT INTO services (service_id, name, price, duration, created_at,description)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	if s.ServiceID == uuid.Nil {
		s.ServiceID = uuid.New()
	}
	_, err := r.db.ExecContext(ctx, query, s.ServiceID, s.Name, s.Price, s.Duration, s.CreatedAt, s.Description)
	return err
}

func (r *PGServiceRepo) GetByID(ctx context.Context, id uuid.UUID) (*entity.Service, error) {
	query := `
		SELECT service_id, name, price, duration, created_at
		FROM services
		WHERE service_id = $1
	`
	var s entity.Service
	err := r.db.QueryRowContext(ctx, query, id).Scan(&s.ServiceID, &s.Name, &s.Price, &s.Duration, &s.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("service not found")
		}
		return nil, err
	}
	return &s, nil
}

func (r *PGServiceRepo) List(ctx context.Context) ([]*entity.Service, error) {
	query := `
		SELECT service_id, name, price, duration, created_at
		FROM services
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []*entity.Service
	for rows.Next() {
		var s entity.Service
		if err := rows.Scan(&s.ServiceID, &s.Name, &s.Price, &s.Duration, &s.CreatedAt); err != nil {
			return nil, err
		}
		services = append(services, &s)
	}
	return services, rows.Err()
}

func (r *PGServiceRepo) Update(ctx context.Context, s *entity.Service) (*entity.Service, error) {
	query := `
		UPDATE services
		SET name = $2, price = $3, duration = $4
		WHERE service_id = $1
		RETURNING service_id, name, price, duration, created_at
	`
	var updated entity.Service
	err := r.db.QueryRowContext(ctx, query, s.ServiceID, s.Name, s.Price, s.Duration).
		Scan(&updated.ServiceID, &updated.Name, &updated.Price, &updated.Duration, &updated.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *PGServiceRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM services
		WHERE service_id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
