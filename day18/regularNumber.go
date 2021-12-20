package day18

import "fmt"

type RegularNumber struct {
	parent Node
	value  int
}

func NewRegularNumber(value int) *RegularNumber {
	return &RegularNumber{
		value: value,
	}
}

func (regularNumber *RegularNumber) GetParent() Node {
	return regularNumber.parent
}

func (regularNumber *RegularNumber) SetParent(parent Node) {
	regularNumber.parent = parent
}

func (regularNumber *RegularNumber) GetLeft() Node {
	return nil
}

func (regularNumber *RegularNumber) GetRight() Node {
	return nil
}

func (regularNumber *RegularNumber) GetValue() int {
	return regularNumber.value
}

func (regularNumber *RegularNumber) Add(other *RegularNumber) {
	regularNumber.value += other.value
}

func (regularNumber RegularNumber) String() string {
	return fmt.Sprintf("%d", regularNumber.value)
}

func (regularNumber *RegularNumber) Split() bool {

	if regularNumber.value > 9 {
		leftValue := regularNumber.value / 2
		rightValue := regularNumber.value - leftValue

		snailfishNumber := &SnailfishNumber{}
		snailfishNumber.SetLeft(NewRegularNumber(leftValue))
		snailfishNumber.SetRight(NewRegularNumber(rightValue))

		if regularNumber.GetParent().(*SnailfishNumber).GetLeft() == regularNumber {
			regularNumber.GetParent().(*SnailfishNumber).SetLeft(snailfishNumber)
		} else {
			regularNumber.GetParent().(*SnailfishNumber).SetRight(snailfishNumber)
		}

		return true
	} else {
		return false
	}

}

func (regularNumber *RegularNumber) Clone() Node {
	clone := NewRegularNumber(regularNumber.value)
	return clone
}
