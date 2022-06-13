package model

import (
	"encoding/json"
	"time"
)

type TimerPreset struct {
	BaseModel
	Name             string          `db:"name"`
	DisplayOrder     int             `db:"display_order"`
	LoopCount        int             `db:"loop_count"`
	WaitsConfirmEach bool            `db:"waits_confirm_each"`
	WaitsConfirmLast bool            `db:"waits_confirm_last"`
	TimerUnits       json.RawMessage `db:"timer_unit"`
}

func NewTimerPreset(
	name string,
	displayOrder, loopCount int,
	waitsConfirmEach, waitsConfirmLast bool,
	timerUnits json.RawMessage) (*TimerPreset, error) {

	now := time.Now()

	tp := &TimerPreset{
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

	return tp, nil
}

func (p *TimerPreset) Set(
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
