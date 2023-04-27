package types

import (
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr/kzg"
)

type Point struct {
	X, Y fr.Element
}

type Message struct {
	Commitment bls12381.G1Affine
	Text       string
	Proof      kzg.OpeningProof
}
