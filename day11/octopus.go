package day11

import "math"

type octopus struct {
	energyLevel       int
	lastFlashedOnStep int
}

func NewOctopus(energylevel int) *octopus {
	return &octopus{
		energyLevel:       energylevel,
		lastFlashedOnStep: math.MinInt,
	}
}

func (octopus *octopus) Charge() {
	octopus.energyLevel++
}

func (octopus *octopus) FlashIfAble(step int) bool {
	if octopus.energyLevel > 9 && step > octopus.lastFlashedOnStep {
		octopus.lastFlashedOnStep = step
		return true
	} else {
		return false
	}
}

func (octopus *octopus) IsFlashing(step int) bool {
	return octopus.energyLevel > 9 && step == octopus.lastFlashedOnStep
}

func (octopus *octopus) StopFlashing() {
	if octopus.energyLevel > 9 {
		octopus.energyLevel = 0
	}
}
