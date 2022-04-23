package repository

import (
    "ultimate_timer/domain/model"
)

type PresetRepository interface {
    Create(preset *model.Preset) (*model.Preset, error)
    Get() ([]*model.Preset, error)
    FindByID(id string) (*model.Preset, error)
    Update(preset *model.Preset) (*model.Preset, error)
    Delete(preset *model.Preset) error

    GetCacheById(id string) (*model.Preset, error)
    SetCache(preset *model.Preset) error
}
