// domain/entity/user.go
package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleClient Role = "client"
	RoleMaster Role = "master"
	RoleAdmin  Role = "admin"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Role         Role      `json:"role"`
	TelegramID   int64     `json:"tg_id"`
	Username     string    `json:"username"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	LanguageCode string    `json:"lang_code"`
	Phone        string    `json:"phone"`
	SessionHash  string    `json:"session_hash"`
	CreatedAt    time.Time `json:"created_at"`
	LoginDate    time.Time `json:"login_date"`
}

func NewTelegramUser(
	tgID int64,
	username string,
	firstName string,
	lastName string,
	languageCode string,
	phone string,
) (*User, error) {
	if tgID == 0 {
		return nil, errors.New("telegram ID is required")
	}

	return &User{
		ID:           uuid.New(),
		Role:         RoleClient,
		TelegramID:   tgID,
		Username:     username,
		FirstName:    firstName,
		LastName:     lastName,
		LanguageCode: languageCode,
		Phone:        phone,
		CreatedAt:    time.Now(),
		LoginDate:    time.Now(),
	}, nil
}

func (u *User) UpdateFromTelegram(
	username string,
	firstName string,
	lastName string,
	languageCode string,
	phone string,
) {
	u.Username = username
	u.FirstName = firstName
	u.LastName = lastName
	u.LanguageCode = languageCode
	u.Phone = phone
	u.LoginDate = time.Now()
}

func (u *User) SetSessionHash(hash string) {
	u.SessionHash = hash
}
