package handler

import (
	"news-app-be23/internal/features/articles"
)

type ArticleResponse struct {
	ID          uint   `json:"id"`
	Tag         string `json:"tag"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkPhoto   string `json:"link_photo"`
	UserID      uint   `json:"user_id"`
}

func ToResponseArticle(a articles.Article) ArticleResponse {
	return ArticleResponse{
		ID:          a.ID,
		Tag:         a.Tag,
		Title:       a.Title,
		Description: a.Description,
		LinkPhoto:   a.LinkPhoto,
	}
}

func ToResponseArticles(articles []articles.Article) []ArticleResponse {
	response := make([]ArticleResponse, len(articles))
	for i, a := range articles {
		response[i] = ToResponseArticle(a)
	}
	return response
}
