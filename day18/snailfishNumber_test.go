package day18

import "testing"

func Test_NewSnailfishNumber(t *testing.T) {
	textIn := "[[1,2],3]"
	snailfishNumber := NewSnailfishNumber(textIn)
	textOut := snailfishNumber.String()

	if textIn != textOut {
		t.Errorf("Expected \"%s\" but got \"%s\".", textIn, textOut)
	}
}

func Test_SnailfishNumber_NextLeftRegularNumber(t *testing.T) {
	textIn := "[7,[6,[5,[4,[3,2]]]]]"
	snailfishNumber := NewSnailfishNumber(textIn)
	start := snailfishNumber.GetRight().GetRight().GetRight().GetRight().(*SnailfishNumber)
	nextLeft := start.NextLeftRegularNumber()

	if nextLeft == nil {
		t.Error("nextLeft is nil.")
	} else if nextLeft.GetValue() != 4 {
		t.Errorf("Expected 4 but got %d.", nextLeft.GetValue())
	}
}

func Test_SnailfishNumber_NextRightRegularNumber(t *testing.T) {
	textIn := "[[[[[9,8],1],2],3],4]"
	snailfishNumber := NewSnailfishNumber(textIn)
	start := snailfishNumber.GetLeft().GetLeft().GetLeft().GetLeft().(*SnailfishNumber)
	nextRight := start.NextRightRegularNumber()

	if nextRight == nil {
		t.Error("nextRight is nil.")
	} else if nextRight.GetValue() != 1 {
		t.Errorf("Expected 1 but got %d.", nextRight.GetValue())
	}
}

func Test_SnailfishNumber_ExplodeFirst(t *testing.T) {
	textIn := "[[[[[9,8],1],2],3],4]"
	snailfishNumber := NewSnailfishNumber(textIn)
	snailfishNumber.Reduce()
	textOut := snailfishNumber.String()

	if textOut != "[[[[0,9],2],3],4]" {
		t.Errorf("Expected \"[[[[0,9],2],3],4]\" but got \"%s\".", textOut)
	}
}

func Test_SnailfishNumber_ExplodeLast(t *testing.T) {
	textIn := "[7,[6,[5,[4,[3,2]]]]]"
	snailfishNumber := NewSnailfishNumber(textIn)
	snailfishNumber.Reduce()
	textOut := snailfishNumber.String()

	if textOut != "[7,[6,[5,[7,0]]]]" {
		t.Errorf("Expected \"[7,[6,[5,[7,0]]]]\" but got \"%s\".", textOut)
	}
}

func Test_SnailfishNumber_ExplodeInterior(t *testing.T) {
	textIn := "[[6,[5,[4,[3,2]]]],1]"
	snailfishNumber := NewSnailfishNumber(textIn)
	snailfishNumber.Reduce()
	textOut := snailfishNumber.String()

	if textOut != "[[6,[5,[7,0]]],3]" {
		t.Errorf("Expected \"[[6,[5,[7,0]]],3]\" but got \"%s\".", textOut)
	}
}

func Test_SnailfishNumber_ExplodeMultiple(t *testing.T) {
	textIn := "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"
	snailfishNumber := NewSnailfishNumber(textIn)
	snailfishNumber.Reduce()
	textOut := snailfishNumber.String()

	if textOut != "[[3,[2,[8,0]]],[9,[5,[7,0]]]]" {
		t.Errorf("Expected \"[[3,[2,[8,0]]],[9,[5,[7,0]]]]\" but got \"%s\".", textOut)
	}
}

func Test_SnailfishNumber_Split(t *testing.T) {
	textIn := "[1,B]"
	snailfishNumber := NewSnailfishNumber(textIn)
	snailfishNumber.Reduce()
	textOut := snailfishNumber.String()

	if textOut != "[1,[5,6]]" {
		t.Errorf("Expected \"[1,[5,6]]\" but got \"%s\".", textOut)
	}
}

func Test_SnailfishNumber_Add(t *testing.T) {
	number1 := NewSnailfishNumber("[1,2]")
	number2 := NewSnailfishNumber("[[3,4],5]")
	sum := number1.Add(number2)
	textOut := sum.String()

	if textOut != "[[1,2],[[3,4],5]]" {
		t.Errorf("Expected \"[[1,2],[[3,4],5]]\" but got \"%s\".", textOut)
	}
}

func Test_SnailfishNumber_AddAndReduce(t *testing.T) {
	number1 := NewSnailfishNumber("[[[[4,3],4],4],[7,[[8,4],9]]]")
	number2 := NewSnailfishNumber("[1,1]")
	sum := number1.Add(number2)
	textOut := sum.String()

	if textOut != "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]" {
		t.Errorf("Expected \"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]\" but got \"%s\".", textOut)
	}
}
