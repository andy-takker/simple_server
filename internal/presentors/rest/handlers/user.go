package handlers

import (
	"net/http"
	"strconv"

	entities "github.com/andy-takker/simple_server/internal/domain/entities"
	services "github.com/andy-takker/simple_server/internal/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	schemas "github.com/andy-takker/simple_server/internal/presentors/rest/schemas"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: *service}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/users", h.CreateUser)
	r.GET("/users", h.FetchUserList)
	r.GET("/users/:id", h.FetchUserByID)
	r.PUT("/users/:id", h.UpdateUserByID)
	r.DELETE("/users/:id", h.DeleteUserByID)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req schemas.CreateUserSchema
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user, err := h.service.CreateUser(c, &entities.CreateUser{
		Username:  req.Username,
		Email:     req.Email,
		Phone:     req.Phone,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		if err == entities.ErrorUserAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, schemas.ConvertToUserSchema(user))
}

func (h *UserHandler) FetchUserList(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "100")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 100
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	users, err := h.service.FetchUserList(c, &entities.UserListParams{
		Limit:  int64(limit),
		Offset: int64(offset),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schemas.ConvertToUserListSchema(users))
}

func (h *UserHandler) FetchUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	user, err := h.service.FetchUserByID(c, id.String())
	if err != nil {
		if err == entities.ErrorUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schemas.ConvertToUserSchema(user))
}

func (h *UserHandler) UpdateUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	var req schemas.CreateUserSchema
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user, err := h.service.UpdateUserByID(c, &entities.UpdateUser{
		ID:        id.String(),
		Username:  req.Username,
		Email:     req.Email,
		Phone:     req.Phone,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		if err == entities.ErrorUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err == entities.ErrorUserAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schemas.ConvertToUserSchema(user))
}

func (h *UserHandler) DeleteUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	if err := h.service.DeleteUserByID(c, id.String()); err != nil {
		if err == entities.ErrorUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
