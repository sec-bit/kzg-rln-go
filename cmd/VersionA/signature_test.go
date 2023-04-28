package main

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
)

func TestHashToPoint(t *testing.T) {
	hash := hashToPoint([]byte("hello"))
	fmt.Printf("hashToPoint = %s", hash.String())
}

func TestSign(t *testing.T) {
	var privateKey fr.Element
	privateKey.SetRandom()

	message, _ := hex.DecodeString("0000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc")
	publicKey, hashedMsg, signature := sign(message, &privateKey)

	fmt.Printf("Public key (G2): %s\n", publicKey.String())
	fmt.Printf("Message: %s\n", string(message))
	fmt.Printf("Hashed message (G1): %s\n", hashedMsg.String())
	fmt.Printf("Signature (G1): %s\n", signature.String())

	// Print values in hex format for use in the Solidity test
	pubKeyBXA0 := publicKey.X.A0.Bytes()
	pubKeyBXA1 := publicKey.X.A1.Bytes()
	pubKeyBYA0 := publicKey.Y.A0.Bytes()
	pubKeyBYA1 := publicKey.Y.A1.Bytes()
	hashedMsgX := hashedMsg.X.Bytes()
	hashedMsgY := hashedMsg.Y.Bytes()
	signatureX := signature.X.Bytes()
	signatureY := signature.Y.Bytes()

	fmt.Printf("Public key (G2):\nX A0: 0x%s\nX A1: 0x%s\n", hex.EncodeToString(pubKeyBXA0[:]), hex.EncodeToString(pubKeyBXA1[:]))
	fmt.Printf("Y A0: 0x%s\nY A1: 0x%s\n", hex.EncodeToString(pubKeyBYA0[:]), hex.EncodeToString(pubKeyBYA1[:]))
	fmt.Printf("Hashed message (G1):\nX: 0x%s\nY: 0x%s\n", hex.EncodeToString(hashedMsgX[:]), hex.EncodeToString(hashedMsgY[:]))
	fmt.Printf("Signature (G1):\nX: 0x%s\nY: 0x%s\n", hex.EncodeToString(signatureX[:]), hex.EncodeToString(signatureY[:]))
}

func TestVerify(t *testing.T) {
	var privateKey fr.Element
	privateKey.SetRandom()

	message, _ := hex.DecodeString("0000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc")
	publicKey, _, signature := sign(message, &privateKey)

	valid, err := verify(message, publicKey, signature)
	if err != nil {
		t.Fatal(err)
	}

	if !valid {
		t.Fatal("Invalid signature")
	}
}
