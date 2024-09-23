package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	setUp()

	app.Listen(":3000")
}

func setUp() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "1234",
	})

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("error to connect database -->", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("enable to ping database")
	}

	log.Println("mongodb connected")
}
