package main

import (
	"math/big"
	"testing"

	"github.com/sec-bit/kzg-rln-go/keypair"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/test"
)

func TestMembershipGroth16(t *testing.T) {
	// private: cofficients, pk
	// public: hash, f(hash), hash(pk)
	// f(0)=cofficients[0] =? pk
	// hash =? mimc(coefficients)
	// f(hash) =? opening.ClaimedValue
	// hash(pk) =? hashPK
	// commitment --> server
	// kzg.Verify(commitment, opening, hash, srs)
	// zkp.Verify(pi, hash, f(hash))
	// register with hashPK, this is also used as nullifier
	// slash with proof of pk avoid interpolating

	const srsSize = 230
	srs, _ := kzg.NewSRS(ecc.NextPowerOfTwo(srsSize), new(big.Int).SetInt64(42))

	assert := test.NewAssert(t)

	var circuit registerCircuit
	var assign registerCircuit

	poly := keypair.RandomPolynomial(DEGREE)
	_, keyPairProof := keypair.GenerateCommitmentAndKeyPairProof(poly, srs)
	privateKey := keyPairProof.PrivateKey

	hasher := mimc.NewMiMC()
	for i := 0; i < DEGREE; i++ {
		tmp := poly[i].Bytes()
		hasher.Write(tmp[:])
	}
	coeffHashB := hasher.Sum(nil)
	assign.CoeffHash = coeffHashB
	hashFr := fr.NewElement(0)
	hashFr.SetBytes(coeffHashB)
	proof, err := kzg.Open(poly, hashFr, srs)
	if err != nil {
		panic(err)
	}
	assign.PolyEvalAtHash = proof.ClaimedValue
	// convert poly to []frontend.Variable
	privateKeyBN := new(big.Int)
	assign.PrivateKey = privateKey.BigInt(privateKeyBN)
	for i := 0; i < DEGREE; i++ {
		assign.Coefficients[i] = poly[i]
	}

	assert.ProverSucceeded(&circuit, &assign, test.WithCurves(ecc.BLS12_381), test.WithBackends(backend.GROTH16))

}
