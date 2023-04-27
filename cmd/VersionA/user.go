package main

import (
	"crypto"
	"fmt"
	"log"

	"github.com/sec-bit/kzg-rln-go/keypair"
	"github.com/sec-bit/kzg-rln-go/types"

	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr/kzg"
)

type User struct {
	srs          *kzg.SRS
	polynomial   []fr.Element
	keyPairProof *keypair.KeyPairProof
	commitment   kzg.Digest
	socket       chan<- interface{}
	nonce        int
	points       []*types.Point
}

func NewUser(socket chan<- interface{}, limit int, srs *kzg.SRS) *User {
	poly := keypair.RandomPolynomial(limit)
	commitment, keyPairProof := keypair.GenerateCommitmentAndKeyPairProof(poly, srs)

	user := new(User)
	user.srs = srs
	user.polynomial = poly
	user.keyPairProof = keyPairProof
	user.commitment = commitment
	user.socket = socket
	log.Printf("User created\n Private key: %s\n", keyPairProof.PrivateKey.String())
	return user
}

func (u *User) Register() error {
	keyPairProofWithoutPrivate := &keypair.KeyPairProof{
		H:              u.keyPairProof.H,
		PublicKeyG1Jac: u.keyPairProof.PublicKeyG1Jac,
	}
	user := &User{
		keyPairProof: keyPairProofWithoutPrivate,
		commitment:   u.commitment,
	}
	u.socket <- user
	return nil
}

func (u *User) SendMessage(text string) error {
	var frMsg fr.Element
	hasher := crypto.SHA256.New()
	hasher.Write([]byte(fmt.Sprint(u.nonce)))
	hasher.Write([]byte(text))
	frMsg.SetBytes(hasher.Sum(nil))
	msgProof, err := kzg.Open(u.polynomial, frMsg, u.srs)
	if err != nil {
		return err
	}
	msg := &types.Message{Commitment: u.commitment, Text: text, Proof: msgProof}
	u.socket <- msg
	u.nonce++
	return nil
}
