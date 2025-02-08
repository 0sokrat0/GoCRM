package handlers

import (
	"GoCRM/internal/usecase"
	"GoCRM/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ServiceHandler struct {
	service *usecase.Service
}

func NewServiceHandler(s *usecase.Service) *ServiceHandler {
	return &ServiceHandler{service: s}
}

type UpdateServiceRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Duration    int     `json:"duration" binding:"required,gt=0"`
}

// DTO для запроса создания сервиса
type CreateServiceRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Duration    int     `json:"duration" binding:"required,gt=0"`
}

func (h *ServiceHandler) CreateServiceHandler(c *gin.Context) {
	var req CreateServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
		return
	}

	ctx := c.Request.Context()
	service, err := h.service.CreateService(ctx, req.Name, req.Description, req.Price, req.Duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to create service", http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse(service))
}

func (h *ServiceHandler) GetServiceHandler(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid service ID", http.StatusBadRequest, err.Error()))
		return
	}

	ctx := c.Request.Context()
	service, err := h.service.GetByID(ctx, id)
	if err != nil {
		if err.Error() == "service not found" { // use case должен возвращать кастомную ошибку
			c.JSON(http.StatusNotFound, response.ErrorResponse("Service not found", http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve service", http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(service))
}

func (h *ServiceHandler) UpdateServiceHandler(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid service ID", http.StatusBadRequest, err.Error()))
		return
	}

	var req UpdateServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
		return
	}

	ctx := c.Request.Context()
	service, err := h.service.UpdateService(ctx, id, req.Name, req.Description, req.Price, req.Duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to update service", http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(service))
}

func (h *ServiceHandler) DeleteServiceHandler(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid service ID", http.StatusBadRequest, err.Error()))
		return
	}

	ctx := c.Request.Context()
	err = h.service.DeleteByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to delete service", http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(map[string]string{"message": "Service deleted successfully"}))
}

func (h *ServiceHandler) ListServicesHandler(c *gin.Context) {
	ctx := c.Request.Context()
	services, err := h.service.ListServices(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to list services", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(services))
}
