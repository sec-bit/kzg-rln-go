package main

import (
	"crypto"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sec-bit/kzg-rln-go/keypair"
	"github.com/sec-bit/kzg-rln-go/types"
	"github.com/sec-bit/kzg-rln-go/utils"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
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
	log.Printf("User created\n Private key: %s\n\nPublicKey: %s", keyPairProof.PrivateKey.String(), keyPairProof.PublicKeyG1Aff.String())
	return user
}

func (u *User) Register(client *ethclient.Client, stake *Stake) error {
	auth := utils.NewTransactor(client, "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")
	auth.Value = big.NewInt(1e18) // in wei
	keyPairProofWithoutPrivate := &keypair.KeyPairProof{
		H:              u.keyPairProof.H,
		PublicKeyG1Aff: u.keyPairProof.PublicKeyG1Aff,
		PublicKeyG2Aff: u.keyPairProof.PublicKeyG2Aff,
	}
	user := &User{
		keyPairProof: keyPairProofWithoutPrivate,
		commitment:   u.commitment,
	}

	pubKeyXA0 := new(big.Int)
	pubKeyXA1 := new(big.Int)
	pubKeyYA0 := new(big.Int)
	pubKeyYA1 := new(big.Int)
	u.keyPairProof.PublicKeyG2Aff.X.A0.BigInt(pubKeyXA0)
	u.keyPairProof.PublicKeyG2Aff.X.A1.BigInt(pubKeyXA1)
	u.keyPairProof.PublicKeyG2Aff.Y.A0.BigInt(pubKeyYA0)
	u.keyPairProof.PublicKeyG2Aff.Y.A1.BigInt(pubKeyYA1)
	var pubKeyG2 BN254HashToG1G2Point
	pubKeyG2.X = [2]*big.Int{pubKeyXA0, pubKeyXA1}
	pubKeyG2.Y = [2]*big.Int{pubKeyYA0, pubKeyYA1}

	_, err := stake.Deposit(auth, pubKeyG2)
	if err != nil {
		log.Fatal(err)
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
