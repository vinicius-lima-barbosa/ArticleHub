package server

import (
	"github.com/gofiber/fiber/v2"

	"articlehub-api/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "articlehub-api",
			AppName:      "articlehub-api",
		}),

		db: database.New(),
	}

	return server
}
