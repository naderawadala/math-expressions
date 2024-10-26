package api

import (
	"math-expressions/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	handler := handlers.NewHandler()

	router.GET("/evaluate", handler.Evaluate)
	router.GET("/validate", handler.Validate)
	router.GET("/errors", handler.GetErrors)
}
