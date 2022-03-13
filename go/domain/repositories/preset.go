package repositories

import (
    "ultimate_timer/domain/model"
)

// TaskRepository task repositoryのinterface
type PresetRepository interface {
    Create(preset *model.Preset) (*model.Preset, error)
    FindByID(id int) (*model.Preset, error)
    Update(preset *model.Preset) (*model.Preset, error)
    Delete(preset *model.Preset) error
}
