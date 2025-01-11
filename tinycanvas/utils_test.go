package tinycanvas

import "testing"

// ----------------------------------------------------------------------------
func Test_abs(t *testing.T) {
	tests := []struct {
		name        string
		inputVal    int
		expectedVal int
	}{
		{"0", 0, 0},
	}

	for _, test := range tests {
		result := abs(test.inputVal)

		if result != test.expectedVal {
			t.Errorf("%s: failed with %d, expected %d", test.name, result, test.expectedVal)
		}
	}
}
