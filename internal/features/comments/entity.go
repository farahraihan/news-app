package comments

import "github.com/labstack/echo/v4"

type Comment struct {
	ID        uint
	UserID    uint
	ArticleID uint
	Text      string
}

type Handler interface {
	InsertComment() echo.HandlerFunc
	GetAllComments() echo.HandlerFunc
	DeleteComment() echo.HandlerFunc
}

type Services interface {
	InsertComment(newComment Comment) error
	GetAllComments() ([]Comment, error)
	DeleteComment(id uint) error
}

type Query interface {
	InsertComment(newComment Comment) error
	GetAllComments() ([]Comment, error)
	DeleteComment(id uint) error
}
