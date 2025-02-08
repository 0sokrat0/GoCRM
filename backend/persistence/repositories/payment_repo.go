package repositories

import (
	"context"
	"database/sql"
	"errors"

	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"

	"github.com/google/uuid"
)

type PGPaymentRepo struct {
	db *sql.DB
}

func NewPGPaymentRepo(db *sql.DB) repository.PaymentRepository {
	return &PGPaymentRepo{db: db}
}

func (r *PGPaymentRepo) Create(ctx context.Context, p *entity.Payment) error {
	query := `
		INSERT INTO payments (payment_id, booking_id, amount, payment_method, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	if p.PaymentID == uuid.Nil {
		p.PaymentID = uuid.New()
	}
	_, err := r.db.ExecContext(ctx, query, p.PaymentID, p.BookingID, p.Amount, p.PaymentMethod, p.Status, p.CreatedAt)
	return err
}

func (r *PGPaymentRepo) GetByID(ctx context.Context, id uuid.UUID) (*entity.Payment, error) {
	query := `
		SELECT payment_id, booking_id, amount, payment_method, status, created_at
		FROM payments
		WHERE payment_id = $1
	`
	var p entity.Payment
	err := r.db.QueryRowContext(ctx, query, id).Scan(&p.PaymentID, &p.BookingID, &p.Amount, &p.PaymentMethod, &p.Status, &p.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}
	return &p, nil
}

func (r *PGPaymentRepo) List(ctx context.Context, filter map[string]interface{}) ([]*entity.Payment, error) {
	query := `
		SELECT payment_id, booking_id, amount, payment_method, status, created_at
		FROM payments
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*entity.Payment
	for rows.Next() {
		var p entity.Payment
		if err := rows.Scan(&p.PaymentID, &p.BookingID, &p.Amount, &p.PaymentMethod, &p.Status, &p.CreatedAt); err != nil {
			return nil, err
		}
		payments = append(payments, &p)
	}
	return payments, rows.Err()
}

func (r *PGPaymentRepo) Update(ctx context.Context, p *entity.Payment) (*entity.Payment, error) {
	query := `
		UPDATE payments
		SET booking_id = $2, amount = $3, payment_method = $4, status = $5
		WHERE payment_id = $1
		RETURNING payment_id, booking_id, amount, payment_method, status, created_at
	`
	var updated entity.Payment
	err := r.db.QueryRowContext(ctx, query, p.PaymentID, p.BookingID, p.Amount, p.PaymentMethod, p.Status).
		Scan(&updated.PaymentID, &updated.BookingID, &updated.Amount, &updated.PaymentMethod, &updated.Status, &updated.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *PGPaymentRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM payments
		WHERE payment_id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
