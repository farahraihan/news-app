package repository

import (
	"news-app-be23/internal/features/comments"

	"gorm.io/gorm"
)

type CommentModel struct {
	db *gorm.DB
}

func NewCommentModel(connection *gorm.DB) comments.Query {
	return &CommentModel{
		db: connection,
	}
}

func (cm *CommentModel) InsertComment(newComment comments.Comment) error {
	return cm.db.Create(&newComment).Error
}

func (cm *CommentModel) GetAllComments() ([]comments.Comment, error) {
	var commentList []comments.Comment
	err := cm.db.Where("deleted_at IS NULL").Find(&commentList).Error
	return commentList, err
}

func (cm *CommentModel) DeleteComment(id uint) error {
	return cm.db.Delete(&Comment{}, id).Error
}
