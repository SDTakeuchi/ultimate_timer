package main

import (
	"ultimate_timer/services"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
    e := echo.New()

	// Middleware
	e.Use(services.Logger)
	e.Use(middleware.Recover())

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}
