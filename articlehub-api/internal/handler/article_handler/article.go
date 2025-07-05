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

func (h *ArticleHandler) GetArticleById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Article ID is required",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	article, err := h.Repo.GetArticleById(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Article not found",
		})
	}

	return c.JSON(article)
}

func (h *ArticleHandler) GetArticlesByAuthorId(c *fiber.Ctx) error {
	authorId := c.Params("authorId")
	if authorId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Author ID is required",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	articles, err := h.Repo.GetArticlesByAuthorId(ctx, authorId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve articles",
		})
	}

	return c.JSON(articles)
}

func (h *ArticleHandler) DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Article ID is required",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.Repo.DeleteArticle(ctx, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete article",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Article deleted successfully",
	})
}
