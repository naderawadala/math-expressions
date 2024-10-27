package main

import (
	"log"
	"math-expressions/api"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	api.SetupRoutes(router)

	server := api.NewApiServer(":8080", router)

	if err := server.Run(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
