package main

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
	"github.com/consensys/gnark/std/hash/mimc"
)

// deposit: hash(pk, kzgCommitment) -> update root
// withdraw: proof of hash(pk) == hashPK, membership(hash(pk, kzgCommitment))
// server tracking hash(pk) on chain
// server know all kzgCommitment, should we be really concerned?
// register with hashPK, this is also used as nullifier
// slash with proof of pk avoid interpolating

// Define the circuit
type withdrawCircuit struct {
	PrivateKey     frontend.Variable  `gnark:",secret"`
	PrivateKeyHash frontend.Variable  `gnark:",public"` //nullifier
	KzgCommitment  frontend.Variable  `gnark:",public"`
	MerkleProof    merkle.MerkleProof `gnark:",secret"`
}

func (circuit *withdrawCircuit) Define(api frontend.API) {
	mimc, _ := mimc.NewMiMC(api)
	mimc.Write(circuit.PrivateKey)
	hash := mimc.Sum()
	api.AssertIsEqual(circuit.PrivateKeyHash, hash)

	mimc.Write(circuit.PrivateKey)
	mimc.Write(circuit.KzgCommitment)
	// commitmentHash is leaf of merkle tree
	commitmentHash := mimc.Sum()

	circuit.MerkleProof.VerifyProof(api, &mimc, commitmentHash)
}
