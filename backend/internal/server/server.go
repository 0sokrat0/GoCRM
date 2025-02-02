package server

import (
	bot_handlers "GoCRM/internal/app/telegram/bot_handlers"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/joho/godotenv/autoload"

	"GoCRM/internal/app"
	"GoCRM/internal/database"
	"GoCRM/internal/domain/repository"
	"GoCRM/internal/infrastructure/persistence/postgres"
	"GoCRM/internal/server/handlers"
)

type Server struct {
	port int

	db database.Service

	bot        *tgbotapi.BotAPI // Добавить
	botHandler *bot_handlers.BotHandler

	userRepo    repository.UserRepository
	userService *app.UserService
	userHandler *handlers.UserHandler

	bookingRepo    repository.BookingRepository
	bookingService *app.BookingService
	bookingHandler *handlers.BookingHandler

	serviceRepo    repository.ServiceRepository
	serviceService *app.ServiceService
	serviceHandler *handlers.ServiceHandler

	masterRepo    repository.MasterProfileRepository
	masterService *app.MasterProfileService
	masterHandler *handlers.MasterProfileHandler

	notificationRepo    repository.NotificationRepository
	notificationService *app.NotificationService
	notificationHandler *handlers.NotificationHandler

	paymentRepo    repository.PaymentRepository
	paymentService *app.PaymentService
	paymentHandler *handlers.PaymentHandler

	auditRepo    repository.AuditLogRepository
	auditService *app.AuditLogService
	auditHandler *handlers.AuditLogHandler
}

func NewServer() *http.Server {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 8080
	}

	mode := os.Getenv("GIN_MODE")
	log.Println("Current GIN_MODE:", mode) // Проверяем, что загружается

	if mode == "" {
		mode = "debug" // Значение по умолчанию, если переменная не установлена
	}

	gin.SetMode(mode)

	s := &Server{
		port: port,
		db:   database.New(),
	}

	// Создаём объект Server

	// Получаем *sql.DB из нашего database.Service
	pgDB := s.db.DB()

	// Инициализируем репозитории / сервисы / обработчики

	// 1. Пользователи (User)
	s.userRepo = postgres.NewPGUserRepo(pgDB)
	s.userService = app.NewUserService(s.userRepo)
	s.userHandler = handlers.NewUserHandler(s.userService)

	// 2. Бронирования (Booking)
	s.bookingRepo = postgres.NewPGBookingRepo(pgDB)
	s.bookingService = app.NewBookingService(s.bookingRepo)
	s.bookingHandler = handlers.NewBookingHandler(s.bookingService)

	// 3. Услуги (Service)
	s.serviceRepo = postgres.NewPGServiceRepo(pgDB)
	s.serviceService = app.NewServiceService(s.serviceRepo)
	s.serviceHandler = handlers.NewServiceHandler(s.serviceService)

	// 4. Профиль мастера (MasterProfile)
	s.masterRepo = postgres.NewPGMasterProfileRepo(pgDB)
	s.masterService = app.NewMasterProfileService(s.masterRepo)
	s.masterHandler = handlers.NewMasterProfileHandler(s.masterService)

	// 5. Уведомления (Notification)
	s.notificationRepo = postgres.NewPGNotificationRepo(pgDB)
	s.notificationService = app.NewNotificationService(s.notificationRepo)
	s.notificationHandler = handlers.NewNotificationHandler(s.notificationService)

	// 6. Платежи (Payment)
	s.paymentRepo = postgres.NewPGPaymentRepo(pgDB)
	s.paymentService = app.NewPaymentService(s.paymentRepo)
	s.paymentHandler = handlers.NewPaymentHandler(s.paymentService)

	// 7. Аудит лог (AuditLog) – если нужно
	s.auditRepo = postgres.NewPGAuditLogRepo(pgDB)
	s.auditService = app.NewAuditLogService(s.auditRepo)
	s.auditHandler = handlers.NewAuditLogHandler(s.auditService)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}
	s.botHandler = bot_handlers.NewBotHandler(
		bot,
		s.userService,
		s.bookingService,
	)

	webhookURL := os.Getenv("WEBHOOK_URL") + "/webhook"
	webhookConfig, err := tgbotapi.NewWebhook(webhookURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to create webhook: %v", err))
	}

	_, err = bot.Request(webhookConfig)
	// Создаём http.Server, передав туда маршруты
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return httpServer
}
