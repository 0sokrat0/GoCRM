package repositories

// import (
// 	"context"
// 	"database/sql"
// 	"encoding/json"
// 	"errors"

// 	"GoCRM/internal/domain/entity"
// 	"GoCRM/internal/domain/repository"

// 	"github.com/google/uuid"
// )

// // PGMasterProfileRepo — реализация репозитория для MasterProfile для PostgreSQL.
// type PGMasterProfileRepo struct {
// 	db *sql.DB
// }

// // NewPGMasterProfileRepo создает новый экземпляр репозитория для MasterProfile.
// func NewPGMasterProfileRepo(db *sql.DB) repository.MasterProfileRepository {
// 	return &PGMasterProfileRepo{db: db}
// }

// // Create сохраняет новый профиль мастера в таблице master_profiles.
// func (r *PGMasterProfileRepo) Create(ctx context.Context, mp *entity.MasterProfile) error {
// 	// Сериализуем расписание в JSON.
// 	scheduleJSON, err := json.Marshal(mp.Schedule)
// 	if err != nil {
// 		return err
// 	}

// 	// Если MasterID не задан, генерируем новый.
// 	if mp.MasterID == uuid.Nil {
// 		mp.MasterID = uuid.New()
// 	}

// 	query := `
// 		INSERT INTO master_profiles (master_id, schedule)
// 		VALUES ($1, $2)
// 	`
// 	_, err = r.db.ExecContext(ctx, query, mp.MasterID, scheduleJSON)
// 	return err
// }

// // GetByID возвращает профиль мастера по его идентификатору.
// func (r *PGMasterProfileRepo) GetByID(ctx context.Context, masterID uuid.UUID) (*entity.MasterProfile, error) {
// 	query := `
// 		SELECT master_id, schedule
// 		FROM master_profiles
// 		WHERE master_id = $1
// 	`
// 	var id uuid.UUID
// 	var scheduleJSON []byte
// 	err := r.db.QueryRowContext(ctx, query, masterID).Scan(&id, &scheduleJSON)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, errors.New("master profile not found")
// 		}
// 		return nil, err
// 	}

// 	var schedule []entity.WorkingSchedule
// 	if err := json.Unmarshal(scheduleJSON, &schedule); err != nil {
// 		return nil, err
// 	}

// 	return &entity.MasterProfile{
// 		MasterID: id,
// 		Schedule: schedule,
// 	}, nil
// }

// // Update обновляет профиль мастера (например, изменение расписания).
// func (r *PGMasterProfileRepo) Update(ctx context.Context, mp *entity.MasterProfile) (*entity.MasterProfile, error) {
// 	if mp == nil {
// 		return nil, errors.New("master profile cannot be nil")
// 	}

// 	// Сериализуем обновленное расписание.
// 	scheduleJSON, err := json.Marshal(mp.Schedule)
// 	if err != nil {
// 		return nil, err
// 	}

// 	query := `
// 		UPDATE master_profiles
// 		SET schedule = $2
// 		WHERE master_id = $1
// 		RETURNING master_id, schedule
// 	`
// 	var id uuid.UUID
// 	var updatedScheduleJSON []byte
// 	err = r.db.QueryRowContext(ctx, query, mp.MasterID, scheduleJSON).
// 		Scan(&id, &updatedScheduleJSON)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var updatedSchedule []entity.WorkingSchedule
// 	if err := json.Unmarshal(updatedScheduleJSON, &updatedSchedule); err != nil {
// 		return nil, err
// 	}

// 	return &entity.MasterProfile{
// 		MasterID: id,
// 		Schedule: updatedSchedule,
// 	}, nil
// }

// // Delete удаляет профиль мастера по его идентификатору.
// func (r *PGMasterProfileRepo) Delete(ctx context.Context, masterID uuid.UUID) error {
// 	query := `
// 		DELETE FROM master_profiles
// 		WHERE master_id = $1
// 	`
// 	_, err := r.db.ExecContext(ctx, query, masterID)
// 	return err
// }

// // List возвращает список всех профилей мастеров.
// func (r *PGMasterProfileRepo) List(ctx context.Context) ([]*entity.MasterProfile, error) {
// 	query := `
// 		SELECT master_id, schedule
// 		FROM master_profiles
// 		ORDER BY master_id
// 	`
// 	rows, err := r.db.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var profiles []*entity.MasterProfile
// 	for rows.Next() {
// 		var id uuid.UUID
// 		var scheduleJSON []byte
// 		if err := rows.Scan(&id, &scheduleJSON); err != nil {
// 			return nil, err
// 		}
// 		var schedule []entity.WorkingSchedule
// 		if err := json.Unmarshal(scheduleJSON, &schedule); err != nil {
// 			return nil, err
// 		}
// 		profiles = append(profiles, &entity.MasterProfile{
// 			MasterID: id,
// 			Schedule: schedule,
// 		})
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return profiles, nil
// }
