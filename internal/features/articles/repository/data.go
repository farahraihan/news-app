package repository

import (
	"news-app-be23/internal/features/articles"
	"news-app-be23/internal/features/comments"
	"news-app-be23/internal/features/users"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UserID      uint               `json:"user_id"`
	User        users.User         `gorm:"foreignKey:UserID"`
	Tag         string             `json:"tag"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	LinkPhoto   string             `json:"link_photo"`
	DeletedAt   gorm.DeletedAt     `gorm:"index"`
	Comments    []comments.Comment `gorm:"foreignKey:ArticleID"`
}

func (a *Article) ToArticleEntity() articles.Article {
	return articles.Article{
		ID:          a.ID,
		UserID:      a.UserID,
		Tag:         a.Tag,
		Title:       a.Title,
		Description: a.Description,
		LinkPhoto:   a.LinkPhoto,
	}
}

func ToArticleData(input articles.Article) Article {
	return Article{
		UserID:      input.UserID,
		Tag:         input.Tag,
		Title:       input.Title,
		Description: input.Description,
		LinkPhoto:   input.LinkPhoto,
	}
}
