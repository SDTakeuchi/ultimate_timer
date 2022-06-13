package infra

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"ultimate_timer/domain/model"
	"ultimate_timer/domain/repository"
	"github.com/google/uuid"
)

type TimerPresetRepository struct {
	Conn  *gorm.DB
	Cache *redis.Client
}

func NewTimerPresetRepository(conn *gorm.DB, cache *redis.Client) repository.TimerPresetRepository {
	return &TimerPresetRepository{
		Conn:  conn,
		Cache: cache,
	}
}

func (pr *TimerPresetRepository) Create(timerPreset *model.TimerPreset) (*model.TimerPreset, error) {
	timerPresets, err := pr.Get()
	if err != nil {
		return nil, err
	}
	var savedNames []string
	for _, p := range timerPresets {
		savedNames = append(savedNames, p.Name)
	}
	CleanName(savedNames, timerPreset)

	if err := pr.Conn.Create(&timerPreset).Error; err != nil {
		return nil, err
	}
	go pr.SetCache(timerPreset)

	return timerPreset, nil
}

func (pr *TimerPresetRepository) Get() (timerPresets []*model.TimerPreset, err error) {
	if err := pr.Conn.Find(&timerPresets).Error; err != nil {
		return nil, err
	}

	return timerPresets, nil
}

func (pr *TimerPresetRepository) FindByID(id string) (*model.TimerPreset, error) {
	timerPreset, err := pr.GetCacheById(id)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if err == redis.Nil {
		// query to Postgres
		idUuid, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}
		timerPreset = &model.TimerPreset{BaseModel: model.BaseModel{ID: idUuid}}
		if err = pr.Conn.Take(&timerPreset).Error; err != nil {
			return nil, err
		}
		go pr.SetCache(timerPreset)
	}

	return timerPreset, nil
}

func (pr *TimerPresetRepository) Update(timerPreset *model.TimerPreset) (*model.TimerPreset, error) {
	timerPresets, err := pr.Get()
	if err != nil {
		return nil, err
	}
	var savedNames []string
	for _, p := range timerPresets {
		if timerPreset.ID != p.ID {
			savedNames = append(savedNames, p.Name)
		}
	}
	CleanName(savedNames, timerPreset)

	if err := pr.Conn.Model(&timerPreset).Updates(&timerPreset).Error; err != nil {
		return nil, err
	}
	go pr.SetCache(timerPreset)

	return timerPreset, nil
}

func (pr *TimerPresetRepository) Delete(timerPreset *model.TimerPreset) error {
	if err := pr.Conn.Delete(&timerPreset).Error; err != nil {
		return err
	}

	return nil
}
