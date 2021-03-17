// package backend implements the backend functalities
package server

import (
	"github.com/gin-gonic/gin"
	"github.com/picaroup/webapp/backend/server/handlers"
)

type Server struct {
	router	*gin.Engine
}

// NewServer creates a new server
func NewServer() *Server{
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server)AddRequestHandlers() {
	handlers.AddHandlers(s.GetRouter())
}

func (s *Server)GetRouter() *gin.Engine {
	return s.router
}

