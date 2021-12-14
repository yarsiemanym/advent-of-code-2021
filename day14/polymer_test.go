package day14

import "testing"

func Test_Polymer_Insert_OneRule_OneInsertion(t *testing.T) {
	polymer := NewPolymer("NNCB")
	rule := NewInsertionRule("NN", "C")

	polymer.Insert(rule)

	if polymer.Render() != "NCNCB" {
		t.Errorf("Expected \"NCNCB\" but got \"%s\".", polymer.Render())
	}
}

func Test_Polymer_Insert_TwoRules_OneInsertionEach(t *testing.T) {
	polymer := NewPolymer("NNCB")
	rules := []*InsertionRule{
		NewInsertionRule("NN", "C"),
		NewInsertionRule("NC", "B"),
	}

	polymer.Insert(rules...)

	if polymer.Render() != "NCNBCB" {
		t.Errorf("Expected \"NCNBCB\" but got \"%s\".", polymer.Render())
	}
}

func Test_Polymer_Insert_OneRule_MultipleInsertions(t *testing.T) {
	polymer := NewPolymer("NNCNNB")
	rules := []*InsertionRule{
		NewInsertionRule("NN", "C"),
	}

	polymer.Insert(rules...)

	if polymer.Render() != "NCNCNCNB" {
		t.Errorf("Expected \"NCNCNCNB\" but got \"%s\".", polymer.Render())
	}
}
