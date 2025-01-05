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
		t.Error("Colour values do not match")
	}
}
