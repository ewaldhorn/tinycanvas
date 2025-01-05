package tinycanvas

import "testing"

// ----------------------------------------------------------------------------
func Test_createNewColour(t *testing.T) {
	const valR uint8 = 10
	const valG uint8 = 20
	const valB uint8 = 30
	const valA uint8 = 40

	tmpC := NewColour(valR, valG, valB, valA)

	if tmpC.a != valA || tmpC.r != valR || tmpC.g != valG || tmpC.b != valB {
		t.Error("Colour values do not match.")
	}
}

// ----------------------------------------------------------------------------
func Test_createWhiteColour(t *testing.T) {
	tmpC := NewColourWhite()

	if tmpC.a != MAX_COLOUR_VALUE || tmpC.b != MAX_COLOUR_VALUE || tmpC.g != MAX_COLOUR_VALUE || tmpC.r != MAX_COLOUR_VALUE {
		t.Error("Colour value for WHITE is wrong.")
	}
}

// ----------------------------------------------------------------------------
func Test_createBlackColour(t *testing.T) {
	tmpC := NewColourBlack()

	if tmpC.a != MAX_COLOUR_VALUE || tmpC.b != 0 || tmpC.g != 0 || tmpC.r != 0 {
		t.Error("Colour value for BLACK is wrong.")
	}
}

// ----------------------------------------------------------------------------
func Test_createEmptyColour(t *testing.T) {
	tmpC := NewColourEmpty()

	if tmpC.a != 0 || tmpC.b != 0 || tmpC.g != 0 || tmpC.r != 0 {
		t.Error("Colour value for EMPTY is wrong.")
	}
}
