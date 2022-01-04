package day22

import (
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n", parseRebootSteps)
	rebootSteps := make([]*RebootStep, len(results))

	for index, result := range results {
		rebootSteps[index] = result.(*RebootStep)
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(rebootSteps),
		Part2: solvePart2(rebootSteps),
	}
}

func solvePart1(rebootSteps []*RebootStep) string {
	log.Info("Solving part 1.")

	reactorCore := NewCuboid(common.New3DPoint(-100000, -100000, -100000), common.New3DPoint(100000, 100000, 100000))

	for _, step := range rebootSteps {
		for x := step.cuboid.span.Start().X(); x <= step.cuboid.span.End().X(); x++ {

			if x < -50 || x > 50 {
				continue
			}

			for y := step.cuboid.span.Start().Y(); y <= step.cuboid.span.End().Y(); y++ {

				if y < -50 || y > 50 {
					continue
				}

				for z := step.cuboid.span.Start().Z(); z <= step.cuboid.span.End().Z(); z++ {

					if z < -50 || z > 50 {
						continue
					}

					reactorCore.SetCube(common.New3DPoint(x, y, z), step.on)
				}
			}
		}
	}

	onCount := reactorCore.CountOnCubes()

	log.Info("Part 1 solved.")
	return strconv.Itoa(onCount)
}

func solvePart2(rebootSteps []*RebootStep) string {
	log.Info("Solving part 2.")

	// TODO: implement part 2 solution

	log.Info("Part 2 solved.")
	return "Not implemented."
}

func parseRebootSteps(text string) interface{} {
	if text == "" {
		return nil
	}

	pattern := regexp.MustCompile(`(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)
	matches := pattern.FindStringSubmatch(text)

	if len(matches) != 8 {
		log.Fatalf("Failed to parse reboot step \"%s\".", text)
	}

	on := matches[1] == "on"

	startX, err := strconv.Atoi(matches[2])

	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer.", matches[2])
	}

	endX, err := strconv.Atoi(matches[3])

	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer.", matches[3])
	}

	startY, err := strconv.Atoi(matches[4])

	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer.", matches[4])
	}

	endY, err := strconv.Atoi(matches[5])

	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer.", matches[5])
	}

	startZ, err := strconv.Atoi(matches[6])

	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer.", matches[7])
	}

	endZ, err := strconv.Atoi(matches[7])

	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer.", matches[7])
	}

	return NewRebootStep(NewCuboid(common.New3DPoint(startX, startY, startZ), common.New3DPoint(endX, endY, endZ)), on)
}
