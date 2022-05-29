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
		now := time.Now().Format("_20060102_150405")
		p.Name += now
	}
}

func contains[T comparable] (elems []T, elem T) bool {
	for _, s := range elems {
		if elem == s { return true }
	}
	return false
}
