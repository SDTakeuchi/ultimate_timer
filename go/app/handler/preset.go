package handler

import (
	"encoding/json"
	"net/http"
	"github.com/labstack/echo"
	"ultimate_timer/usecase"
)

type PresetHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	FindByID() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type presetHandler struct {
	presetUsecase usecase.PresetUsecase
}

func NewPresetHandler(presetUsecase usecase.PresetUsecase) PresetHandler {
	return &presetHandler{presetUsecase: presetUsecase}
}

type requestPreset struct {
	Name             string `json:"name"`
	DisplayOrder     int    `json:"display_order"`
	LoopCount        int    `json:"loop_count"`
	WaitsConfirmEach bool   `json:"waits_confirm_each"`
	WaitsConfirmLast bool   `json:"waits_confirm_last"`
	TimerUnits       []structTimerUnit `json:"timer_unit"`
}

type responsePreset struct {
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	DisplayOrder     int               `json:"display_order"`
	LoopCount        int               `json:"loop_count"`
	WaitsConfirmEach bool              `json:"waits_confirm_each"`
	WaitsConfirmLast bool              `json:"waits_confirm_last"`
	TimerUnits       []structTimerUnit `json:"timer_unit"`
}

type structTimerUnit struct {
	Order    int `json:"order"`
	Duration int `json:"duration"`
}

func (ph *presetHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestPreset
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		tuJson, err := json.Marshal(req.TimerUnits)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdPreset, err := ph.presetUsecase.Create(
			req.Name,
			req.DisplayOrder,
			req.LoopCount,
			req.WaitsConfirmEach,
			req.WaitsConfirmLast,
			tuJson,
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res := &responsePreset{
			ID:               createdPreset.ID.String(),
			Name:             createdPreset.Name,
			LoopCount:        createdPreset.LoopCount,
			WaitsConfirmEach: createdPreset.WaitsConfirmEach,
			WaitsConfirmLast: createdPreset.WaitsConfirmLast,
		}
		var stu []structTimerUnit
		err = json.Unmarshal(createdPreset.TimerUnits, &stu)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res.TimerUnits = stu

		return c.JSON(http.StatusCreated, res)
	}
}

func (ph *presetHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		foundPresets, err := ph.presetUsecase.Get()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		var res []responsePreset
		for _, fp := range foundPresets {
			p := responsePreset{
				ID:               fp.ID.String(),
				Name:             fp.Name,
				LoopCount:        fp.LoopCount,
				WaitsConfirmEach: fp.WaitsConfirmEach,
				WaitsConfirmLast: fp.WaitsConfirmLast,
			}
			var stu []structTimerUnit
			err = json.Unmarshal(fp.TimerUnits, &stu)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			p.TimerUnits = stu
			res = append(res, p)
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (ph *presetHandler) FindByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		foundPreset, err := ph.presetUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res := responsePreset{
			ID:               foundPreset.ID.String(),
			Name:             foundPreset.Name,
			LoopCount:        foundPreset.LoopCount,
			WaitsConfirmEach: foundPreset.WaitsConfirmEach,
			WaitsConfirmLast: foundPreset.WaitsConfirmLast,
		}
		var stu []structTimerUnit
		err = json.Unmarshal(foundPreset.TimerUnits, &stu)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res.TimerUnits = stu
		return c.JSON(http.StatusOK, res)
	}
}

func (ph *presetHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		var req requestPreset
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		tuJson, err := json.Marshal(req.TimerUnits)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedPreset, err := ph.presetUsecase.Update(
			id,
			req.Name,
			req.DisplayOrder,
			req.LoopCount,
			req.WaitsConfirmEach,
			req.WaitsConfirmLast,
			tuJson,
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responsePreset{
			ID:               updatedPreset.ID.String(),
			Name:             updatedPreset.Name,
			LoopCount:        updatedPreset.LoopCount,
			WaitsConfirmEach: updatedPreset.WaitsConfirmEach,
			WaitsConfirmLast: updatedPreset.WaitsConfirmLast,
		}
		var stu []structTimerUnit
		err = json.Unmarshal(updatedPreset.TimerUnits, &stu)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res.TimerUnits = stu

		return c.JSON(http.StatusOK, res)
	}
}

func (th *presetHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		if err := th.presetUsecase.Delete(id); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
