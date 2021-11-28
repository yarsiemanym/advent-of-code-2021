package common

import (
	log "github.com/sirupsen/logrus"
)

type Puzzle struct {
	Day       int
	InputFile string
	solution  Solution
}

func (puzzle *Puzzle) SetSolution(solution Solution) {
	puzzle.solution = solution
}

func (puzzle Puzzle) Solve() Answer {
	if puzzle.solution == nil {
		panic("Solution is not set.")
	}

	log.Info("Solving puzzle.")
	answer := puzzle.solution(puzzle)
	log.Info("Puzzle solved!")

	return answer
}
