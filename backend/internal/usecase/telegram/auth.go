package telegram

import (
	"GoCRM/internal/config"
	"errors"
	"fmt"
	"log"
	"time"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

// UserData хранит данные пользователя из Telegram init-data.
type UserData struct {
	ID           int64
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	Phone        string
	Hash         string
}

func ValidateTelegramInitDataWithThirdParty(initData string) (*UserData, error) {
	if initData == "" {
		log.Println("🚨 Ошибка: initData пустой!")
		return nil, errors.New("initData is empty")
	}
	cfg := config.GetConfig()
	token := cfg.Telegram.BotToken
	if token == "" {
		log.Println("🚨 Ошибка: TELEGRAM_BOT_TOKEN не установлен!")
		return nil, errors.New("TELEGRAM_BOT_TOKEN is not set")
	}

	parsed, err := initdata.Parse(initData)
	if err != nil {
		log.Printf("🚨 Ошибка парсинга init-data: %v", err)
		return nil, fmt.Errorf("failed to parse init data: %w", err)
	}

	if err := initdata.Validate(initData, token, 24*time.Hour); err != nil {
		log.Printf("🚨 Ошибка валидации init-data: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	return &UserData{
		ID:           parsed.User.ID,
		FirstName:    parsed.User.FirstName,
		LastName:     parsed.User.LastName,
		Username:     parsed.User.Username,
		LanguageCode: parsed.User.LanguageCode,
		Phone:        "",
		Hash:         parsed.Hash,
	}, nil
}
