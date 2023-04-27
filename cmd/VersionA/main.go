package main

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr/kzg"
)

var (
	// srs re-used accross tests of the KZG scheme
	GLOBAL_SRS *kzg.SRS
)

func init() {
	const srsSize = 230
	GLOBAL_SRS, _ = kzg.NewSRS(ecc.NextPowerOfTwo(srsSize), new(big.Int).SetInt64(42))
}

const (
	MESSAGE_LIMIT = 100
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	server := NewServer(MESSAGE_LIMIT, GLOBAL_SRS)
	go server.Run(wg)

	user := NewUser(server.socket, MESSAGE_LIMIT, GLOBAL_SRS)
	user.Register()
	for i := 0; i < MESSAGE_LIMIT; i++ {
		user.SendMessage(fmt.Sprintf("Spam %d!", i))
	}
	close(server.socket)
	wg.Wait()
}
