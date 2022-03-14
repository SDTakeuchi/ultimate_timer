package handler

import (
	"net/http"
	"strconv"
	"time"

	"app/usecase"

	"github.com/labstack/echo"
	null "gopkg.in/guregu/null.v4"
)

// PresetHandler preset handlerのinterface
type PresetHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
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
	Name             string      `json:"name"`
	DisplayOrder     int         `json:"display_order"`
	LoopCount        int         `json:"loop_count"`
	WaitsConfirmEach bool        `json:"waits_confirm_each"`
	WaitsConfirmLast bool        `json:"waits_confirm_last"`
	TimerUnits       []struct{
		Order int `json:"order"`
		Duration time.Duration `json:"duration"`
	} `json:"timer_unit"`
}

type responsePreset struct {
	ID string `json:"id"`
	Name             string      `json:"name"`
	DisplayOrder     int         `json:"display_order"`
	LoopCount        int         `json:"loop_count"`
	WaitsConfirmEach bool        `json:"waits_confirm_each"`
	WaitsConfirmLast bool        `json:"waits_confirm_last"`
	TimerUnits       []struct{
		Order int `json:"order"`
		Duration time.Duration `json:"duration"`
	} `json:"timer_unit"`
}

// Post presetを保存するときのハンドラー
func (th *presetHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestPreset
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdPreset, err := th.presetUsecase.Create(
			req.Name,
			req.Memo,
			req.Price,
			req.PurchaseDate,
			req.SmallCategoryId,
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responsePreset{
			ID:              createdPreset.ID,
			Name:            createdPreset.Name,
			Memo:            createdPreset.Memo,
			Price:           createdPreset.Price,
			PurchaseDate:    createdPreset.PurchaseDate,
			SmallCategoryId: createdPreset.SmallCategoryId,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get presetを取得するときのハンドラー
func (ih *presetHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := c.Param("id")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
        if id != "" {
            foundPreset, err := ih.presetUsecase.FindByID(id)
            if err != nil {
                return c.JSON(http.StatusBadRequest, err.Error())
            }
        }

		res := responsePreset{
			ID:      foundPreset.ID,
			Name:            foundPreset.Name,
			Memo:            foundPreset.Memo,
			Price:           foundPreset.Price,
			PurchaseDate:    foundPreset.PurchaseDate,
			SmallCategoryId: foundPreset.SmallCategoryId,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Put presetを更新するときのハンドラー
func (th *presetHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestPreset
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedPreset, err := th.presetUsecase.Update(id, req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responsePreset{
			ID:      updatedPreset.ID,
			Title:   updatedPreset.Title,
			Content: updatedPreset.Content,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete presetを削除するときのハンドラー
func (th *presetHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = th.presetUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
