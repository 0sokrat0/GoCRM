package handlers

// import (
// 	"GoCRM/internal/app"
// 	"GoCRM/internal/domain/entity"
// 	"GoCRM/pkg/response"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// type MasterProfileHandler struct {
// 	service *app.MasterProfileService
// }

// func NewMasterProfileHandler(s *app.MasterProfileService) *MasterProfileHandler {
// 	return &MasterProfileHandler{service: s}
// }

// func (h *MasterProfileHandler) CreateMasterProfileHandler(c *gin.Context) {
// 	var mp entity.MasterProfile
// 	if err := c.ShouldBindJSON(&mp); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
// 		return
// 	}
// 	ctx := c.Request.Context()
// 	if err := h.service.CreateMasterProfile(ctx, &mp); err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to create master profile", http.StatusInternalServerError, err.Error()))
// 		return
// 	}
// 	c.JSON(http.StatusCreated, response.SuccessResponse(mp))
// }

// func (h *MasterProfileHandler) GetMasterProfileHandler(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := uuid.Parse(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid master profile ID", http.StatusBadRequest, err.Error()))
// 		return
// 	}
// 	ctx := c.Request.Context()
// 	mp, err := h.service.GetMasterProfile(ctx, id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, response.ErrorResponse("Master profile not found", http.StatusNotFound, err.Error()))
// 		return
// 	}
// 	c.JSON(http.StatusOK, response.SuccessResponse(mp))
// }

// func (h *MasterProfileHandler) UpdateMasterProfileHandler(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := uuid.Parse(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid master profile ID", http.StatusBadRequest, err.Error()))
// 		return
// 	}
// 	var mp entity.MasterProfile
// 	if err := c.ShouldBindJSON(&mp); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request payload", http.StatusBadRequest, err.Error()))
// 		return
// 	}
// 	mp.MasterID = id
// 	ctx := c.Request.Context()
// 	updated, err := h.service.UpdateMasterProfile(ctx, &mp)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to update master profile", http.StatusInternalServerError, err.Error()))
// 		return
// 	}
// 	c.JSON(http.StatusOK, response.SuccessResponse(updated))
// }

// func (h *MasterProfileHandler) DeleteMasterProfileHandler(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := uuid.Parse(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid master profile ID", http.StatusBadRequest, err.Error()))
// 		return
// 	}
// 	ctx := c.Request.Context()
// 	if err := h.service.DeleteMasterProfile(ctx, id); err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to delete master profile", http.StatusInternalServerError, err.Error()))
// 		return
// 	}
// 	c.JSON(http.StatusOK, response.SuccessResponse(map[string]string{"message": "Master profile deleted successfully"}))
// }
