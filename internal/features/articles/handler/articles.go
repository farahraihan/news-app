package handler

import (
	"news-app-be23/internal/features/articles"
	"news-app-be23/internal/helper"
	"news-app-be23/internal/utils"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	srv articles.Services
}

func NewArticleController(s articles.Services) articles.Handler {
	return &ArticleController{
		srv: s,
	}
}

func (ac *ArticleController) InsertArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input InsertArticleRequest
		err := c.Bind(&input)
		if err != nil {
			c.Logger().Error("insert article parse error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		newArticle := ToModelArticles(input)
		err = ac.srv.InsertArticle(newArticle)
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "terjadi kesalahan pada server") {
				errCode = 400
			}
			return c.JSON(500, helper.ResponseFormat(errCode, "server error", nil))
		}
		return c.JSON(201, helper.ResponseFormat(201, "success insert article", nil))
	}
}

func (ac *ArticleController) GetArticleByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		articleID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.Logger().Error("get article parse error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		article, err := ac.srv.GetArticleByID(uint(articleID))
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "terjadi kesalahan pada server") {
				errCode = 400
			}
			return c.JSON(500, helper.ResponseFormat(errCode, "server error", nil))
		}
		return c.JSON(200, helper.ResponseFormat(200, "success get article", ToResponseArticle(*article)))
	}
}

func (ac *ArticleController) GetAllArticles() echo.HandlerFunc {
	return func(c echo.Context) error {
		articles, err := ac.srv.GetAllArticles()
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "terjadi kesalahan pada server") {
				errCode = 400
			}
			return c.JSON(500, helper.ResponseFormat(errCode, "server error", nil))
		}
		return c.JSON(200, helper.ResponseFormat(200, "success get all articles", ToResponseArticles(articles)))
	}
}

func (ac *ArticleController) UpdateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateArticleRequest
		err := c.Bind(&input)
		if err != nil {
			c.Logger().Error("update article parse error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		id := c.Param("id")
		getId, err := utils.StringToUint(id)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		updatedArticle := ToModelArticle(input)
		updatedArticle.ID = getId
		err = ac.srv.UpdateArticle(updatedArticle)
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "terjadi kesalahan pada server") {
				errCode = 400
			}
			return c.JSON(500, helper.ResponseFormat(errCode, "server error", nil))
		}
		return c.JSON(200, helper.ResponseFormat(200, "success update article", nil))
	}
}

func (ac *ArticleController) DeleteArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		articleID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.Logger().Error("delete article parse error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		err = ac.srv.DeleteArticle(uint(articleID))
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "terjadi kesalahan pada server") {
				errCode = 400
			}
			return c.JSON(500, helper.ResponseFormat(errCode, "server error", nil))
		}
		return c.JSON(200, helper.ResponseFormat(200, "success delete article", nil))
	}
}
