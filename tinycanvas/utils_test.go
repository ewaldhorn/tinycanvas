package tinycanvas

import (
	"math"
	"testing"
)

// ----------------------------------------------------------------------------
func Test_abs(t *testing.T) {
	tests := []struct {
		name        string
		inputVal    int
		expectedVal int
	}{
		{"0", 0, 0},
		{"-10", -10, 10},
		{"-1", -1, 1},
		{"2", 2, 2},
		{"maxInt", math.MaxInt, math.MaxInt},
		{"minInt", math.MinInt + 1, math.MaxInt}, // Extreme edge case
	}

	for _, test := range tests {
		result := abs(test.inputVal)

		if result != test.expectedVal {
			t.Errorf("%s: failed with %d, expected %d", test.name, result, test.expectedVal)
		}
	}
}
