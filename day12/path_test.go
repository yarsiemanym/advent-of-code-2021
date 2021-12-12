package day12

import "testing"

func Test_Path_Clone(t *testing.T) {
	path1 := NewPath([]*Cave{NewCave("A")})
	path2 := path1.Clone()
	path1.Add(NewCave("B"))

	if path2.Length() != 1 {
		t.Errorf("Expected 1 but got %v.", path2.Length())
	}

	if path1.nodes[0] != path2.nodes[0] {
		t.Error("Pointers point to different structs.")
	}
}
