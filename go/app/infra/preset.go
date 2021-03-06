package infra

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"ultimate_timer/domain/model"
	"ultimate_timer/domain/repository"
	"github.com/google/uuid"
)

type PresetRepository struct {
	Conn  *gorm.DB
	Cache *redis.Client
}

func NewPresetRepository(conn *gorm.DB, cache *redis.Client) repository.PresetRepository {
	return &PresetRepository{
		Conn:  conn,
		Cache: cache,
	}
}

func (pr *PresetRepository) Create(preset *model.Preset) (*model.Preset, error) {
	presets, err := pr.Get()
	if err != nil {
		return nil, err
	}
	var savedNames []string
	for _, p := range presets {
		savedNames = append(savedNames, p.Name)
	}
	CleanName(savedNames, preset)

	if err := pr.Conn.Create(&preset).Error; err != nil {
		return nil, err
	}
	go pr.SetCache(preset)

	return preset, nil
}

func (pr *PresetRepository) Get() (presets []*model.Preset, err error) {
	if err := pr.Conn.Find(&presets).Error; err != nil {
		return nil, err
	}

	return presets, nil
}

func (pr *PresetRepository) FindByID(id string) (*model.Preset, error) {
	preset, err := pr.GetCacheById(id)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if err == redis.Nil {
		// query to Postgres
		idUuid, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}
		preset = &model.Preset{BaseModel: model.BaseModel{ID: idUuid}}
		if err = pr.Conn.Take(&preset).Error; err != nil {
			return nil, err
		}
		go pr.SetCache(preset)
	}

	return preset, nil
}

func (pr *PresetRepository) Update(preset *model.Preset) (*model.Preset, error) {
	presets, err := pr.Get()
	if err != nil {
		return nil, err
	}
	var savedNames []string
	for _, p := range presets {
		if preset.ID != p.ID {
			savedNames = append(savedNames, p.Name)
		}
	}
	CleanName(savedNames, preset)

	if err := pr.Conn.Model(&preset).Updates(&preset).Error; err != nil {
		return nil, err
	}
	go pr.SetCache(preset)

	return preset, nil
}

func (pr *PresetRepository) Delete(preset *model.Preset) error {
	if err := pr.Conn.Delete(&preset).Error; err != nil {
		return err
	}

	return nil
}
