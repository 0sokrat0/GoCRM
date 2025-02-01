package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

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

	userRepo    repository.UserRepository
	userService *app.UserService
	userHandler *handlers.UserHandler
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

	s := &Server{
		port: port,
		db:   database.New(),
	}

	pgDB := s.db.DB()

	s.userRepo = postgres.NewPGUserRepo(pgDB)
	s.userService = app.NewUserService(s.userRepo)
	s.userHandler = handlers.NewUserHandler(s.userService)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return srv
}
