package repository

import (
    "ultimate_timer/domain/model"
)

type TimerPresetRepository interface {
    Create(preset *model.TimerPreset) (*model.TimerPreset, error)
    Get() ([]*model.TimerPreset, error)
    FindByID(id string) (*model.TimerPreset, error)
    Update(preset *model.TimerPreset) (*model.TimerPreset, error)
    Delete(preset *model.TimerPreset) error

    GetCacheById(id string) (*model.TimerPreset, error)
    SetCache(preset *model.TimerPreset) error
}
