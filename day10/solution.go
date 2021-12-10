package day10

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n", parseLine)
	lines := make([]string, len(results))

	for index, result := range results {
		lines[index] = result.(string)
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(lines),
		Part2: solvePart2(lines),
	}
}

var matches = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func solvePart1(lines []string) string {
	log.Info("Solving part 1.")

	totalScore := 0

	for i, line := range lines {
		hasSyntaxError := false
		score := 0
		stack := common.NewStack()

		for j, char := range line {
			switch char {
			case '(', '[', '{', '<':
				stack.Push(char)
			case ')', ']', '}', '>':
				opening := stack.Pop().(rune)
				if char != matches[opening] {
					log.Debugf("Line %v has a syntax error at position %v.", i, j)
					log.Tracef("line = \"%v\"", line)
					log.Tracef("char = '%c'", char)
					hasSyntaxError = true
					score = scoreSyntaxError(char)
					totalScore += score
				}
			default:
				log.Fatalf("Invalid character '%c'.", char)
			}

			if hasSyntaxError {
				break
			}
		}
	}

	log.Tracef("totalScore = %v", totalScore)

	log.Info("Part 1 solved.")
	return strconv.Itoa(totalScore)
}

func scoreSyntaxError(illegalChar rune) int {
	log.Tracef("Scoring illegal character '%c'", illegalChar)

	score := 0

	switch illegalChar {
	case ')':
		score = 3
	case ']':
		score = 57
	case '}':
		score = 1197
	case '>':
		score = 25137
	default:
		log.Warningf("Character '%c' does not have a score.", illegalChar)
	}

	log.Tracef("score = %v", score)

	return score
}

func solvePart2(lines []string) string {
	log.Info("Solving part 2.")

	var scores []int

	for i, line := range lines {
		isCorrupted := false
		stack := common.NewStack()

		for j, char := range line {
			switch char {
			case '(', '[', '{', '<':
				stack.Push(char)
			case ')', ']', '}', '>':
				opening := stack.Pop().(rune)
				if char != matches[opening] {
					log.Debugf("Line %v has a syntax error at position %v.", i, j)
					log.Tracef("line = \"%v\"", line)
					log.Tracef("char = '%c'", char)
					isCorrupted = true
				}
			default:
				log.Fatalf("Invalid character '%c'.", char)
			}

			if isCorrupted {
				break
			}
		}

		if !isCorrupted {
			log.Debugf("Line %v is incomplete.", i)
			log.Tracef("line = \"%v\"", line)
			lineScore := autoCompleteAndScore(line, stack)
			scores = append(scores, lineScore)
		}
	}

	finalScore := common.MedianInt(scores...)

	log.Info("Part 2 solved.")
	return strconv.Itoa(finalScore)
}

func autoCompleteAndScore(incompleteLine string, stack *common.Stack) int {
	log.Tracef("Autocompleting line.")
	log.Tracef("incompleteLine = \"%v\"", incompleteLine)
	score := 0
	completedLine := incompleteLine

	for value := stack.Pop(); value != nil; value = stack.Pop() {
		nextChar := matches[value.(rune)]
		completedLine += string(nextChar)
		score = (score * 5) + scoreAutoCompleteChar(nextChar)
	}

	log.Tracef("completedLine = \"%v\"", completedLine)
	log.Tracef("score = %v", score)

	return score
}

func scoreAutoCompleteChar(autoCompletedChar rune) int {
	log.Tracef("Scoring autocompleted character '%c'", autoCompletedChar)

	score := 0

	switch autoCompletedChar {
	case ')':
		score = 1
	case ']':
		score = 2
	case '}':
		score = 3
	case '>':
		score = 4
	default:
		log.Warningf("Character '%c' does not have a score.", autoCompletedChar)
	}

	log.Tracef("score = %v", score)

	return score
}

func parseLine(text string) interface{} {
	if text == "" {
		return nil
	}

	return text
}
