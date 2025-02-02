package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ServiceID   uuid.UUID `json:"service_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Duration    int       `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewService(name, description string, price float64, duration int) (*Service, error) {
	if name == "" {
		return nil, errors.New("service name is required")
	}
	if price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}
	if duration <= 0 {
		return nil, errors.New("duration must be greater than zero")
	}

	return &Service{
		ServiceID:   uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
		Duration:    duration,
		CreatedAt:   time.Now(),
	}, nil
}

func (s *Service) ChangePrice(newPrice float64) error {
	if newPrice <= 0 {
		return errors.New("new price must be greater than zero")
	}
	s.Price = newPrice
	return nil
}

func (s *Service) ChangeDuration(newDuration int) error {
	if newDuration <= 0 {
		return errors.New("new duration must be greater than zero")
	}
	s.Duration = newDuration
	return nil
}
