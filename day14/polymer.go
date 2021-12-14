package day14

import (
	"container/list"
	"math"

	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type Polymer struct {
	composition *list.List
}

func NewPolymer(elements string) *Polymer {
	list := list.New()

	for _, element := range elements {
		list.PushBack(string(element))
	}

	return &Polymer{
		composition: list,
	}
}

func (polymer *Polymer) Render() string {
	output := ""

	for node := polymer.composition.Front(); node != nil; node = node.Next() {
		output += node.Value.(string)
	}

	return output
}

func (polymer *Polymer) Insert(rules ...*InsertionRule) {
	for node := polymer.composition.Front().Next(); node != nil; node = node.Next() {
		for _, rule := range rules {
			sequence := node.Prev().Value.(string) + node.Value.(string)
			if sequence == rule.Target() {
				polymer.composition.InsertBefore(rule.NewElement(), node)
				break
			}
		}
	}
}

func (polymer *Polymer) Anaylze() int {
	elementCounts := map[string]int{}

	for node := polymer.composition.Front(); node != nil; node = node.Next() {
		element := node.Value.(string)
		count := elementCounts[element]
		count++
		elementCounts[element] = count
	}

	max := math.MinInt
	min := math.MaxInt

	for _, count := range elementCounts {
		max = common.MaxInt(count, max)
		min = common.MinInt(count, min)
	}

	return max - min
}
