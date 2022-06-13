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

// timerPresetRepoインターフェースを満たすtestTimerPresetUsecaseを作る
type testTimerPresetUsecase struct {
	repository.TimerPresetRepository
}

type timerPresetTableStruct struct {
	name   string
	id    string
	timerPreset model.TimerPreset
	err    error
}

var testString1 string = `[{"order": 1, "duration": 60}, {"order": 2, "duration": 120}]`
var json1, _ = json.Marshal([]byte(testString1))
var id1, _ = uuid.Parse("f9b1303e-76e6-4071-8fb0-0599a6247376")

var testString2 string = `[{"order": 1, "duration": 540}, {"order": 2, "duration": 30}, {"order": 3, "duration": 630}]`
var json2, _ = json.Marshal([]byte(testString2))
var id2, _ = uuid.Parse("b0987480-5d8b-4e2e-940b-dafc75c2a5a0")

var (
	timerPresetTable = []timerPresetTableStruct {
		{
			"test_1",
			"f9b1303e-76e6-4071-8fb0-0599a6247376",
			model.TimerPreset{
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
			model.TimerPreset{
				BaseModel: model.BaseModel{
					ID:        id2,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Name: "second_test_timerPreset",
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
func (tpu *testTimerPresetUsecase) FindByID(id string) (*model.TimerPreset, error) {
	for _, t := range timerPresetTable {
		if t.timerPreset.ID.String() == id {
			return &t.timerPreset, nil
		}
	}
	return &model.TimerPreset{}, errors.New("not found")
}

func (tpu *testTimerPresetUsecase) Get() ([]*model.TimerPreset, error) {
	var timerPresets []*model.TimerPreset
	for _, t := range timerPresetTable {
		timerPresets = append(timerPresets, &t.timerPreset)
	}

	return timerPresets, nil
}

func (tpu *testTimerPresetUsecase) Create(timerPreset *model.TimerPreset) (*model.TimerPreset, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	pt := timerPresetTableStruct {
		name: "",
		id: id.String(),
		timerPreset: *timerPreset,
		err: nil,
	}
	timerPresetTable = append(timerPresetTable, pt)
	return timerPreset, nil
}

func (tpu *testTimerPresetUsecase) Update(timerPreset *model.TimerPreset) (*model.TimerPreset, error) {
	for i, p := range timerPresetTable {
		if p.timerPreset.ID == timerPreset.ID {
			(&timerPresetTable[i]).timerPreset = *timerPreset
			return &timerPresetTable[i].timerPreset, nil
		}
	}
	return nil, errors.New("not found")
}

func (tpu *testTimerPresetUsecase) Delete(timerPreset *model.TimerPreset) error {
	for i, p := range timerPresetTable {
		if p.timerPreset.ID == timerPreset.ID {
			timerPresetTable[i] = timerPresetTable[len(timerPresetTable)-1]
			timerPresetTable = timerPresetTable[:len(timerPresetTable)-1]
			return nil
		}
	}
	return errors.New("not found")
}

func (tpu *testTimerPresetUsecase) GetCacheById(id string) (*model.TimerPreset, error) {
	return &model.TimerPreset{}, nil
}

func (tpu *testTimerPresetUsecase) SetCache(timerPreset *model.TimerPreset) error {
	return nil
}

// test
func TestFindByID(t *testing.T) {
	mockTimerPresetRepo := new(testTimerPresetUsecase)
	usecase_ := usecase.NewTimerPresetUsecase(mockTimerPresetRepo)

	for _, pt := range timerPresetTable {
		t.Run(pt.name, func(t *testing.T) {
			p, err := usecase_.FindByID(pt.id)

			assert.Equal(t, pt.err, err)
			assert.Equal(t, pt.timerPreset.Name, p.Name)
		})
	}
}

func TestGet(t *testing.T) {
	mockTimerPresetRepo := new(testTimerPresetUsecase)
	usecase_ := usecase.NewTimerPresetUsecase(mockTimerPresetRepo)

	for _, pt := range timerPresetTable {
		t.Run(pt.name, func(t *testing.T) {
			p, err := usecase_.Get()

			assert.Equal(t, pt.err, err)
			assert.Equal(t, len(timerPresetTable), len(p))
		})
	}
}

func TestCreate(t *testing.T) {
	mockTimerPresetRepo := new(testTimerPresetUsecase)
	usecase_ := usecase.NewTimerPresetUsecase(mockTimerPresetRepo)

	for _, pt := range timerPresetTable {
		t.Run(pt.name, func(t *testing.T) {
			p, err := usecase_.Create(
				pt.timerPreset.Name,
				pt.timerPreset.DisplayOrder,
				pt.timerPreset.LoopCount,
				pt.timerPreset.WaitsConfirmEach,
				pt.timerPreset.WaitsConfirmLast,
				pt.timerPreset.TimerUnits,
			)

			assert.Equal(t, pt.err, err)
			assert.Equal(t, pt.timerPreset.TimerUnits, p.TimerUnits)
		})
	}
}

func TestUpdate(t *testing.T) {
	mockTimerPresetRepo := new(testTimerPresetUsecase)
	usecase_ := usecase.NewTimerPresetUsecase(mockTimerPresetRepo)
	testLoopCount := 1234

	for _, pt := range timerPresetTable {
		t.Run(pt.name, func(t *testing.T) {
			p, err := usecase_.Update(
				pt.timerPreset.ID.String(),
				pt.timerPreset.Name,
				pt.timerPreset.DisplayOrder,
				testLoopCount,
				pt.timerPreset.WaitsConfirmEach,
				pt.timerPreset.WaitsConfirmLast,
				pt.timerPreset.TimerUnits,
			)
			assert.Equal(t, pt.err, err)

			pUpdated, err := usecase_.FindByID(p.ID.String())
			assert.Equal(t, pt.err, err)
			assert.Equal(t, testLoopCount, pUpdated.LoopCount)
		})
	}
}

func TestDelete(t *testing.T) {
	mockTimerPresetRepo := new(testTimerPresetUsecase)
	usecase_ := usecase.NewTimerPresetUsecase(mockTimerPresetRepo)

	for _, pt := range timerPresetTable {
		t.Run(pt.name, func(t *testing.T) {
			err := usecase_.Delete(pt.timerPreset.ID.String())
			assert.Equal(t, pt.err, err)
		})
	}
}
