package day19

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Test_Scanner_RotateAroundXAxis(t *testing.T) {
	scanner := NewScanner(0, []*common.Point{
		common.New3DPoint(1, 2, 3),
	})

	rotatedScanner := scanner.RotateX()

	if rotatedScanner == nil {
		t.Error("rotatedScanner is nil.")
	}

	if rotatedScanner.Beacons()[0].X() != 1 {
		t.Errorf("Expected 1 but got %d.", rotatedScanner.Beacons()[0].X())
	}

	if rotatedScanner.Beacons()[0].Y() != 3 {
		t.Errorf("Expected 3 but got %d.", rotatedScanner.Beacons()[0].Y())
	}

	if rotatedScanner.Beacons()[0].Z() != -2 {
		t.Errorf("Expected -2 but got %d.", rotatedScanner.Beacons()[0].Z())
	}
}

func Test_Scanner_RotateAroundYAxis(t *testing.T) {
	scanner := NewScanner(0, []*common.Point{
		common.New3DPoint(1, 2, 3),
	})

	rotatedScanner := scanner.RotateY()

	if rotatedScanner == nil {
		t.Error("rotatedScanner is nil.")
	}

	if rotatedScanner.Beacons()[0].X() != -3 {
		t.Errorf("Expected -3 but got %d.", rotatedScanner.Beacons()[0].X())
	}

	if rotatedScanner.Beacons()[0].Y() != 2 {
		t.Errorf("Expected 2 but got %d.", rotatedScanner.Beacons()[0].Y())
	}

	if rotatedScanner.Beacons()[0].Z() != 1 {
		t.Errorf("Expected 1 but got %d.", rotatedScanner.Beacons()[0].Z())
	}
}

func Test_Scanner_RotateAroundZAxis(t *testing.T) {
	scanner := NewScanner(0, []*common.Point{
		common.New3DPoint(1, 2, 3),
	})

	rotatedScanner := scanner.RotateZ()

	if rotatedScanner == nil {
		t.Error("rotatedScanner is nil.")
	}

	if rotatedScanner.Beacons()[0].X() != 2 {
		t.Errorf("Expected 2 but got %d.", rotatedScanner.Beacons()[0].X())
	}

	if rotatedScanner.Beacons()[0].Y() != -1 {
		t.Errorf("Expected -1 but got %d.", rotatedScanner.Beacons()[0].Y())
	}

	if rotatedScanner.Beacons()[0].Z() != 3 {
		t.Errorf("Expected 3 but got %d.", rotatedScanner.Beacons()[0].Z())
	}
}

func Test_Scanner_GetOverlappingBeacons(t *testing.T) {
	scanner0 := NewScanner(0, []*common.Point{
		common.New3DPoint(404, -588, -901),
		common.New3DPoint(528, -643, 409),
		common.New3DPoint(-838, 591, 734),
		common.New3DPoint(390, -675, -793),
		common.New3DPoint(-537, -823, -458),
		common.New3DPoint(-485, -357, 347),
		common.New3DPoint(-345, -311, 381),
		common.New3DPoint(-661, -816, -575),
		common.New3DPoint(-876, 649, 763),
		common.New3DPoint(-618, -824, -621),
		common.New3DPoint(553, 345, -567),
		common.New3DPoint(474, 580, 667),
		common.New3DPoint(-447, -329, 318),
		common.New3DPoint(-584, 868, -557),
		common.New3DPoint(544, -627, -890),
		common.New3DPoint(564, 392, -477),
		common.New3DPoint(455, 729, 728),
		common.New3DPoint(-892, 524, 684),
		common.New3DPoint(-689, 845, -530),
		common.New3DPoint(423, -701, 434),
		common.New3DPoint(7, -33, -71),
		common.New3DPoint(630, 319, -379),
		common.New3DPoint(443, 580, 662),
		common.New3DPoint(-789, 900, -551),
		common.New3DPoint(459, -707, 401),
	})

	scanner1 := NewScanner(1, []*common.Point{
		common.New3DPoint(686, 422, 578),
		common.New3DPoint(605, 423, 415),
		common.New3DPoint(515, 917, -361),
		common.New3DPoint(-336, 658, 858),
		common.New3DPoint(95, 138, 22),
		common.New3DPoint(-476, 619, 847),
		common.New3DPoint(-340, -569, -846),
		common.New3DPoint(567, -361, 727),
		common.New3DPoint(-460, 603, -452),
		common.New3DPoint(669, -402, 600),
		common.New3DPoint(729, 430, 532),
		common.New3DPoint(-500, -761, 534),
		common.New3DPoint(-322, 571, 750),
		common.New3DPoint(-466, -666, -811),
		common.New3DPoint(-429, -592, 574),
		common.New3DPoint(-355, 545, -477),
		common.New3DPoint(703, -491, -529),
		common.New3DPoint(-328, -685, 520),
		common.New3DPoint(413, 935, -424),
		common.New3DPoint(-391, 539, -444),
		common.New3DPoint(586, -435, 557),
		common.New3DPoint(-364, -763, -893),
		common.New3DPoint(807, -499, -711),
		common.New3DPoint(755, -354, -619),
		common.New3DPoint(553, 889, -390),
	})

	overlaps, difference := scanner0.DetectBeaconOverlap(scanner1)

	if overlaps {
		t.Errorf("Expected false but got true.")
	}

	if difference != nil {
		t.Errorf("Expected nil but got %s.", difference)
	}

	scanner0 = scanner0.RotateY().RotateY()
	overlaps, difference = scanner0.DetectBeaconOverlap(scanner1)

	if !overlaps {
		t.Error("Expected true but got false.")
	} else if *difference == *common.New3DPoint(68, -1246, -43) {
		t.Error("difference is  nil.")
	}
}

