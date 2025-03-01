package httpserver

import (
	"GoCRM/internal/config"
	"GoCRM/internal/delivery/http_server/handlers"
	"GoCRM/persistence/db"
	"GoCRM/persistence/repositories"
	"GoCRM/pkg/logger"

	"GoCRM/internal/usecase"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	port           int
	db             db.Database
	userHandler    *handlers.UserHandler
	serviceHandler *handlers.ServiceHandler
}

func NewServer(cfg *config.Config) *http.Server {

	logger.Info("🔹 Приложение:", zap.String("name", cfg.App.Name))
	logger.Info("| Среда:", zap.String("env", cfg.App.Env))
	logger.Info("| Порт:", zap.Int("port", cfg.App.Port))

	gin.SetMode(gin.ReleaseMode)

	dbService, err := db.NewDatabase(&cfg.Database)
	if err != nil {
		logger.Fatal("❌ Ошибка подключения к БД: %v", zap.Error(err))
	}
	gormDB := dbService.DB()

	userRepo := repositories.NewDBUserRepo(gormDB)
	userService := usecase.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	serviceRepo := repositories.NewPGServiceRepo(gormDB)
	serviceService := usecase.NewService(serviceRepo)
	serviceHandler := handlers.NewServiceHandler(serviceService)

	s := &Server{
		port:           cfg.App.Port,
		db:             dbService,
		userHandler:    userHandler,
		serviceHandler: serviceHandler,
	}

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return httpServer
}
