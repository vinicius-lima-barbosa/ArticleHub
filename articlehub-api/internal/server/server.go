package server

import (
	"github.com/gofiber/fiber/v2"

	"articlehub-api/internal/database"
	"articlehub-api/internal/handler"
)

type FiberServer struct {
	*fiber.App

	db      database.Service
	handler *handler.UserHandler
}

func New() *FiberServer {
	db := database.New()
	userHandler := handler.NewUserHandler(db.UserRepo())

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "articlehub-api",
			AppName:      "articlehub-api",
		}),

		db:      db,
		handler: userHandler,
	}

	return server
}
