package structures_test

import (
	"reflect"
	"testing"

	"github.com/mmaxemm/internship_tasks/structures"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		sl   []string
		want []string
	}{
		{"empty slice", []string{}, []string{}},
		{"1 element slice", []string{"a"}, []string{"a"}},
		{"no repeated values", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"2 repeated values",[]string{"a", "b", "c", "c"}, []string{"a", "b", "c"}},
		{"3 repeated values", []string{"a", "b", "c", "c", "c"}, []string{"a", "b", "c"}},
		{"2 values are repeated", []string{"a", "a", "b", "b", "c"}, []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := structures.Unique(tt.sl)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

