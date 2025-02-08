package entity

import (
	"time"

	"github.com/google/uuid"
)

type Master struct {
	MasterID    uuid.UUID   `json:"master_id"`
	Username    string      `json:"username"`
	TelegramID  int64       `json:"telegram_id"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Phone       string      `json:"phone"`
	ServiceIDs  []uuid.UUID `json:"service_ids"` // Только ID услуг
	Description string      `json:"description"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

func NewMaster(username, firstName, lastName, phone string, services []uuid.UUID, description string, telegramID int64) *Master {
	return &Master{
		MasterID:    uuid.New(),
		Username:    username,
		TelegramID:  telegramID,
		FirstName:   firstName,
		LastName:    lastName,
		Phone:       phone,
		ServiceIDs:  services,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (m *Master) UpdateProfile(username, firstName, lastName, phone, description string) {
	m.Username = username
	m.FirstName = firstName
	m.LastName = lastName
	m.Phone = phone
	m.Description = description
	m.UpdatedAt = time.Now()
}
