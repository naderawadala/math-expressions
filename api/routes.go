package api

import (
	"math-expressions/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	handler := handlers.NewHandler()

	router.POST("/evaluate", handler.Evaluate)
	router.POST("/validate", handler.Validate)
	router.GET("/errors", handler.GetErrors)
}
