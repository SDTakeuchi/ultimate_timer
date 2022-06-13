package handler

import (
	"encoding/json"
	"net/http"
	"github.com/labstack/echo"
	"ultimate_timer/usecase"
)

type TimerPresetHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	FindByID() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type timerPresetHandler struct {
	timerPresetUsecase usecase.TimerPresetUsecase
}

func NewTimerPresetHandler(timerPresetUsecase usecase.TimerPresetUsecase) TimerPresetHandler {
	return &timerPresetHandler{timerPresetUsecase: timerPresetUsecase}
}

type requestTimerPreset struct {
	Name             string `json:"name"`
	DisplayOrder     int    `json:"display_order"`
	LoopCount        int    `json:"loop_count"`
	WaitsConfirmEach bool   `json:"waits_confirm_each"`
	WaitsConfirmLast bool   `json:"waits_confirm_last"`
	TimerUnits       []structTimerUnit `json:"timer_unit"`
}

type responseTimerPreset struct {
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

func (ph *timerPresetHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestTimerPreset
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		tuJson, err := json.Marshal(req.TimerUnits)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdTimerPreset, err := ph.timerPresetUsecase.Create(
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
		res := &responseTimerPreset{
			ID:               createdTimerPreset.ID.String(),
			Name:             createdTimerPreset.Name,
			LoopCount:        createdTimerPreset.LoopCount,
			WaitsConfirmEach: createdTimerPreset.WaitsConfirmEach,
			WaitsConfirmLast: createdTimerPreset.WaitsConfirmLast,
		}
		var stu []structTimerUnit
		err = json.Unmarshal(createdTimerPreset.TimerUnits, &stu)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res.TimerUnits = stu

		return c.JSON(http.StatusCreated, res)
	}
}

func (ph *timerPresetHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		foundTimerPresets, err := ph.timerPresetUsecase.Get()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		var res []responseTimerPreset
		for _, fp := range foundTimerPresets {
			p := responseTimerPreset{
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

func (ph *timerPresetHandler) FindByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		foundTimerPreset, err := ph.timerPresetUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res := responseTimerPreset{
			ID:               foundTimerPreset.ID.String(),
			Name:             foundTimerPreset.Name,
			LoopCount:        foundTimerPreset.LoopCount,
			WaitsConfirmEach: foundTimerPreset.WaitsConfirmEach,
			WaitsConfirmLast: foundTimerPreset.WaitsConfirmLast,
		}
		var stu []structTimerUnit
		err = json.Unmarshal(foundTimerPreset.TimerUnits, &stu)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res.TimerUnits = stu
		return c.JSON(http.StatusOK, res)
	}
}

func (ph *timerPresetHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		var req requestTimerPreset
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		tuJson, err := json.Marshal(req.TimerUnits)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedTimerPreset, err := ph.timerPresetUsecase.Update(
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

		res := responseTimerPreset{
			ID:               updatedTimerPreset.ID.String(),
			Name:             updatedTimerPreset.Name,
			LoopCount:        updatedTimerPreset.LoopCount,
			WaitsConfirmEach: updatedTimerPreset.WaitsConfirmEach,
			WaitsConfirmLast: updatedTimerPreset.WaitsConfirmLast,
		}
		var stu []structTimerUnit
		err = json.Unmarshal(updatedTimerPreset.TimerUnits, &stu)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res.TimerUnits = stu

		return c.JSON(http.StatusOK, res)
	}
}

func (th *timerPresetHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		if err := th.timerPresetUsecase.Delete(id); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
