package day15

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Test_Solve_Test1(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       15,
		InputFile: "test1.txt",
	}

	expectedPart1 := "2"
	expectedPart2 := "Not implemented."

	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test2(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       15,
		InputFile: "test2.txt",
	}

	expectedPart1 := "10"
	expectedPart2 := "Not implemented."

	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test3(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       15,
		InputFile: "test3.txt",
	}

	expectedPart1 := "17"
	expectedPart2 := "Not implemented."

	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}

/* func Test_Solve_Test4(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       15,
		InputFile: "test4.txt",
	}

	expectedPart1 := "40"
	expectedPart2 := "Not implemented."

	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}
*/
