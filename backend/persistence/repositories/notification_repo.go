package repositories

import (
	"context"
	"database/sql"
	"errors"

	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"

	"github.com/google/uuid"
)

type PGNotificationRepo struct {
	db *sql.DB
}

func NewPGNotificationRepo(db *sql.DB) repository.NotificationRepository {
	return &PGNotificationRepo{db: db}
}

func (r *PGNotificationRepo) Create(ctx context.Context, n *entity.Notification) error {
	query := `
		INSERT INTO notifications (notification_id, user_id, type, message, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	if n.NotificationID == uuid.Nil {
		n.NotificationID = uuid.New()
	}
	_, err := r.db.ExecContext(ctx, query, n.NotificationID, n.UserID, n.Type, n.Message, n.Status, n.CreatedAt)
	return err
}

func (r *PGNotificationRepo) GetByID(ctx context.Context, id uuid.UUID) (*entity.Notification, error) {
	query := `
		SELECT notification_id, user_id, type, message, status, created_at
		FROM notifications
		WHERE notification_id = $1
	`
	var n entity.Notification
	err := r.db.QueryRowContext(ctx, query, id).Scan(&n.NotificationID, &n.UserID, &n.Type, &n.Message, &n.Status, &n.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("notification not found")
		}
		return nil, err
	}
	return &n, nil
}

func (r *PGNotificationRepo) List(ctx context.Context, filter map[string]interface{}) ([]*entity.Notification, error) {
	query := `
		SELECT notification_id, user_id, type, message, status, created_at
		FROM notifications
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*entity.Notification
	for rows.Next() {
		var n entity.Notification
		if err := rows.Scan(&n.NotificationID, &n.UserID, &n.Type, &n.Message, &n.Status, &n.CreatedAt); err != nil {
			return nil, err
		}
		notifications = append(notifications, &n)
	}
	return notifications, rows.Err()
}

func (r *PGNotificationRepo) Update(ctx context.Context, n *entity.Notification) (*entity.Notification, error) {
	query := `
		UPDATE notifications
		SET user_id = $2, type = $3, message = $4, status = $5
		WHERE notification_id = $1
		RETURNING notification_id, user_id, type, message, status, created_at
	`
	var updated entity.Notification
	err := r.db.QueryRowContext(ctx, query, n.NotificationID, n.UserID, n.Type, n.Message, n.Status).
		Scan(&updated.NotificationID, &updated.UserID, &updated.Type, &updated.Message, &updated.Status, &updated.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *PGNotificationRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM notifications
		WHERE notification_id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
