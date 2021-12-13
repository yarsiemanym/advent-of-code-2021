package day13

import (
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
	"github.com/yarsiemanym/advent-of-code-2021/vt100"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	text := common.ReadFile(puzzle.InputFile)
	points, creases := parseInput(text)

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(points, creases),
		Part2: solvePart2(points, creases),
	}
}

func solvePart1(points []*common.Point, creases []*Crease) string {
	log.Info("Solving part 1.")

	paper := NewPaperFromPoints(points)
	paper = paper.Fold(creases[0])
	markedPoints := paper.GetMarkedPoints()
	markCount := len(markedPoints)

	log.Info("Part 1 solved.")
	return strconv.Itoa(markCount)
}

func solvePart2(points []*common.Point, creases []*Crease) string {
	log.Info("Solving part 2.")

	paper := NewPaperFromPoints(points)

	for _, crease := range creases {
		paper = paper.Fold(crease)
	}

	output := "\n" + paper.Render()
	output = strings.Replace(output, "#", "\x1b["+vt100.YellowBackgroundAttribute+"m \x1b[0m", -1)
	output = strings.Replace(output, ".", " ", -1)

	log.Info("Part 2 solved.")
	return output
}

func parseInput(text string) ([]*common.Point, []*Crease) {
	chunks := common.Split(text, "\n\n")

	if len(chunks) != 2 {
		log.Fatalf("Splitting \"%s\" on \"\\n\\n\" yielded %v tokens.", common.Peek(text, common.PEEK_MAX_DEFAULT), len(chunks))
	}

	points := parsePoints(chunks[0])
	creases := parseCreases(chunks[1])

	return points, creases
}

func parsePoints(text string) []*common.Point {
	lines := common.Split(text, "\n")
	points := []*common.Point{}

	for _, line := range lines {
		point := parsePoint(line)

		if point != nil {
			points = append(points, parsePoint(line))
		}
	}

	log.Debugf("%v points parsed.", len(points))
	return points
}

func parsePoint(text string) *common.Point {
	if text == "" {
		return nil
	}

	log.Debugf("Parsing point from \"%s\".", text)

	tokens := common.Split(text, ",")
	if len(tokens) != 2 {
		log.Fatalf("Splitting \"%s\" on ',' yielded %v tokens.", text, len(tokens))
	}

	x, err := strconv.Atoi(tokens[0])
	if err != nil {
		log.Fatalf("\"%s\" is not an integer.", tokens[0])
	}

	y, err := strconv.Atoi(tokens[1])
	if err != nil {
		log.Fatalf("\"%s\" is not an integer.", tokens[1])
	}

	point := common.NewPoint(x, y)
	log.Tracef("Point (%d, %d) parsed.", point.X(), point.Y())

	return point
}

func parseCreases(text string) []*Crease {
	lines := common.Split(text, "\n")
	creases := []*Crease{}

	for _, line := range lines {
		crease := parseCrease(line)

		if crease != nil {
			creases = append(creases, crease)
		}
	}

	log.Debugf("%v creases parsed.", len(creases))
	return creases
}

func parseCrease(text string) *Crease {
	if text == "" {
		return nil
	}

	log.Debugf("Parsing crease from \"%s\".", text)

	pattern := regexp.MustCompile(`^fold along (x|y)=(\d+)$`)
	match := pattern.FindStringSubmatch(text)

	axis := rune(match[1][0])

	position, err := strconv.Atoi(match[2])
	if err != nil {
		log.Fatalf("\"%s\" is not an integer.", match[1])
	}

	crease := &Crease{
		Axis:     axis,
		Position: position,
	}
	log.Tracef("Crease %c=%d parsed.", crease.Axis, crease.Position)
	return crease
}
