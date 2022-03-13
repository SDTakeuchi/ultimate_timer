package models

import (
	"time"
	"ultimate_timer/services"
)

type TimerUnit struct {
	BaseModel
	Duration time.Duration `db:"duration" json:"duration"`
	Order    int           `db:"order" json:"order"`
	PresetID string        `db:"preset_id" json:"preset_id"`
}

// constructor
func NewTimerUnit(duration time.Duration, order int, presetID string) (*TimerUnit, error) {
	now := time.Now()
	id := services.GenUuid()

	tu := &TimerUnit{
		BaseModel: BaseModel{
			ID:        id,
			CreatedAt: now,
			UpdatedAt: now,
		},
		Duration: duration,
		Order:    order,
		PresetID: presetID,
	}

	return tu, nil
}

// setter
func (tu *TimerUnit) Set(duration time.Duration, order int) error {
	now := time.Now()

	tu.UpdatedAt = now
	tu.Duration = duration
	tu.Order = order

	return nil
}
