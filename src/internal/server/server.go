package server

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	app *fiber.App
	db  *mongo.Database
}

func (s *Server) Run() {

}
