package day05

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	results := common.ParseFile(puzzle.InputFile, "\n", parseLineSegment)
	lines := make([]*common.LineSegment, len(results))

	for index, result := range results {
		lines[index] = result.(*common.LineSegment)
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(lines),
		Part2: solvePart2(lines),
	}
}

func solvePart1(lines []*common.LineSegment) string {
	log.Info("Solving part 1.")

	ventMap := ventMap{}
	ventMap.Init(lines)

	log.Info("Applying lines to vent map.")
	for _, line := range lines {

		log.Debugf("Inspecting line \"%s -> %s\" to the vent map.", line.Start(), line.End())

		if line.IsHorizontal() || line.IsVertical() {
			ventMap.ApplyLine(line)
		} else {
			log.Debug("Line is diagonal. Skipping.")
		}
	}

	log.Info("Counting locations with more than 2 lines overlapping.")
	count := ventMap.CountOverlaps(2)
	log.Tracef("count = %d", count)

	log.Info("Part 1 solved.")
	return strconv.Itoa(count)
}

func solvePart2(lines []*common.LineSegment) string {
	log.Info("Solving part 2.")

	ventMap := ventMap{}
	ventMap.Init(lines)

	log.Info("Applying lines to vent map.")
	for _, line := range lines {

		ventMap.ApplyLine(line)
	}

	log.Info("Counting locations with more than 2 lines overlapping.")
	count := ventMap.CountOverlaps(2)
	log.Tracef("count = %d", count)

	log.Info("Part 2 solved.")
	return strconv.Itoa(count)
}

func parseLineSegment(text string) interface{} {
	if text == "" {
		return nil
	}

	tokens := common.Split(text, " -> ")

	if len(tokens) != 2 {
		log.Fatalf("Splitting \"%s\" on \" -> \" yielded %d tokens.", text, len(tokens))
	}

	start := parsePoint(tokens[0])
	end := parsePoint(tokens[1])
	line := common.NewLineSegment(start, end)

	return line
}

func parsePoint(text string) *common.Point {
	tokens := common.Split(text, ",")

	if len(tokens) != 2 {
		log.Fatalf("Splitting \"%s\" on \",\" yielded %d tokens.", text, len(tokens))
	}

	x, err := strconv.Atoi(tokens[0])

	if err != nil {
		log.Fatalf("\"%s\" is not an integer.", tokens[0])
	}

	y, err := strconv.Atoi(tokens[1])

	if err != nil {
		log.Fatalf("\"%s\" is not an integer.", tokens[1])
	}

	return common.New2DPoint(x, y)
}
