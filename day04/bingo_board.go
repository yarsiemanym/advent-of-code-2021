package day04

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type bingoBoardSquare struct {
	Number   int
	IsMarked bool
}

type bingoBoard struct {
	rowCount int
	colCount int
	Squares  [][]*bingoBoardSquare
	isWinner *bool
	score    *int
}

func (board *bingoBoard) Init(numbers [][]int) {
	board.rowCount = len(numbers)
	board.colCount = len(numbers[0]) // Assume the board is a square
	board.Squares = make([][]*bingoBoardSquare, len(numbers))

	for row := range numbers {
		board.Squares[row] = make([]*bingoBoardSquare, len(numbers[row]))

		for col := range numbers[row] {
			board.Squares[row][col] = &bingoBoardSquare{
				Number:   numbers[row][col],
				IsMarked: false,
			}
		}
	}

	isWinner := false
	board.isWinner = &isWinner
}

func (board *bingoBoard) MarkNumber(number int) {
	for _, row := range board.Squares {
		for _, square := range row {
			if square.Number == number {
				log.Debugf("Number %v found! Marking.", number)
				square.IsMarked = true
				break
			}
		}
	}

	board.isWinner = nil
}

func (board *bingoBoard) IsWinner() bool {
	if board.isWinner == nil {
		log.Debug("Checking board state for win condition.")
		board.isWinner = new(bool)
		*board.isWinner = board.checkRowsForWinner() || board.checkColsForWinner()
	} else {
		log.Debug("Using cached winner check.")
	}

	return *board.isWinner
}

func (board *bingoBoard) checkRowsForWinner() bool {
	log.Debug("Checking board state for row win condition.")
	var isWinner bool
	for i := 0; i < board.rowCount; i++ {
		isWinner = true
		for j := 0; j < board.colCount; j++ {
			square := board.Squares[i][j]
			isWinner = isWinner && square.IsMarked
		}

		if isWinner {
			break
		}
	}

	if isWinner {
		log.Debug("Board is a winner!")
	} else {
		log.Debug("Board is not a winner.")
	}

	return isWinner
}

func (board *bingoBoard) checkColsForWinner() bool {
	log.Debug("Checking board state for column win condition.")
	var isWinner bool
	for i := 0; i < board.colCount; i++ {
		isWinner = true
		for j := 0; j < board.rowCount; j++ {
			square := board.Squares[j][i]
			isWinner = isWinner && square.IsMarked
		}

		if isWinner {
			break
		}
	}

	if isWinner {
		log.Debug("Board is a winner!")
	} else {
		log.Debug("Board is not a winner.")
	}

	return isWinner
}

func (board *bingoBoard) Score(lastNumber int) int {

	if board.score == nil {
		log.Debug("Calculating score.")
		var unmarkedNumbers []int

		for _, row := range board.Squares {
			for _, square := range row {
				if !square.IsMarked {
					unmarkedNumbers = append(unmarkedNumbers, square.Number)
				}
			}
		}

		sum := common.SumInt(unmarkedNumbers...)
		board.score = new(int)
		*board.score = sum * lastNumber
	} else {
		log.Debug("Using cached score.")
	}

	return *board.score
}

func (board *bingoBoard) Print() {
	message := "Bingo Board\n"

	for _, row := range board.Squares {
		for _, square := range row {
			if square.IsMarked {
				message = message + fmt.Sprintf("[%2v] ", strconv.Itoa(square.Number))
			} else {
				message = message + fmt.Sprintf(" %2v  ", strconv.Itoa(square.Number))
			}

		}
		message += "\n"
	}

	log.Info(message)
}
