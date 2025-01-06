package colour

import "testing"

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
