package usecase_test

import (
	"encoding/json"
	"errors"
	"testing"
	"time"
	"ultimate_timer/domain/model"
	"ultimate_timer/domain/repository"
	"ultimate_timer/usecase"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// presetRepoインターフェースを満たすtestPresetUsecaseを作る
type testPresetUsecase struct {
	repository.PresetRepository
}

var testString string = `[{"order": 1, "duration": 60}, {"order": 2, "duration": 120}]`
var bytes = []byte(testString)
var json1, _ = json.Marshal(bytes)

var id1, _ = uuid.Parse("f9b1303e-76e6-4071-8fb0-0599a6247376")

var (
	presetTable = []struct {
		name   string
		arg    string
		preset model.Preset
		err    error
	}{
		{
			"success",
			"f9b1303e-76e6-4071-8fb0-0599a6247376",
			model.Preset{
				BaseModel: model.BaseModel{
					ID:        id1,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Name: "test_name",
				DisplayOrder: 1,
				LoopCount: 2,
				WaitsConfirmEach: false,
				WaitsConfirmLast: true,
				TimerUnits: json1,
			},
			nil,
		},
	}
)

// FindByIdメソッドの実装を定義
func (tpu *testPresetUsecase) FindByID(id string) (*model.Preset, error) {
	for _, t := range presetTable {
		if t.preset.ID.String() == id {
			return &t.preset, nil
		}
	}
	return &model.Preset{}, errors.New("not found")
}

func (tpu *testPresetUsecase) Get() ([]*model.Preset, error) {
	var presets []*model.Preset
	for _, t := range presetTable {
		presets = append(presets, &t.preset)
	}

	return presets, nil
}

func (tpu *testPresetUsecase) Create(preset *model.Preset) (*model.Preset, error) {
	return preset , nil
}

func (tpu *testPresetUsecase) Update(preset *model.Preset) (*model.Preset, error) {
	return preset, nil
}

func (tpu *testPresetUsecase) Delete(preset *model.Preset) error {
	return nil
}

func (tpu *testPresetUsecase) GetCacheById(id string) (*model.Preset, error) {
	return &model.Preset{}, nil
}

func (tpu *testPresetUsecase) SetCache(preset *model.Preset) error {
	return nil
}

// test
func TestFindByID(t *testing.T) {
	mockPresetRepo := new(testPresetUsecase)
	usecase_ := usecase.NewPresetUsecase(mockPresetRepo)

	for _, pt := range presetTable {
		t.Run(pt.name, func(t *testing.T) {
			p, err := usecase_.FindByID(pt.arg)

			assert.Equal(t, pt.err, err)
			assert.Equal(t, pt.preset.Name, p.Name)
		})
	}
}
