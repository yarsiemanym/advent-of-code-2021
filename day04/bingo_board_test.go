package day04

import "testing"

func Test_Board_Init(t *testing.T) {
	board := bingoBoard{}
	numbers := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	board.Init(numbers)

	if board.rowCount != len(numbers) {
		t.Errorf("Expected %v but got %v.", len(numbers), board.rowCount)
	}

	if board.colCount != len(numbers[0]) {
		t.Errorf("Expected %v but got %v.", len(numbers), board.colCount)
	}

	for i := 0; i < board.rowCount; i++ {
		for j := 0; j < board.colCount; j++ {
			square := board.Squares[i][j]

			if square.IsMarked {
				t.Errorf("Square (%v, %v) is marked.", i, j)
			}

			if square.Number != numbers[i][j] {
				t.Errorf("Expected %v but got %v.", numbers[i][j], square.Number)
			}
		}
	}
}

func Test_Board_MarkNumber(t *testing.T) {
	square1 := &bingoBoardSquare{
		Number:   1,
		IsMarked: false,
	}
	square2 := &bingoBoardSquare{
		Number:   2,
		IsMarked: false,
	}
	square3 := &bingoBoardSquare{
		Number:   3,
		IsMarked: false,
	}
	square4 := &bingoBoardSquare{
		Number:   4,
		IsMarked: false,
	}

	board := bingoBoard{
		rowCount: 2,
		colCount: 2,
		Squares: [][]*bingoBoardSquare{
			{
				square1,
				square2,
			},
			{
				square3,
				square4,
			},
		},
	}

	board.MarkNumber(4)

	if square1.IsMarked {
		t.Error("Square with number 1 should not be marked.")
	}

	if square2.IsMarked {
		t.Error("Square with number 2 should not be marked.")
	}

	if square3.IsMarked {
		t.Error("Square with number 3 should not be marked.")
	}

	if !square4.IsMarked {
		t.Error("Square with number 4 should be marked.")
	}
}

func Test_Board_IsWinner_NotWinner(t *testing.T) {
	square1 := &bingoBoardSquare{
		Number:   1,
		IsMarked: true,
	}
	square2 := &bingoBoardSquare{
		Number:   2,
		IsMarked: false,
	}
	square3 := &bingoBoardSquare{
		Number:   3,
		IsMarked: false,
	}
	square4 := &bingoBoardSquare{
		Number:   4,
		IsMarked: true,
	}

	board := bingoBoard{
		rowCount: 2,
		colCount: 2,
		Squares: [][]*bingoBoardSquare{
			{
				square1,
				square2,
			},
			{
				square3,
				square4,
			},
		},
	}

	isWinner := board.IsWinner()

	if isWinner {
		t.Error("Board is a winner.")
	}
}

func Test_Board_IsWinner_Row1Winner(t *testing.T) {
	square1 := &bingoBoardSquare{
		Number:   1,
		IsMarked: true,
	}
	square2 := &bingoBoardSquare{
		Number:   2,
		IsMarked: true,
	}
	square3 := &bingoBoardSquare{
		Number:   3,
		IsMarked: false,
	}
	square4 := &bingoBoardSquare{
		Number:   4,
		IsMarked: false,
	}

	board := bingoBoard{
		rowCount: 2,
		colCount: 2,
		Squares: [][]*bingoBoardSquare{
			{
				square1,
				square2,
			},
			{
				square3,
				square4,
			},
		},
	}

	isWinner := board.IsWinner()

	if !isWinner {
		t.Error("Board is not a winner.")
	}
}

func Test_Board_IsWinner_Row2Winner(t *testing.T) {
	square1 := &bingoBoardSquare{
		Number:   1,
		IsMarked: false,
	}
	square2 := &bingoBoardSquare{
		Number:   2,
		IsMarked: false,
	}
	square3 := &bingoBoardSquare{
		Number:   3,
		IsMarked: true,
	}
	square4 := &bingoBoardSquare{
		Number:   4,
		IsMarked: true,
	}

	board := bingoBoard{
		rowCount: 2,
		colCount: 2,
		Squares: [][]*bingoBoardSquare{
			{
				square1,
				square2,
			},
			{
				square3,
				square4,
			},
		},
	}

	isWinner := board.IsWinner()

	if !isWinner {
		t.Error("Board is not a winner.")
	}
}

func Test_Board_IsWinner_Col1Winner(t *testing.T) {
	square1 := &bingoBoardSquare{
		Number:   1,
		IsMarked: true,
	}
	square2 := &bingoBoardSquare{
		Number:   2,
		IsMarked: false,
	}
	square3 := &bingoBoardSquare{
		Number:   3,
		IsMarked: true,
	}
	square4 := &bingoBoardSquare{
		Number:   4,
		IsMarked: false,
	}

	board := bingoBoard{
		rowCount: 2,
		colCount: 2,
		Squares: [][]*bingoBoardSquare{
			{
				square1,
				square2,
			},
			{
				square3,
				square4,
			},
		},
	}

	isWinner := board.IsWinner()

	if !isWinner {
		t.Error("Board is not a winner.")
	}
}

func Test_Board_IsWinner_Col2Winner(t *testing.T) {
	square1 := &bingoBoardSquare{
		Number:   1,
		IsMarked: false,
	}
	square2 := &bingoBoardSquare{
		Number:   2,
		IsMarked: true,
	}
	square3 := &bingoBoardSquare{
		Number:   3,
		IsMarked: false,
	}
	square4 := &bingoBoardSquare{
		Number:   4,
		IsMarked: true,
	}

	board := bingoBoard{
		rowCount: 2,
		colCount: 2,
		Squares: [][]*bingoBoardSquare{
			{
				square1,
				square2,
			},
			{
				square3,
				square4,
			},
		},
	}

	isWinner := board.IsWinner()

	if !isWinner {
		t.Error("Board is not a winner.")
	}
}

func Test_Board_Score(t *testing.T) {
	square1 := &bingoBoardSquare{
		Number:   1,
		IsMarked: false,
	}
	square2 := &bingoBoardSquare{
		Number:   2,
		IsMarked: true,
	}
	square3 := &bingoBoardSquare{
		Number:   3,
		IsMarked: false,
	}
	square4 := &bingoBoardSquare{
		Number:   4,
		IsMarked: true,
	}

	board := bingoBoard{
		rowCount: 2,
		colCount: 2,
		Squares: [][]*bingoBoardSquare{
			{
				square1,
				square2,
			},
			{
				square3,
				square4,
			},
		},
	}

	expectedScore := 8

	actualScore := board.Score(2)

	if actualScore != expectedScore {
		t.Errorf("Expected %v but got %v.", expectedScore, actualScore)
	}

}
