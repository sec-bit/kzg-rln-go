package keypair

import (
	"log"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
)

func RandomPolynomial(size int) []fr.Element {
	f := make([]fr.Element, size)
	for i := 0; i < size; i++ {
		f[i].SetRandom()
	}
	return f
}

func pRandomPolynomial(size int) []fr.Element {
	f := make([]fr.Element, size)
	for i := 0; i < size; i++ {
		f[i].SetInt64(int64(size - i))
	}
	return f
}

type KeyPairProof struct {
	H              bn254.G1Affine
	PrivateKey     fr.Element
	PublicKeyG1Aff bn254.G1Affine
	PublicKeyG2Aff bn254.G2Affine
}

func GenerateCommitmentAndKeyPairProof(poly []fr.Element, srs *kzg.SRS) (kzg.Digest, *KeyPairProof) {
	//commit the polynomial
	commitment, err := kzg.Commit(poly, srs)
	if err != nil {
		panic(err)
	}
	log.Printf("commitment: \nX: %02x\nY: %02x\n", commitment.X, commitment.Y)

	// compute opening proof at a random point
	var point fr.Element
	point.SetInt64(0)
	proof, err := kzg.Open(poly, point, srs)
	if err != nil {
		panic(err)
	}

	// claimed value is private key.
	// derive public key from private key
	privateKey := new(big.Int)
	proof.ClaimedValue.BigInt(privateKey)

	publicKey := new(bn254.G1Affine)
	publicKey.ScalarMultiplication(&srs.G1[0], privateKey)

	publicKeyG2 := new(bn254.G2Affine)
	publicKeyG2.ScalarMultiplication(&srs.G2[0], privateKey)

	pubKeyProof := new(KeyPairProof)
	pubKeyProof.PrivateKey = proof.ClaimedValue
	pubKeyProof.H = proof.H
	pubKeyProof.PublicKeyG1Aff = *publicKey
	pubKeyProof.PublicKeyG2Aff = *publicKeyG2
	return commitment, pubKeyProof
}

// VerifyPubKey verifies the public key of a user, copy from kzg.Verify
func VerifyPubKey(commitment *kzg.Digest, proof *KeyPairProof, srs *kzg.SRS) error {
	var point fr.Element
	point.SetInt64(0)
	// [f(a)]G₁
	// var claimedValueG1Aff bn254.G1Jac
	// var claimedValueBigInt big.Int
	// proof.ClaimedValue.BigInt(&claimedValueBigInt)
	// claimedValueG1Aff.ScalarMultiplicationAffine(&srs.G1[0], &claimedValueBigInt)

	// [f(α) - f(a)]G₁
	var fminusfaG1Jac bn254.G1Jac
	fminusfaG1Jac.FromAffine(commitment)
	pubKeyJac := new(bn254.G1Jac)
	pubKeyJac.FromAffine(&proof.PublicKeyG1Aff)
	fminusfaG1Jac.SubAssign(pubKeyJac)

	// [-H(α)]G₁
	var negH bn254.G1Affine
	negH.Neg(&proof.H)

	// [α-a]G₂
	var alphaMinusaG2Jac, genG2Jac, alphaG2Jac bn254.G2Jac
	var pointBigInt big.Int
	point.BigInt(&pointBigInt)
	genG2Jac.FromAffine(&srs.G2[0])
	alphaG2Jac.FromAffine(&srs.G2[1])
	alphaMinusaG2Jac.ScalarMultiplication(&genG2Jac, &pointBigInt).
		Neg(&alphaMinusaG2Jac).
		AddAssign(&alphaG2Jac)

	// [α-a]G₂
	var xminusaG2Aff bn254.G2Affine
	xminusaG2Aff.FromJacobian(&alphaMinusaG2Jac)

	// [f(α) - f(a)]G₁
	var fminusfaG1Aff bn254.G1Affine
	fminusfaG1Aff.FromJacobian(&fminusfaG1Jac)

	// e([f(α) - f(a)]G₁, G₂).e([-H(α)]G₁, [α-a]G₂) ==? 1
	check, err := bn254.PairingCheck(
		[]bn254.G1Affine{fminusfaG1Aff, negH},
		[]bn254.G2Affine{srs.G2[0], xminusaG2Aff},
	)
	if err != nil {
		return err
	}
	if !check {
		return kzg.ErrVerifyOpeningProof
	}
	return nil
}
