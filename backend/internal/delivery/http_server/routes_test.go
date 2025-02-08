package httpserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"GoCRM/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	// Для тестов создаем новый экземпляр Gin и регистрируем только базовые маршруты.
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.SuccessResponse(map[string]string{"message": "Hello World"}))
	})
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.SuccessResponse(map[string]string{"status": "up"}))
	})
	// Дополнительно можно добавить тестовые эндпоинты для других обработчиков.
	return router
}

func TestHelloWorldHandler(t *testing.T) {
	router := setupTestRouter()

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var resp response.APIResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	assert.Contains(t, resp.Data.(map[string]interface{})["message"], "Hello World")
}

func TestHealthHandler(t *testing.T) {
	router := setupTestRouter()

	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var resp response.APIResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	assert.Equal(t, "up", resp.Data.(map[string]interface{})["status"])
}

func TestUserEndpointsH(t *testing.T) {
	// Предположим, что у вас есть мок или тестовая реализация UserHandler.
	// Для примера мы протестируем POST /users.
	router := gin.New()
	// Создадим фиктивный UserHandler, который просто возвращает входные данные.
	router.POST("/users", func(c *gin.Context) {
		var u map[string]interface{}
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
			return
		}
		c.JSON(http.StatusCreated, response.SuccessResponse(u))
	})

	// Тест запроса на создание пользователя.
	userJSON := `{"name": "John Doe", "email": "john@example.com", "phone": "+123456789", "tgID": "johndoe", "password": "secret", "role": "client"}`
	req, err := http.NewRequest("POST", "/users", strings.NewReader(userJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var resp response.APIResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
}

func TestBookingEndpointsH(t *testing.T) {
	// Аналогично можно написать тесты для бронирований.
	// Здесь мы создадим фиктивный обработчик CreateBooking, который возвращает входные данные.
	router := gin.New()
	router.POST("/booking", func(c *gin.Context) {
		var b map[string]interface{}
		if err := c.ShouldBindJSON(&b); err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
			return
		}
		c.JSON(http.StatusCreated, response.SuccessResponse(b))
	})

	bookingJSON := `{
		"user_id": "11111111-1111-1111-1111-111111111111",
		"master_id": "22222222-2222-2222-2222-222222222222",
		"service_id": "33333333-3333-3333-3333-333333333333",
		"booking_time": "2025-01-01T10:00:00Z",
		"status": "pending"
	}`
	req, err := http.NewRequest("POST", "/booking", strings.NewReader(bookingJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var resp response.APIResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
}

func TestServiceEndpointsH(t *testing.T) {
	// Тестируем фиктивный обработчик создания услуги.
	router := gin.New()
	router.POST("/service", func(c *gin.Context) {
		var s map[string]interface{}
		if err := c.ShouldBindJSON(&s); err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
			return
		}
		c.JSON(http.StatusCreated, response.SuccessResponse(s))
	})

	serviceJSON := `{
		"name": "Haircut",
		"description": "Standard haircut service",
		"price": 30.50,
		"duration": 45
	}`
	req, err := http.NewRequest("POST", "/service", strings.NewReader(serviceJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var resp response.APIResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
}

func TestPaymentEndpointsH(t *testing.T) {
	// Тестируем фиктивный обработчик создания платежа.
	router := gin.New()
	router.POST("/payment", func(c *gin.Context) {
		var p map[string]interface{}
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
			return
		}
		c.JSON(http.StatusCreated, response.SuccessResponse(p))
	})

	paymentJSON := `{
		"booking_id": "44444444-4444-4444-4444-444444444444",
		"amount": 30.50,
		"payment_method": "card",
		"status": "pending"
	}`
	req, err := http.NewRequest("POST", "/payment", strings.NewReader(paymentJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var resp response.APIResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
}

func TestNotificationEndpointsH(t *testing.T) {
	// Тестируем фиктивный обработчик создания уведомления.
	router := gin.New()
	router.POST("/notification", func(c *gin.Context) {
		var n map[string]interface{}
		if err := c.ShouldBindJSON(&n); err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
			return
		}
		c.JSON(http.StatusCreated, response.SuccessResponse(n))
	})

	notificationJSON := `{
		"user_id": "55555555-5555-5555-5555-555555555555",
		"type": "email",
		"message": "Your booking is confirmed",
		"status": "pending"
	}`
	req, err := http.NewRequest("POST", "/notification", strings.NewReader(notificationJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var resp response.APIResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
}

func TestAuditLogEndpointsH(t *testing.T) {
	// Тестируем фиктивный обработчик создания аудиторской записи.
	router := gin.New()
	router.POST("/audit_logs", func(c *gin.Context) {
		var a map[string]interface{}
		if err := c.ShouldBindJSON(&a); err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
			return
		}
		c.JSON(http.StatusCreated, response.SuccessResponse(a))
	})

	auditLogJSON := `{
		"user_id": "66666666-6666-6666-6666-666666666666",
		"action": "user_created",
		"details": "{\"name\": \"John Doe\"}"
	}`
	req, err := http.NewRequest("POST", "/audit_logs", strings.NewReader(auditLogJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var resp response.APIResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
}
