package httpserver

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func (s *Server) RegisterRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger()) // ✅ Логируем все запросы

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	api := r.Group("/api/v1")
	{
		api.POST("/auth/telegram", s.userHandler.TelegramAuthHandler)

		api.GET("/users/telegram/:tgID", s.userHandler.GetUserByTelegramIDHandler)
		api.GET("/users/:id", s.userHandler.GetUserByIDHandler)

		api.POST("/services", s.serviceHandler.CreateServiceHandler)
		api.GET("/services/:id", s.serviceHandler.GetServiceHandler)
		api.GET("/services", s.serviceHandler.ListServicesHandler)
		api.PATCH("/services/:id", s.serviceHandler.UpdateServiceHandler)
		api.DELETE("/services/:id", s.serviceHandler.DeleteServiceHandler)

	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Message: "Hello World"})
}

func (s *Server) healthHandler(c *gin.Context) {
	status := s.db.Health()
	if status["status"] != "up" {
		c.JSON(http.StatusServiceUnavailable, status)
		return
	}
	c.JSON(http.StatusOK, status)
}
