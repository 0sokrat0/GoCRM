package usecase

import (
	"GoCRM/internal/domain/user/entity"
	"GoCRM/internal/domain/user/repo"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo repo.UserRepository
}

// ✅ Конструктор сервиса
func NewUserService(r repo.UserRepository) *UserService {
	return &UserService{userRepo: r}
}

// ✅ Получить или создать пользователя// ✅ Получить или создать пользователя
func (us *UserService) GetOrCreateTelegramUser(
	ctx context.Context,
	tgID int64,
	username string,
	clientName string,
	firstName string,
	lastName string,
	languageCode string,
	phone string,
) (*entity.User, error) {
	// 🔹 Проверяем, есть ли пользователь по Telegram ID
	user, err := us.userRepo.GetByTelegramID(ctx, tgID)
	if err != nil {
		return nil, err
	}

	if user != nil {
		// 🔹 Обновляем данные пользователя
		user.UpdateFromTelegram(username, firstName, lastName, languageCode, phone)
		_, err = us.userRepo.Update(ctx, user)
		return user, err
	}

	// 🔹 Проверяем, есть ли другой пользователь с таким номером
	if phone != "" {
		existingUser, err := us.userRepo.GetByPhone(ctx, phone)
		if err != nil {
			return nil, err
		}
		if existingUser != nil {
			return nil, fmt.Errorf("user with this phone number already exists")
		}
	}

	// 🔹 Создаем нового пользователя
	newUser, err := entity.NewTelegramUser(tgID, username, clientName, firstName, lastName, languageCode, phone, true)
	if err != nil {
		return nil, err
	}

	err = us.userRepo.Create(ctx, newUser)
	return newUser, err
}

// ✅ Получение пользователя по Telegram ID
func (us *UserService) GetUserByTelegramID(ctx context.Context, tgID int64) (*entity.User, error) {
	return us.userRepo.GetByTelegramID(ctx, tgID)
}

// ✅ Обновление сессии пользователя
func (us *UserService) UpdateUserSession(ctx context.Context, tgID int64, sessionHash string) error {
	user, err := us.userRepo.GetByTelegramID(ctx, tgID)
	if err != nil {
		return err
	}

	// 🔹 Обновляем сессию
	user.UpdateSession(sessionHash)

	// 🔹 Игнорируем возвращаемого пользователя и берем только ошибку
	_, err = us.userRepo.Update(ctx, user)
	return err
}

// ✅ Получение пользователя по ID
func (us *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return us.userRepo.GetByID(ctx, id)
}
