package server

import (
	"articlehub-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (s *FiberServer) RegisterFiberRoutes() {
	// Apply CORS middleware
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false, // credentials require explicit origins
		MaxAge:           300,
	}))

	s.App.Get("/", s.HelloWorldHandler)
	s.App.Get("/health", s.healthHandler)

	users := s.App.Group("/users")
	users.Get("/", s.user_handler.GetUsers)
	users.Post("/", s.user_handler.CreateUser)
	users.Post("/login", s.user_handler.Login)
	users.Get("/:id", s.user_handler.GetUserById)
	users.Put("/:id", middleware.Middleware(), s.user_handler.UpdateUser)
	users.Put("/update-avatar/:id", middleware.Middleware(), s.user_handler.UpdateUserAvatar)
	users.Delete("/:id", s.user_handler.DeleteUser)

	articles := s.App.Group("/articles")
	articles.Post("/", middleware.Middleware(), s.article_handler.CreateArticle)
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
