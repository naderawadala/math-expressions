package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Evaluate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "test"})
}

func (h *Handler) Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "test"})
}

func (h *Handler) GetErrors(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "test"})
}
