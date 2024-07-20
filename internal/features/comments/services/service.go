package services

import (
	"errors"
	"news-app-be23/internal/features/comments"
)

type CommentServices struct {
	qry comments.Services
}

func NewCommentService(q comments.Services) comments.Services {
	return &CommentServices{
		qry: q,
	}
}

func (cs *CommentServices) InsertComment(newComment comments.Comment) error {
	err := cs.qry.InsertComment(newComment)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat menambah komentar")
	}
	return nil
}

func (cs *CommentServices) GetAllComments() ([]comments.Comment, error) {
	comments, err := cs.qry.GetAllComments()
	if err != nil {
		return nil, errors.New("terjadi kesalahan pada server saat mengambil daftar komentar")
	}
	return comments, nil
}

func (cs *CommentServices) DeleteComment(id uint) error {
	err := cs.qry.DeleteComment(id)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat menghapus komentar")
	}
	return nil
}
