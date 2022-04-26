package main

import (
	"fmt"
	"os"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/joho/godotenv"
	"ultimate_timer/config"
	"ultimate_timer/handler"
	"ultimate_timer/infra"
	"ultimate_timer/services"
	"ultimate_timer/usecase"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	port := fmt.Sprintf(":%s", os.Getenv("GO_PORT"))

	presetRepository := infra.NewPresetRepository(config.NewDB(), config.NewRedis())
	presetUsecase := usecase.NewPresetUsecase(presetRepository)
	presetHandler := handler.NewPresetHandler(presetUsecase)

	e := echo.New()
	e.Use(services.Logger)
	e.Use(middleware.Recover())
	handler.InitRouting(e, presetHandler)
	e.Logger.Fatal(e.Start(port))
}
