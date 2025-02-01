package handlers

import (
	"GoCRM/internal/app"
	"GoCRM/internal/domain/entity"
	"GoCRM/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service *app.UserService
}

func NewUserHandler(s *app.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to create user", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.SuccessResponse(user))

}

func (h *UserHandler) GetUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid user ID", http.StatusBadRequest, err.Error()))
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse("User not found", http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(user))
}

func (h *UserHandler) UpdateUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid user ID", http.StatusBadRequest, err.Error()))
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
		return
	}

	user.ID = id

	updatedUser, err := h.service.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to update user", http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(updatedUser))
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid user ID", http.StatusBadRequest, err.Error()))
		return
	}

	user := &entity.User{ID: id}

	if err := h.service.DeleteUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to delete user", http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(map[string]string{"message": "User deleted successfully"}))
}
