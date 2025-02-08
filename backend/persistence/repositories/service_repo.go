package repositories

import (
	"context"
	"errors"
	"fmt"

	"GoCRM/internal/domain/service/entity"
	"GoCRM/internal/domain/service/repo"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PGServiceRepo struct {
	db *gorm.DB
}

func NewPGServiceRepo(db *gorm.DB) repo.ServiceRepository {
	return &PGServiceRepo{db: db}
}

func (r *PGServiceRepo) Create(ctx context.Context, s *entity.Service) error {
	return r.db.WithContext(ctx).Create(s).Error

}

func (r *PGServiceRepo) GetByID(ctx context.Context, id uuid.UUID) (*entity.Service, error) {
	var s entity.Service
	err := r.db.WithContext(ctx).First(&s, "service_id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &s, nil
}

func (r *PGServiceRepo) List(ctx context.Context) ([]*entity.Service, error) {
	var s []*entity.Service
	err := r.db.WithContext(ctx).Find(&s).Error
	if err != nil {
		return nil, fmt.Errorf("error retrieving services: %w", err)
	}
	return s, nil
}

func (r *PGServiceRepo) Update(ctx context.Context, s *entity.Service) (*entity.Service, error) {
	err := r.db.WithContext(ctx).Model(s).
		Where("service_id = ?", s.ServiceID).
		Updates(s).Error

	if err != nil {
		return nil, fmt.Errorf("error updating service: %w", err)
	}

	return r.GetByID(ctx, s.ServiceID)
}

func (r *PGServiceRepo) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.db.WithContext(ctx).Where("service_id = ?", id).Delete(&entity.Service{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return nil
	}
	return nil
}
