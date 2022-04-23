package infra

import (
	// "context"
	// "encoding/json"
	// "time"
	"github.com/jinzhu/gorm"
	"github.com/go-redis/redis/v8"
	"ultimate_timer/domain/model"
	"ultimate_timer/domain/repository"
)

type PresetRepository struct {
	Conn *gorm.DB
	Cache *redis.Client
}

func NewPresetRepository(conn *gorm.DB, cache *redis.Client) repository.PresetRepository {
	return &PresetRepository{Conn: conn, Cache: cache}
}

func (pr *PresetRepository) Create(preset *model.Preset) (*model.Preset, error) {
	go pr.SetCache(preset)
	if err := pr.Conn.Create(&preset).Error; err != nil {
		return nil, err
	}

	return preset, nil
}

func (pr *PresetRepository) Get() (presets []*model.Preset, err error) {
	if err := pr.Conn.Preload("TimerUnits").Find(&presets).Error; err != nil {
		return nil, err
	}

	return presets, nil
}

func (pr *PresetRepository) FindByID(id string) (*model.Preset, error) {
	// p, err := pr.Cache.Get(context.Background(), id).Result()
	// if err == redis.Nil {
	// 	preset := &model.Preset{BaseModel: model.BaseModel{ID: id}}
	// 	if err := pr.Conn.First(&preset).Related(&preset.TimerUnits).Error; err != nil {
	// 		return nil, err
	// 	}
	// 	p, err := json.Marshal(&preset)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	err = pr.Cache.Set(context.Background(), preset.ID, p, time.Hour*24).Err()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	
	// 	return preset, nil
	// } else if err != nil {
	// 	return nil, err
	// } else {
	// 	preset := &model.Preset{}
	// 	err = json.Unmarshal([]byte(p), preset)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return preset, nil
	// }
	preset, err := pr.GetCacheById(id)
	if err != nil && err != redis.Nil {
		return nil , err
	}
	if err == redis.Nil {
		preset = &model.Preset{BaseModel: model.BaseModel{ID: id}}
		if err = pr.Conn.First(&preset).Related(&preset.TimerUnits).Error; err != nil {
			return nil, err
		}
		go pr.SetCache(preset)
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
	if err := pr.Conn.Where("preset_id = ?", id).Delete(preset.TimerUnits).Error; err != nil {
		return err
	}
	if err := pr.Conn.Delete(&preset).Error; err != nil {
		return err
	}

	return nil
}
