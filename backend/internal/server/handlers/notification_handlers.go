package handlers

import (
	"GoCRM/internal/app"
	"GoCRM/internal/domain/entity"
	"GoCRM/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NotificationHandler struct {
	service *app.NotificationService
}

func NewNotificationHandler(s *app.NotificationService) *NotificationHandler {
	return &NotificationHandler{service: s}
}

func (h *NotificationHandler) CreateNotificationHandler(c *gin.Context) {
	var n entity.Notification
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
		return
	}
	ctx := c.Request.Context()
	if err := h.service.CreateNotification(ctx, &n); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to create notification", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.SuccessResponse(n))
}

func (h *NotificationHandler) GetNotificationHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid notification ID", http.StatusBadRequest, err.Error()))
		return
	}
	ctx := c.Request.Context()
	n, err := h.service.GetNotification(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse("Notification not found", http.StatusNotFound, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(n))
}

func (h *NotificationHandler) UpdateNotificationHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid notification ID", http.StatusBadRequest, err.Error()))
		return
	}
	var n entity.Notification
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
		return
	}
	n.NotificationID = id
	ctx := c.Request.Context()
	updated, err := h.service.UpdateNotification(ctx, &n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to update notification", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(updated))
}

func (h *NotificationHandler) DeleteNotificationHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid notification ID", http.StatusBadRequest, err.Error()))
		return
	}
	ctx := c.Request.Context()
	if err := h.service.DeleteNotification(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to delete notification", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(map[string]string{"message": "Notification deleted successfully"}))
}

func (h *NotificationHandler) ListNotificationsHandler(c *gin.Context) {
	// Здесь можно извлечь параметры фильтрации из запроса (например, user_id, status)
	// Для простоты используем пустой фильтр
	ctx := c.Request.Context()
	notifications, err := h.service.ListNotifications(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to list notifications", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(notifications))
}
