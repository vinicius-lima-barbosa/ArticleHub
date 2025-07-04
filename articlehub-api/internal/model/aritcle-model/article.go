package article_model

import "time"

type Article struct {
	ID        string    `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	AuthorID  string    `json:"author_id" db:"author_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateArticleRequest struct {
	Title   string `json:"title" validate:"required,min=2,max=200"`
	Content string `json:"content" validate:"required,min=10"`
}
