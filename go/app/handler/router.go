package handler

import (
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, timerPresetHandler TimerPresetHandler) {
	g := e.Group("/api")
	g.POST("/timer-preset/", timerPresetHandler.Post())
	g.GET("/timer-preset/", timerPresetHandler.Get())
	g.GET("/timer-preset/:id", timerPresetHandler.FindByID())
	g.PUT("/timer-preset/:id", timerPresetHandler.Put())
	g.DELETE("/timer-preset/:id", timerPresetHandler.Delete())
}
