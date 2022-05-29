package infra

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
	"ultimate_timer/domain/model"
	"ultimate_timer/services"
)

func (pr *PresetRepository) GetCacheById(id string) (*model.Preset, error) {
	p, err := pr.Cache.Get(context.Background(), id).Result()
	if err == redis.Nil {
		return nil, redis.Nil
	} else if err != nil {
		services.Logf("%v", err)
		return nil, err
	}
	preset := &model.Preset{}
	err = json.Unmarshal([]byte(p), preset)
	if err != nil {
		services.Logf("%v", err)
		return nil, err
	}
	return preset, nil
}

func (pr *PresetRepository) SetCache(preset *model.Preset) error {
	p, err := json.Marshal(preset)
	if err != nil {
		services.Logf("%v", err)
		return err
	}
	err = pr.Cache.Set(context.Background(), preset.ID.String(), p, time.Hour*24).Err()
	if err != nil {
		services.Logf("%v", err)
		return err
	}
	return nil
}
