package tinycanvas

import "testing"

// ----------------------------------------------------------------------------
func Test_getSlopes(t *testing.T) {
	tests := []struct {
		name                           string
		x1, y1, x2, y2                 int
		expectedSlopeX, expectedSlopeY int
	}{
		{"0", 0, 0, 0, 0, -1, -1},
		{"horizontal", 10, 5, 100, 5, -1, -1},
		{"vertical", 10, 5, 10, 300, -1, 1},
		{"left down right", 10, 10, 100, 100, -1, 1},
		{"left down left", 100, 10, 10, 100, -1, 1},
	}

	for _, test := range tests {
		slopeX, slopeY := getSlopes(test.x2, test.y1, test.x2, test.y2)
		if slopeX != test.expectedSlopeX || slopeY != test.expectedSlopeY {
			t.Errorf("%s failed with %d,%d : expected %d,%d", test.name, slopeX, slopeY, test.expectedSlopeX, test.expectedSlopeY)
		}
	}
}
