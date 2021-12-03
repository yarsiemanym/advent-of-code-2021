package day02

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n", parseCommand)
	commands := make([]command, len(results))

	for index, result := range results {
		commands[index] = result.(command)
	}

	answer := common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(commands),
		Part2: solvePart2(commands),
	}

	return answer
}

func solvePart1(commands []command) string {
	log.Debug("Solving part 1.")
	log.Tracef("commands = %v", commands)

	position := 0
	depth := 0

	log.Tracef("position = %v", position)
	log.Tracef("depth = %v", depth)

	for _, command := range commands {
		log.Debug("Inspecting command.")
		log.Tracef("command.Name = \"%v\"", command.Name)
		log.Tracef("command.Value = %v", command.Value)

		log.Debug("Executing command.")
		switch command.Name {
		case "forward":
			position += command.Value
		case "down":
			depth += command.Value
		case "up":
			depth -= command.Value
		default:
			log.Warnf("Unsupported command \"%v %v\".", command.Name, command.Value)
		}

		log.Tracef("position = %v", position)
		log.Tracef("depth = %v", depth)
	}

	log.Debug("Part 1 solved.")
	return strconv.Itoa(position * depth)
}

func solvePart2(commands []command) string {
	log.Debug("Solving part 2.")
	log.Tracef("commands = %v", commands)

	position := 0
	depth := 0
	aim := 0

	log.Tracef("position = %v", position)
	log.Tracef("depth = %v", depth)
	log.Tracef("aim = %v", aim)

	for _, command := range commands {
		log.Debug("Inspecting command.")
		log.Tracef("command.Name = \"%v\"", command.Name)
		log.Tracef("command.Value = %v", command.Value)

		log.Debug("Executing command.")
		switch command.Name {
		case "forward":
			position += command.Value
			depth += command.Value * aim
		case "down":
			aim += command.Value
		case "up":
			aim -= command.Value
		default:
			log.Warnf("Skipping unsupported command \"%v %v\".", command.Name, command.Value)
		}

		log.Tracef("position = %v", position)
		log.Tracef("depth = %v", depth)
		log.Tracef("aim = %v", aim)
	}

	log.Debug("Part 2 solved.")
	return strconv.Itoa(position * depth)
}

func parseCommand(text string) interface{} {
	if text == "" {
		return nil
	}

	tokens := common.Split(text, " ")

	if len(tokens) != 2 {
		log.Fatalf("When split on \" \", text \"%v\" does not yield 2 tokens.", text)
	}

	value, err := strconv.Atoi(tokens[1])
	common.Check(err)

	command := command{
		Name:  tokens[0],
		Value: value,
	}

	return command
}
