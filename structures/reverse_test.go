package structures_test

import (
	"testing"

	"github.com/mmaxemm/internship_tasks/structures"
)

var reverseTests = []struct {
	name string // description of this test case
	// Named input parameters for target function.
	str  string
	want string
}{
	{"0 chars", "", ""},
	{"1 char", "a", "a"},
	{"even number of chars", "123456", "654321"},
	{"odd number of chars", "12345", "54321"},
	{">1 words", "hello world", "olleh dlrow"},
	{"all spaces", "    ",  "    "},
	{"spaces at start", "   some words", "   emos sdrow"},
	{"spaces at the end", "some words   ", "emos sdrow   "},
	{"tab instead of space", "some	words", "emos	sdrow"},
}

func TestReverse1(t *testing.T) {
	for _, tt := range reverseTests {
		t.Run(tt.name, func(t *testing.T) {
			got := structures.Reverse1(tt.str)
			if got != tt.want {
				t.Errorf("Reverse1() = %v, want %v on test '%s'", got, tt.want, tt.name)
			}
		})
	}
}


func TestReverse2(t *testing.T) {
	for _, tt := range reverseTests {
		t.Run(tt.name, func(t *testing.T) {
			got := structures.Reverse2(tt.str)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("Reverse2() = %v, want %v on test '%s'", got, tt.want, tt.name)
			}
		})
	}
}

