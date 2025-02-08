package entity

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ServiceID   uuid.UUID // Уникальный идентификатор
	Name        string    // Название сервиса
	Description string    // Описание
	Price       float64   // Цена
	Duration    int       // Продолжительность (в минутах)
	CreatedAt   time.Time // Дата создания
}

func NewService(name, description string, price float64, duration int) *Service {
	return &Service{
		ServiceID:   uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
		Duration:    duration,
		CreatedAt:   time.Now(),
	}
}
