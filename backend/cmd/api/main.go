package main

import (
	"GoCRM/internal/config"
	httpserver "GoCRM/internal/delivery/http_server"
	"GoCRM/pkg/logger"
	"context"
	"fmt"
	"log"
	"net/http"

	"os/signal"
	"syscall"
	"time"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {

	logger.Shutdown()

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

// @title GoCRM API
// @version 1.0
// @description REST API для системы управления клиентами
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()

	logger.Init(logger.Config{
		Environment: cfg.Logger.Level,
		Color:       true,
		AddCaller:   true,
		CallerSkip:  5,
	})

	server := httpserver.NewServer(cfg)

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}
