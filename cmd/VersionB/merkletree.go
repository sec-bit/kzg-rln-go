package main

import (
	"github.com/consensys/gnark/frontend"
	// "github.com/consensys/gnark/std/hash/mimc"
	"github.com/liyue201/gnark-circomlib/circuits"
)

// HashLeftRight computes MiMC([left, right])
func HashLeftRight(api frontend.API, left, right frontend.Variable) frontend.Variable {
	hasher := circuits.NewMimcSpongeHash(api)
	hasher.Write(left, right)
	hash := hasher.Sum()

	return hash
}

// DualMux returns [in[0], in[1]] if s == 0 and [in[1], in[0]] if s == 1
func DualMux(api frontend.API, in [2]frontend.Variable, s frontend.Variable) [2]frontend.Variable {
	var out [2]frontend.Variable

	api.AssertIsBoolean(s)
	out[0] = api.Add(api.Mul(api.Sub(in[1], in[0]), s), in[0])
	out[1] = api.Add(api.Mul(api.Sub(in[0], in[1]), s), in[1])

	return out
}

// MerkleTreeChecker verifies that merkle proof is correct for given merkle root and a leaf
// pathIndices input is an array of 0/1 selectors telling whether given pathElement is on the left or right side of merkle path
func MerkleTreeChecker(api frontend.API, levels int, leaf, root frontend.Variable, pathElements, pathIndices []frontend.Variable) {
	selectors := make([]struct {
		in, out [2]frontend.Variable
		s       frontend.Variable
	}, levels)

	hashers := make([]struct {
		left, right, hash frontend.Variable
	}, levels)

	for i := 0; i < levels; i++ {
		if i == 0 {
			selectors[i].in[0] = leaf
		} else {
			selectors[i].in[0] = hashers[i-1].hash
		}
		selectors[i].in[1] = pathElements[i]
		selectors[i].s = pathIndices[i]

		selectors[i].out[0] = api.Add(api.Mul(api.Sub(selectors[i].in[1], selectors[i].in[0]), selectors[i].s), selectors[i].in[0])
		selectors[i].out[1] = api.Add(api.Mul(api.Sub(selectors[i].in[0], selectors[i].in[1]), selectors[i].s), selectors[i].in[1])

		hashers[i].left = selectors[i].out[0]
		hashers[i].right = selectors[i].out[1]
		hashers[i].hash = HashLeftRight(api, hashers[i].left, hashers[i].right)
	}

	api.AssertIsEqual(root, hashers[levels-1].hash)
}

func generateMerkleProof(api frontend.API, leaves []frontend.Variable, index int, levels int) ([]frontend.Variable, []frontend.Variable) {
	if len(leaves) == 0 {
		return nil, nil
	}

	nodes := make([][]frontend.Variable, levels+1)
	nodes[0] = leaves

	for i := 1; i <= levels; i++ {
		levelSize := len(nodes[i-1]) / 2
		nodes[i] = make([]frontend.Variable, levelSize)

		for j := 0; j < levelSize; j++ {
			left := nodes[i-1][2*j]
			right := nodes[i-1][2*j+1]
			nodes[i][j] = HashLeftRight(api, left, right)
		}
	}

	proof := make([]frontend.Variable, levels)
	pathIndices := make([]frontend.Variable, levels)
	idx := index

	for i := 0; i < levels; i++ {
		siblingIndex := idx ^ 1

		proof[i] = nodes[i][siblingIndex]
		pathIndices[i], _ = api.ConstantValue(idx & 1)
		idx = idx / 2
	}

	return proof, pathIndices
}
