package usecase

import (
	"encoding/json"
	"ultimate_timer/domain/model"
	"ultimate_timer/domain/repository"
)

type TimerPresetUsecase interface {
	Create(
		name string,
		displayOrder, loopCount int,
		waitsConfirmEach, waitsConfirmLast bool,
		timerUnits json.RawMessage) (*model.TimerPreset, error)
	Get() ([]*model.TimerPreset, error)
	FindByID(id string) (*model.TimerPreset, error)
	Update(
		id string,
		name string,
		displayOrder, loopCount int,
		waitsConfirmEach, waitsConfirmLast bool,
		timerUnits json.RawMessage) (*model.TimerPreset, error)
	Delete(id string) error
}

type timerPresetUsecase struct {
	timerPresetRepo repository.TimerPresetRepository
}

func NewTimerPresetUsecase(timerPresetRepo repository.TimerPresetRepository) TimerPresetUsecase {
	return &timerPresetUsecase{timerPresetRepo: timerPresetRepo}
}

func (pr *timerPresetUsecase) Create(
	name string,
	displayOrder, loopCount int,
	waitsConfirmEach, waitsConfirmLast bool,
	timerUnits json.RawMessage) (*model.TimerPreset, error) {

	timerPreset, err := model.NewTimerPreset(
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

	createdTimerPreset, err := pr.timerPresetRepo.Create(timerPreset)
	if err != nil {
		return nil, err
	}

	return createdTimerPreset, nil
}

func (pr *timerPresetUsecase) FindByID(id string) (*model.TimerPreset, error) {
	timerPreset, err := pr.timerPresetRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return timerPreset, nil
}

func (pr *timerPresetUsecase) Get() ([]*model.TimerPreset, error) {
	timerPresets, err := pr.timerPresetRepo.Get()
	if err != nil {
		return nil, err
	}

	return timerPresets, nil
}

func (pr *timerPresetUsecase) Update(
	id, name string,
	displayOrder, loopCount int,
	waitsConfirmEach, waitsConfirmLast bool,
	timerUnits json.RawMessage) (*model.TimerPreset, error) {

	timerPreset, err := pr.timerPresetRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = timerPreset.Set(name, displayOrder, loopCount, waitsConfirmEach, waitsConfirmLast, timerUnits)
	if err != nil {
		return nil, err
	}

	updatedTimerPreset, err := pr.timerPresetRepo.Update(timerPreset)
	if err != nil {
		return nil, err
	}

	return updatedTimerPreset, nil
}

func (pr *timerPresetUsecase) Delete(id string) error {
	timerPreset, err := pr.timerPresetRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = pr.timerPresetRepo.Delete(timerPreset)
	if err != nil {
		return err
	}

	return nil
}
