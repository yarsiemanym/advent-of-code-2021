package day15

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Test_CaveMap_GetRiskLevelAt(t *testing.T) {
	caveMap := NewCaveMapFromValues([][]int{
		{1, 2},
		{3, 4},
	})
	riskLevel := caveMap.GetRiskLevelAt(common.NewPoint(0, 1))

	if riskLevel != 3 {
		t.Errorf("Expected 3 but got %d.", riskLevel)
	}
}

func Test_CaveMap_GetRiskLevelOf(t *testing.T) {
	caveMap := NewCaveMapFromValues([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{4, 3, 2, 1},
		{8, 7, 6, 5},
	})

	path := NewPath()
	path.Append(common.NewPoint(0, 0))
	path.Append(common.NewPoint(0, 1))
	path.Append(common.NewPoint(0, 2))
	path.Append(common.NewPoint(1, 2))
	path.Append(common.NewPoint(2, 2))
	path.Append(common.NewPoint(3, 2))
	path.Append(common.NewPoint(3, 3))

	riskLevel := caveMap.GetRiskLevelOf(path)

	if riskLevel != 20 {
		t.Errorf("Expected 20 but got %d.", riskLevel)
	}
}

func Test_CaveMap_Expand(t *testing.T) {
	caveMap := NewCaveMapFromValues([][]int{
		{6, 7},
		{8, 9},
	})

	expandedCaveMap := caveMap.Expand(2)

	// 6 7 7 8
	// 8 9 9 1
	// 7 8 8 9
	// 9 1 1 2

	if expandedCaveMap == nil {
		t.Error("expandedCaveMap is nil.")
	}

	riskLevel := expandedCaveMap.GetRiskLevelAt(common.NewPoint(0, 0))
	if riskLevel != 6 {
		t.Errorf("Exepcted 6 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(1, 0))
	if riskLevel != 7 {
		t.Errorf("Exepcted 7 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(2, 0))
	if riskLevel != 7 {
		t.Errorf("Exepcted 7 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(3, 0))
	if riskLevel != 8 {
		t.Errorf("Exepcted 8 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(0, 1))
	if riskLevel != 8 {
		t.Errorf("Exepcted 8 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(1, 1))
	if riskLevel != 9 {
		t.Errorf("Exepcted 9 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(2, 1))
	if riskLevel != 9 {
		t.Errorf("Exepcted 9 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(3, 1))
	if riskLevel != 1 {
		t.Errorf("Exepcted 1 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(0, 2))
	if riskLevel != 7 {
		t.Errorf("Exepcted 7 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(1, 2))
	if riskLevel != 8 {
		t.Errorf("Exepcted 8 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(2, 2))
	if riskLevel != 8 {
		t.Errorf("Exepcted 1 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(3, 2))
	if riskLevel != 9 {
		t.Errorf("Exepcted 9 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(0, 3))
	if riskLevel != 9 {
		t.Errorf("Exepcted 9 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(1, 3))
	if riskLevel != 1 {
		t.Errorf("Exepcted 1 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(2, 3))
	if riskLevel != 1 {
		t.Errorf("Exepcted 1 but got %d.", riskLevel)
	}

	riskLevel = expandedCaveMap.GetRiskLevelAt(common.NewPoint(3, 3))
	if riskLevel != 2 {
		t.Errorf("Exepcted 2 but got %d.", riskLevel)
	}
}
