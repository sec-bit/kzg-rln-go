package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"sync"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	// read contract address from arguments
	contractAddress := common.HexToAddress(os.Args[1])

	stake, err := NewStake(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	server := NewServer(MESSAGE_LIMIT, GLOBAL_SRS)
	go server.Run(client, stake, wg)

	user := NewUser(server.socket, MESSAGE_LIMIT, GLOBAL_SRS)
	user.Register(client, stake)
	for i := 0; i < MESSAGE_LIMIT; i++ {
		// read from stdin

		user.SendMessage(fmt.Sprintf("message %d", i))
	}
	close(server.socket)
	wg.Wait()
}
