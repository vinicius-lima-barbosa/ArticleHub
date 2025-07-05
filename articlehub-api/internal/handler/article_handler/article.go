package article_handler

import (
	"context"
	"time"

	article_model "articlehub-api/internal/model/aritcle-model"
	article_repository "articlehub-api/internal/repository/article-repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ArticleHandler struct {
	Repo article_repository.ArticleRepository
}

func NewArticleHandler(repo article_repository.ArticleRepository) *ArticleHandler {
	return &ArticleHandler{Repo: repo}
}

func (h *ArticleHandler) CreateArticle(c *fiber.Ctx) error {
	var req article_model.CreateArticleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title is required",
		})
	}
	if req.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Content is required",
		})
	}

	id, err := uuid.NewV7()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate user ID",
		})
	}

	art := &article_model.Article{
		ID:       id.String(),
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: c.Locals("id").(string),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.Repo.CreateArticle(ctx, art); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create article",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Article created successfully",
		"article": art,
	})
}
