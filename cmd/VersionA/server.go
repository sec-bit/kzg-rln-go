package main

import (
	"crypto"
	"fmt"
	"log"
	"sync"

	"github.com/sec-bit/kzg-rln-go/interpolation"
	"github.com/sec-bit/kzg-rln-go/keypair"
	"github.com/sec-bit/kzg-rln-go/types"

	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr/kzg"
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

func (s *Server) Run(wg *sync.WaitGroup) {
	log.Printf("Server is running")
	defer wg.Done()
	// check s.socket type User or Message
	for packet := range s.socket {
		switch packet := packet.(type) {
		case *User:
			log.Printf("Server received user registration")
			user := packet
			err := keypair.VerifyPubKey(&user.commitment, user.keyPairProof, s.srs)
			if err != nil {
				panic(err)
			}
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
			}
		}
	}
}
