package repository

import (
	"news-app-be23/internal/features/articles"
	"time"

	"gorm.io/gorm"
)

type ArticleModel struct {
	db *gorm.DB
}

func NewArticleModel(connection *gorm.DB) articles.Query {
	return &ArticleModel{
		db: connection,
	}
}

func (am *ArticleModel) InsertArticle(newArticle articles.Article) error {
	return am.db.Create(&newArticle).Error
}

func (am *ArticleModel) GetAllArticles() ([]articles.Article, error) {
	var articleList []articles.Article
	err := am.db.Where("deleted_at IS NULL").Find(&articleList).Error
	return articleList, err
}

func (am *ArticleModel) GetArticleByID(id uint) (*articles.Article, error) {
	var article articles.Article
	err := am.db.Where("id = ? AND deleted_at IS NULL", id).First(&article).Error
	return &article, err
}

func (am *ArticleModel) UpdateArticle(updatedArticle articles.Article) error {
	return am.db.Save(&updatedArticle).Error
}

func (am *ArticleModel) DeleteArticle(id uint) error {
	var article articles.Article
	err := am.db.First(&article, id).Error
	if err != nil {
		return err
	}
	now := time.Now()
	article.DeletedAt = &now
	return am.db.Save(&article).Error // Soft delete
}
