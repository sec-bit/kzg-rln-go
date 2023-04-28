package main

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

const (
	DEGREE = 100
)

// Define the circuit
type registerCircuit struct {
	// PublicInput  frontend.Variable
	// N            frontend.Variable
	// FAlpha       frontend.Variable
	PrivateKey     frontend.Variable `gnark:",secret"`
	CoeffHash      frontend.Variable `gnark:",public"`
	PolyEvalAtHash frontend.Variable `gnark:",public"`
	// PolyF        frontend.Variable   `gnark:",secret"`
	Coefficients [DEGREE]frontend.Variable `gnark:",secret"`
}

// Define and implement the constraints
func (circuit *registerCircuit) Define(api frontend.API) error {
	// Constraint 1: f(0) = private key
	api.Println(circuit.PrivateKey)
	api.Println(circuit.Coefficients[0])
	api.AssertIsEqual(circuit.PrivateKey, circuit.Coefficients[0])

	// Constraint 2: membership proof of g^f(0)
	// api.AssertIsEqual(circuit.PublicInput, circuit.FAlpha)

	// Constraint 3: check the hash of all coefficients, mimc([]coefficient) == coeffHash
	mimc, _ := mimc.NewMiMC(api)
	for i := 0; i < DEGREE; i++ {
		mimc.Write(circuit.Coefficients[i])
	}
	hash := mimc.Sum()
	api.Println(circuit.CoeffHash)
	api.Println(hash)
	api.AssertIsEqual(circuit.CoeffHash, hash)

	// Constraint 4: evaluate the polynomial at hash
	api.Println(circuit.PolyEvalAtHash)
	polyEvalAtH := circuit.Coefficients[DEGREE-1]
	for i := DEGREE - 2; i >= 0; i-- {
		polyEvalAtH = api.Add(api.Mul(polyEvalAtH, hash), circuit.Coefficients[i])
	}
	api.Println(polyEvalAtH)
	api.AssertIsEqual(circuit.PolyEvalAtHash, polyEvalAtH)

	return nil
}
