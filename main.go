package main

import (
	"news-app-be23/internal/factory"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	factory.InitFactory(e)
	e.Logger.Error(e.Start(":5000"))
}
