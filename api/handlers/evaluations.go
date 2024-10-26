package handlers

import (
	"math-expressions/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}
type ExpressionRequest struct {
	Expression string `json:"expression"`
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Evaluate(c *gin.Context) {
	var req ExpressionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if valid, reason := utils.ValidateExpression(req.Expression); !valid {
		c.JSON(http.StatusBadRequest, gin.H{"valid": false, "reason": reason})
		return
	}

	result, err := utils.CalculateExpression(req.Expression)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (h *Handler) Validate(c *gin.Context) {
	var req ExpressionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if valid, reason := utils.ValidateExpression(req.Expression); !valid {
		c.JSON(http.StatusBadRequest, gin.H{"valid": false, "reason": reason})
		return
	}

	c.JSON(http.StatusOK, gin.H{"valid": true})
}

func (h *Handler) GetErrors(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "test"})
}
