package factory

import (
	"fmt"
	"news-app-be23/configs"
	articleHandler "news-app-be23/internal/features/articles/handler"
	articleRepository "news-app-be23/internal/features/articles/repository"
	articleServices "news-app-be23/internal/features/articles/services"
	commentHandler "news-app-be23/internal/features/comments/handler"
	commentRepository "news-app-be23/internal/features/comments/repository"
	commentServices "news-app-be23/internal/features/comments/services"
	userHandler "news-app-be23/internal/features/users/handler"
	userRepository "news-app-be23/internal/features/users/repository"
	userServices "news-app-be23/internal/features/users/services"
	"news-app-be23/internal/routes"
	"news-app-be23/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitFactory(e *echo.Echo) {
	cfg := configs.ImportSetting()
	db, err := configs.ConnectDB(cfg)
	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
	}
	if err := db.AutoMigrate(&userRepository.User{}, &articleRepository.Article{}, &commentRepository.Comment{}); err != nil {
		fmt.Println("Ada yg bermasalah saat memasukan table user", err.Error())
	}

	pu := utils.NewPasswordUtility()
	jt := utils.NewJwtUtility()

	um := userRepository.NewUserModel(db)
	us := userServices.NewUserService(um, pu, jt)
	uc := userHandler.NewUserController(us)

	am := articleRepository.NewArticleModel(db)
	as := articleServices.NewArticleService(am)
	ac := articleHandler.NewArticleController(as)

	cm := commentRepository.NewCommentModel(db)
	cs := commentServices.NewCommentService(cm)
	cc := commentHandler.NewCommentController(cs)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	routes.InitRoute(e, uc, ac, cc)
}
