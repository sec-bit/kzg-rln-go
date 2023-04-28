package interpolation

import (
	"fmt"
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
)

func TestInterpolate(t *testing.T) {
	xValues := []*fr.Element{
		fromString("1"),
		fromString("2"),
		fromString("3"),
	}
	yValues := []*fr.Element{
		fromString("1"),
		fromString("4"),
		fromString("9"),
	}

	xNew := fromString("99")
	weights := barycentricInterpolation(xValues)
	yNew := evaluateBarycentricPolynomial(xNew, xValues, yValues, weights)
	fmt.Printf("Value of the interpolated polynomial at x=%s: %s\n", xNew.String(), yNew.String())
}
