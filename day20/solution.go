package day20

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	text := common.ReadFile(puzzle.InputFile)
	enhancer, image := parseInput(text)

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(enhancer, image),
		Part2: solvePart2(),
	}
}

func solvePart1(enhancer *ImageEnhancer, image *Image) string {
	log.Info("Solving part 1.")

	log.Debugf("Image\n%s", image)
	image = enhancer.EnhanceImage(image)
	log.Debugf("Image\n%s", image)
	image = enhancer.EnhanceImage(image)
	log.Debugf("Image\n%s", image)

	log.Info("Part 1 solved.")
	return strconv.FormatUint(image.CountIlluminatedPixels(), 10)
}

func solvePart2() string {
	log.Info("Solving part 2.")

	// TODO: implement part 2 solution

	log.Info("Part 2 solved.")
	return "Not implemented."
}

func parseInput(text string) (*ImageEnhancer, *Image) {
	chunks := common.Split(text, "\n\n")

	if len(chunks) != 2 {
		log.Fatalf("Splitting \"%s\" on \"\\n\\n\" did not yield 2 chunks.", common.Peek(text, common.PEEK_MAX_DEFAULT))
	}

	return NewImageEnhancer(chunks[0]), NewImageFromString(chunks[1])
}
