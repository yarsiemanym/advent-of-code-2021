package day18

type Node interface {
	GetParent() Node
	SetParent(parent Node)
	String() string
	GetLeft() Node
	GetRight() Node
	GetValue() int
	Clone() Node
}
