package handler

import "news-app-be23/internal/features/articles"

type InsertArticleRequest struct {
	Tag         string `json:"tag"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkPhoto   string `json:"link_photo"`
	UserID      uint   `json:"user_id"`
}

type UpdateArticleRequest struct {
	ID          uint   `json:"id"`
	Tag         string `json:"tag"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkPhoto   string `json:"link_photo"`
	UserID      uint   `json:"user_id"`
}

func ToModelArticles(r InsertArticleRequest) articles.Article {
	return articles.Article{
		Tag:         r.Tag,
		Title:       r.Title,
		Description: r.Description,
		LinkPhoto:   r.LinkPhoto,
		UserID:      r.UserID,
	}
}

func ToModelArticle(r UpdateArticleRequest) articles.Article {
	return articles.Article{
		ID:          r.ID,
		Tag:         r.Tag,
		Title:       r.Title,
		Description: r.Description,
		LinkPhoto:   r.LinkPhoto,
		UserID:      r.UserID,
	}
}