func Test_Scanner_Merge(t *testing.T) {
	scanner0 := NewScanner(0, []*common.Point{
		common.New3DPoint(404, -588, -901),
		common.New3DPoint(528, -643, 409),
		common.New3DPoint(-838, 591, 734),
		common.New3DPoint(390, -675, -793),
		common.New3DPoint(-537, -823, -458),
		common.New3DPoint(-485, -357, 347),
		common.New3DPoint(-345, -311, 381),
		common.New3DPoint(-661, -816, -575),
		common.New3DPoint(-876, 649, 763),
		common.New3DPoint(-618, -824, -621),
		common.New3DPoint(553, 345, -567),
		common.New3DPoint(474, 580, 667),
		common.New3DPoint(-447, -329, 318),
		common.New3DPoint(-584, 868, -557),
		common.New3DPoint(544, -627, -890),
		common.New3DPoint(564, 392, -477),
		common.New3DPoint(455, 729, 728),
		common.New3DPoint(-892, 524, 684),
		common.New3DPoint(-689, 845, -530),
		common.New3DPoint(423, -701, 434),
		common.New3DPoint(7, -33, -71),
		common.New3DPoint(630, 319, -379),
		common.New3DPoint(443, 580, 662),
		common.New3DPoint(-789, 900, -551),
		common.New3DPoint(459, -707, 401),
	}).RotateY().RotateY()

	scanner1 := NewScanner(1, []*common.Point{
		common.New3DPoint(686, 422, 578),
		common.New3DPoint(605, 423, 415),
		common.New3DPoint(515, 917, -361),
		common.New3DPoint(-336, 658, 858),
		common.New3DPoint(95, 138, 22),
		common.New3DPoint(-476, 619, 847),
		common.New3DPoint(-340, -569, -846),
		common.New3DPoint(567, -361, 727),
		common.New3DPoint(-460, 603, -452),
		common.New3DPoint(669, -402, 600),
		common.New3DPoint(729, 430, 532),
		common.New3DPoint(-500, -761, 534),
		common.New3DPoint(-322, 571, 750),
		common.New3DPoint(-466, -666, -811),
		common.New3DPoint(-429, -592, 574),
		common.New3DPoint(-355, 545, -477),
		common.New3DPoint(703, -491, -529),
		common.New3DPoint(-328, -685, 520),
		common.New3DPoint(413, 935, -424),
		common.New3DPoint(-391, 539, -444),
		common.New3DPoint(586, -435, 557),
		common.New3DPoint(-364, -763, -893),
		common.New3DPoint(807, -499, -711),
		common.New3DPoint(755, -354, -619),
		common.New3DPoint(553, 889, -390),
	})

	overlaps, difference := scanner0.DetectBeaconOverlap(scanner1)
	merged := scanner0.Merge(scanner1, difference)

	if !overlaps {
		t.Error("Expected true but got false.")
	}

	if len(merged.Beacons()) != 38 {
		t.Errorf("Expected 38 but got %d.", len(merged.Beacons()))
	}

	if *merged.Beacons()[0] != *scanner0.Beacons()[0] {
		t.Error("Unexpected beacon.")
	}

	if *merged.Beacons()[25] != *scanner1.Beacons()[4].Move(difference) {
		t.Error("Unexpected beacon.")
	}
}
