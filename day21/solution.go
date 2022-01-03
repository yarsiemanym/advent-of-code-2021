package day21

import (
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n", parsePlayers)
	players := make([]*Player, len(results))

	for index, result := range results {
		players[index] = result.(*Player)
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(players),
		Part2: solvePart2(),
	}
}

func solvePart1(players []*Player) string {
	log.Info("Solving part 1.")

	dice := NewDeterministicDie(100)
	var winner *Player

	for winner == nil {
		for i := 0; i < len(players) && winner == nil; i++ {
			player := players[i]
			player.TakeTurn(dice)

			if player.Score() >= 1000 {
				winner = player
			}
		}
	}

	result := 0

	if *players[0] == *winner {
		result = players[1].Score() * dice.NumberOfRolls()
	} else {
		result = players[0].Score() * dice.NumberOfRolls()
	}

	log.Info("Part 1 solved.")
	return strconv.Itoa(result)
}

func solvePart2() string {
	log.Info("Solving part 2.")

	// TODO: implement part 2 solution

	log.Info("Part 2 solved.")
	return "Not implemented."
}

func parsePlayers(text string) interface{} {
	if text == "" {
		return nil
	}

	pattern := regexp.MustCompile(`Player (\d+) starting position: (\d+)`)
	matches := pattern.FindStringSubmatch(text)

	if len(matches) != 3 {
		log.Fatalf("Failed to parse text \"%s\".", common.Peek(text, common.PEEK_MAX_DEFAULT))
	}

	id, err := strconv.Atoi(matches[1])

	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer.", matches[1])
	}

	position, err := strconv.Atoi(matches[2])

	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer.", matches[2])
	}

	return NewPlayer(id, position)
}
