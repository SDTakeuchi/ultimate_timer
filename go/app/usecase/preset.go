package usecase

import (
    // "time"
	"ultimate_timer/domain/model"
	"ultimate_timer/domain/repository"
)

// interface of PresetUsecase preset usecase
type PresetUsecase interface {
	Create(
		name string,
		displayOrder, loopCount int,
		waitsConfirmEach, waitsConfirmLast bool,
		timerUnits []struct{}) (*model.Preset, error)
    Get() ([]*model.Preset, error)
	FindByID(id string) (*model.Preset, error)
	Update(
        id string,
		name string,
		displayOrder, loopCount int,
		waitsConfirmEach, waitsConfirmLast bool,
		timerUnits []model.TimerUnit) (*model.Preset, error)
	Delete(id string) error
}

type presetUsecase struct {
	presetRepo repository.PresetRepository
}

// NewPresetUsecase preset usecaseのコンストラクタ
func NewPresetUsecase(presetRepo repository.PresetRepository) PresetUsecase {
	return &presetUsecase{presetRepo: presetRepo}
}

// Create presetを保存するときのユースケース
func (pr *presetUsecase) Create(
	name string,
	displayOrder, loopCount int,
	waitsConfirmEach, waitsConfirmLast bool,
	timerUnits []struct{}) (*model.Preset, error) {

	newTu := []model.TimerUnit{}
	for tu := range timerUnits {
		newTu = append(newTu, model.TimerUnit{Duration: tu.Duration, Order: tu.Order})
	}

	preset, err := model.NewPreset(
		name,
		displayOrder,
		loopCount,
		waitsConfirmEach,
		waitsConfirmLast,
		newTu,
	)
	if err != nil {
		return nil, err
	}

	createdPreset, err := pr.presetRepo.Create(preset)
	if err != nil {
		return nil, err
	}

	return createdPreset, nil
}

// FindByID presetをIDで取得するときのユースケース
func (iu *presetUsecase) FindByID(id string) (*model.Preset, error) {
	preset, err := iu.presetRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return preset, nil
}

func (pr *presetUsecase) Get() ([]*model.Preset, error) {
	presets, err := pr.presetRepo.Get()
	if err != nil {
		return nil, err
	}

	return presets, nil
}

// Update presetを更新するときのユースケース
func (pr *presetUsecase) Update(
	id,	name string,
	displayOrder, loopCount int,
	waitsConfirmEach, waitsConfirmLast bool,
	timerUnits []model.TimerUnit) (*model.Preset, error) {
	preset, err := pr.presetRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = preset.Set(name, displayOrder, loopCount, waitsConfirmEach, waitsConfirmLast, timerUnits)
	if err != nil {
		return nil, err
	}

	updatedPreset, err := pr.presetRepo.Update(preset)
	if err != nil {
		return nil, err
	}

	return updatedPreset, nil
}

// Delete presetを削除するときのユースケース
func (pr *presetUsecase) Delete(id string) error {
	preset, err := pr.presetRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = pr.presetRepo.Delete(preset)
	if err != nil {
		return err
	}

	return nil
}
