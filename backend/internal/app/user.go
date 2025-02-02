package app

import (
	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{userRepo: r}
}

func (us *UserService) GetOrCreateTelegramUser(
	ctx context.Context,
	tgID int64,
	username string,
	firstName string,
	lastName string,
	languageCode string,
	phone string,
) (*entity.User, error) {
	existingUser, err := us.userRepo.GetByTelegramID(ctx, tgID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if existingUser != nil {
		existingUser.UpdateFromTelegram(username, firstName, lastName, languageCode, phone)
		_, err = us.userRepo.Update(ctx, existingUser)
		return existingUser, err
	}

	newUser, err := entity.NewTelegramUser(tgID, username, firstName, lastName, languageCode, phone)
	if err != nil {
		return nil, err
	}

	if err := us.userRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (us *UserService) GetUserByTelegramID(ctx context.Context, tgID int64) (*entity.User, error) {
	return us.userRepo.GetByTelegramID(ctx, tgID)
}

func (us *UserService) UpdateUserSession(ctx context.Context, tgID int64, sessionHash string) error {
	user, err := us.userRepo.GetByTelegramID(ctx, tgID)
	if err != nil {
		return err
	}

	user.SetSessionHash(sessionHash)
	_, err = us.userRepo.Update(ctx, user)
	return err
}

func (us *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return us.userRepo.GetByID(ctx, id)
}
