package services

import (
	"errors"
	"news-app-be23/internal/features/articles"
)

type ArticleServices struct {
	qry articles.Services
}

func NewArticleService(q articles.Services) articles.Services {
	return &ArticleServices{
		qry: q,
	}
}

func (as *ArticleServices) InsertArticle(newArticle articles.Article) error {
	err := as.qry.InsertArticle(newArticle)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat menambah artikel")
	}
	return nil
}

func (as *ArticleServices) GetArticleByID(id uint) (*articles.Article, error) {
	article, err := as.qry.GetArticleByID(id)
	if err != nil {
		return nil, errors.New("terjadi kesalahan pada server saat mengambil artikel")
	}
	return article, nil
}

func (as *ArticleServices) GetAllArticles() ([]articles.Article, error) {
	articles, err := as.qry.GetAllArticles()
	if err != nil {
		return nil, errors.New("terjadi kesalahan pada server saat mengambil daftar artikel")
	}
	return articles, nil
}

func (as *ArticleServices) UpdateArticle(updatedArticle articles.Article) error {
	err := as.qry.UpdateArticle(updatedArticle)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat mengupdate artikel")
	}
	return nil
}

func (as *ArticleServices) DeleteArticle(id uint) error {
	err := as.qry.DeleteArticle(id)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat menghapus artikel")
	}
	return nil
}
