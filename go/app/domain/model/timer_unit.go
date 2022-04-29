package model
import (
	"time"
	"ultimate_timer/services"
)
type TimerUnit struct {
	BaseModel
	Duration int    `db:"order" json:"duration"`
	Order    int    `db:"duration" json:"order"`
	PresetID string `db:"preset_id" json:"-"`
}

func NewTimerUnit(presetID string, duration, order int) (*TimerUnit, error) {
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

func (tu *TimerUnit) Set(duration int, order int) error {
	tu.UpdatedAt = time.Now()
	tu.Duration = duration
	tu.Order = order

	return nil
}
