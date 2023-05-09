package keypair

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
)

func TestFakeKeyPairProof(t *testing.T) {
	srs, err := kzg.NewSRS(16, big.NewInt(42))
	if err != nil {
		panic(err)
	}
	commitment, proof := GenerateCommitmentAndKeyPairProof(pRandomPolynomial(10), srs)
	proof.PublicKeyG2Aff = srs.G2[0]
	gamma := big.NewInt(114514)
	g1Gamma := new(bn254.G1Affine)
	g1Gamma.ScalarMultiplication(&srs.G1[0], gamma)
	g1AlphaGamma := new(bn254.G1Affine)
	g1AlphaGamma.ScalarMultiplication(&srs.G1[1], gamma)
	fmt.Printf("H: \nX: %02x\nY: %02x\n", proof.H.X, proof.H.Y)
	proof.H.Add(&proof.H, g1Gamma)
	fmt.Printf("H(adj): \nX: %02x\nY: %02x\n", proof.H.X, proof.H.Y)
	fmt.Printf("public key: \nX: %02x\nY: %02x\n", proof.PublicKeyG1Aff.X, proof.PublicKeyG1Aff.Y)
	var negG1AlphaGamma bn254.G1Affine
	negG1AlphaGamma.Neg(g1AlphaGamma)
	proof.PublicKeyG1Aff.Add(&proof.PublicKeyG1Aff, &negG1AlphaGamma)
	fmt.Printf("public key(adj): \nX: %02x\nY: %02x\n", proof.PublicKeyG1Aff.X, proof.PublicKeyG1Aff.Y)
	err = VerifyPubKey(&commitment, proof, srs)
	if err != nil {
		panic(err)
	}
	priv := proof.PrivateKey
	fmt.Printf("private key: %s\n", priv.String())
	gammaFr := new(fr.Element)
	gammaFr.SetBigInt(gamma)
	priv.Sub(priv, gammaFr)
	fmt.Printf("private key(adj): %s\n", priv.String())
	var pubKey bn254.G1Affine
	privBN := new(big.Int)
	priv.BigInt(privBN)
	pubKey.ScalarMultiplication(&srs.G1[0], privBN)
	fmt.Printf("public key(adj): \nX: %02x\nY: %02x\n", pubKey.X, pubKey.Y)
}
