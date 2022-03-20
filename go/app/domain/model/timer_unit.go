package model

type TimerUnit struct {
	Duration int `db:"order" json:"duration"`
	Order    int           `db:"duration" json:"order"`
	PresetID string          `db:"preset_id" json:"-"`			//hides in json response
}

// constructor
func NewTimerUnit(duration int, order int, presetID string) (*TimerUnit, error) {
	tu := &TimerUnit{
		Duration: duration,
		Order:    order,
		PresetID: presetID,
	}

	return tu, nil
}

// setter
func (tu *TimerUnit) Set(duration int, order int) error {
	tu.Duration = duration
	tu.Order = order

	return nil
}
