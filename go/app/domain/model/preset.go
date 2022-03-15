package model

import (
	"errors"
	"time"
	"encoding/json"
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

type TimerUnit struct {
	Order    int           `db:"order"`
	Duration time.Duration `db:"duration"`
	PresetID string          `db:"preset_id"`
}

/*
constructor is for preparing objects (returns struct and error)
setter is for updating new objects (returns error if failed to udpate)
*/

// constructor
func NewPreset(
	name string,
	displayOrder, loopCount int,
	waitsConfirmEach, waitsConfirmLast bool,
	timerUnits []TimerUnit) (*Preset, error) {

	if name == "" {
		return nil, errors.New("項目名を入力してください")
	}

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
		TimerUnit:        timerUnits,
	}

	return preset, nil
}

// setter
func (p *Preset) Set(
	name string,
	displayOrder, loopCount int,
	waitsConfirmEach, waitsConfirmLast bool,
	timerUnits []TimerUnit) error {

	now := time.Now()

	p.UpdatedAt = now
	p.DisplayOrder = displayOrder
	p.LoopCount = loopCount
	p.WaitsConfirmEach = waitsConfirmEach
	p.WaitsConfirmLast = waitsConfirmLast
	p.TimerUnit = timerUnits

	return nil
}
