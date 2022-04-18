package repository

import (
    "ultimate_timer/domain/model"
)

type TimerUnitRepository interface {
    Create(timerUnit *model.TimerUnit) (*model.TimerUnit, error)
    FindByPresetID(id string) ([]*model.TimerUnit, error)
    FindByID(id string) (*model.TimerUnit, error)
    Update(timerUnit *model.TimerUnit) (*model.TimerUnit, error)
    Delete(timerUnit *model.TimerUnit) error
}
