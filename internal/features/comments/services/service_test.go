package services_test

import (
	"news-app-be23/internal/features/comments"
	"news-app-be23/internal/features/comments/mocks"
	"news-app-be23/internal/features/comments/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestInsertComment(t *testing.T) {
	qry := mocks.NewQuery(t)
	srv := services.NewCommentService(qry)
	input := comments.Comment{UserID: 1, ArticleID: 2, Text: "first comment"}

	t.Run("Success Insert Comment", func(t *testing.T) {
		qry.On("InsertComment", input).Return(nil).Once()
		err := srv.InsertComment(input)
		qry.AssertExpectations(t)

		assert.Nil(t, err)
	})

	t.Run("Error From Query", func(t *testing.T) {
		qry.On("InsertComment", input).Return(gorm.ErrInvalidData).Once()
		err := srv.InsertComment(input)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server saat menambah komentar")
	})
}

func TestGetAllComment(t *testing.T) {
	qry := mocks.NewQuery(t)
	srv := services.NewCommentService(qry)
	id := uint(1)
	result := []comments.Comment{
		{
			ID:        id,
			UserID:    1,
			ArticleID: 2,
			Text:      "first comment",
		},
	}

	t.Run("Success Get All Comment", func(t *testing.T) {
		qry.On("GetAllComments").Return(result, nil).Once()
		data, err := srv.GetAllComments()

		qry.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, result, data)
	})

	t.Run("Error From Query", func(t *testing.T) {
		qry.On("GetAllComments").Return(result, gorm.ErrInvalidData).Once()
		_, err := srv.GetAllComments()

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server saat mengambil daftar komentar")
	})
}

func TestDeleteComment(t *testing.T) {
	qry := mocks.NewQuery(t)
	srv := services.NewCommentService(qry)
	id := uint(1)

	t.Run("Success Delete Comment", func(t *testing.T) {
		qry.On("DeleteComment", id).Return(nil).Once()
		err := srv.DeleteComment(id)

		qry.AssertExpectations(t)

		assert.Nil(t, err)
	})

	t.Run("Error From Query", func(t *testing.T) {
		qry.On("DeleteComment", id).Return(gorm.ErrInvalidData).Once()
		err := srv.DeleteComment(id)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server saat menghapus komentar")
	})
}
