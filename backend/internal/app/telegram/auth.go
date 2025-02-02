package telegram

import (
	"fmt"
	"os"
	"time"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

// UserData хранит данные пользователя, полученные из init-data.
type UserData struct {
	ID           int64
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	Phone        string
	Hash         string
}

// ValidateTelegramInitDataWithThirdParty использует пакет init-data-golang для валидации и парсинга initData.
// Если переменная окружения SKIP_VALIDATION установлена в "true", то валидация пропускается (только для разработки!).
func ValidateTelegramInitDataWithThirdParty(initData string) (*UserData, error) {
	// Получаем токен бота из переменной окружения.
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("telegram bot token is not set")
	}

	// Определяем допустимый интервал жизни данных.
	// Для продакшена рекомендуется использовать, например, 24 часа.
	expIn := 24 * time.Hour
	if os.Getenv("SKIP_AUTH_DATE_CHECK") == "true" {
		// Для разработки можно отключить проверку срока жизни
		expIn = 0
	}

	// Если режим пропуска валидации включен, сразу распарсим данные без вызова Validate.
	if os.Getenv("SKIP_VALIDATION") == "true" {
		parsed, err := initdata.Parse(initData)
		if err != nil {
			return nil, fmt.Errorf("failed to parse init data: %w", err)
		}
		return &UserData{
			ID:           parsed.User.ID,
			FirstName:    parsed.User.FirstName,
			LastName:     parsed.User.LastName,
			Username:     parsed.User.Username,
			LanguageCode: parsed.User.LanguageCode,
			Phone:        "", // Phone отсутствует, задаем пустую строку
			Hash:         parsed.Hash,
		}, nil
	}

	// Выполняем стандартную валидацию init-data.
	if err := initdata.Validate(initData, token, expIn); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// Парсим данные, если валидация прошла успешно.
	parsed, err := initdata.Parse(initData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse init data: %w", err)
	}

	return &UserData{
		ID:           parsed.User.ID,
		FirstName:    parsed.User.FirstName,
		LastName:     parsed.User.LastName,
		Username:     parsed.User.Username,
		LanguageCode: parsed.User.LanguageCode,
		Phone:        "", // Если данных телефона нет, оставляем пустую строку.
		Hash:         parsed.Hash,
	}, nil
}
