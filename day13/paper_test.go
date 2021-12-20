package day13

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Test_Paper_GetMarkAt(t *testing.T) {
	paper := NewPaper(2, 1)
	point00 := common.New2DPoint(0, 0)
	point01 := common.New2DPoint(0, 1)
	paper.DrawMark(point01)

	mark := paper.GetMarkAt(point00)
	if mark != '.' {
		t.Errorf("Expected '.' but got '%c'.", mark)
	}

	mark = paper.GetMarkAt(point01)
	if mark != '#' {
		t.Errorf("Expected '#' but got '%c'.", mark)
	}
}

func Test_Paper_Fold_Up(t *testing.T) {
	paper := NewPaper(5, 5)
	paper.DrawMark(common.New2DPoint(1, 1))
	paper.DrawMark(common.New2DPoint(2, 1))
	paper.DrawMark(common.New2DPoint(2, 3))
	paper.DrawMark(common.New2DPoint(3, 3))
	crease := &Crease{
		Axis:     'y',
		Position: 2,
	}
	foldedPaper := paper.Fold(crease)

	if foldedPaper == nil {
		t.Error("foldedPaper is nil.")
	} else if foldedPaper.Height() != 2 {
		t.Errorf("Expected 2 but got %v.", foldedPaper.Height())
	} else if foldedPaper.Width() != 5 {
		t.Errorf("Expected 5 but got %v.", foldedPaper.Width())
	} else {
		for _, point := range foldedPaper.plane.GetAllPoints() {
			mark := foldedPaper.GetMarkAt(point)
			if point.X() == 1 && point.Y() == 1 {
				if mark != '#' {
					t.Errorf("Expected '#' but got '%c'.", mark)
				}
			} else if point.X() == 2 && point.Y() == 1 {
				if mark != '#' {
					t.Errorf("Expected '#' but got '%c'.", mark)
				}
			} else if point.X() == 3 && point.Y() == 1 {
				if mark != '#' {
					t.Errorf("Expected '#' but got '%c'.", mark)
				}
			} else if mark == '#' {
				t.Error("Expected '.' but got '#'.")
			}
		}
	}
}

func Test_Paper_Fold_Left(t *testing.T) {
	paper := NewPaper(5, 5)
	paper.DrawMark(common.New2DPoint(1, 1))
	paper.DrawMark(common.New2DPoint(1, 2))
	paper.DrawMark(common.New2DPoint(3, 2))
	paper.DrawMark(common.New2DPoint(3, 3))
	crease := &Crease{
		Axis:     'x',
		Position: 2,
	}
	foldedPaper := paper.Fold(crease)

	if foldedPaper == nil {
		t.Error("foldedPaper is nil.")
	} else if foldedPaper.Width() != 2 {
		t.Errorf("Expected 2 but got %v.", foldedPaper.Width())
	} else if foldedPaper.Height() != 5 {
		t.Errorf("Expected 5 but got %v.", foldedPaper.Height())
	} else {
		for _, point := range foldedPaper.plane.GetAllPoints() {
			mark := foldedPaper.GetMarkAt(point)
			if point.X() == 1 && point.Y() == 1 {
				if mark != '#' {
					t.Errorf("Expected '#' but got '%c'.", mark)
				}
			} else if point.X() == 1 && point.Y() == 2 {
				if mark != '#' {
					t.Errorf("Expected '#' but got '%c'.", mark)
				}
			} else if point.X() == 1 && point.Y() == 3 {
				if mark != '#' {
					t.Errorf("Expected '#' but got '%c'.", mark)
				}
			} else if mark == '#' {
				t.Error("Expected '.' but got '#'.")
			}
		}
	}
}
