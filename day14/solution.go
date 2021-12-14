package day14

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	text := common.ReadFile(puzzle.InputFile)
	polymer, rules := parseInput(text)

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(polymer, rules),
		Part2: solvePart2(polymer, rules),
	}
}

func solvePart1(polymer *Polymer, rules []*InsertionRule) string {
	log.Info("Solving part 1.")

	for i := 1; i <= 10; i++ {
		log.Debugf("Beginning step %d.", i)
		polymer.Insert(rules...)
	}

	log.Debug("Analyzing polymer.")
	analysis := polymer.Anaylze()

	log.Info("Part 1 solved.")
	return strconv.Itoa(analysis)
}

func solvePart2(polymer *Polymer, rules []*InsertionRule) string {
	log.Info("Solving part 2.")

	/* for i := 1; i <= 40; i++ {
		log.Debugf("Beginning step %d.", i)
		polymer.Insert(rules...)
	}

	log.Debug("Analyzing polymer.")
	analysis := polymer.Anaylze()

	log.Info("Part 2 solved.")
	return strconv.Itoa(analysis) */

	return "Not implemented."
}

func parseInput(text string) (*Polymer, []*InsertionRule) {
	chunks := common.Split(text, "\n\n")

	if len(chunks) != 2 {
		log.Fatalf("Splitting \"%s\" on \"\\n\\n\" yielded %d chunks.", common.Peek(text, common.PEEK_MAX_DEFAULT), len(chunks))
	}

	polymer := NewPolymer(chunks[0])

	rules := []*InsertionRule{}
	lines := common.Split(chunks[1], "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		tokens := common.Split(line, " -> ")

		if len(tokens) != 2 {
			log.Fatalf("Splitting \"%s\" on \" -> \" yielded %d tokens.", common.Peek(line, common.PEEK_MAX_DEFAULT), len(tokens))
		}

		rules = append(rules, NewInsertionRule(tokens[0], tokens[1]))
	}

	return polymer, rules
}
