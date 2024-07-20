package repository

import (
	"news-app-be23/internal/features/articles"
	"news-app-be23/internal/features/comments"
	"news-app-be23/internal/features/users"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID    uint             `json:"user_id"`
	User      users.User       `gorm:"foreignKey:UserID"`
	ArticleID uint             `json:"article_id"`
	Article   articles.Article `gorm:"foreignKey:ArticleID"`
	Text      string           `json:"text"`
	DeletedAt gorm.DeletedAt   `gorm:"index"`
}

func (c *Comment) ToCommentEntity() comments.Comment {
	return comments.Comment{
		ID:        c.ID,
		UserID:    c.UserID,
		ArticleID: c.ArticleID,
		Text:      c.Text,
	}
}

func ToCommentData(input comments.Comment) Comment {
	return Comment{
		UserID:    input.UserID,
		ArticleID: input.ArticleID,
		Text:      input.Text,
	}
}
