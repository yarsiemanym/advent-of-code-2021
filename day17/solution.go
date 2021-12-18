package day17

import (
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	text := common.ReadFile(puzzle.InputFile)
	text = strings.Trim(text, " \n")
	targetLeft, targetRight, targetBottom, targetTop := parseInput(text)

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(targetLeft, targetRight, targetBottom, targetTop),
		Part2: solvePart2(targetLeft, targetRight, targetBottom, targetTop),
	}
}

func solvePart1(targetLeft int, targetRight int, targetBottom int, targetTop int) string {
	log.Info("Solving part 1.")

	// There is no drag on the Y axis so what goes up must reach Y=0 with the same Y velocity as it
	// started with and the next step must not overshoot the target area.
	maximumVelocity := common.AbsInt(targetBottom) - 1
	maxHeight := triangularSum(maximumVelocity)

	log.Info("Part 1 solved.")
	return strconv.Itoa(maxHeight)
}

func solvePart2(targetLeft int, targetRight int, targetBottom int, targetTop int) string {
	log.Info("Solving part 2.")

	minXVelocity := 0
	for i := 1; triangularSum(i) <= targetLeft; i++ {
		minXVelocity = i
	}
	log.Tracef("minXVelocity = %d", minXVelocity)

	maxXVelocity := targetRight
	log.Tracef("maxXVelocity = %d", maxXVelocity)

	minYVelocity := targetBottom
	log.Tracef("minYVelocity = %d", minYVelocity)

	maxYVelocity := common.AbsInt(targetBottom) - 1
	log.Tracef("maxYVelocity = %d", maxYVelocity)

	trajectoriesThatHit := 0

	for x := minXVelocity; x <= maxXVelocity; x++ {
		for y := minYVelocity; y <= maxYVelocity; y++ {
			if isHit(x, y, targetLeft, targetRight, targetBottom, targetTop) {
				trajectoriesThatHit++
			}
		}
	}

	log.Debugf("%d trajectories hit.", trajectoriesThatHit)

	log.Info("Part 2 solved.")
	return strconv.Itoa(trajectoriesThatHit)
}

// https://en.wikipedia.org/wiki/Triangular_number
func triangularSum(number int) int {
	return (number * (number + 1)) / 2
}

func isHit(xVelocity int, yVelocity int, targetLeft int, targetRight int, targetBottom int, targetTop int) bool {
	log.Debugf("Determining if trajectory (%d,%d) is a hit.", xVelocity, yVelocity)
	for step, xPosition, yPosition := 1, 0, 0; ; step++ {
		xPosition += xVelocity
		xVelocity = common.MaxInt(xVelocity-1, 0)
		yPosition += yVelocity
		yVelocity--

		log.Tracef("step = %d", step)
		log.Tracef("xPosition = %d", xPosition)
		log.Tracef("yPosition = %d", yPosition)

		if xPosition >= targetLeft && xPosition <= targetRight && yPosition >= targetBottom && yPosition <= targetTop {
			log.Debugf("Hit!")
			return true
		} else if xPosition > targetRight || yPosition < targetBottom {
			log.Debugf("Miss.")
			return false
		}
	}
}

func parseInput(text string) (int, int, int, int) {
	pattern := regexp.MustCompile(`^target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)$`)
	matches := pattern.FindStringSubmatch(text)

	targetRight, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer.", matches[1])
	}

	targetLeft, err := strconv.Atoi(matches[2])
	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer.", matches[2])
	}

	targetBottom, err := strconv.Atoi(matches[3])
	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer.", matches[3])
	}

	targetTop, err := strconv.Atoi(matches[4])
	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer.", matches[4])
	}

	return targetRight, targetLeft, targetBottom, targetTop
}
