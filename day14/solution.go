package day14

import (
	"math"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	text := common.ReadFile(puzzle.InputFile)
	template, rules := parseInput(text)

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(template, rules),
		Part2: solvePart2(template, rules),
	}
}

func solvePart1(template string, rules map[string]rune) string {
	log.Info("Solving part 1.")
	result := iterate(template, rules, 10)
	log.Info("Part 1 solved.")
	return strconv.Itoa(result)
}

func solvePart2(template string, rules map[string]rune) string {
	log.Info("Solving part 2.")
	result := iterate(template, rules, 40)
	log.Info("Part 2 solved.")
	return strconv.Itoa(result)
}

func iterate(template string, rules map[string]rune, numIterations int) int {
	pairCounts := map[string]int{}

	log.Debugf("Counting pairs in template \"%s\".", template)
	for i := 1; i < len(template); i++ {
		pairKey := template[i-1 : i+1]
		pairCounts[pairKey]++
		log.Tracef("\"%s\": %d", pairKey, pairCounts[pairKey])
	}

	log.Debug("Applying insertion rules.")
	for i := 1; i <= numIterations; i++ {
		log.Debugf("Step %d.", i)
		newPairCounts := map[string]int{}
		for pairKey := range pairCounts {
			ruleValue := rules[pairKey]
			log.Debugf("Inserting '%c' into \"%s\".", ruleValue, pairKey)
			newRuleKey1 := string(pairKey[0]) + string(ruleValue)
			newPairCounts[newRuleKey1] = newPairCounts[newRuleKey1] + pairCounts[pairKey]
			log.Debugf("Pair count for \"%s\" is %d.", newRuleKey1, newPairCounts[newRuleKey1])
			newRuleKey2 := string(ruleValue) + string(pairKey[1])
			newPairCounts[newRuleKey2] = newPairCounts[newRuleKey2] + pairCounts[pairKey]
			log.Debugf("Pair count for \"%s\" is %d.", newRuleKey2, newPairCounts[newRuleKey2])
		}

		pairCounts = newPairCounts
	}

	elementCounts := map[rune]int{
		rune(template[0]):               1,
		rune(template[len(template)-1]): 1,
	}

	for pair, pairCount := range pairCounts {
		elementCounts[rune(pair[0])] += pairCount
		elementCounts[rune(pair[1])] += pairCount
	}

	max := math.MinInt
	min := math.MaxInt

	for _, elementCount := range elementCounts {
		max = common.MaxInt(max, elementCount)
		min = common.MinInt(min, elementCount)
	}

	for k, v := range pairCounts {
		log.Tracef("\"%s\" = %d", k, v)
	}

	max, min = max/2, min/2

	log.Tracef("max = %d", max)
	log.Tracef("min = %d", min)

	return max - min
}

func parseInput(text string) (string, map[string]rune) {
	chunks := common.Split(text, "\n\n")

	if len(chunks) != 2 {
		log.Fatalf("Splitting \"%s\" on \"\\n\\n\" yielded %d chunks.", common.Peek(text, common.PEEK_MAX_DEFAULT), len(chunks))
	}

	template := chunks[0]

	rules := map[string]rune{}
	lines := common.Split(chunks[1], "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		tokens := common.Split(line, " -> ")

		if len(tokens) != 2 {
			log.Fatalf("Splitting \"%s\" on \" -> \" yielded %d tokens.", common.Peek(line, common.PEEK_MAX_DEFAULT), len(tokens))
		}

		target := tokens[0]
		if !regexp.MustCompile(`[A-Z]{2}`).MatchString(target) {
			log.Fatalf("\"%s\" is not a valid insertion target.", target)
		}

		insertedElement := tokens[1]
		if !regexp.MustCompile(`[A-Z]`).MatchString(insertedElement) {
			log.Fatalf("\"%s\" is not a valid element.", insertedElement)
		}

		rules[target] = rune(insertedElement[0])
	}

	log.Tracef("template = \"%s\"", template)

	for k, v := range rules {
		log.Tracef("\"%s\" => '%c'", k, v)
	}

	return template, rules
}
