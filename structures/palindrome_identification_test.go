package structures_test

import (
	"testing"

	"github.com/mmaxemm/internship_tasks/structures"
)

var tests = []struct {
	name string // description of this test case
	// Named input parameters for target function.
	str  string
	want bool
}{
	{"0 characters", "", true},
	{"1 character", "a", true},
	{"sentance", "Mr. Owl ate my metal worm", true},
	{"odd number of chars", "radar", true},
	{"even number of chars", "ABBA", true},
	{"not a palindrome even", "aabb", false},
	{"not a palindrome odd", "aab", false},
	{"some digits", "22\\2\\22", true},
}

func TestIsPalindrome1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := structures.IsPalindrome1(tt.str)
			if got != tt.want {
				t.Errorf("IsPalindrome1() = %v, want %v on test '%s'", got, tt.want, tt.name)
			}
		})
	}
}


func TestIsPalindrome2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := structures.IsPalindrome2(tt.str)
			if got != tt.want {
				t.Errorf("IsPalindrome1() = %v, want %v on test '%s'", got, tt.want, tt.name)
			}
		})
	}
}


func TestIsPalindrome3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := structures.IsPalindrome3(tt.str)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("IsPalindrome1() = %v, want %v on test '%s'", got, tt.want, tt.name)
			}
		})
	}
}

