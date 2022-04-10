package handler

import (
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, presetHandler PresetHandler) {
	g := e.Group("/api")
	g.POST("/presets/", presetHandler.Post())
	g.GET("/presets/", presetHandler.Get())
	g.GET("/presets/:id", presetHandler.FindByID())
	g.PUT("/presets/:id", presetHandler.Put())
	g.DELETE("/presets/:id", presetHandler.Delete())
}
