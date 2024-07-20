package articles

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Article struct {
	ID          uint
	UserID      uint
	Tag         string
	Title       string
	Description string
	LinkPhoto   string
	DeletedAt   *time.Time `sql:"index"`
}

type Handler interface {
	InsertArticle() echo.HandlerFunc
	GetArticleByID() echo.HandlerFunc
	GetAllArticles() echo.HandlerFunc
	UpdateArticle() echo.HandlerFunc
	DeleteArticle() echo.HandlerFunc
}

type Services interface {
	InsertArticle(newArticle Article) error
	GetArticleByID(id uint) (*Article, error)
	GetAllArticles() ([]Article, error)
	UpdateArticle(updatedArticle Article) error
	DeleteArticle(id uint) error
}

type Query interface {
	InsertArticle(newArticle Article) error
	GetAllArticles() ([]Article, error)
	GetArticleByID(id uint) (*Article, error)
	UpdateArticle(updatedArticle Article) error
	DeleteArticle(id uint) error
}
