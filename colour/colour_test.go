package colour

import (
	"testing"
)

// ----------------------------------------------------------------------------
func Test_createNewColour(t *testing.T) {
	const valR uint8 = 10
	const valG uint8 = 20
	const valB uint8 = 30
	const valA uint8 = 40

	tmpC := NewColour(valR, valG, valB, valA)

	if tmpC.A != valA || tmpC.R != valR || tmpC.G != valG || tmpC.B != valB {
		t.Error("Colour values do not match.")
	}
}

// ----------------------------------------------------------------------------
func Test_createWhiteColour(t *testing.T) {
	tmpC := NewColourWhite()

	if tmpC.A != MAX_COLOUR_VALUE || tmpC.B != MAX_COLOUR_VALUE || tmpC.G != MAX_COLOUR_VALUE || tmpC.R != MAX_COLOUR_VALUE {
		t.Error("Colour value for WHITE is wrong.")
	}
}

// ----------------------------------------------------------------------------
func Test_IsEmpty(t *testing.T) {
	// Test empty colour
	emptyColour := NewColourEmpty()
	if !emptyColour.IsEmpty() {
		t.Error("Empty colour should return true for IsEmpty()")
	}

	// Test non-empty colour
	nonEmptyColour := NewColour(10, 20, 30, 40)
	if nonEmptyColour.IsEmpty() {
		t.Error("Non-empty colour should return false for IsEmpty()")
	}

	// Test colour with only alpha
	alphaOnly := NewColour(0, 0, 0, 100)
	if alphaOnly.IsEmpty() {
		t.Error("Colour with only alpha should return false for IsEmpty()")
	}
}

// ----------------------------------------------------------------------------
func Test_NewRandomColour(t *testing.T) {
	// Generate a few random colours to ensure they're valid
	for _ = range 5 {
		randomColour := NewRandomColour()

		// Alpha should always be MAX_COLOUR_VALUE
		if randomColour.A != MAX_COLOUR_VALUE {
			t.Errorf("Random colour alpha should be %d, got %d", MAX_COLOUR_VALUE, randomColour.A)
		}

		// All values should be within valid range (0-255)
		if randomColour.R > MAX_COLOUR_VALUE || randomColour.G > MAX_COLOUR_VALUE || randomColour.B > MAX_COLOUR_VALUE {
			t.Error("Random colour components should be within valid range (0-255)")
		}
	}
}

// ----------------------------------------------------------------------------
func Test_ConvertToGrayscale(t *testing.T) {
	// Test with a known colour
	colour := NewColour(100, 150, 200, 255)
	colour.ConvertToGrayscale()

	// All RGB values should be the same after conversion
	if colour.R != colour.G || colour.G != colour.B {
		t.Error("After grayscale conversion, R, G, and B should be equal")
	}

	// Alpha should remain unchanged
	if colour.A != 255 {
		t.Error("Alpha should remain unchanged after grayscale conversion")
	}

	// Test with white (should remain white)
	white := NewColourWhite()
	white.ConvertToGrayscale()
	if white.R != MAX_COLOUR_VALUE || white.G != MAX_COLOUR_VALUE || white.B != MAX_COLOUR_VALUE {
		t.Error("White should remain white after grayscale conversion")
	}

	// Test with black (should remain black)
	black := NewColourBlack()
	black.ConvertToGrayscale()
	if black.R != 0 || black.G != 0 || black.B != 0 {
		t.Error("Black should remain black after grayscale conversion")
	}
}

// ----------------------------------------------------------------------------
func Test_createBlackColour(t *testing.T) {
	tmpC := NewColourBlack()

	if tmpC.A != MAX_COLOUR_VALUE || tmpC.B != 0 || tmpC.G != 0 || tmpC.R != 0 {
		t.Error("Colour value for BLACK is wrong.")
	}
}

// ----------------------------------------------------------------------------
func Test_createEmptyColour(t *testing.T) {
	tmpC := NewColourEmpty()

	if tmpC.A != 0 || tmpC.B != 0 || tmpC.G != 0 || tmpC.R != 0 {
		t.Error("Colour value for EMPTY is wrong.")
	}
}
