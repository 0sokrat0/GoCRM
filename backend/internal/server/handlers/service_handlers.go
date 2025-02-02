package handlers

import (
	"GoCRM/internal/app"
	"GoCRM/internal/domain/entity"
	"GoCRM/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ServiceHandler struct {
	service *app.ServiceService
}

func NewServiceHandler(s *app.ServiceService) *ServiceHandler {
	return &ServiceHandler{service: s}
}

func (h *ServiceHandler) CreateServiceHandler(c *gin.Context) {
	var svc entity.Service
	if err := c.ShouldBindJSON(&svc); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
		return
	}
	ctx := c.Request.Context()
	if err := h.service.CreateService(ctx, &svc); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to create service", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.SuccessResponse(svc))
}

func (h *ServiceHandler) GetServiceHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid service ID", http.StatusBadRequest, err.Error()))
		return
	}
	ctx := c.Request.Context()
	svc, err := h.service.GetService(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse("Service not found", http.StatusNotFound, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(svc))
}

func (h *ServiceHandler) UpdateServiceHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid service ID", http.StatusBadRequest, err.Error()))
		return
	}

	var svc entity.Service
	if err := c.ShouldBindJSON(&svc); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
		return
	}
	svc.ServiceID = id
	ctx := c.Request.Context()
	updated, err := h.service.UpdateService(ctx, &svc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to update service", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(updated))
}

func (h *ServiceHandler) DeleteServiceHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid service ID", http.StatusBadRequest, err.Error()))
		return
	}
	ctx := c.Request.Context()
	if err := h.service.DeleteService(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to delete service", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(map[string]string{"message": "Service deleted successfully"}))
}
