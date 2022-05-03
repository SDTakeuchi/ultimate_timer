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

type presetTableStruct struct {
	name   string
	id    string
	preset model.Preset
	err    error
}

var testString1 string = `[{"order": 1, "duration": 60}, {"order": 2, "duration": 120}]`
var json1, _ = json.Marshal([]byte(testString1))
var id1, _ = uuid.Parse("f9b1303e-76e6-4071-8fb0-0599a6247376")

var testString2 string = `[{"order": 1, "duration": 540}, {"order": 2, "duration": 30}, {"order": 3, "duration": 630}]`
var json2, _ = json.Marshal([]byte(testString2))
var id2, _ = uuid.Parse("b0987480-5d8b-4e2e-940b-dafc75c2a5a0")

var (
	presetTable = []presetTableStruct {
		{
			"test_1",
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
		{
			"test_2",
			"b0987480-5d8b-4e2e-940b-dafc75c2a5a0",
			model.Preset{
				BaseModel: model.BaseModel{
					ID:        id2,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Name: "second_test_preset",
				DisplayOrder: 2,
				LoopCount: 100,
				WaitsConfirmEach: true,
				WaitsConfirmLast: true,
				TimerUnits: json2,
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
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	pt := presetTableStruct {
		name: "",
		id: id.String(),
		preset: *preset,
		err: nil,
	}
	presetTable = append(presetTable, pt)
	return preset, nil
}

func (tpu *testPresetUsecase) Update(preset *model.Preset) (*model.Preset, error) {
	for i, p := range presetTable {
		if p.preset.ID == preset.ID {
			(&presetTable[i]).preset = *preset
			return &presetTable[i].preset, nil
		}
	}
	return nil, errors.New("not found")
}

func (tpu *testPresetUsecase) Delete(preset *model.Preset) error {
	for i, p := range presetTable {
		if p.preset.ID == preset.ID {
			presetTable[i] = presetTable[len(presetTable)-1]
			presetTable = presetTable[:len(presetTable)-1]
			return nil
		}
	}
	return errors.New("not found")
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
			p, err := usecase_.FindByID(pt.id)

			assert.Equal(t, pt.err, err)
			assert.Equal(t, pt.preset.Name, p.Name)
		})
	}
}

func TestGet(t *testing.T) {
	mockPresetRepo := new(testPresetUsecase)
	usecase_ := usecase.NewPresetUsecase(mockPresetRepo)

	for _, pt := range presetTable {
		t.Run(pt.name, func(t *testing.T) {
			p, err := usecase_.Get()

			assert.Equal(t, pt.err, err)
			assert.Equal(t, len(presetTable), len(p))
		})
	}
}

func TestCreate(t *testing.T) {
	mockPresetRepo := new(testPresetUsecase)
	usecase_ := usecase.NewPresetUsecase(mockPresetRepo)

	for _, pt := range presetTable {
		t.Run(pt.name, func(t *testing.T) {
			p, err := usecase_.Create(
				pt.preset.Name,
				pt.preset.DisplayOrder,
				pt.preset.LoopCount,
				pt.preset.WaitsConfirmEach,
				pt.preset.WaitsConfirmLast,
				pt.preset.TimerUnits,
			)

			assert.Equal(t, pt.err, err)
			assert.Equal(t, pt.preset.TimerUnits, p.TimerUnits)
		})
	}
}

func TestUpdate(t *testing.T) {
	mockPresetRepo := new(testPresetUsecase)
	usecase_ := usecase.NewPresetUsecase(mockPresetRepo)
	testLoopCount := 1234

	for _, pt := range presetTable {
		t.Run(pt.name, func(t *testing.T) {
			p, err := usecase_.Update(
				pt.preset.ID.String(),
				pt.preset.Name,
				pt.preset.DisplayOrder,
				testLoopCount,
				pt.preset.WaitsConfirmEach,
				pt.preset.WaitsConfirmLast,
				pt.preset.TimerUnits,
			)
			assert.Equal(t, pt.err, err)

			pUpdated, err := usecase_.FindByID(p.ID.String())
			assert.Equal(t, pt.err, err)
			assert.Equal(t, testLoopCount, pUpdated.LoopCount)
		})
	}
}

func TestDelete(t *testing.T) {
	mockPresetRepo := new(testPresetUsecase)
	usecase_ := usecase.NewPresetUsecase(mockPresetRepo)

	for _, pt := range presetTable {
		t.Run(pt.name, func(t *testing.T) {
			err := usecase_.Delete(pt.preset.ID.String())
			assert.Equal(t, pt.err, err)
		})
	}
}
