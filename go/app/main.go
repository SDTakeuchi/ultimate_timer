package main

import (
	"ultimate_timer/config"
	"ultimate_timer/handler"
	"ultimate_timer/infra"
	"ultimate_timer/services"
	"ultimate_timer/usecase"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func main() {
    presetRepository := infra.NewPresetRepository(config.NewDB(), config.NewRedis())
    presetUsecase := usecase.NewPresetUsecase(presetRepository)
    presetHandler := handler.NewPresetHandler(presetUsecase)

    e := echo.New()
	e.Use(services.Logger)
	e.Use(middleware.Recover())
    handler.InitRouting(e, presetHandler)
    e.Logger.Fatal(e.Start(":8080"))
}
