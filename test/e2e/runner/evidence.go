package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/cometbft/cometbft/crypto"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/types"
)

// 1 in 4 evidence is light client evidence, the rest is duplicate vote evidence
const lightClientEvidenceRatio = 4

// InjectEvidence takes a running testnet and generates an amount of valid/invalid
// evidence and broadcasts it to a random node through the rpc endpoint `/broadcast_evidence`.
// Evidence is random and can be a mixture of LightClientAttackEvidence and
// DuplicateVoteEvidence.
func InjectEvidence(ctx context.Context, r *rand.Rand, testnet *e2e.Testnet, amount int) error {
	_ = "STUB: not implemented"
	// select a random node
	return nil
}

// request the latest block and validator set from the node

// get the private keys of all the validators in the network

// wait for the node to reach the height above the forged height so that
// it is able to validate the evidence

// Alternate valid and invalid evidence

// The tests will count committed evidences later on,
// and only valid evidences will make it

// wait for the node to reach the height above the forged height so that
// it is able to validate the evidence

func getPrivateValidatorKeys(testnet *e2e.Testnet) ([]types.MockPV, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create mock private validators from the validators private key. MockPV is
// stateless which means we can double vote and do other funky stuff

// creates evidence of a lunatic attack. The height provided is the common height.
// The forged height happens 2 blocks later.
func generateLightClientAttackEvidence(
	ctx context.Context,
	privVals []types.MockPV,
	height int64,
	vals *types.ValidatorSet,
	chainID string,
	evTime time.Time,
	validEvidence bool,
) (*types.LightClientAttackEvidence, error) {
	_ = "STUB: not implemented"
	// forge a random header
	return nil, nil
}

// add a new bogus validator and remove an existing one to
// vary the validator set slightly

// create a commit for the forged header

// malleate the last signature of the commit by adding one to its first byte

// generateDuplicateVoteEvidence picks a random validator from the val set and
// returns duplicate vote evidence against the validator
func generateDuplicateVoteEvidence(
	privVals []types.MockPV,
	height int64,
	vals *types.ValidatorSet,
	chainID string,
	time time.Time,
) (*types.DuplicateVoteEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getRandomValidatorIndex picks a random validator from a slice of mock PrivVals that's
// also part of the validator set, returning the PrivVal and its index in the validator set
func getRandomValidatorIndex(privVals []types.MockPV, vals *types.ValidatorSet) (types.MockPV, int32, error) {
	_ = "STUB: not implemented"
	return *new(types.MockPV), 0, nil
}

func readPrivKey(keyFilePath string) (crypto.PrivKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivKey), nil
}

func makeHeaderRandom(chainID string, height int64) *types.Header {
	_ = "STUB: not implemented"
	return nil
}

func makeRandomBlockID() types.BlockID { _ = "STUB: not implemented"; return *new(types.BlockID) }

func makeBlockID(hash []byte, partSetSize uint32, partSetHash []byte) types.BlockID {
	_ = "STUB: not implemented"
	return *new(types.BlockID)
}

func mutateValidatorSet(
	ctx context.Context,
	privVals []types.MockPV,
	vals *types.ValidatorSet,
	nop bool,
) ([]types.PrivValidator, *types.ValidatorSet, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// we need to sort the priv validators with the same index as the validator set
