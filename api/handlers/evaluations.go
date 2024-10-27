package handlers

import (
	"math-expressions/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExpressionRequest struct {
	Expression string `json:"expression"`
}

type Handler struct {
	db *gorm.DB
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Evaluate(c *gin.Context) {
	var req ExpressionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		utils.LogError(req.Expression, "/evaluate", err.Error())
		return
	}

	if valid, err := utils.ValidateExpression(req.Expression); !valid {
		c.JSON(http.StatusBadRequest, gin.H{"valid": false, "reason": err.Error()})
		utils.LogError(req.Expression, "/evaluate", err.Error())
		return
	}

	result, err := utils.CalculateExpression(req.Expression)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.LogError(req.Expression, "/evaluate", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (h *Handler) Validate(c *gin.Context) {
	var req ExpressionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		utils.LogError(req.Expression, "/validate", err.Error())
		return
	}

	if valid, err := utils.ValidateExpression(req.Expression); !valid {
		c.JSON(http.StatusBadRequest, gin.H{"valid": false, "reason": err.Error()})
		utils.LogError(req.Expression, "/validate", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"valid": true})
}

func (h *Handler) GetErrors(c *gin.Context) {
	result := utils.GetErrors()

	c.JSON(http.StatusOK, gin.H{"result": result})
}
