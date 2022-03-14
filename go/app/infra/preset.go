package infra

import (
	"ultimate_timer/domein/model"
	"ultimate_timer/domein/repository"
	"github.com/jinzhu/gorm"
)

type PresetRepository struct {
	Conn *gorm.DB
}

func NewPresetRepository(conn *gorm.DB) repository.PresetRepository {
	return &PresetRepository{Conn: conn}
}

func (pr *PresetRepository) Create(preset *model.Preset) (*model.Preset, error) {
	if err := pr.Conn.NewRecord(&preset); err != nil {
		return nil, err
	}
	if err := pr.Conn.Create(&preset).Error; err != nil {
		return nil, err
	}

	return preset, nil
}

func (pr *PresetRepository) Get() (presets []*model.Preset, err error) {
	if err := pr.Conn.Find(&presets).Error; err != nil {
		return nil, err
	}

	return presets, nil
}

func (pr *PresetRepository) FindByID(id string) (*model.Preset, error) {
	preset := &model.Preset{ID: id}
	if err := pr.Conn.First(&preset).Related(&preset.TimerUnits).Error; err != nil {
		return nil, err
	}

	return preset, nil
}

func (pr *PresetRepository) Update(preset *model.Preset) (*model.Preset, error) {
	if err := pr.Conn.Model(&preset).Update(&preset).Error; err != nil {
		return nil, err
	}

	return preset, nil
}

func (pr *PresetRepository) Delete(preset *model.Preset) error {
	id := preset.ID
	if err := pr.Conn.Where("preset_id = ?", id).Delete(&preset.TimerUnit{}).Error; err != nil {
		return err
	}
	if err := pr.Conn.Delete(&preset).Error; err != nil {
		return err
	}

	return nil
}
