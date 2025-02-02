package handlers

import (
	"GoCRM/internal/app"
	"GoCRM/internal/domain/entity"
	"GoCRM/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuditLogHandler struct {
	service *app.AuditLogService
}

func NewAuditLogHandler(s *app.AuditLogService) *AuditLogHandler {
	return &AuditLogHandler{service: s}
}

func (h *AuditLogHandler) CreateAuditLogHandler(c *gin.Context) {
	var a entity.AuditLog
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
		return
	}
	ctx := c.Request.Context()
	if err := h.service.CreateAuditLog(ctx, &a); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to create audit log", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.SuccessResponse(a))
}

func (h *AuditLogHandler) GetAuditLogHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid audit log ID", http.StatusBadRequest, err.Error()))
		return
	}
	ctx := c.Request.Context()
	a, err := h.service.GetAuditLog(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse("Audit log not found", http.StatusNotFound, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(a))
}

func (h *AuditLogHandler) ListAuditLogsHandler(c *gin.Context) {
	// Для простоты передадим пустой фильтр. В реальной реализации извлекайте параметры из запроса.
	ctx := c.Request.Context()
	logs, err := h.service.ListAuditLogs(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to list audit logs", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(logs))
}

func (h *AuditLogHandler) DeleteAuditLogHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid audit log ID", http.StatusBadRequest, err.Error()))
		return
	}
	ctx := c.Request.Context()
	if err := h.service.DeleteAuditLog(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to delete audit log", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(map[string]string{"message": "Audit log deleted successfully"}))
}
