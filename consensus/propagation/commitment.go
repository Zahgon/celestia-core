package propagation

import (
	proptypes "github.com/cometbft/cometbft/consensus/propagation/types"
	"github.com/cometbft/cometbft/crypto/merkle"
	"github.com/cometbft/cometbft/libs/bits"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/proto/tendermint/propagation"
	"github.com/cometbft/cometbft/types"
)

var _ Propagator = (*Reactor)(nil)

const (
	CompactBlockUID = "compactBlock"
)

// ProposeBlock is called when the consensus routine has created a new proposal,
// and it needs to be gossiped to the rest of the network.
func (blockProp *Reactor) ProposeBlock(proposal *types.Proposal, block *types.PartSet, txs []proptypes.TxMetaData) error {
	_ = "STUB: not implemented"
	// create the parity data and the compact block
	return nil
}

// sign the hash of the compact block NOTE: p2p message sign bytes are
// prepended with the chain id and UID

// save the compact block locally and broadcast it to the connected peers

// distribute equal portions of haves to each of the proposer's peers

// since this node is proposing, it already has the data and
// there's not a lot of reason to update this peer's have
// state other than to be consistent atm.

// skip saving the old routine's state if the state here
// cannot also be saved

// this might not get set depending on when the consensus peer
// state is getting updated. This will result in sending the
// peer redundant parts :shrug:

func extractHashes(blocks ...*types.PartSet) [][]byte { _ = "STUB: not implemented"; return nil }

// Preallocate capacity

func extractProofs(blocks ...*types.PartSet) []*merkle.Proof { _ = "STUB: not implemented"; return nil }

// Preallocate capacity

func chunkToPartMetaData(chunk *bits.BitArray, partSet *types.PartSet) []*propagation.PartMetaData {
	_ = "STUB: not implemented"
	return nil
}

// handleCompactBlock adds a proposal to the data routine. This should be called any
// time a proposal is received from a peer or when a proposal is created. If the
// proposal is new, it will be stored and broadcast to the relevant peers.
func (blockProp *Reactor) handleCompactBlock(cb *proptypes.CompactBlock, peer p2p.ID, proposer bool) {
	_ = "STUB: not implemented"
	// Proposers skip validation since they created the block
	return
}

// Try to validate the compact block

// Validation failed - cache for later if it's for current or future height.
// We cache because:
// 1. Future height: we don't know the proposer key yet
// 2. Current height, wrong round: we may advance to that round later
// 3. Wrong proposer: proposer key may not be set yet (catchup scenario)

// processValidatedCompactBlock handles a compact block that has passed validation.
func (blockProp *Reactor) processValidatedCompactBlock(cb *proptypes.CompactBlock, peer p2p.ID, proposer bool) {
	_ = "STUB: not implemented"
	// generate (and cache) the proofs from the partset hashes in the compact block
	return
}

// check if we have any transactions that are in the compact block

// recoverPartsFromMempool queries the mempool to see if we can recover any block parts locally.
func (blockProp *Reactor) recoverPartsFromMempool(cb *proptypes.CompactBlock) {
	_ = "STUB: not implemented"
	return

	// find the compact block transactions that exist in our mempool
}

// todo: investigate why this could get hit, it shouldn't ever get hit

// broadcastProposal gossips the provided proposal to all peers. This should
// only be called upon receiving a proposal for the first time or after creating
// a proposal block.
func (blockProp *Reactor) broadcastCompactBlock(cb *proptypes.CompactBlock, from p2p.ID) {
	_ = "STUB: not implemented"
	return
}

// todo: we need to avoid sending this peer anything else until we can queue this message.

// chunkParts takes a bit array then returns an array of chunked bit arrays.
func chunkParts(p *bits.BitArray, peerCount, redundancy int) []*bits.BitArray {
	_ = "STUB: not implemented"
	return nil
}

// round up to use the ceil

// Create empty bit arrays for each peer

// chunkIndexes creates a nested slice of starting and ending indexes for each
// chunk. totalSize indicates the number of chunks. chunkSize indicates the size
// of each chunk.
func chunkIndexes(totalSize, chunkSize int) [][2]int { _ = "STUB: not implemented"; return nil }

// return nil // Handle invalid input gracefully

// Ensure the last chunk doesn't exceed the total size

// validateCompactBlock stateful validation of the compact block.
func (blockProp *Reactor) validateCompactBlock(cb *proptypes.CompactBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify POLRound, which must be -1 or in range [0, proposal.Round).

// Verify proposal signature

// Validate the proposed block size, derived from its PartSetHeader

// Calculate the minimum number of block parts needed to store a block of maxBytes size.
// Each part is BlockPartSizeBytes (64KB) in size. The formula (maxBytes-1)/partSize + 1
// ensures we round up to the next integer when maxBytes is not evenly divisible by partSize.

// validate the compact block
