package model

import (
	"time"
	"ultimate_timer/service"
)

type TimerUnit struct {
	Duration time.Duration `json:"duration"`
	Order    int           `json:"order"`
	PresetID string          `json:"-"`			//hides in json response
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
