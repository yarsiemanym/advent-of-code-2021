package day12

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	text := common.ReadFile(puzzle.InputFile)
	caveMap := parseCaveSystem(text)

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(caveMap),
		Part2: solvePart2(caveMap),
	}
}

func solvePart1(caveMap map[string]*Cave) string {
	log.Info("Solving part 1.")

	start := caveMap["start"]
	end := caveMap["end"]
	scanner := NewScanner()
	paths := scanner.Scan(NewPath([]*Cave{start}), end, false)

	if log.GetLevel() >= log.DebugLevel {
		for _, path := range paths {
			log.Info(path.Render())
		}
	}

	pathCount := len(paths)

	log.Info("Part 1 solved.")
	return strconv.Itoa(pathCount)
}

func solvePart2(caveMap map[string]*Cave) string {
	log.Info("Solving part 2.")

	start := caveMap["start"]
	end := caveMap["end"]
	scanner := NewScanner()
	paths := scanner.Scan(NewPath([]*Cave{start}), end, true)

	if log.GetLevel() >= log.DebugLevel {
		for _, path := range paths {
			log.Debug(path.Render())
		}
	}

	pathCount := len(paths)

	log.Info("Part 2 solved.")
	return strconv.Itoa(pathCount)
}

func parseCaveSystem(text string) map[string]*Cave {
	lines := common.Split(text, "\n")
	caveMap := map[string]*Cave{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		tokens := common.Split(line, "-")

		if len(tokens) != 2 {
			log.Fatalf("Splitting \"%v\" on \"-\" yielded %v tokens.", line, len(tokens))
		}

		cave1Name, cave2Name := tokens[0], tokens[1]

		cave1, exists := caveMap[cave1Name]
		if !exists {
			cave1 = NewCave(cave1Name)
			caveMap[cave1Name] = cave1
		}

		cave2, exists := caveMap[cave2Name]
		if !exists {
			cave2 = NewCave(cave2Name)
			caveMap[cave2Name] = cave2
		}

		if cave2.Name() != "start" && cave1.Name() != "end" {
			cave1.Connect(cave2)
		}

		if cave1.Name() != "start" && cave2.Name() != "end" {
			cave2.Connect(cave1)
		}
	}

	log.Debug(renderCaveMap(caveMap))

	return caveMap
}

func renderCaveMap(caveMap map[string]*Cave) string {
	output := "Cave Map\n"

	for k, v := range caveMap {
		output += fmt.Sprintf("\"%v\" -> [", k)

		for _, connectedCave := range v.ConnectedCaves() {
			output += fmt.Sprintf("\"%v\", ", connectedCave.Name())
		}

		output = strings.Trim(output, ", ") + "]\n"
	}

	return output
}
