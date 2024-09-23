package server

import (
	"quizClone/config"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Run() {
	app := fiber.New()

	app.Listen(s.config.Port)
}
