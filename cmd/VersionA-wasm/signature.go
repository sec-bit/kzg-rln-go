package main

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"golang.org/x/crypto/sha3"
)

func hashToPoint(data []byte) bn254.G1Affine {
	var X, Y fp.Element
	X.SetString("0x059dac1925a1d0bee704dd2ae3836a3d8e76a4c4249f17860ce1d0a530c5f8f7")
	Y.SetString("0x03870b29cb77fab35c1394ac29e19344465046309674e8d138da412f834ecaee")
	nothingUpMySleeve := &bn254.G1Affine{
		X: X,
		Y: Y,
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	hashedData := hash.Sum(nil)
	hashedBigInt := new(big.Int).SetBytes(hashedData)

	var result bn254.G1Affine
	result.ScalarMultiplication(nothingUpMySleeve, hashedBigInt)

	return result
}

func sign(message []byte, privateKey *fr.Element) (bn254.G2Affine, bn254.G1Affine, bn254.G1Affine) {
	// Hash the message to a G1 point
	hashedMsg := hashToPoint(message)
	// log.Printf("msgHash = %s", hashedMsg.String())

	_, _, _, g2Aff := bn254.Generators()

	// Compute the signature (in G1) by multiplying the private key with the generator of G1
	var signature bn254.G1Affine
	signature.ScalarMultiplication(&hashedMsg, privateKey.BigInt(new(big.Int)))

	// Compute the public key (in G2) by multiplying the private key with the generator of G2
	var publicKey bn254.G2Affine
	privateKeyBN := new(big.Int)
	privateKey.BigInt(privateKeyBN)
	// fmt.Printf("privateKey = %s\n", privateKey.String())
	// fmt.Printf("privateKeyBN = %s\n", privateKeyBN.String())
	publicKey.ScalarMultiplication(&g2Aff, privateKeyBN)

	// var negG2Aff bn254.G2Affine
	// negG2Aff.Neg(&g2Aff)

	// negXA0 := negG2Aff.X.A0
	// negXA1 := negG2Aff.X.A1
	// negYA0 := negG2Aff.Y.A0
	// negYA1 := negG2Aff.Y.A1
	// fmt.Printf("G2Affine: X A0: %02X, X A1: %02X, Y A0: %02X, Y A1: %02X\n", negXA0.Bytes(), negXA1.Bytes(), negYA0.Bytes(), negYA1.Bytes())

	return publicKey, hashedMsg, signature
}

func verify(message []byte, publicKey bn254.G2Affine, signature bn254.G1Affine) (bool, error) {
	// Hash the message to a G1 point
	hashedMsg := hashToPoint(message)
	isOnCurve := hashedMsg.IsOnCurve()
	fmt.Printf("isOnCurve = %t\n", isOnCurve)
	// log.Printf("msgHash = %s", hashedMsg.String())

	// Pairing check: e(hashedMsg, publicKey) == e(signature, g2)
	eLeft, err := bn254.Pair([]bn254.G1Affine{hashedMsg}, []bn254.G2Affine{publicKey})
	if err != nil {
		return false, err
	}

	_, _, _, g2Aff := bn254.Generators()
	var g2AffNeg bn254.G2Affine
	g2AffNeg.Neg(&g2Aff)
	fmt.Printf("g2AffNeg.X.A0 = %02X\n", g2AffNeg.X.A0.Bytes())
	fmt.Printf("g2AffNeg.X.A1 = %02X\n", g2AffNeg.X.A1.Bytes())
	fmt.Printf("g2AffNeg.Y.A0 = %02X\n", g2AffNeg.Y.A0.Bytes())
	fmt.Printf("g2AffNeg.Y.A1 = %02X\n", g2AffNeg.Y.A1.Bytes())
	eRight, err := bn254.Pair([]bn254.G1Affine{signature}, []bn254.G2Affine{g2Aff})
	if err != nil {
		return false, err
	}

	return eLeft.Equal(&eRight), nil
}
