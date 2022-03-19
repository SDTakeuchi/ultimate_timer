package services

import (
	"strconv"
	// "ultimate_timer/usecase"
)

func AutoName() string {
	maxTimerCount := 999
	var n string
	var timer interface{}
	for i := 1; i < maxTimerCount; i++ {
		n = "タイマー " + strconv.Itoa(i)
		// TODO no such function as GetByName
		// timer = usecase.GetByName(n)
		if timer == nil {
			return n
		}
	}
}
