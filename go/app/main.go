package main

import (
	"ultimate_timer/config"
	"ultimate_timer/handler"
	"ultimate_timer/infra"
	"ultimate_timer/services"
	"ultimate_timer/usecase"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	// "errors"

	// "encoding/json"
	// "ultimate_timer/services"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func main() {
    taskRepository := infra.NewPresetRepository(config.NewDB(), config.SetupRedis())
    taskUsecase := usecase.NewPresetUsecase(taskRepository)
    taskHandler := handler.NewPresetHandler(taskUsecase)

    e := echo.New()
	// Middleware
	e.Use(services.Logger)
	e.Use(middleware.Recover())
    handler.InitRouting(e, taskHandler)
    e.Logger.Fatal(e.Start(":8080"))
}
