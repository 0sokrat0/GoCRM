package httpserver

import (
	"GoCRM/internal/config"
	"GoCRM/internal/delivery/http_server/handlers"
	"GoCRM/persistence/db"
	"GoCRM/persistence/repositories"

	"GoCRM/internal/usecase"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port           int
	db             db.Database
	userHandler    *handlers.UserHandler
	serviceHandler *handlers.ServiceHandler
}

func NewServer() *http.Server {

	cfg := config.GetConfig()

	fmt.Println("🔹 Приложение:", cfg.App.Name, "| Среда:", cfg.App.Env, "| Порт:", cfg.App.Port)
	fmt.Println("🔹 База данных:", cfg.Database.User, "@", cfg.Database.Host)

	log.Println("Current GIN_MODE:", cfg.App.GinMode)
	gin.SetMode(cfg.App.GinMode)

	dbService, err := db.NewDatabase(&cfg.Database)
	if err != nil {
		log.Fatalf("❌ Ошибка подключения к БД: %v", err)
	}
	gormDB := dbService.DB()

	userRepo := repositories.NewPGUserRepo(gormDB)
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
