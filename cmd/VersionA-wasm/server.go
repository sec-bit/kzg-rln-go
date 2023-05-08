package main

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"syscall/js"
	"time"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
	"github.com/sec-bit/kzg-rln-go/interpolation"
	"github.com/sec-bit/kzg-rln-go/keypair"
	"github.com/sec-bit/kzg-rln-go/types"
	"golang.org/x/crypto/sha3"
)

type Server struct {
	limit int
	srs   *kzg.SRS
	users map[string]*Client
}

func NewServer(limit int, srs *kzg.SRS) *Server {
	users := make(map[string]*Client)
	return &Server{limit, srs, users}
}

func (s *Server) registerOnServer(this js.Value, args []js.Value) interface{} {
	commitmentStr := args[0].String()
	keyPairProofStr := args[1].String()
	fmt.Printf("Server received user registration, commitment: %s, keyPairProof: %s\n", commitmentStr, keyPairProofStr)
	var commitment kzg.Digest
	commitmentB, err := hex.DecodeString(commitmentStr)
	if err != nil {
		fmt.Println("[E]", commitmentStr, err)
		return nil
	}
	commitment.SetBytes(commitmentB)
	if s.users[string(commitmentB)] != nil {
		fmt.Println("[E] user already registered")
		return nil
	}
	// spew.Dump(commitment)
	var keyPairProof keypair.KeyPairProof
	err = json.Unmarshal([]byte(keyPairProofStr), &keyPairProof)
	if err != nil {
		fmt.Println("[E]", keyPairProofStr, err)
		return nil
	}
	// spew.Dump(keyPairProof)

	hash := sha3.NewLegacyKeccak256()
	pubKeyXA0 := new(big.Int)
	pubKeyXA1 := new(big.Int)
	pubKeyYA0 := new(big.Int)
	pubKeyYA1 := new(big.Int)
	keyPairProof.PublicKeyG2Aff.X.A0.BigInt(pubKeyXA0)
	keyPairProof.PublicKeyG2Aff.X.A1.BigInt(pubKeyXA1)
	keyPairProof.PublicKeyG2Aff.Y.A0.BigInt(pubKeyYA0)
	keyPairProof.PublicKeyG2Aff.Y.A1.BigInt(pubKeyYA1)
	hash.Write(pubKeyXA0.Bytes())
	hash.Write(pubKeyXA1.Bytes())
	hash.Write(pubKeyYA0.Bytes())
	hash.Write(pubKeyYA1.Bytes())
	var keyHash [32]byte
	copy(keyHash[:], hash.Sum(nil))
	log.Printf("PubKeyHash: 0x%02x", keyHash)

	start := time.Now()
	err = keypair.VerifyPubKey(&commitment, &keyPairProof, s.srs)
	if err != nil {
		fmt.Printf("VerifyPubKey failed: %s\n", err)
		return nil
	}
	// pairing check for publicKeyG1 and publicKeyG2
	eLeft, err := bn254.Pair([]bn254.G1Affine{keyPairProof.PublicKeyG1Aff}, []bn254.G2Affine{s.srs.G2[0]})
	if err != nil {
		panic(err)
	}
	eRight, err := bn254.Pair([]bn254.G1Affine{s.srs.G1[0]}, []bn254.G2Affine{keyPairProof.PublicKeyG2Aff})
	if err != nil {
		panic(err)
	}
	if !eLeft.Equal(&eRight) {
		fmt.Printf("pubkey G1 G2 pairing check failed: %s\n", err)
		return nil
	}
	elapsed := time.Since(start)
	log.Printf("VerifyPubKey time: %s", elapsed)
	response := fmt.Sprintf("Register success, commitment: %s, verify pubkey took: %s", commitmentStr, elapsed)
	log.Println(response)
	Print("leftTextArea", response)
	client := NewClient(s.limit, s.srs)
	client.commitment = commitment
	client.keyPairProof = &keyPairProof
	s.users[string(commitmentB)] = client
	return nil
}

func (s *Server) receiveMessage(this js.Value, args []js.Value) interface{} {
	msgStr := args[0].String()
	var msg types.Message
	err := json.Unmarshal([]byte(msgStr), &msg)
	if err != nil {
		log.Printf("unmarshal failed: %s\n", err)
		return nil
	}
	commitmentB := msg.Commitment.Bytes()
	user, ok := s.users[string(commitmentB[:])]
	if !ok {
		log.Printf("user not registered")
		return nil
	}
	// if user.nonce >= 95 {
	// 	userInput <- msg.Text
	// } else {
	log.Printf("Server received message: %s", msg.Text)
	// }
	// verify message
	hasher := crypto.SHA256.New()
	hasher.Write([]byte(fmt.Sprint(user.nonce)))
	hasher.Write([]byte(msg.Text))
	var frMsg fr.Element
	frMsg.SetBytes(hasher.Sum(nil))
	start := time.Now()
	err = kzg.Verify(&msg.Commitment, &msg.Proof, frMsg, s.srs)
	if err != nil {
		log.Printf("verify failed: %s\n", err)
		return nil
	}
	elapsed := time.Since(start)
	log.Printf("verify time: %s", elapsed)

	user.nonce++
	response := fmt.Sprintf("Recv: [%s] (verify took: %s) used(%d/%d)", msg.Text, elapsed, user.nonce, user.limit-1)
	Print("leftTextArea", response)

	user.points = append(user.points, &types.Point{X: frMsg, Y: msg.Proof.ClaimedValue})
	// update nonce
	fmt.Printf("user nonce: %d, limit: %d\n", user.nonce, user.limit)
	if user.nonce >= user.limit {
		privateKey := interpolation.RecoverPrivateKeyByPoints(user.points[:])
		log.Printf("private key recovered: %s", privateKey.String())
		Print("leftTextArea", fmt.Sprintf("private key recovered: %s", privateKey.String()))
		// auth := utils.NewTransactor(client, "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a")
		// message := auth.From.Hash().Bytes()
		// log.Printf("message: %02X", message)
		// pubKey, _, signature := sign(message, privateKey)
		// var pubKeyPoint BN254HashToG1G2Point
		// var sigPoint BN254HashToG1G1Point

		// pubKeyXA0 := new(big.Int)
		// pubKeyXA1 := new(big.Int)
		// pubKeyYA0 := new(big.Int)
		// pubKeyYA1 := new(big.Int)
		// pubKey.X.A0.BigInt(pubKeyXA0)
		// pubKey.X.A1.BigInt(pubKeyXA1)
		// pubKey.Y.A0.BigInt(pubKeyYA0)
		// pubKey.Y.A1.BigInt(pubKeyYA1)

		// signatureX := new(big.Int)
		// signatureY := new(big.Int)
		// signature.X.BigInt(signatureX)
		// signature.Y.BigInt(signatureY)

		// pubKeyPoint.X = [2]*big.Int{pubKeyXA0, pubKeyXA1}
		// pubKeyPoint.Y = [2]*big.Int{pubKeyYA0, pubKeyYA1}
		// sigPoint.X = signatureX
		// sigPoint.Y = signatureY
		// tx, err := stake.Withdraw90Percent(auth, pubKeyPoint, sigPoint)
		// if err != nil {
		// 	panic(err)
		// }
		// log.Printf("tx sent: %s", tx.Hash().String())
		log.Printf("Slash 90 percent of stake 💸💸💸")
		// log.Printf("Reward to the user: %s", auth.From.Hex())
	}
	return nil
}
