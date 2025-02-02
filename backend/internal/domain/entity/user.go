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
	ID         uuid.UUID `json:"id"`
	Role       Role      `json:"role"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	TelegramID int64     `json:"tgID"`
	Password   string    `json:"password"`

	LoginDate time.Time `json:"login_date"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(name, email, phone, password string, tgID int64, role Role) (*User, error) {
	if name == "" || email == "" || phone == "" || password == "" {
		return nil, errors.New("name, email and password must not be empty")
	}
	if tgID == 0 {
		return nil, errors.New("telegram ID is required")
	}

	return &User{
		ID:         uuid.New(),
		Role:       role,
		Name:       name,
		Phone:      phone,
		Email:      email,
		TelegramID: tgID,
		Password:   password,
		CreatedAt:  time.Now(),
	}, nil

}

func (u *User) ChangePassword(newPassword string) error {
	if newPassword == "" {
		return errors.New("new password cannot be empty")
	}

	u.Password = newPassword
	return nil
}

func (u *User) UpdateProfile(name, email, phone string) error {
	if name == "" || email == "" || phone == "" {
		return errors.New("profile fields cannot be empty")
	}
	u.Name = name
	u.Email = email
	u.Phone = phone
	return nil
}
