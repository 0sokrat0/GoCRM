package models

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ServiceID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `gorm:"type:varchar(20);not null;uniqueIndex"`
	Description string    `gorm:"type:varchar(100)"`
	Price       float64   `gorm:"not null"`
	Duration    int       `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
