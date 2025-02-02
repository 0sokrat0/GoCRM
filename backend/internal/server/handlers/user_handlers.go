package handlers

import (
	"GoCRM/internal/app"
	"GoCRM/internal/app/telegram"
	"GoCRM/pkg/response"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service *app.UserService
}

func NewUserHandler(s *app.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) TelegramAuthHandler(c *gin.Context) {
	type TelegramAuthRequest struct {
		InitData string `json:"init_data"`
	}

	var req TelegramAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Invalid request body:", err)
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request", http.StatusBadRequest, err.Error()))
		return
	}

	log.Println("Received init_data:", req.InitData)

	// Используем стороннюю валидацию через init-data-golang.
	userData, err := telegram.ValidateTelegramInitDataWithThirdParty(req.InitData)
	if err != nil {
		log.Println("Validation failed:", err)
		c.JSON(http.StatusUnauthorized, response.ErrorResponse("Invalid Telegram data", http.StatusUnauthorized, err.Error()))
		return
	}

	log.Println("User validated:", userData.ID, userData.Username)

	// Создаем или получаем пользователя.
	user, err := h.service.GetOrCreateTelegramUser(
		c.Request.Context(),
		userData.ID,
		userData.Username,
		userData.FirstName,
		userData.LastName,
		userData.LanguageCode,
		userData.Phone,
	)
	if err != nil {
		log.Println("Error creating/getting user:", err)
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to process user", http.StatusInternalServerError, err.Error()))
		return
	}

	// Обновляем сессию пользователя.
	if err := h.service.UpdateUserSession(c.Request.Context(), user.TelegramID, userData.Hash); err != nil {
		log.Println("Error updating session:", err)
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to update session", http.StatusInternalServerError, err.Error()))
		return
	}

	log.Println("User authenticated:", user.TelegramID)
	c.JSON(http.StatusOK, response.SuccessResponse(user))
}

func (h *UserHandler) GetUserByTelegramIDHandler(c *gin.Context) {
	tgIDStr := c.Param("tgID")
	tgID, err := strconv.ParseInt(tgIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid Telegram ID", http.StatusBadRequest, err.Error()))
		return
	}

	user, err := h.service.GetUserByTelegramID(c.Request.Context(), tgID)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse("User not found", http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(user))
}

func (h *UserHandler) GetUserByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid user ID", http.StatusBadRequest, err.Error()))
		return
	}

	user, err := h.service.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse("User not found", http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(user))
}
