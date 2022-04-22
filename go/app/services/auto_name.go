package services

import (
	"fmt"
	"strings"
	"ultimate_timer/domain/model"
)

func AutoName(presets []model.Preset) (presetName string) {
	maxTimerCount := 999
	var names []string
	for _, preset := range presets {
		if strings.Contains("タイマー ", preset.Name) {
			names = append(names, preset.Name)
		}
	}
	// check possible smallest number to give to the new preset
	var n string
	var timer interface{}
	for i := 1; i < maxTimerCount; i++ {
		n = fmt.Sprintf("タイマー %d", i)
		// TODO no such function as GetByName
		// timer = usecase.GetByName(n)
		if timer == nil {
			return n
		}
	}
}
