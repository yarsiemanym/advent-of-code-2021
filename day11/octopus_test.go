package day11

import "testing"

func Test_octopus_Charge(t *testing.T) {
	octopus := &octopus{
		energyLevel: 0,
	}

	octopus.Charge()

	if octopus.energyLevel != 1 {
		t.Errorf("Expected 1 but got %v.", octopus.energyLevel)
	}
}

func Test_octopus_FlashIfAble_True(t *testing.T) {
	octopus := &octopus{
		energyLevel:       10,
		lastFlashedOnStep: 0,
	}

	flashed := octopus.FlashIfAble(1)

	if !flashed {
		t.Error("Expected true but got false.")
	}

	if octopus.energyLevel != 10 {
		t.Errorf("Expected 10 but got %v.", octopus.energyLevel)
	}
}

func Test_octopus_FlashIfAble_NotEnoughEnergy(t *testing.T) {
	octopus := &octopus{
		energyLevel:       9,
		lastFlashedOnStep: 0,
	}

	flashed := octopus.FlashIfAble(1)

	if flashed {
		t.Error("Expected false but got true.")
	}

	if octopus.energyLevel != 9 {
		t.Errorf("Expected 9 but got %v.", octopus.energyLevel)
	}
}

func Test_octopus_FlashIfAble_AlreadyFlashed(t *testing.T) {
	octopus := &octopus{
		energyLevel:       10,
		lastFlashedOnStep: 1,
	}

	flashed := octopus.FlashIfAble(1)

	if flashed {
		t.Error("Expected false but got true.")
	}

	if octopus.energyLevel != 10 {
		t.Errorf("Expected 10 but got %v.", octopus.energyLevel)
	}
}

func Test_octopus_IsFlashing_True(t *testing.T) {
	octopus := &octopus{
		energyLevel:       10,
		lastFlashedOnStep: 1,
	}

	isFlashing := octopus.IsFlashing(1)

	if !isFlashing {
		t.Error("Expected true but got false.")
	}

	if octopus.energyLevel != 10 {
		t.Errorf("Expected 1 but got %v.", octopus.energyLevel)
	}
}

func Test_octopus_IsFlashing_False(t *testing.T) {
	octopus := &octopus{
		energyLevel:       9,
		lastFlashedOnStep: 1,
	}

	isFlashing := octopus.IsFlashing(1)

	if isFlashing {
		t.Error("Expected false but got true.")
	}

	if octopus.energyLevel != 9 {
		t.Errorf("Expected 1 but got %v.", octopus.energyLevel)
	}
}

func Test_octopus_StopFlashing_WasFlashing(t *testing.T) {
	octopus := &octopus{
		energyLevel:       10,
		lastFlashedOnStep: 1,
	}

	octopus.StopFlashing()

	if octopus.energyLevel != 0 {
		t.Errorf("Expected 0 but got %v.", octopus.energyLevel)
	}
}

func Test_octopus_StopFlashing_WasNotFlashing(t *testing.T) {
	octopus := &octopus{
		energyLevel:       9,
		lastFlashedOnStep: 1,
	}

	octopus.StopFlashing()

	if octopus.energyLevel != 9 {
		t.Errorf("Expected 0 but got %v.", octopus.energyLevel)
	}
}
