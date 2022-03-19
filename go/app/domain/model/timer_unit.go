package model

import (
	"time"
)

type TimerUnit struct {
	Duration time.Duration `db:"order" json:"duration"`
	Order    int           `db:"duration" json:"order"`
	PresetID string          `db:"preset_id" json:"-"`			//hides in json response
}

// constructor
func NewTimerUnit(duration time.Duration, order int, presetID string) (*TimerUnit, error) {
	tu := &TimerUnit{
		Duration: duration,
		Order:    order,
		PresetID: presetID,
	}

	return tu, nil
}

// setter
func (tu *TimerUnit) Set(duration time.Duration, order int) error {
	tu.Duration = duration
	tu.Order = order

	return nil
}
