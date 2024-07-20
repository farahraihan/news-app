package routes

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"news-app-be23/internal/features/articles"
	"news-app-be23/internal/features/comments"
	"news-app-be23/internal/features/users"
)

func InitRoute(e *echo.Echo, uc users.Handler, ac articles.Handler, cc comments.Handler) {
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, "hello world")
	})

	e.POST("/register", uc.SignUp())
	e.POST("/login", uc.Login())

	// JWT Middleware
	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		fmt.Println("JWT_SECRET environment variable not set")
	}

	articleGroup := e.Group("/articles")
	articleGroup.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return jwt.MapClaims{}
		},
		SigningKey:    []byte(jwtKey),
		SigningMethod: jwt.SigningMethodHS256.Name,
	}))

	// Article routes
	articleGroup.POST("", ac.InsertArticle())
	articleGroup.PUT("/:id", ac.UpdateArticle())
	articleGroup.DELETE("/:id", ac.DeleteArticle())

	// Public routes
	e.GET("/articles", ac.GetAllArticles())
	e.GET("/articles/:id", ac.GetArticleByID())

	commentGroup := e.Group("/comments")
	commentGroup.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return jwt.MapClaims{}
		},
		SigningKey:    []byte(jwtKey),
		SigningMethod: jwt.SigningMethodHS256.Name,
	}))

	// Comment routes
	commentGroup.POST("", cc.InsertComment())
	commentGroup.DELETE("/:id", cc.DeleteComment())

	// Public routes
	e.GET("/comments", cc.GetAllComments())
}
