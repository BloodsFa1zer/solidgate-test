package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"solidgate-test/router"
)

func main() {
	e := echo.New()
	//	e.HidePort = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.UserRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
