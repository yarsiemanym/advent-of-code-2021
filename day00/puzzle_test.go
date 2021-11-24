package day00

import "testing"

func Test_Solve_Input1(t *testing.T) {
	input := "test1.txt"
	expected := "User 0 = { name: Joe Schmoe, email: joe@email.com, birthday: 1983-11-24 }\nUser 1 = { name: John Q Public, email: johnyq@email.com, birthday: 1997-04-03 }\n"
	actual := Solve(input)

	if actual != expected {
		t.Errorf("Expected:\n%v\nActual:\n%v", expected, actual)
	}
}
