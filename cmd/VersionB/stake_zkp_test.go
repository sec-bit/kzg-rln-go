package main

import (
	"testing"

	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
)

func TestMerkleProof(t *testing.T) {
	merkletree.New(mimc.NewMiMC())
}
