package day04

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	log.Info("Reading input.")
	text := common.ReadFile(puzzle.InputFile)
	chunks := common.Split(text, "\n\n")

	log.Info("Setting up part 1.")
	generator1 := parseNumberGenerator(chunks[0])
	boards1 := parseBingoBoards(chunks[1:])
	answer1 := solvePart1(boards1, generator1)

	log.Info("Setting up part 2.")
	generator2 := parseNumberGenerator(chunks[0])
	boards2 := parseBingoBoards(chunks[1:])
	answer2 := solvePart2(boards2, generator2)

	answer := common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: answer1,
		Part2: answer2,
	}

	return answer
}

func solvePart1(boards []*bingoBoard, generator numberGenerator) string {
	log.Info("Solving part 1.")

	var winner *bingoBoard
	var calledNumber int

	for winner == nil {
		calledNumber = generator.Next()
		log.Infof("Calling number %v.", calledNumber)

		for i, board := range boards {
			log.Debugf("Checking for number %v on board %v.", calledNumber, i+1)
			board.MarkNumber(calledNumber)

			log.Debugf("Checking board %v for win conditions.", i+1)
			if board.IsWinner() {
				winner = board
				log.Infof("Winner detected on board %v!", i+1)

				log.Infof("Scoring board %v.", i+1)
				score := winner.Score(calledNumber)
				log.Infof("Score is %v.", score)
				break
			}
		}
	}

	score := winner.Score(calledNumber)

	log.Info("Part 1 solved.")
	return strconv.Itoa(score)
}

func solvePart2(boards []*bingoBoard, generator numberGenerator) string {
	log.Info("Solving part 2.")

	var winners []*bingoBoard
	var calledNumber int

	for len(winners) < len(boards) {
		calledNumber = generator.Next()
		log.Infof("Calling number %v.", calledNumber)

		for i, board := range boards {

			log.Debugf("Checking if board %v has already won.", i+1)
			if board.IsWinner() {
				log.Debugf("Board %v has already won. Skipping.", i+1)
				continue
			} else {
				log.Debugf("Board %v has not yet won.", i+1)
			}

			log.Debugf("Checking for number %v on board %v.", calledNumber, i+1)
			board.MarkNumber(calledNumber)

			if board.IsWinner() {
				log.Infof("Winner %v detected on board %v!", len(winners)+1, i+1)
				winners = append(winners, board)

				log.Infof("Scoring board %v.", i+1)
				score := board.Score(calledNumber)
				log.Infof("Score is %v.", score)
			}
		}
	}

	lastWinner := winners[len(winners)-1]

	log.Info("Scoring last winner.")
	score := lastWinner.Score(calledNumber)
	log.Infof("Score is %v.", score)

	log.Info("Part 2 solved.")
	return strconv.Itoa(score)
}

func parseNumberGenerator(text string) numberGenerator {
	tokens := common.Split(text, ",")
	numbers := make([]int, len(tokens))

	for i, token := range tokens {
		number, err := strconv.Atoi(token)
		common.Check(err)
		numbers[i] = number
	}

	generator := numberGenerator{}
	generator.Init(numbers)
	return generator
}

func parseBingoBoards(chunks []string) []*bingoBoard {
	log.Debug("Parsing bingo boards.")
	log.Tracef("len(chunks) = %v", len(chunks))
	boards := make([]*bingoBoard, len(chunks))

	for i, chunk := range chunks {
		boards[i] = parseBingoBoard(chunk)
	}

	log.Infof("%v bingo boards detected.", len(boards))

	return boards
}

func parseBingoBoard(text string) *bingoBoard {
	log.Debug("Parsing bingo board.")
	log.Tracef("text = \"%v\"", common.Peek(text, common.PEEK_MAX_DEFAULT))
	lines := common.Split(text, "\n")
	numbers := make([][]int, len(lines))

	for i, line := range lines {
		// Clean up extra spaces so we can split correctly
		line = strings.Trim(strings.ReplaceAll(line, "  ", " "), " ")

		if line == "" {
			continue
		}

		tokens := common.Split(line, " ")
		numbers[i] = make([]int, len(tokens))

		for j, token := range tokens {
			number, err := strconv.Atoi(token)

			if err != nil {
				log.Fatalf("\"%v\" is not a number.", token)
			}

			numbers[i][j] = number
		}
	}

	board := &bingoBoard{}
	board.Init(numbers)
	board.Print()

	return board
}
