package day16

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Test_Solve_Test01(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       16,
		InputFile: "test01.txt",
	}

	expectedPart1 := "16"

	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart1, answer.Part1)
	}
}

func Test_Solve_Test02(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       16,
		InputFile: "test02.txt",
	}

	expectedPart1 := "12"

	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart1, answer.Part1)
	}
}
func Test_Solve_Test03(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       16,
		InputFile: "test03.txt",
	}

	expectedPart1 := "23"

	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart1, answer.Part1)
	}
}
func Test_Solve_Test04(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       16,
		InputFile: "test04.txt",
	}

	expectedPart1 := "31"

	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart1, answer.Part1)
	}
}

func Test_Solve_Test05(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       16,
		InputFile: "test05.txt",
	}

	expectedPart2 := "3"

	answer := Solve(puzzle)

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test06(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       16,
		InputFile: "test06.txt",
	}

	expectedPart2 := "54"

	answer := Solve(puzzle)

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test07(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       16,
		InputFile: "test07.txt",
	}

	expectedPart2 := "7"

	answer := Solve(puzzle)

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test08(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       16,
		InputFile: "test08.txt",
	}

	expectedPart2 := "9"

	answer := Solve(puzzle)

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test09(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       16,
		InputFile: "test09.txt",
	}

	expectedPart2 := "1"

	answer := Solve(puzzle)

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test10(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       16,
		InputFile: "test10.txt",
	}

	expectedPart2 := "0"

	answer := Solve(puzzle)

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test11(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       16,
		InputFile: "test11.txt",
	}

	expectedPart2 := "0"

	answer := Solve(puzzle)

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test12(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       16,
		InputFile: "test12.txt",
	}

	expectedPart2 := "1"

	answer := Solve(puzzle)

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}
