package handler

import "news-app-be23/internal/features/comments"

type CommentResponse struct {
	ID        uint   `json:"id"`
	Text      string `json:"text"`
	UserID    uint   `json:"user_id"`
	ArticleID uint   `json:"article_id"`
}

func ToResponseComment(c comments.Comment) CommentResponse {
	return CommentResponse{
		ID:        c.ID,
		Text:      c.Text,
		UserID:    c.UserID,
		ArticleID: c.ArticleID,
	}
}

func ToResponseComments(comments []comments.Comment) []CommentResponse {
	response := make([]CommentResponse, len(comments))
	for i, c := range comments {
		response[i] = ToResponseComment(c)
	}
	return response
}
