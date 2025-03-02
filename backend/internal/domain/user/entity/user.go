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

type UserLevel string

const (
	LevelNew      UserLevel = "new"       // 0-4 записей
	LevelRegular  UserLevel = "regular"   // 5-14 записей
	LevelLoyal    UserLevel = "loyal"     // 15+ записей
	LevelVIP      UserLevel = "vip"       // 30+ записей
	LevelSuperVIP UserLevel = "super_vip" // 50+ записей
)

type User struct {
	ID           uuid.UUID // Уникальный идентификатор
	Role         Role      // Роль пользователя (client, master, admin)
	Level        UserLevel // Уровень лояльности клиента
	TelegramID   int64     // ID в Telegram
	Username     string    // Telegram username
	ClientName   string    // Имя клиента реальное (не тг username)
	FirstName    string    // Имя
	LastName     string    // Фамилия
	LanguageCode string    // Язык
	Phone        string    // Номер телефона
	// DateOfBirth  time.Time  // Дата рождения
	IsVerified  bool       // Подтверждённый пользователь
	IsBot       bool       // Является ли ботом
	SessionHash string     // Хеш сессии для авторизации
	ReferrerID  *uuid.UUID // ID пригласившего пользователя
	CreatedAt   time.Time  // Дата регистрации
	UpdatedAt   time.Time  // Последнее обновление профиля
	LoginDate   time.Time  // Дата последнего входа
}

func NewTelegramUser(
	tgID int64,
	username, clientName, firstName, lastName, languageCode, phone string,
	isBot bool,
) (*User, error) {
	if tgID == 0 {
		return nil, errors.New("telegram ID is required")
	}

	return &User{
		ID:           uuid.New(),
		Role:         RoleClient,
		Level:        LevelNew,
		TelegramID:   tgID,
		Username:     username,
		ClientName:   "",
		FirstName:    firstName,
		LastName:     lastName,
		LanguageCode: languageCode,
		Phone:        phone,
		IsVerified:   false,
		IsBot:        isBot,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		LoginDate:    time.Now(),
	}, nil
}

func (u *User) UpdateLevel(bookingCount int) {
	switch {
	case bookingCount >= 50:
		u.Level = LevelSuperVIP
	case bookingCount >= 30:
		u.Level = LevelVIP
	case bookingCount >= 15:
		u.Level = LevelLoyal
	case bookingCount >= 5:
		u.Level = LevelRegular
	default:
		u.Level = LevelNew
	}
}

func (u *User) UpdateFromTelegram(username, firstName, lastName, languageCode, phone string) {
	u.Username = username
	u.FirstName = firstName
	u.LastName = lastName
	u.LanguageCode = languageCode
	u.Phone = phone
}

func (u *User) UpdateSession(sessionHash string) {
	u.SessionHash = sessionHash
}
