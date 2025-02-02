package app

import (
	"context"
	"errors"
	"time"

	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"

	"github.com/google/uuid"
)

// PaymentService инкапсулирует бизнес-логику для агрегата Payment.
type PaymentService struct {
	paymentRepo repository.PaymentRepository
}

// NewPaymentService создаёт новый экземпляр PaymentService.
func NewPaymentService(repo repository.PaymentRepository) *PaymentService {
	return &PaymentService{
		paymentRepo: repo,
	}
}

// CreatePayment создаёт новый платёж с базовой валидацией.
func (ps *PaymentService) CreatePayment(ctx context.Context, p *entity.Payment) error {
	if p == nil {
		return errors.New("payment cannot be nil")
	}
	if p.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	// Устанавливаем время создания, генерируем идентификатор и устанавливаем начальный статус.
	p.CreatedAt = time.Now()
	p.PaymentID = uuid.New()
	// Предположим, что PaymentPending определён в entity.PaymentStatus.
	p.Status = entity.PaymentPending
	return ps.paymentRepo.Create(ctx, p)
}

// GetPayment возвращает платёж по его идентификатору.
func (ps *PaymentService) GetPayment(ctx context.Context, id uuid.UUID) (*entity.Payment, error) {
	p, err := ps.paymentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, errors.New("payment not found")
	}
	return p, nil
}

// UpdatePayment обновляет данные платежа.
func (ps *PaymentService) UpdatePayment(ctx context.Context, p *entity.Payment) (*entity.Payment, error) {
	if p == nil {
		return nil, errors.New("payment cannot be nil")
	}
	return ps.paymentRepo.Update(ctx, p)
}

// DeletePayment удаляет платёж по его идентификатору.
func (ps *PaymentService) DeletePayment(ctx context.Context, id uuid.UUID) error {
	return ps.paymentRepo.Delete(ctx, id)
}
