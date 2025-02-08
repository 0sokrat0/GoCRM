package handlers

// import (
// 	"GoCRM/internal/app"
// 	"GoCRM/internal/domain/entity"
// 	"GoCRM/pkg/response"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// // BookingHandler обрабатывает HTTP-запросы, связанные с бронированиями.
// type BookingHandler struct {
// 	service *app.BookingService
// }

// // NewBookingHandler создает новый экземпляр BookingHandler.
// // Обратите внимание, что мы возвращаем объект, используя переданный параметр service.
// func NewBookingHandler(service *app.BookingService) *BookingHandler {
// 	return &BookingHandler{service: service}
// }

// // CreateBookingHandler обрабатывает POST-запрос для создания бронирования.
// func (b *BookingHandler) CreateBookingHandler(c *gin.Context) {
// 	var booking entity.Booking

// 	// Привязываем JSON из тела запроса к структуре Booking.
// 	if err := c.ShouldBindJSON(&booking); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	// Пример: можно проверить, что время бронирования не в прошлом.
// 	if booking.BookingTime.Before(time.Now()) {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Booking time cannot be in the past", http.StatusBadRequest, ""))
// 		return
// 	}

// 	// Извлекаем контекст из запроса.
// 	ctx := c.Request.Context()

// 	// Вызываем сервис для создания бронирования.
// 	if err := b.service.CreateBooking(ctx, &booking); err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to create booking", http.StatusInternalServerError, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusCreated, response.SuccessResponse(booking))
// }

// // GetByIDHandler обрабатывает GET-запрос для получения бронирования по его идентификатору.
// func (b *BookingHandler) GetByIDHandler(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := uuid.Parse(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid booking ID", http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	ctx := c.Request.Context()

// 	booking, err := b.service.GetBooking(ctx, id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, response.ErrorResponse("Booking not found", http.StatusNotFound, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.SuccessResponse(booking))
// }

// // UpdateBookingHandler обрабатывает PUT-запрос для обновления бронирования.
// func (b *BookingHandler) UpdateBookingHandler(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := uuid.Parse(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid booking ID", http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	var booking entity.Booking
// 	if err := c.ShouldBindJSON(&booking); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	// Устанавливаем ID бронирования из URL, чтобы гарантировать корректное обновление.
// 	booking.BookingID = id

// 	ctx := c.Request.Context()

// 	updatedBooking, err := b.service.UpdateBooking(ctx, &booking)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to update booking", http.StatusInternalServerError, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.SuccessResponse(updatedBooking))
// }

// // DeleteBookingHandler обрабатывает DELETE-запрос для отмены бронирования.
// // Здесь мы используем операцию CancelBooking, которая меняет статус бронирования на "canceled".
// func (b *BookingHandler) CancelBookingHandler(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := uuid.Parse(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid booking ID", http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	ctx := c.Request.Context()

// 	if err := b.service.CancelBooking(ctx, id); err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to cancel booking", http.StatusInternalServerError, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.SuccessResponse(map[string]string{"message": "Booking canceled successfully"}))
// }

// func (b *BookingHandler) RescheduleBookingHandler(c *gin.Context) {
// 	// Извлекаем идентификатор бронирования из URL
// 	idStr := c.Param("id")
// 	id, err := uuid.Parse(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid booking ID", http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	// Извлекаем новый временной параметр из query-параметра "new_time"
// 	newTimeStr := c.Query("new_time")
// 	if newTimeStr == "" {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Missing new booking time", http.StatusBadRequest, "Parameter 'new_time' is required"))
// 		return
// 	}

// 	// Парсим новый временной параметр, предполагаем формат RFC3339 (например, "2025-01-01T10:00:00Z")
// 	newTime, err := time.Parse(time.RFC3339, newTimeStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid booking time format", http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	// Извлекаем контекст из запроса
// 	ctx := c.Request.Context()

// 	// Вызываем сервисный метод для переноса бронирования
// 	rescheduledBooking, err := b.service.RescheduleBooking(ctx, id, newTime)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to reschedule booking", http.StatusInternalServerError, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.SuccessResponse(rescheduledBooking))
// }
