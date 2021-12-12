package day12

import "testing"

func Test_Cave_IsBig_A(t *testing.T) {
	cave := NewCave("A")
	isBig := cave.IsBig()

	if !isBig {
		t.Error("Expected true but got false.")
	}
}

func Test_Cave_IsBig_Z(t *testing.T) {
	cave := NewCave("Z")
	isBig := cave.IsBig()

	if !isBig {
		t.Error("Expected true but got false.")
	}
}

func Test_Cave_IsBig_z(t *testing.T) {
	cave := NewCave("z")
	isBig := cave.IsBig()

	if isBig {
		t.Error("Expected false but got true.")
	}
}

func Test_Cave_GetExplorableConnectedCaves_DoNotRevisitSmallCaves(t *testing.T) {
	start := NewCave("start")
	caveA := NewCave("A")
	caveB := NewCave("B")
	caveC := NewCave("C")
	caveD := NewCave("d")
	caveE := NewCave("e")

	start.Connect(caveA)
	caveA.Connect(caveB)
	caveA.Connect(caveC)
	caveA.Connect(caveD)
	caveA.Connect(caveE)

	path := NewPath([]*Cave{start, caveB, caveD})

	explorableCaves := caveA.GetExplorableConnectedCaves(path, false)

	if len(explorableCaves) != 3 {
		t.Errorf("Expected 3 but got %v.", len(explorableCaves))
	}

	if explorableCaves[0].name != "B" {
		t.Errorf("Expected \"B\" but got %v.", explorableCaves[0].name)
	}

	if explorableCaves[1].name != "C" {
		t.Errorf("Expected \"C\" but got %v.", explorableCaves[1].name)
	}

	if explorableCaves[2].name != "e" {
		t.Errorf("Expected \"e\" but got %v.", explorableCaves[2].name)
	}
}

func Test_Cave_GetExplorableConnectedCaves_RevisitSmallCaves(t *testing.T) {
	start := NewCave("start")
	caveA := NewCave("A")
	caveB := NewCave("B")
	caveC := NewCave("C")
	caveD := NewCave("d")
	caveE := NewCave("e")

	start.Connect(caveA)
	caveA.Connect(caveB)
	caveA.Connect(caveC)
	caveA.Connect(caveD)
	caveA.Connect(caveE)

	path := NewPath([]*Cave{start, caveB, caveD})

	explorableCaves := caveA.GetExplorableConnectedCaves(path, true)

	if len(explorableCaves) != 4 {
		t.Errorf("Expected 4 but got %v.", len(explorableCaves))
	}

	if explorableCaves[0].name != "B" {
		t.Errorf("Expected \"B\" but got %v.", explorableCaves[0].name)
	}

	if explorableCaves[1].name != "C" {
		t.Errorf("Expected \"C\" but got %v.", explorableCaves[1].name)
	}

	if explorableCaves[2].name != "d" {
		t.Errorf("Expected \"d\" but got %v.", explorableCaves[2].name)
	}

	if explorableCaves[3].name != "e" {
		t.Errorf("Expected \"e\" but got %v.", explorableCaves[3].name)
	}
}
