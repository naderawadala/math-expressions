package api

import (
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr   string
	router *gin.Engine
}

func NewApiServer(addr string, router *gin.Engine) *APIServer {
	return &APIServer{
		addr:   addr,
		router: router,
	}
}

func (s *APIServer) Run() error {
	return s.router.Run(s.addr)
}
