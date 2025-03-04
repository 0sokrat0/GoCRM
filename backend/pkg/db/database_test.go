package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	tcpostgres "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDB Database // ✅ Было `Service`, теперь `Database`

func mustStartPostgresContainer() (func(), string, error) {
	dbName := "testdb"
	dbUser := "testuser"
	dbPwd := "testpass"

	// Запуск контейнера с Postgres
	dbContainer, err := tcpostgres.Run(
		context.Background(),
		"postgres:latest",
		tcpostgres.WithDatabase(dbName),
		tcpostgres.WithUsername(dbUser),
		tcpostgres.WithPassword(dbPwd),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		return nil, "", err
	}

	// Получаем параметры подключения
	dbHost, err := dbContainer.Host(context.Background())
	if err != nil {
		return func() { dbContainer.Terminate(context.Background()) }, "", err
	}

	dbPort, err := dbContainer.MappedPort(context.Background(), "5432/tcp")
	if err != nil {
		return func() { dbContainer.Terminate(context.Background()) }, "", err
	}

	// ✅ Формируем корректный DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPwd, dbName, dbPort.Port(),
	)

	return func() { dbContainer.Terminate(context.Background()) }, dsn, nil
}

// TestMain — основной тестовый раннер.
func TestMain(m *testing.M) {
	teardown, dsn, err := mustStartPostgresContainer()
	if err != nil {
		log.Fatalf("❌ Ошибка запуска контейнера с Postgres: %v", err)
	}

	testDB = newTestDatabase(dsn)

	code := m.Run()

	testDB.Close()
	teardown()

	os.Exit(code)
}

func newTestDatabase(dsn string) Database { // ✅ Было `Service`, теперь `Database`
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Ошибка подключения к БД: %v", err)
	}

	return &database{db: db} // ✅ Было `&service{db: db}`
}

func TestNew(t *testing.T) {
	if testDB == nil {
		t.Fatal("❌ New() вернул nil")
	}
}

func TestHealth(t *testing.T) {
	stats := testDB.Health()

	if stats["status"] != "up" {
		t.Fatalf("❌ Ожидался статус 'up', получено: %s", stats["status"])
	}

	if _, ok := stats["error"]; ok {
		t.Fatalf("❌ Не ожидали ошибку в health check")
	}

	if stats["message"] != "БД в рабочем состоянии" {
		t.Fatalf("❌ Ожидали 'БД в рабочем состоянии', получили: %s", stats["message"])
	}
}

func TestClose(t *testing.T) {
	err := testDB.Close()
	if err != nil {
		t.Fatalf("❌ Ошибка при закрытии БД: %v", err)
	}

	sqlDB, _ := testDB.DB().DB() // ✅ Получаем `sql.DB` из GORM
	if err := sqlDB.Ping(); err == nil {
		t.Fatalf("❌ База данных не закрылась, Ping() не должен работать")
	}
}
