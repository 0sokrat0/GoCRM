package repositories

import (
	"context"
	"database/sql"
	"errors"

	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"

	"github.com/google/uuid"
)

type PGBookingRepo struct {
	db *sql.DB
}

func NewPGBookingRepo(db *sql.DB) repository.BookingRepository {
	return &PGBookingRepo{db: db}
}

func (r *PGBookingRepo) Create(ctx context.Context, b *entity.Booking) error {
	query := `
		INSERT INTO bookings (booking_id, user_id, master_id, service_id, booking_time, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	if b.BookingID == uuid.Nil {
		b.BookingID = uuid.New()
	}
	_, err := r.db.ExecContext(ctx, query, b.BookingID, b.UserID, b.MasterID, b.ServiceID, b.BookingTime, b.Status, b.CreatedAt)
	return err
}

func (r *PGBookingRepo) GetByID(ctx context.Context, id uuid.UUID) (*entity.Booking, error) {
	query := `
		SELECT booking_id, user_id, master_id, service_id, booking_time, status, created_at
		FROM bookings
		WHERE booking_id = $1
	`
	var b entity.Booking
	err := r.db.QueryRowContext(ctx, query, id).Scan(&b.BookingID, &b.UserID, &b.MasterID, &b.ServiceID, &b.BookingTime, &b.Status, &b.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("booking not found")
		}
		return nil, err
	}
	return &b, nil
}

func (r *PGBookingRepo) List(ctx context.Context, filter map[string]interface{}) ([]*entity.Booking, error) {
	// Простой пример: возвращаем все бронирования, без динамической фильтрации
	query := `
		SELECT booking_id, user_id, master_id, service_id, booking_time, status, created_at
		FROM bookings
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*entity.Booking
	for rows.Next() {
		var b entity.Booking
		if err := rows.Scan(&b.BookingID, &b.UserID, &b.MasterID, &b.ServiceID, &b.BookingTime, &b.Status, &b.CreatedAt); err != nil {
			return nil, err
		}
		bookings = append(bookings, &b)
	}
	return bookings, rows.Err()
}

func (r *PGBookingRepo) Update(ctx context.Context, b *entity.Booking) (*entity.Booking, error) {
	query := `
		UPDATE bookings
		SET user_id = $2, master_id = $3, service_id = $4, booking_time = $5, status = $6
		WHERE booking_id = $1
		RETURNING booking_id, user_id, master_id, service_id, booking_time, status, created_at
	`
	var updated entity.Booking
	err := r.db.QueryRowContext(ctx, query, b.BookingID, b.UserID, b.MasterID, b.ServiceID, b.BookingTime, b.Status).
		Scan(&updated.BookingID, &updated.UserID, &updated.MasterID, &updated.ServiceID, &updated.BookingTime, &updated.Status, &updated.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *PGBookingRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM bookings
		WHERE booking_id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
