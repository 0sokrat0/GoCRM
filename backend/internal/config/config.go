package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig      `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
	Telegram TelegramConfig `yaml:"telegram"`
}

type AppConfig struct {
	Name           string `yaml:"name" env:"APP_NAME" env-default:"GoCRM"`
	Env            string `yaml:"env" env:"APP_ENV" env-default:"development"`
	Port           int    `yaml:"port" env:"APP_PORT" env-default:"8080"`
	GinMode        string `yaml:"gin_mode" env:"GIN_MODE" env-default:"debug"`
	SkipValidation bool   `yaml:"skip_validation" env:"SKIP_VALIDATION" env-default:"true"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host" env:"POSTGRES_HOST" env-default:"localhost"`
	Port     int    `yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
	User     string `yaml:"user" env:"POSTGRES_USER" env-default:"sokrat"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-default:"1234"`
	Name     string `yaml:"name" env:"POSTGRES_DB" env-default:"crm"`
	Schema   string `yaml:"schema" env:"POSTGRES_SCHEMA" env-default:"public"`
	SSLMode  string `yaml:"sslmode" env:"POSTGRES_SSLMODE" env-default:"disable"`
}

type TelegramConfig struct {
	BotToken string `yaml:"bot_token" env:"TELEGRAM_BOT_TOKEN"`
}

var cfg *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{}

		if err := godotenv.Load(); err != nil {
			log.Println("⚠️ Не удалось загрузить .env, продолжаем...")
		}

		// Загружаем переменные в структуру конфигурации
		if err := cleanenv.ReadEnv(cfg); err != nil {
			log.Fatalf("❌ Ошибка загрузки конфигурации: %v", err)
		}

		log.Printf("✅ TELEGRAM_BOT_TOKEN загружен: %s", cfg.Telegram.BotToken)
	})
	return cfg
}
