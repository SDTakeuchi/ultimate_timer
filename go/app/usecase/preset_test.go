package usecase_test

import (
	"encoding/json"
	"ultimate_timer/domain/repository"
	"ultimate_timer/domain/model"
	"ultimate_timer/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BlogRepositoryインターフェースを満たすfakeBlogRepositoryを作る
type testPresetUsecase struct {
	presetRepo repository.PresetRepository
}


type timerUnits struct {
	order int
	duration int
}

var (
	presetTable = []struct {
		id string
		name             string
		displayOrder     int
		loopCount        int
		waitsConfirmEach bool
		waitsConfirmLast bool
		timerUnits       []timerUnits
	} {
		{
			"f9b1303e-76e6-4071-8fb0-0599a6247376",
			"test_name",
			1,
			2,
			false,
			true,
			[]timerUnits{
				timerUnits{order: 1, duration: 60},
				timerUnits{order: 2, duration: 120},
			},
		},
	}
)

// FindByIdメソッドの実装を定義
func (tpu *testPresetUsecase) FindById(id string) (model.Preset, error) {
	for _, t := range presetTable {
		if t.id == id {
			return t.blog, nil
		}
	}
	return model.Preset{}, errors.New("not found")
}

// test
func TestDetailBlog(t *testing.T) {
	mockPresetRepo := new(testPresetUsecase)
	usecase_ := usecase.NewPresetUsecase(mockPresetRepo)

	for _, pt := range presetTable {
		t.Run(pt.name, func(t *testing.T) {
			p, err := usecase_.FindByID(pt.id)

			assert.Equal(t, pt.err, err)
			assert.Equal(t, pt.name, p.Name)
		})
	}
}
