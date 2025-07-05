package server

import (
	"github.com/gofiber/fiber/v2"

	"articlehub-api/internal/database"
	"articlehub-api/internal/handler/article_handler"
	"articlehub-api/internal/handler/user_handler"
)

type FiberServer struct {
	*fiber.App

	db              database.Service
	user_handler    *user_handler.UserHandler
	article_handler *article_handler.ArticleHandler
}

func New() *FiberServer {
	db := database.New()
	userHandler := user_handler.NewUserHandler(db.UserRepo())
	articleHandler := article_handler.NewArticleHandler(db.ArticleRepo())

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "articlehub-api",
			AppName:      "articlehub-api",
		}),

		db:              db,
		user_handler:    userHandler,
		article_handler: articleHandler,
	}

	return server
}
