package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes настраивает маршруты для сервера.
func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	// Настройка CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Добавьте URL вашего фронтенда
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Базовые эндпоинты
	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)
	// В routes.go
	r.POST("/webhook", s.botHandler.HandleWebhookUpdate)

	// Группа API (например, версия 1)
	api := r.Group("/api/v1")
	{
		// User endpoints
		api.POST("/users", s.userHandler.CreateUserHandler)
		api.GET("/users/:id", s.userHandler.GetUserHandler)
		api.PUT("/users/:id", s.userHandler.UpdateUserHandler)
		api.DELETE("/users/:id", s.userHandler.DeleteUserHandler)
		api.GET("/users/telegram/:tgID", s.userHandler.GetUserByTelegramIDHandler) // Новый эндпоинт
		api.POST("/users/telegram", s.userHandler.CreateOrUpdateUserHandler)

		// Booking endpoints
		api.POST("/booking", s.bookingHandler.CreateBookingHandler)
		api.GET("/booking/:id", s.bookingHandler.GetByIDHandler)
		api.PUT("/booking/:id", s.bookingHandler.UpdateBookingHandler)
		api.DELETE("/booking/:id", s.bookingHandler.CancelBookingHandler)
		api.PUT("/booking/:id/reschedule", s.bookingHandler.RescheduleBookingHandler)

		// Service endpoints
		api.POST("/service", s.serviceHandler.CreateServiceHandler)
		api.GET("/service/:id", s.serviceHandler.GetServiceHandler)
		api.PUT("/service/:id", s.serviceHandler.UpdateServiceHandler)
		api.DELETE("/service/:id", s.serviceHandler.DeleteServiceHandler)

		// MasterProfile endpoints
		api.POST("/master", s.masterHandler.CreateMasterProfileHandler)
		api.GET("/master/:id", s.masterHandler.GetMasterProfileHandler)
		api.PUT("/master/:id", s.masterHandler.UpdateMasterProfileHandler)
		api.DELETE("/master/:id", s.masterHandler.DeleteMasterProfileHandler)

		// Notification endpoints
		api.POST("/notification", s.notificationHandler.CreateNotificationHandler)
		api.GET("/notification/:id", s.notificationHandler.GetNotificationHandler)
		api.GET("/notification/all", s.notificationHandler.ListNotificationsHandler)
		api.PUT("/notification/:id", s.notificationHandler.UpdateNotificationHandler)
		api.DELETE("/notification/:id", s.notificationHandler.DeleteNotificationHandler)

		// Payment endpoints
		api.POST("/payment", s.paymentHandler.CreatePaymentHandler)
		api.GET("/payment/:id", s.paymentHandler.GetPaymentHandler)
		api.PUT("/payment/:id", s.paymentHandler.UpdatePaymentHandler)
		api.DELETE("/payment/:id", s.paymentHandler.DeletePaymentHandler)
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
