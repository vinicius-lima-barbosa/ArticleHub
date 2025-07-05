package article_repository

import (
	"context"
	"database/sql"
	"fmt"

	article_model "articlehub-api/internal/model/aritcle-model"
)

type ArticleRepository interface {
	CreateArticle(ctx context.Context, article *article_model.Article) error
	GetArticleById(ctx context.Context, id string) (*article_model.Article, error)
	GetArticlesByAuthorId(ctx context.Context, authorId string) ([]article_model.Article, error)
	DeleteArticle(ctx context.Context, id string) error
}

type articleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) CreateArticle(ctx context.Context, article *article_model.Article) error {
	query := `INSERT INTO articles (id, title, content, author_id, created_at) VALUES ($1, $2, $3, $4, NOW()) RETURNING id, created_at`

	err := r.db.QueryRowContext(ctx, query, article.ID, article.Title, article.Content, article.AuthorID).Scan(&article.ID, &article.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create article: %w", err)
	}

	return nil
}

func (r *articleRepository) GetArticleById(ctx context.Context, id string) (*article_model.Article, error) {
	query := `SELECT id, title, content, author_id, created_at FROM articles WHERE id = $1`

	var article article_model.Article
	err := r.db.QueryRowContext(ctx, query, id).Scan(&article.ID, &article.Title, &article.Content, &article.AuthorID, &article.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("article not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get article by id: %w", err)
	}

	return &article, nil
}

func (r *articleRepository) GetArticlesByAuthorId(ctx context.Context, authorId string) ([]article_model.Article, error) {
	query := `SELECT id, title, content, author_id, created_at FROM articles WHERE author_id = $1 ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, authorId)
	if err != nil {
		return nil, fmt.Errorf("failed to get articles by author id: %w", err)
	}
	defer rows.Close()

	var articles []article_model.Article
	for rows.Next() {
		var article article_model.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.AuthorID, &article.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan article: %w", err)
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (r *articleRepository) DeleteArticle(ctx context.Context, id string) error {
	query := `DELETE FROM articles WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete article: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no article found with id: %s", id)
	}

	return nil
}
