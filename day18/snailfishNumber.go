package day18

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type SnailfishNumber struct {
	parent Node
	left   Node
	right  Node
}

func NewSnailfishNumber(text string) *SnailfishNumber {
	log.Debug("Parsing snailfish number.")
	log.Tracef("text = \"%s\"", text)
	text = text[1 : len(text)-1]
	root := &SnailfishNumber{}
	current := root

	for _, character := range text {
		switch character {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F':
			value, err := strconv.ParseInt(string(character), 16, 8)

			if err != nil {
				log.Fatalf("\"%c\" is not a valid digit.", character)
			}

			regularNumber := NewRegularNumber(int(value))

			if current.left == nil {
				current.SetLeft(regularNumber)
			} else {
				current.SetRight(regularNumber)
			}
		case '[':
			newNode := &SnailfishNumber{
				parent: current,
			}

			if current.left == nil {
				current.left = newNode
			} else {
				current.right = newNode
			}
			current = newNode
		case ']':
			current = current.parent.(*SnailfishNumber)
		case ',':

		default:
			log.Fatalf("Expected character '%c'.", character)
		}
	}

	return root
}

func (snailfishNumber *SnailfishNumber) GetParent() Node {
	return snailfishNumber.parent
}

func (snailfishNumber *SnailfishNumber) SetParent(parent Node) {
	snailfishNumber.parent = parent
}

func (snailfishNumber *SnailfishNumber) GetLeft() Node {
	return snailfishNumber.left
}

func (snailfishNumber *SnailfishNumber) GetRight() Node {
	return snailfishNumber.right
}

func (snailfishNumber *SnailfishNumber) SetLeft(child Node) {
	if snailfishNumber.left != nil {
		snailfishNumber.left.SetParent(nil)
	}

	child.SetParent(snailfishNumber)
	snailfishNumber.left = child
}

func (snailfishNumber *SnailfishNumber) SetRight(child Node) {
	if snailfishNumber.right != nil {
		snailfishNumber.right.SetParent(nil)
	}

	child.SetParent(snailfishNumber)
	snailfishNumber.right = child
}

func (snailfishNumber SnailfishNumber) String() string {
	return fmt.Sprintf("[%s,%s]", snailfishNumber.left, snailfishNumber.right)
}

func (snailfishNumber *SnailfishNumber) Reduce() {
	for somethingChanged := true; somethingChanged; {
		if snailfishNumber.Explode(0) {
			continue
		}

		somethingChanged = snailfishNumber.Split()
	}
}

func (snailfishNumber *SnailfishNumber) Explode(depth int) bool {
	if depth == 4 {
		parent := snailfishNumber.GetParent().(*SnailfishNumber)
		leftValue := snailfishNumber.GetLeft().(*RegularNumber)
		rightValue := snailfishNumber.GetRight().(*RegularNumber)
		nextLeft := snailfishNumber.NextLeftRegularNumber()
		nextRight := snailfishNumber.NextRightRegularNumber()

		if nextLeft != nil {
			nextLeft.Add(leftValue)
		}

		if nextRight != nil {
			nextRight.Add(rightValue)
		}

		if parent.GetLeft() == snailfishNumber {
			parent.SetLeft(NewRegularNumber(0))
		} else {
			parent.SetRight(NewRegularNumber(0))
		}

		return true
	} else {
		exploded := false

		switch left := snailfishNumber.GetLeft().(type) {
		case *SnailfishNumber:
			exploded = left.Explode(depth + 1)
		case *RegularNumber:
			exploded = false
		}

		if !exploded {
			switch right := snailfishNumber.GetRight().(type) {
			case *SnailfishNumber:
				exploded = right.Explode(depth + 1)
			case *RegularNumber:
				exploded = false
			}
		}

		return exploded
	}
}

func (snailfishNumber *SnailfishNumber) NextLeftRegularNumber() *RegularNumber {
	var previous Node = snailfishNumber
	var current Node = snailfishNumber

	for current.GetParent() != nil {
		previous = current
		current = current.GetParent()
		if current.GetLeft() != previous {
			break
		}
	}

	if current.GetParent() == nil && current.GetLeft() == previous {
		return nil
	}

	var nextLeftRegularNumber *RegularNumber
	current = current.GetLeft()

	if current != nil {
		for nextLeftRegularNumber == nil {
			switch node := current.(type) {
			case *SnailfishNumber:
				current = node.GetRight()
			case *RegularNumber:
				nextLeftRegularNumber = node
			}
		}
	}

	return nextLeftRegularNumber
}

func (snailfishNumber *SnailfishNumber) NextRightRegularNumber() *RegularNumber {
	var previous Node = snailfishNumber
	var current Node = snailfishNumber

	for current.GetParent() != nil {
		previous = current
		current = current.GetParent()
		if current.GetRight() != previous {
			break
		}
	}

	if current.GetParent() == nil && current.GetRight() == previous {
		return nil
	}

	var nextRightRegularNumber *RegularNumber
	current = current.GetRight()

	if current != nil {
		for nextRightRegularNumber == nil {
			switch node := current.(type) {
			case *SnailfishNumber:
				current = node.GetLeft()
			case *RegularNumber:
				nextRightRegularNumber = node
			}
		}
	}

	return nextRightRegularNumber
}

func (snailfishNumber *SnailfishNumber) Split() bool {
	split := false

	switch node := snailfishNumber.left.(type) {
	case *RegularNumber:
		split = node.Split()
	case *SnailfishNumber:
		split = node.Split()
	default:
		log.Fatal("Unsupported node type.")
	}

	if !split {
		switch node := snailfishNumber.right.(type) {
		case *RegularNumber:
			split = node.Split()
		case *SnailfishNumber:
			split = node.Split()
		default:
			log.Fatal("Unsupported node type.")
		}
	}

	return split
}

func (snailfishNumber *SnailfishNumber) Add(other *SnailfishNumber) *SnailfishNumber {
	log.Debugf("Adding \"%s\" and \"%s\".", snailfishNumber, other)
	sum := &SnailfishNumber{}
	sum.SetLeft(snailfishNumber.Clone())
	sum.SetRight(other.Clone())
	log.Debugf("Sum before reducing is \"%s\".", sum)
	sum.Reduce()
	log.Debugf("Sum after reducing is \"%s\".", sum)
	return sum
}

func (snailfishNumber *SnailfishNumber) GetValue() int {
	return 3*snailfishNumber.GetLeft().GetValue() + 2*snailfishNumber.GetRight().GetValue()
}

func (snailfishNumber *SnailfishNumber) Clone() Node {
	clone := &SnailfishNumber{}
	clone.SetLeft(snailfishNumber.GetLeft().Clone())
	clone.SetRight(snailfishNumber.GetRight().Clone())
	return clone
}
