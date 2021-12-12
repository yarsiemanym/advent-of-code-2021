package day12

import (
	"fmt"
	"strings"
)

type Path struct {
	nodes []*Cave
}

func NewPath(nodes []*Cave) *Path {
	return &Path{
		nodes: nodes,
	}
}

func (path *Path) Contains(cave *Cave) bool {
	found := false

	for _, node := range path.nodes {
		if node == cave {
			found = true
			break
		}
	}

	return found
}

func (path *Path) Nodes() []*Cave {
	return path.nodes
}

func (path *Path) Add(newNodes ...*Cave) {
	path.nodes = append(path.nodes, newNodes...)
}

func (path *Path) Length() int {
	return len(path.nodes)
}

func (path *Path) Start() *Cave {
	if len(path.nodes) <= 0 {
		return nil
	}

	return path.nodes[0]
}

func (path *Path) NodeAt(index int) *Cave {
	if len(path.nodes) <= index {
		return nil
	}

	return path.nodes[index]
}

func (path *Path) End() *Cave {
	if len(path.nodes) <= 0 {
		return nil
	}

	return path.nodes[len(path.nodes)-1]
}

func (path *Path) Clone() *Path {
	nodes := make([]*Cave, path.Length())
	copy(nodes, path.nodes)
	clone := NewPath(nodes)
	return clone
}

func (path *Path) Render() string {
	output := "Path: "

	for _, node := range path.nodes {
		output += fmt.Sprintf("\"%v\" -> ", node.Name())
	}

	output = strings.Trim(output, " ->")
	return output
}
