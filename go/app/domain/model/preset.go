package model

import (
	"encoding/json"
	"time"
)

type Preset struct {
	BaseModel
	Name             string          `db:"name"`
	DisplayOrder     int             `db:"display_order"`
	LoopCount        int             `db:"loop_count"`
	WaitsConfirmEach bool            `db:"waits_confirm_each"`
	WaitsConfirmLast bool            `db:"waits_confirm_last"`
	TimerUnits       json.RawMessage `db:"timer_unit"`
}

func NewPreset(
	name string,
	displayOrder, loopCount int,
	waitsConfirmEach, waitsConfirmLast bool,
	timerUnits json.RawMessage) (*Preset, error) {

	now := time.Now()

	preset := &Preset{
		BaseModel: BaseModel{
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
	timerUnits json.RawMessage) error {

	p.UpdatedAt = time.Now()
	p.Name = name
	p.DisplayOrder = displayOrder
	p.LoopCount = loopCount
	p.WaitsConfirmEach = waitsConfirmEach
	p.WaitsConfirmLast = waitsConfirmLast
	p.TimerUnits = timerUnits

	return nil
}
