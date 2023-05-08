package main

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"syscall/js"
	"time"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/sec-bit/kzg-rln-go/keypair"
	"github.com/sec-bit/kzg-rln-go/types"
)

var (
	// load abi from string
	STAKE_ABI_STRING = `[ { "anonymous": false, "inputs": [ { "indexed": true, "internalType": "bytes32", "name": "publicKeyHash", "type": "bytes32" } ], "name": "Deposit", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "internalType": "bytes32", "name": "publicKeyHash", "type": "bytes32" } ], "name": "Withdraw", "type": "event" }, { "inputs": [], "name": "DEPOSIT_AMOUNT", "outputs": [ { "internalType": "uint256", "name": "", "type": "uint256" } ], "stateMutability": "view", "type": "function" }, { "inputs": [], "name": "WAIT_TIME", "outputs": [ { "internalType": "uint256", "name": "", "type": "uint256" } ], "stateMutability": "view", "type": "function" }, { "inputs": [ { "components": [ { "internalType": "uint256[2]", "name": "X", "type": "uint256[2]" }, { "internalType": "uint256[2]", "name": "Y", "type": "uint256[2]" } ], "internalType": "struct BN254HashToG1.G2Point", "name": "publicKey", "type": "tuple" } ], "name": "deposit", "outputs": [], "stateMutability": "payable", "type": "function" }, { "inputs": [ { "internalType": "bytes", "name": "data", "type": "bytes" } ], "name": "hashToPoint", "outputs": [ { "internalType": "uint256[2]", "name": "result", "type": "uint256[2]" } ], "stateMutability": "view", "type": "function" }, { "inputs": [ { "internalType": "bytes32", "name": "", "type": "bytes32" } ], "name": "publicKeys", "outputs": [ { "internalType": "bool", "name": "exists", "type": "bool" }, { "internalType": "uint256", "name": "timestamp", "type": "uint256" } ], "stateMutability": "view", "type": "function" }, { "inputs": [ { "components": [ { "internalType": "uint256", "name": "X", "type": "uint256" }, { "internalType": "uint256", "name": "Y", "type": "uint256" } ], "internalType": "struct BN254HashToG1.G1Point", "name": "point", "type": "tuple" }, { "internalType": "uint256", "name": "scalar", "type": "uint256" } ], "name": "scalarMul", "outputs": [ { "components": [ { "internalType": "uint256", "name": "X", "type": "uint256" }, { "internalType": "uint256", "name": "Y", "type": "uint256" } ], "internalType": "struct BN254HashToG1.G1Point", "name": "", "type": "tuple" } ], "stateMutability": "view", "type": "function" }, { "inputs": [ { "components": [ { "internalType": "uint256", "name": "X", "type": "uint256" }, { "internalType": "uint256", "name": "Y", "type": "uint256" } ], "internalType": "struct BN254HashToG1.G1Point", "name": "message", "type": "tuple" }, { "components": [ { "internalType": "uint256[2]", "name": "X", "type": "uint256[2]" }, { "internalType": "uint256[2]", "name": "Y", "type": "uint256[2]" } ], "internalType": "struct BN254HashToG1.G2Point", "name": "pubKey", "type": "tuple" }, { "components": [ { "internalType": "uint256", "name": "X", "type": "uint256" }, { "internalType": "uint256", "name": "Y", "type": "uint256" } ], "internalType": "struct BN254HashToG1.G1Point", "name": "signature", "type": "tuple" } ], "name": "verify", "outputs": [ { "internalType": "bool", "name": "", "type": "bool" } ], "stateMutability": "view", "type": "function" }, { "inputs": [ { "internalType": "bytes", "name": "message", "type": "bytes" }, { "components": [ { "internalType": "uint256[2]", "name": "X", "type": "uint256[2]" }, { "internalType": "uint256[2]", "name": "Y", "type": "uint256[2]" } ], "internalType": "struct BN254HashToG1.G2Point", "name": "pubKey", "type": "tuple" }, { "components": [ { "internalType": "uint256", "name": "X", "type": "uint256" }, { "internalType": "uint256", "name": "Y", "type": "uint256" } ], "internalType": "struct BN254HashToG1.G1Point", "name": "signature", "type": "tuple" } ], "name": "verifyMessage", "outputs": [ { "internalType": "bool", "name": "", "type": "bool" } ], "stateMutability": "view", "type": "function" }, { "inputs": [ { "components": [ { "internalType": "uint256[2]", "name": "X", "type": "uint256[2]" }, { "internalType": "uint256[2]", "name": "Y", "type": "uint256[2]" } ], "internalType": "struct BN254HashToG1.G2Point", "name": "publicKey", "type": "tuple" }, { "components": [ { "internalType": "uint256", "name": "X", "type": "uint256" }, { "internalType": "uint256", "name": "Y", "type": "uint256" } ], "internalType": "struct BN254HashToG1.G1Point", "name": "signature", "type": "tuple" } ], "name": "withdraw90Percent", "outputs": [], "stateMutability": "nonpayable", "type": "function" }, { "inputs": [ { "components": [ { "internalType": "uint256[2]", "name": "X", "type": "uint256[2]" }, { "internalType": "uint256[2]", "name": "Y", "type": "uint256[2]" } ], "internalType": "struct BN254HashToG1.G2Point", "name": "publicKey", "type": "tuple" }, { "components": [ { "internalType": "uint256", "name": "X", "type": "uint256" }, { "internalType": "uint256", "name": "Y", "type": "uint256" } ], "internalType": "struct BN254HashToG1.G1Point", "name": "signature", "type": "tuple" } ], "name": "withdrawWaitFor1day", "outputs": [], "stateMutability": "nonpayable", "type": "function" } ]`
	STAKE_ABI, _     = abi.JSON(strings.NewReader(STAKE_ABI_STRING))
)

