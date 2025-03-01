package repositories

import (
	"gorm.io/gorm"
)

type PGBookingRepo struct {
	db *gorm.DB
}

// func NewPGBookingRepo(db *gorm.DB) repo.BookingRepository {
// 	return &PGBookingRepo{db: db}
// }
