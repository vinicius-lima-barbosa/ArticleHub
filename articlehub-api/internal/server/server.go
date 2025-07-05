package server

import (
	"github.com/gofiber/fiber/v2"

	"articlehub-api/internal/database"
	"articlehub-api/internal/handler/user_handler"
)

type FiberServer struct {
	*fiber.App

	db      database.Service
	handler *user_handler.UserHandler
}

func New() *FiberServer {
	db := database.New()
	userHandler := user_handler.NewUserHandler(db.UserRepo())

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
