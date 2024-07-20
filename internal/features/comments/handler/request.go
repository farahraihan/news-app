package handler

import "news-app-be23/internal/features/comments"

type InsertCommentRequest struct {
	Text      string `json:"text"`
	UserID    uint   `json:"user_id"`
	ArticleID uint   `json:"article_id"`
}

func ToModelComments(r InsertCommentRequest) comments.Comment {
	return comments.Comment{
		Text:      r.Text,
		UserID:    r.UserID,
		ArticleID: r.ArticleID,
	}
}
