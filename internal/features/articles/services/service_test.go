package services_test

import (
	"news-app-be23/internal/features/articles"
	"news-app-be23/internal/features/articles/mocks"
	"news-app-be23/internal/features/articles/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestInsertArticle(t *testing.T) {
	qry := mocks.NewQuery(t)
	srv := services.NewArticleService(qry)
	input := articles.Article{Tag: "test", Title: "berita baru 1", Description: "description 1", LinkPhoto: "localhost:500/photo/1"}

	t.Run("Success Insert Article", func(t *testing.T) {
		qry.On("InsertArticle", input).Return(nil).Once()
		err := srv.InsertArticle(input)
		qry.AssertExpectations(t)

		assert.Nil(t, err)
	})

	t.Run("Error From Query", func(t *testing.T) {
		qry.On("InsertArticle", input).Return(gorm.ErrInvalidData).Once()
		err := srv.InsertArticle(input)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server saat menambah artikel")
	})
}

func TestGetArticleByID(t *testing.T) {
	qry := mocks.NewQuery(t)
	srv := services.NewArticleService(qry)
	id := uint(1)
	result := &articles.Article{ID: id, Tag: "test", Title: "berita baru 1", Description: "description 1", LinkPhoto: "localhost:500/photo/1"}

	t.Run("Success Get Article By ID", func(t *testing.T) {
		qry.On("GetArticleByID", id).Return(result, nil).Once()
		data, err := srv.GetArticleByID(id)

		qry.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, result, data)
	})

	t.Run("Error From Query", func(t *testing.T) {
		qry.On("GetArticleByID", id).Return(result, gorm.ErrInvalidData).Once()
		_, err := srv.GetArticleByID(id)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server saat mengambil artikel")
	})
}

func TestGetAllArticle(t *testing.T) {
	qry := mocks.NewQuery(t)
	srv := services.NewArticleService(qry)
	id := uint(1)
	result := []articles.Article{
		{
			ID:          id,
			Tag:         "test",
			Title:       "berita baru 1",
			Description: "description 1",
			LinkPhoto:   "localhost:500/photo/1",
		},
	}

	t.Run("Success Get All Article", func(t *testing.T) {
		qry.On("GetAllArticles").Return(result, nil).Once()
		data, err := srv.GetAllArticles()

		qry.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, result, data)
	})

	t.Run("Error From Query", func(t *testing.T) {
		qry.On("GetAllArticles").Return(result, gorm.ErrInvalidData).Once()
		_, err := srv.GetAllArticles()

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server saat mengambil daftar artikel")
	})
}

func TestUpdateArticle(t *testing.T) {
	qry := mocks.NewQuery(t)
	srv := services.NewArticleService(qry)
	id := uint(1)
	result := articles.Article{ID: id, Tag: "test update", Title: "berita baru 1 update", Description: "description 1 update", LinkPhoto: "localhost:500/photo/1"}

	t.Run("Success Update Article", func(t *testing.T) {
		qry.On("UpdateArticle", result).Return(nil).Once()
		err := srv.UpdateArticle(result)

		qry.AssertExpectations(t)

		assert.Nil(t, err)
	})

	t.Run("Error From Query", func(t *testing.T) {
		qry.On("UpdateArticle", result).Return(gorm.ErrInvalidData).Once()
		err := srv.UpdateArticle(result)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server saat mengupdate artikel")
	})
}

func TestDeleteArticle(t *testing.T) {
	qry := mocks.NewQuery(t)
	srv := services.NewArticleService(qry)
	id := uint(1)

	t.Run("Success Delete Article", func(t *testing.T) {
		qry.On("DeleteArticle", id).Return(nil).Once()
		err := srv.DeleteArticle(id)

		qry.AssertExpectations(t)

		assert.Nil(t, err)
	})

	t.Run("Error From Query", func(t *testing.T) {
		qry.On("DeleteArticle", id).Return(gorm.ErrInvalidData).Once()
		err := srv.DeleteArticle(id)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server saat menghapus artikel")
	})
}
