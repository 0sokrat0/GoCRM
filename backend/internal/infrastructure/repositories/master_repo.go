package repositories

// import (
// 	"GoCRM/internal/domain/master/entity"
// 	"GoCRM/internal/domain/master/repo"
// 	"context"
// 	"errors"
// 	"fmt"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// type DBMasterRepo struct {
// 	db *gorm.DB
// }

// func NewDBMasterRepo(db *gorm.DB) repo.MasterRepository {
// 	return &DBMasterRepo{db: db}
// }

// func (r *DBMasterRepo) CreateMaster(ctx context.Context, m *entity.Master) error {
// 	return r.db.WithContext(ctx).Create(m).Error
// }

// func (r *DBMasterRepo) GetByID(ctx context.Context, masterID uuid.UUID) (*entity.Master, error) {
// 	var m entity.Master
// 	err := r.db.WithContext(ctx).First(&m, "master_id = ?", masterID).Error

// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, nil
// 	}
// 	if err != nil {
// 		return nil, fmt.Errorf("error getting master by ID: %w", err)
// 	}

// 	return &m, err
// }

// func (r *DBMasterRepo) GetByTelegramID(ctx context.Context, tgID int64) (*entity.Master, error) {
// 	var u entity.Master
// 	err := r.db.WithContext(ctx).Where("telegram_id = ?", tgID).First(&u).Error

// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, nil
// 	}
// 	if err != nil {
// 		return nil, fmt.Errorf("error getting master by Telegram ID: %w", err)
// 	}

// 	return &u, nil
// }

// func (r *DBMasterRepo) GetAll(ctx context.Context) ([]*entity.Master, error) {
// 	var m []*entity.Master

// 	err := r.db.WithContext(ctx).Find(&m).Error
// 	if err != nil {
// 		return nil, fmt.Errorf("error retrieving masters: %w", err)
// 	}

// 	return m, nil
// }
