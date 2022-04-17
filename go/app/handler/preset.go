package handler

import (
	"net/http"
	// "strconv"

	"ultimate_timer/domain/model"
	"ultimate_timer/services"
	"ultimate_timer/usecase"

	// "github.com/fatih/structs"
	"github.com/labstack/echo"
	// null "gopkg.in/guregu/null.v4"
)

// PresetHandler preset handlerのinterface
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

// NewPresetHandler preset handlerのコンストラクタ
func NewPresetHandler(presetUsecase usecase.PresetUsecase) PresetHandler {
	return &presetHandler{presetUsecase: presetUsecase}
}

type requestPreset struct {
	Name             string `json:"name"`
	DisplayOrder     int    `json:"display_order"`
	LoopCount        int    `json:"loop_count"`
	WaitsConfirmEach bool   `json:"waits_confirm_each"`
	WaitsConfirmLast bool   `json:"waits_confirm_last"`
	TimerUnits       []struct {
		Order    int `json:"order"`
		Duration int `json:"duration"`
	} `json:"timer_unit"`
}

type responsePreset struct {
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	DisplayOrder     int               `json:"display_order"`
	LoopCount        int               `json:"loop_count"`
	WaitsConfirmEach bool              `json:"waits_confirm_each"`
	WaitsConfirmLast bool              `json:"waits_confirm_last"`
	TimerUnits       []model.TimerUnit `json:"timer_unit"`
}

// Post presetを保存するときのハンドラー
func (ph *presetHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestPreset
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		var tuMap []map[string]int
		for _, tu := range req.TimerUnits {
			m := services.StructToMapInt(tu)
			tuMap = append(tuMap, m)
		}

		createdPreset, err := ph.presetUsecase.Create(
			req.Name,
			req.DisplayOrder,
			req.LoopCount,
			req.WaitsConfirmEach,
			req.WaitsConfirmLast,
			tuMap,
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res := &responsePreset{
			ID:               createdPreset.ID,
			Name:             createdPreset.Name,
			LoopCount:        createdPreset.LoopCount,
			WaitsConfirmEach: createdPreset.WaitsConfirmEach,
			WaitsConfirmLast: createdPreset.WaitsConfirmLast,
			TimerUnits:       createdPreset.TimerUnits,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get presetを取得するときのハンドラー
func (ph *presetHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		foundPresets, err := ph.presetUsecase.Get()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		var res []responsePreset
		for _, fp := range foundPresets {
			p := responsePreset{
				ID:               fp.ID,
				Name:             fp.Name,
				LoopCount:        fp.LoopCount,
				WaitsConfirmEach: fp.WaitsConfirmEach,
				WaitsConfirmLast: fp.WaitsConfirmLast,
				TimerUnits:       fp.TimerUnits,
			}
			res = append(res, p)
		}
		return c.JSON(http.StatusOK, res)
	}
}

// Get presetを取得するときのハンドラー
func (ph *presetHandler) FindByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		foundPreset, err := ph.presetUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res := responsePreset{
			ID:               foundPreset.ID,
			Name:             foundPreset.Name,
			LoopCount:        foundPreset.LoopCount,
			WaitsConfirmEach: foundPreset.WaitsConfirmEach,
			WaitsConfirmLast: foundPreset.WaitsConfirmLast,
			TimerUnits:       foundPreset.TimerUnits,
		}
		return c.JSON(http.StatusOK, res)
	}
}

// Put presetを更新するときのハンドラー
func (ph *presetHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		var req requestPreset
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		var tuMap []map[string]int
		for _, tu := range req.TimerUnits {
			m := services.StructToMapInt(tu)
			tuMap = append(tuMap, m)
		}

		updatedPreset, err := ph.presetUsecase.Update(
			id,
			req.Name,
			req.DisplayOrder,
			req.LoopCount,
			req.WaitsConfirmEach,
			req.WaitsConfirmLast,
			tuMap,
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responsePreset{
			ID:               updatedPreset.ID,
			Name:             updatedPreset.Name,
			LoopCount:        updatedPreset.LoopCount,
			WaitsConfirmEach: updatedPreset.WaitsConfirmEach,
			WaitsConfirmLast: updatedPreset.WaitsConfirmLast,
			TimerUnits:       updatedPreset.TimerUnits,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete presetを削除するときのハンドラー
func (th *presetHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		err := th.presetUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
