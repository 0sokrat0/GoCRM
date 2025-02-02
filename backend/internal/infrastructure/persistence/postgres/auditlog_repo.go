package postgres

import (
	"context"
	"database/sql"
	"errors"

	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"

	"github.com/google/uuid"
)

type PGAuditLogRepo struct {
	db *sql.DB
}

func NewPGAuditLogRepo(db *sql.DB) repository.AuditLogRepository {
	return &PGAuditLogRepo{db: db}
}

func (r *PGAuditLogRepo) Create(ctx context.Context, a *entity.AuditLog) error {
	query := `
		INSERT INTO audit_logs (log_id, user_id, action, details, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	if a.LogID == uuid.Nil {
		a.LogID = uuid.New()
	}
	_, err := r.db.ExecContext(ctx, query, a.LogID, a.UserID, a.Action, a.Details, a.CreatedAt)
	return err
}

func (r *PGAuditLogRepo) GetByID(ctx context.Context, id uuid.UUID) (*entity.AuditLog, error) {
	query := `
		SELECT log_id, user_id, action, details, created_at
		FROM audit_logs
		WHERE log_id = $1
	`
	var a entity.AuditLog
	err := r.db.QueryRowContext(ctx, query, id).Scan(&a.LogID, &a.UserID, &a.Action, &a.Details, &a.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("audit log not found")
		}
		return nil, err
	}
	return &a, nil
}

func (r *PGAuditLogRepo) List(ctx context.Context, filter map[string]interface{}) ([]*entity.AuditLog, error) {
	query := `
		SELECT log_id, user_id, action, details, created_at
		FROM audit_logs
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []*entity.AuditLog
	for rows.Next() {
		var a entity.AuditLog
		if err := rows.Scan(&a.LogID, &a.UserID, &a.Action, &a.Details, &a.CreatedAt); err != nil {
			return nil, err
		}
		logs = append(logs, &a)
	}
	return logs, rows.Err()
}

func (r *PGAuditLogRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM audit_logs
		WHERE log_id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
