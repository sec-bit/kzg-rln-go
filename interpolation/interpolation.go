package interpolation

import (
	"math/big"

	"github.com/sec-bit/kzg-rln-go/types"

	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
)

func fromString(s string) *fr.Element {
	bi := new(big.Int)
	bi.SetString(s, 10)
	result := &fr.Element{}
	result.SetBigInt(bi)
	return result
}

func getWeightsByPoints(points []*types.Point) []*fr.Element {
	xValues := make([]*fr.Element, len(points))
	for i := 0; i < len(points); i++ {
		xValues[i] = &points[i].X
	}
	return barycentricInterpolation(xValues)
}

// barycentricInterpolation calculates the barycentric weights for the given xValues.
func barycentricInterpolation(xValues []*fr.Element) []*fr.Element {
	n := len(xValues)
	weights := make([]*fr.Element, n)

	for i := 0; i < n; i++ {
		one := fr.One()
		weights[i] = &one
		for j := 0; j < n; j++ {
			if i != j {
				var difference fr.Element
				difference.Sub(xValues[i], xValues[j])
				weights[i].Mul(weights[i], &difference)
			}
		}
		one2 := fr.One()
		weights[i].Div(&one2, weights[i])
	}

	return weights
}

// RecoverPrivateKeyByPoints recovers the private key from the given points.
func RecoverPrivateKeyByPoints(points []*types.Point) *fr.Element {
	zeroX := fr.NewElement(0)
	return evaluateBarycentricPolynomialByPoints(&zeroX, points, getWeightsByPoints(points))
}

func evaluateBarycentricPolynomialByPoints(x *fr.Element, points []*types.Point, weights []*fr.Element) *fr.Element {
	xValues := make([]*fr.Element, len(points))
	yValues := make([]*fr.Element, len(points))
	for i := 0; i < len(points); i++ {
		xValues[i] = &points[i].X
		yValues[i] = &points[i].Y
	}
	return evaluateBarycentricPolynomial(x, xValues, yValues, weights)
}

// evaluateBarycentricPolynomial evaluates the interpolated polynomial at a given x value.
func evaluateBarycentricPolynomial(x *fr.Element, xValues, yValues, weights []*fr.Element) *fr.Element {
	n := len(xValues)
	numerator := fr.NewElement(0)
	denominator := fr.NewElement(0)

	for i := 0; i < n; i++ {
		if x.Equal(xValues[i]) {
			return yValues[i]
		}

		temp := fr.NewElement(0)
		temp.Sub(x, xValues[i])
		temp.Div(weights[i], &temp)
		var temp2 fr.Element
		temp2.Mul(&temp, yValues[i])
		numerator.Add(&numerator, &temp2)
		denominator.Add(&denominator, &temp)
	}
	var result fr.Element
	result.Div(&numerator, &denominator)
	return &result
}
