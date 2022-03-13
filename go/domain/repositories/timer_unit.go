package repositories

import (
    "ultimate_timer/domain/model"
)

// TaskRepository task repositoryのinterface
type TimerUnitRepository interface {
    Create(timerUnit *model.TimerUnit) (*model.TimerUnit, error)
    FindByID(id int) (*model.TimerUnit, error)
    Update(timerUnit *model.TimerUnit) (*model.TimerUnit, error)
    Delete(timerUnit *model.TimerUnit) error
}
