package handlers

import (
	"GoCRM/internal/app"
	"GoCRM/internal/domain/entity"
	"GoCRM/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PaymentHandler struct {
	service *app.PaymentService
}

func NewPaymentHandler(s *app.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: s}
}

func (h *PaymentHandler) CreatePaymentHandler(c *gin.Context) {
	var p entity.Payment
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
		return
	}
	ctx := c.Request.Context()
	if err := h.service.CreatePayment(ctx, &p); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to create payment", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.SuccessResponse(p))
}

func (h *PaymentHandler) GetPaymentHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid payment ID", http.StatusBadRequest, err.Error()))
		return
	}
	ctx := c.Request.Context()
	p, err := h.service.GetPayment(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse("Payment not found", http.StatusNotFound, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(p))
}

func (h *PaymentHandler) UpdatePaymentHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid payment ID", http.StatusBadRequest, err.Error()))
		return
	}
	var p entity.Payment
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
		return
	}
	p.PaymentID = id
	ctx := c.Request.Context()
	updated, err := h.service.UpdatePayment(ctx, &p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to update payment", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(updated))
}

func (h *PaymentHandler) DeletePaymentHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid payment ID", http.StatusBadRequest, err.Error()))
		return
	}
	ctx := c.Request.Context()
	if err := h.service.DeletePayment(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to delete payment", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(map[string]string{"message": "Payment deleted successfully"}))
}
