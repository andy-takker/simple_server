package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthcheckHandler struct {
}

func NewHealthcheckHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/health", h.GetHealthcheck)
}

func (h *HealthcheckHandler) GetHealthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
