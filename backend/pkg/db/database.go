package db

import (
	"GoCRM/internal/config"
	"GoCRM/migrations/models"

	zlogger "GoCRM/pkg/logger"
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database interface {
	Health() map[string]string
	Close() error
	DB() *gorm.DB
}

type database struct {
	db *gorm.DB
}

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

	if config.GetConfig().App.Env != "production" {
		zlogger.Warn("⚠️  Запускаем AutoMigrate (dev mode)...")
		if err := db.AutoMigrate(
			&models.User{},
			&models.Service{},
		); err != nil {
			zlogger.Fatal("❌ Ошибка миграции: %v", zap.Error(err))
		}
		zlogger.Info("✅ AutoMigrate успешно завершён")
	} else {
		zlogger.Warn("🚀 Продакшен-режим: AutoMigrate отключён")
	}

	zlogger.Info("✅ Подключение к базе данных установлено")
	return &database{db: db}, nil
}

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

func (d *database) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		zlogger.Warn("❌ Ошибка при получении SQL DB:", zap.Error(err))
		return err
	}
	zlogger.Warn("⛔ Закрытие соединения с БД")
	return sqlDB.Close()
}

func (d *database) DB() *gorm.DB {
	return d.db
}
