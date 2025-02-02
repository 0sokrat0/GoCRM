// internal/app/telegram/bot_handlers/bot_handler.go
package bot_handlers

import (
	"GoCRM/internal/app"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler struct {
	bot            *tgbotapi.BotAPI
	userService    *app.UserService
	bookingService *app.BookingService
}

func NewBotHandler(
	bot *tgbotapi.BotAPI,
	userService *app.UserService,
	bookingService *app.BookingService,
) *BotHandler {
	return &BotHandler{
		bot:            bot,
		userService:    userService,
		bookingService: bookingService,
	}
}

func (h *BotHandler) HandleWebhookUpdate(c *gin.Context) {
	var update tgbotapi.Update
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(400, gin.H{"error": "Invalid update"})
		return
	}

	// if update.Message != nil {
	// 	switch update.Message.Text {
	// 	case "/start":
	// 		h.handleStartCommand(update.Message)
	// 		// Добавьте другие команды
	// 	}
	// }

	c.JSON(200, gin.H{"status": "ok"})
}

// func (h *BotHandler) handleStartCommand(msg *tgbotapi.Message) {
// 	user := &entity.User{
// 		TelegramID: msg.From.ID,
// 	}

// 	if err := h.userService.CreateOrUpdateUser(user); err != nil {
// 		// Обработка ошибки
// 		h.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Произошла ошибка"))
// 		return
// 	}

// 	response := "Добро пожаловать! Используйте:\n/mybookings - ваши записи"
// 	h.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, response))
// }
