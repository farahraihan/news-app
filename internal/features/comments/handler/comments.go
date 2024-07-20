package handler

import (
	"news-app-be23/internal/features/comments"
	"news-app-be23/internal/helper"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
	srv comments.Services
}

func NewCommentController(s comments.Services) comments.Handler {
	return &CommentController{
		srv: s,
	}
}

func (cc *CommentController) InsertComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input InsertCommentRequest
		err := c.Bind(&input)
		if err != nil {
			c.Logger().Error("insert comment parse error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		newArticle := ToModelComments(input)
		err = cc.srv.InsertComment(newArticle)
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "terjadi kesalahan pada server") {
				errCode = 400
			}
			return c.JSON(500, helper.ResponseFormat(errCode, "server error", nil))
		}
		return c.JSON(201, helper.ResponseFormat(201, "success insert comment", nil))
	}
}

func (cc *CommentController) GetAllComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		comments, err := cc.srv.GetAllComments()
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "terjadi kesalahan pada server") {
				errCode = 400
			}
			return c.JSON(500, helper.ResponseFormat(errCode, "server error", nil))
		}
		return c.JSON(200, helper.ResponseFormat(200, "success get all comments", ToResponseComments(comments)))
	}
}

func (cc *CommentController) DeleteComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		commentID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.Logger().Error("delete comment parse error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		err = cc.srv.DeleteComment(uint(commentID))
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "terjadi kesalahan pada server") {
				errCode = 400
			}
			return c.JSON(500, helper.ResponseFormat(errCode, "server error", nil))
		}
		return c.JSON(200, helper.ResponseFormat(200, "success delete comment", nil))
	}
}
