package app

import (
	"GoCRM/internal/domain/entity"
	"GoCRM/internal/domain/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{
		userRepo: r,
	}
}

func (us *UserService) CreateUser(u *entity.User) error {

	if u == nil {
		return errors.New("user cannot be nil")
	}

	if u.Name == "" {
		return errors.New("user name is required")
	}
	if u.Email == "" {
		return errors.New("user email is required")
	}

	err := us.userRepo.Create(context.Background(), u)
	if err != nil {
		return err
	}
	return nil

}

func (us *UserService) GetUser(id uuid.UUID) (*entity.User, error) {
	u, err := us.userRepo.GetByID(context.Background(), id)

	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func (us *UserService) UpdateUser(u *entity.User) (*entity.User, error) {
	if u == nil {
		return nil, errors.New("user cannot be nil")
	}

	updatedUser, err := us.userRepo.Update(context.Background(), u)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (us *UserService) DeleteUser(u *entity.User) error {
	if u == nil {
		return errors.New("user cannot be nil")
	}
	err := us.userRepo.Delete(context.Background(), u)
	if err != nil {
		return err
	}
	return nil
}
