package day16

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	text := common.ReadFile(puzzle.InputFile)

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(text),
		Part2: solvePart2(text),
	}
}

func solvePart1(text string) string {
	log.Info("Solving part 1.")

	bits := hexStringToBitArray(text)

	var packet Packet
	packet, _ = ParsePacket(bits)
	checkSum := packet.CheckSum()

	log.Info("Part 1 solved.")
	return strconv.FormatUint(checkSum, 10)
}

func solvePart2(text string) string {
	log.Info("Solving part 2.")

	bits := hexStringToBitArray(text)

	var packet Packet
	packet, _ = ParsePacket(bits)
	value := packet.Value()

	log.Info("Part 2 solved.")
	return strconv.FormatUint(value, 10)
}
