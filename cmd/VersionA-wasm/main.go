package main

import (
	"fmt"
	"math/big"
	"syscall/js"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
)

var (
	// srs re-used accross tests of the KZG scheme
	GLOBAL_SRS *kzg.SRS
)

const (
	MESSAGE_LIMIT = 10
)

func init() {
	const srsSize = MESSAGE_LIMIT
	GLOBAL_SRS, _ = kzg.NewSRS(ecc.NextPowerOfTwo(srsSize), new(big.Int).SetInt64(42))
}

func registerCallbacks() {
	fmt.Printf("srs length: %d\n", len(GLOBAL_SRS.G1))
	server := NewServer(len(GLOBAL_SRS.G1), GLOBAL_SRS)
	client := NewClient(len(GLOBAL_SRS.G1), GLOBAL_SRS)
	client.server = server
	js.Global().Set("sendMessage", js.FuncOf(client.sendMessage))
	js.Global().Set("genNewPoly", js.FuncOf(client.genNewPoly))
	js.Global().Set("registerPubKeyOnChain", js.FuncOf(client.registerPubKeyOnChain))
	js.Global().Set("registerOnServer", js.FuncOf(server.registerOnServer))
	js.Global().Set("receiveMessage", js.FuncOf(server.receiveMessage))

}

func Print(elementId, msg string) {
	currentText := js.Global().Get("document").Call("getElementById", elementId).Get("value").String()
	if currentText == "" {
		currentText = msg
	} else {
		currentText += "\n" + msg
	}
	js.Global().Get("document").Call("getElementById", elementId).Set("value", currentText)
}

func main() {
	fmt.Printf("WASM Go Initialized\n")
	c := make(chan struct{})
	registerCallbacks()
	<-c
}
