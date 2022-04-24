package infra

import (
	"time"
	"ultimate_timer/domain/model"
)

func CleanName(savedNames []string, p *model.Preset) {
	if p.Name == "" {
		p.Name = "Timer"
	}
	nameIsDuplicate := contains[string](savedNames, p.Name)
	if nameIsDuplicate {
		timeNow := time.Now().Format("_20060102_150405")
		p.Name += timeNow
	}
	return
}

func contains[T comparable] (elems []T, v T) bool {
	for _, s := range elems {
		if v == s { return true }
	}
	return false
}
