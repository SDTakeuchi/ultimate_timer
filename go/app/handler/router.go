package handler

import (
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, presetHandler PresetHandler) {
	e.POST("/presets/", presetHandler.Post())
	e.GET("/presets/", presetHandler.Get())
	e.GET("/presets/:id", presetHandler.FindByID())
	e.PUT("/presets/:id", presetHandler.Put())
	e.DELETE("/presets/:id", presetHandler.Delete())
}
