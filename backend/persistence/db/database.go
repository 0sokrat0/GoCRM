package db

import (
	"fmt"
	"log"
	"time"

	"GoCRM/internal/config"
	"GoCRM/persistence/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database представляет интерфейс работы с БД.
type Database interface {
	Health() map[string]string // ✅ Добавляем метод Health()
	Close() error
	DB() *gorm.DB
}

type database struct {
	db *gorm.DB
}

// NewDatabase создает подключение к БД
func NewDatabase(cfg *config.DatabaseConfig) (Database, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode, cfg.Schema,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, fmt.Errorf("❌ Ошибка подключения к БД: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("❌ Ошибка получения SQL DB: %w", err)
	}
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if config.GetConfig().App.Env != "production" { // Только в dev-режиме
		log.Println("⚠️  Запускаем AutoMigrate (dev mode)...")
		if err := db.AutoMigrate(
			&models.User{},
			&models.Service{},
		); err != nil {
			log.Fatalf("❌ Ошибка миграции: %v", err)
		}
		log.Println("✅ AutoMigrate успешно завершён")
	} else {
		log.Println("🚀 Продакшен-режим: AutoMigrate отключён")
	}

	log.Println("✅ Подключение к базе данных установлено")
	return &database{db: db}, nil
}

// ✅ Реализация метода Health()
func (d *database) Health() map[string]string {
	stats := make(map[string]string)
	sqlDB, err := d.db.DB()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = "Ошибка получения SQL DB"
		log.Println(stats["error"])
		return stats
	}

	err = sqlDB.Ping()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("БД недоступна: %v", err)
		log.Println(stats["error"])
		return stats
	}

	stats["status"] = "up"
	stats["message"] = "БД в рабочем состоянии"
	return stats
}

// Close закрывает соединение с БД
func (d *database) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		log.Println("❌ Ошибка при получении SQL DB:", err)
		return err
	}
	log.Println("⛔ Закрытие соединения с БД")
	return sqlDB.Close()
}

// DB возвращает GORM-подключение
func (d *database) DB() *gorm.DB {
	return d.db
}
