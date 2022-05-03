package usecase

import (
	"encoding/json"
	"ultimate_timer/domain/model"
	"ultimate_timer/domain/repository"
)

type PresetUsecase interface {
	Create(
		name string,
		displayOrder, loopCount int,
		waitsConfirmEach, waitsConfirmLast bool,
		timerUnits json.RawMessage) (*model.Preset, error)
	Get() ([]*model.Preset, error)
	FindByID(id string) (*model.Preset, error)
	Update(
		id string,
		name string,
		displayOrder, loopCount int,
		waitsConfirmEach, waitsConfirmLast bool,
		timerUnits json.RawMessage) (*model.Preset, error)
	Delete(id string) error
}

type presetUsecase struct {
	presetRepo repository.PresetRepository
}

func NewPresetUsecase(presetRepo repository.PresetRepository) PresetUsecase {
	return &presetUsecase{presetRepo: presetRepo}
}

func (pr *presetUsecase) Create(
	name string,
	displayOrder, loopCount int,
	waitsConfirmEach, waitsConfirmLast bool,
	timerUnits json.RawMessage) (*model.Preset, error) {

	preset, err := model.NewPreset(
		name,
		displayOrder,
		loopCount,
		waitsConfirmEach,
		waitsConfirmLast,
		timerUnits,
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

func (pr *presetUsecase) FindByID(id string) (*model.Preset, error) {
	preset, err := pr.presetRepo.FindByID(id)
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

func (pr *presetUsecase) Update(
	id, name string,
	displayOrder, loopCount int,
	waitsConfirmEach, waitsConfirmLast bool,
	timerUnits json.RawMessage) (*model.Preset, error) {

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
