package main

import (
	"crypto"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sec-bit/kzg-rln-go/interpolation"
	"github.com/sec-bit/kzg-rln-go/keypair"
	"github.com/sec-bit/kzg-rln-go/types"
	"github.com/sec-bit/kzg-rln-go/utils"
	"golang.org/x/crypto/sha3"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
)

type Server struct {
	limit  int
	srs    *kzg.SRS
	users  map[string]*User
	socket chan interface{}
}

func NewServer(limit int, srs *kzg.SRS) *Server {
	socket := make(chan interface{}, 10)
	users := make(map[string]*User)
	return &Server{limit, srs, users, socket}
}

func (s *Server) Run(client *ethclient.Client, stake *Stake, wg *sync.WaitGroup) {
	log.Printf("Server is running")
	defer wg.Done()
	// check s.socket type User or Message
	for packet := range s.socket {
		switch packet := packet.(type) {
		case *User:
			user := packet
			log.Printf("Server received user registration")
			log.Printf("Wait 1 sec to confirm the tx on chain")
			hash := sha3.NewLegacyKeccak256()
			pubKeyXA0 := new(big.Int)
			pubKeyXA1 := new(big.Int)
			pubKeyYA0 := new(big.Int)
			pubKeyYA1 := new(big.Int)
			user.keyPairProof.PublicKeyG2Aff.X.A0.BigInt(pubKeyXA0)
			user.keyPairProof.PublicKeyG2Aff.X.A1.BigInt(pubKeyXA1)
			user.keyPairProof.PublicKeyG2Aff.Y.A0.BigInt(pubKeyYA0)
			user.keyPairProof.PublicKeyG2Aff.Y.A1.BigInt(pubKeyYA1)
			hash.Write(pubKeyXA0.Bytes())
			hash.Write(pubKeyXA1.Bytes())
			hash.Write(pubKeyYA0.Bytes())
			hash.Write(pubKeyYA1.Bytes())
			var keyHash [32]byte
			copy(keyHash[:], hash.Sum(nil))
			log.Printf("PubKeyHash: %02X", keyHash)
			exists, err := stake.PublicKeys(nil, keyHash)
			if err != nil {
				panic(err)
			}
			if !exists.Exists {
				panic("key not exists")
			}
			err = keypair.VerifyPubKey(&user.commitment, user.keyPairProof, s.srs)
			if err != nil {
				panic(err)
			}
			// pairing check for publicKeyG1 and publicKeyG2
			eLeft, err := bn254.Pair([]bn254.G1Affine{user.keyPairProof.PublicKeyG1Aff}, []bn254.G2Affine{s.srs.G2[0]})
			if err != nil {
				panic(err)
			}
			eRight, err := bn254.Pair([]bn254.G1Affine{s.srs.G1[0]}, []bn254.G2Affine{user.keyPairProof.PublicKeyG2Aff})
			if err != nil {
				panic(err)
			}
			if !eLeft.Equal(&eRight) {
				panic("pairing check failed")
			}
			// ok, err := bn254.PairingCheck(
			// 	[]bn254.G1Affine{user.keyPairProof.PublicKeyG1Aff, s.srs.G1[0]},
			// 	[]bn254.G2Affine{s.srs.G2[0], user.keyPairProof.PublicKeyG2Aff})
			// if err != nil {
			// 	panic(err)
			// }
			// if !ok {
			// 	panic("pairing check failed")
			// }
			user.points = make([]*types.Point, 0)
			s.users[user.commitment.String()] = user
		case *types.Message:
			log.Printf("Server received message: %s", packet.Text)
			msg := packet
			user := s.users[msg.Commitment.String()]
			// verify message
			hasher := crypto.SHA256.New()
			hasher.Write([]byte(fmt.Sprint(user.nonce)))
			hasher.Write([]byte(msg.Text))
			var frMsg fr.Element
			frMsg.SetBytes(hasher.Sum(nil))
			err := kzg.Verify(&msg.Commitment, &msg.Proof, frMsg, s.srs)
			if err != nil {
				panic(err)
			}
			user.points = append(user.points, &types.Point{X: frMsg, Y: msg.Proof.ClaimedValue})
			// update nonce
			user.nonce++
			if user.nonce >= s.limit {
				privateKey := interpolation.RecoverPrivateKeyByPoints(user.points[:])
				log.Printf("private key recovered: %s", privateKey.String())
				auth := utils.NewTransactor(client, "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a")
				message := auth.From.Hash().Bytes()
				log.Printf("message: %02X", message)
				pubKey, _, signature := sign(message, privateKey)
				var pubKeyPoint BN254HashToG1G2Point
				var sigPoint BN254HashToG1G1Point

				pubKeyXA0 := new(big.Int)
				pubKeyXA1 := new(big.Int)
				pubKeyYA0 := new(big.Int)
				pubKeyYA1 := new(big.Int)
				pubKey.X.A0.BigInt(pubKeyXA0)
				pubKey.X.A1.BigInt(pubKeyXA1)
				pubKey.Y.A0.BigInt(pubKeyYA0)
				pubKey.Y.A1.BigInt(pubKeyYA1)

				signatureX := new(big.Int)
				signatureY := new(big.Int)
				signature.X.BigInt(signatureX)
				signature.Y.BigInt(signatureY)

				pubKeyPoint.X = [2]*big.Int{pubKeyXA0, pubKeyXA1}
				pubKeyPoint.Y = [2]*big.Int{pubKeyYA0, pubKeyYA1}
				sigPoint.X = signatureX
				sigPoint.Y = signatureY
				tx, err := stake.Withdraw90Percent(auth, pubKeyPoint, sigPoint)
				if err != nil {
					panic(err)
				}
				log.Printf("tx sent: %s", tx.Hash().String())
				log.Printf("Slash 90 percent of stake")
				log.Printf("Reward to the user: %s", auth.From.Hex())
			}
		}
	}
}
