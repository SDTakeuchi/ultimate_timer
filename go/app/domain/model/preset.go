package model

import (
	// "errors"
	"time"
	"ultimate_timer/services"
)

type Preset struct {
	BaseModel
	Name             string      `db:"name" json:"name"`
	DisplayOrder     int         `db:"display_order" json:"display_order"`
	LoopCount        int         `db:"loop_count" json:"loop_count"`
	WaitsConfirmEach bool        `db:"waits_confirm_each" json:"waits_confirm_each"`
	WaitsConfirmLast bool        `db:"waits_confirm_last" json:"waits_confirm_last"`
	TimerUnits       []TimerUnit `db:"timer_unit" json:"timer_unit"`
}

func NewPreset(
	name string,
	displayOrder, loopCount int,
	waitsConfirmEach, waitsConfirmLast bool,
	timerUnits []TimerUnit) (*Preset, error) {

	// allows blank
	// if name == "" {
	// 	return nil, errors.New("項目名を入力してください")
	// }

	now := time.Now()
	id := services.GenUuid()

	preset := &Preset{
		BaseModel: BaseModel{
			ID:        id,
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name:             name,
		DisplayOrder:     displayOrder,
		LoopCount:        loopCount,
		WaitsConfirmEach: waitsConfirmEach,
		WaitsConfirmLast: waitsConfirmLast,
		TimerUnits:       timerUnits,
	}

	return preset, nil
}

func (p *Preset) Set(
	name string,
	displayOrder, loopCount int,
	waitsConfirmEach, waitsConfirmLast bool,
	timerUnits []TimerUnit) error {

	p.UpdatedAt = time.Now()
	p.Name = name
	p.DisplayOrder = displayOrder
	p.LoopCount = loopCount
	p.WaitsConfirmEach = waitsConfirmEach
	p.WaitsConfirmLast = waitsConfirmLast
	p.TimerUnits = timerUnits

	return nil
}
