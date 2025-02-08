package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Role         string         `gorm:"type:varchar(20);not null;default:'client'"`
	Level        string         `gorm:"type:varchar(20);not null;default:'new'"`
	TelegramID   int64          `gorm:"uniqueIndex;not null"`
	Username     string         `gorm:"type:varchar(100);uniqueIndex"`
	ClientName   string         `gorm:"type:varchar(100)"`
	FirstName    string         `gorm:"type:varchar(100)"`
	LastName     string         `gorm:"type:varchar(100)"`
	LanguageCode string         `gorm:"type:varchar(10)"`
	Phone        string         `gorm:"type:varchar(20)"`
	IsVerified   bool           `gorm:"not null;default:false"`
	IsBot        bool           `gorm:"not null;default:false"`
	SessionHash  string         `gorm:"type:varchar(255)"`
	ReferrerID   *uuid.UUID     `gorm:"type:uuid"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	LoginDate    time.Time      `gorm:"not null"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
