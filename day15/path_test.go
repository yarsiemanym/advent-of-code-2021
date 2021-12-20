package day15

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Test_Path_Clone(t *testing.T) {
	path1 := NewPath()
	path1.Append(common.New2DPoint(0, 0))
	path2 := path1.Clone()
	path1.Append(common.New2DPoint(1, 1))

	if path2.Length() != 1 {
		t.Errorf("Expected 1 but got %v.", path2.Length())
	}

	if path2.Points()[0].X() != 0 {
		t.Errorf("Expected 0 but got %v.", path2.Length())
	}

	if path2.Points()[0].Y() != 0 {
		t.Errorf("Expected 0 but got %v.", path2.Length())
	}

	if path1.Length() != 2 {
		t.Errorf("Expected 2 but got %v.", path2.Length())
	}

	if path1.Points()[0].X() != 0 {
		t.Errorf("Expected 0 but got %v.", path2.Length())
	}

	if path1.Points()[0].Y() != 0 {
		t.Errorf("Expected 0 but got %v.", path2.Length())
	}

	if path1.Points()[1].X() != 1 {
		t.Errorf("Expected 1 but got %v.", path2.Length())
	}

	if path1.Points()[1].Y() != 1 {
		t.Errorf("Expected 1 but got %v.", path2.Length())
	}
}