type Client struct {
	limit        int
	srs          *kzg.SRS
	polynomial   []fr.Element
	keyPairProof *keypair.KeyPairProof
	commitment   kzg.Digest
	nonce        int
	points       []*types.Point
	server       *Server
}

func NewClient(limit int, srs *kzg.SRS) *Client {
	return &Client{limit, srs, nil, nil, kzg.Digest{}, 0, nil, nil}
}

func (c *Client) genNewPoly(this js.Value, args []js.Value) interface{} {
	start := time.Now()
	poly, commitment, proof := GenerateRandomPolynomialAndPrivateKey(c.srs, c.limit)
	elapsed := time.Since(start)
	c.polynomial = poly
	c.commitment = commitment
	c.keyPairProof = proof

	proof.PrivateKey = nil // discard private key
	commitmentHex := fmt.Sprintf("%02x", commitment.Bytes())

	polyB, err := json.Marshal(poly)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	proofB, err := json.Marshal(proof)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	js.Global().Get("document").Call("getElementById", "commitment").Set("value", commitmentHex)
	js.Global().Get("document").Call("getElementById", "keyProof").Set("value", string(proofB))
	js.Global().Get("document").Call("getElementById", "coefficients").Set("value", string(polyB))

	response := fmt.Sprintf("Generate Polynomial And Proof took %s", elapsed)
	js.Global().Get("document").Call("getElementById", "rightTextArea").Set("value", response)
	return nil
}

func (c *Client) registerPubKeyOnChain(this js.Value, args []js.Value) interface{} {
	response := fmt.Sprintf("Server received: %s", "register")
	js.Global().Get("document").Call("getElementById", "leftTextArea").Set("value", response)
	pubKeyProofStr := args[0].String()
	fmt.Printf("pubKeyProofStr: %s\n", pubKeyProofStr)
	var keyPairProof keypair.KeyPairProof
	err := json.Unmarshal([]byte(pubKeyProofStr), &keyPairProof)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	spew.Dump(keyPairProof)

	pubKeyXA0 := new(big.Int)
	pubKeyXA1 := new(big.Int)
	pubKeyYA0 := new(big.Int)
	pubKeyYA1 := new(big.Int)
	keyPairProof.PublicKeyG2Aff.X.A0.BigInt(pubKeyXA0)
	keyPairProof.PublicKeyG2Aff.X.A1.BigInt(pubKeyXA1)
	keyPairProof.PublicKeyG2Aff.Y.A0.BigInt(pubKeyYA0)
	keyPairProof.PublicKeyG2Aff.Y.A1.BigInt(pubKeyYA1)
	var pubKeyG2 BN254HashToG1G2Point
	pubKeyG2.X = [2]*big.Int{pubKeyXA0, pubKeyXA1}
	pubKeyG2.Y = [2]*big.Int{pubKeyYA0, pubKeyYA1}
	target := "0x5FbDB2315678afecb367f032d93F642f64180aa3"
	calldata := GenerateDepositArgs(pubKeyG2)
	resultMap := struct {
		Target   string `json:"target"`
		CallData string `json:"calldata"`
	}{
		Target:   target,
		CallData: "0x" + hex.EncodeToString(calldata),
	}

	jsonString, err := json.Marshal(resultMap)
	if err != nil {
		fmt.Println("Error marshaling resultMap:", err)
		return nil
	}

	return string(jsonString)
}

func (c *Client) sendMessage(this js.Value, args []js.Value) interface{} {
	message := args[0].String()

	var frMsg fr.Element
	hasher := crypto.SHA256.New()
	hasher.Write([]byte(fmt.Sprint(c.nonce)))
	hasher.Write([]byte(message))
	frMsg.SetBytes(hasher.Sum(nil))
	start := time.Now()
	msgProof, err := kzg.Open(c.polynomial, frMsg, c.srs)
	if err != nil {
		return err
	}
	elapsed := time.Since(start)
	log.Printf("Message proof generation took %s", elapsed)

	response := fmt.Sprintf("Msg: [%s] (proof took: %s)", message, elapsed)
	Print("rightTextArea", response)
	var msg types.Message
	msg.Commitment = c.commitment
	msg.Text = message
	msg.Proof = msgProof
	msgJson, err := json.Marshal(&msg)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// spew.Dump(msgJson)
	// c.server.ReceiveMessage(msg)
	c.nonce++
	return string(msgJson)
}

func GenerateRandomPolynomialAndPrivateKey(srs *kzg.SRS, limit int) (poly []fr.Element, commitment kzg.Digest, keyPairProof *keypair.KeyPairProof) {
	if limit > len(srs.G1) {
		log.Fatalf("Limit %d is greater than srs size %d", limit, len(srs.G1))
	}
	poly = keypair.RandomPolynomial(limit)
	commitment, keyPairProof = keypair.GenerateCommitmentAndKeyPairProof(poly, srs)
	log.Printf("commitment: %s", commitment.String())
	log.Printf("User created\n Private key: %s\n\nPublicKey: %s", keyPairProof.PrivateKey.String(), keyPairProof.PublicKeyG1Aff.String())
	return
}

func GenerateDepositArgs(pubkeyG2 BN254HashToG1G2Point) []byte {
	ret, _ := STAKE_ABI.Pack("deposit", pubkeyG2)
	return ret
}
