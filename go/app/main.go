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
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
	port := fmt.Sprintf(":%s", os.Getenv("GO_PORT"))

	timerPresetRepository := infra.NewTimerPresetRepository(config.NewDB(), config.NewRedis())
	timerPresetUsecase := usecase.NewTimerPresetUsecase(timerPresetRepository)
	timerPresetHandler := handler.NewTimerPresetHandler(timerPresetUsecase)

	e := echo.New()
	e.Use(services.Logger)
	e.Use(middleware.Recover())
	handler.InitRouting(e, timerPresetHandler)
	e.Logger.Fatal(e.Start(port))
}
