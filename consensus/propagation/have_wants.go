package propagation

import (
	"github.com/cometbft/cometbft/crypto/merkle"

	proptypes "github.com/cometbft/cometbft/consensus/propagation/types"
	"github.com/cometbft/cometbft/p2p"
)

// handleHaves is called when a peer sends a have message. This is used to
// determine if the sender has or is getting portions of the proposal that this
// node doesn't have. If the sender has parts that this node doesn't have, this
// node will request those parts. The peer must always send the proposal before
// sending parts. If they did not, this node must disconnect from them.
func (blockProp *Reactor) handleHaves(peer p2p.ID, haves *proptypes.HaveParts) {
	_ = "STUB: not implemented"

	// TODO handle the disconnection case
	return
}

// blockProp.Switch.StopPeerForError(blockProp.getPeer(peer).peer, errors.New("received part for unknown proposal"))

// we can't process haves for a compact block we don't have

// this is a catchup block, we shouldn't receive haves for

//blockProp.Switch.StopPeerForError(p.peer, err, blockProp.String())

// Check if the sender has parts that we don't have.

// avoid blocking if a single peer is backed up. This means that they
// are sending us too many haves

// ReqLimit limits the number of requests per part.
// It allows requesting the small blocks multiple times to avoid relying only on a few/single
// peer to upload the whole block.
// The provided partsCount is the number of parts in the block and parity data.
func ReqLimit(partsCount int) int { _ = "STUB: not implemented"; return 0 }

func (blockProp *Reactor) requestFromPeer(ps *PeerState) { _ = "STUB: not implemented"; return }

// should never be below zero

// haves for a new height, resetting

// we can ignore this have in this case

// don't request a part that is already downloaded

// don't request a part that has already hit the request limit

// don't request the part from this peer if we've already requested it
// from them.

// p == peer means we have already requested the part from this peer.

// if none of the requests were relevant, then wants will still be
// nil

// no need to send the want in this case

// countRemainingParts counts the remaining parts to decode a block.
func countRemainingParts(totalParts, existingParts int) int32 { _ = "STUB: not implemented"; return 0 }

func (blockProp *Reactor) sendWantsThenBroadcastHaves(ps *PeerState, wants *proptypes.WantParts) error {
	_ = "STUB: not implemented"
	return nil
}

func (blockProp *Reactor) convertWantToHave(want *proptypes.WantParts) (*proptypes.HaveParts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (blockProp *Reactor) sendWant(ps *PeerState, want *proptypes.WantParts) {
	_ = "STUB: not implemented"
	return
}

// keep track of the parts that this node has requested.

// countRequests returns the number of requests for a given part.
func (blockProp *Reactor) countRequests(height int64, round int32, part int) []p2p.ID {
	_ = "STUB: not implemented"
	return nil
}

// broadcastHaves gossips the provided have msg to all peers except to the
// original sender. This should only be called upon receiving a new have for the
// first time.
func (blockProp *Reactor) broadcastHaves(haves *proptypes.HaveParts, from p2p.ID, partSetSize int) {
	_ = "STUB: not implemented"
	return
}

// todo(evan): don't rely strictly on try, however since we're using
// pull based gossip, this isn't as big as a deal since if someone asks
// for data, they must already have the proposal.
// TODO: use retry and logs

// handleWants is called when a peer sends a want message. This is used to send
// peers data that this node already has and store the wants to send them data
// in the future.
func (blockProp *Reactor) handleWants(peer p2p.ID, wants *proptypes.WantParts) {
	_ = "STUB: not implemented"
	return
}

// get data, use the prove as a proxy for determining if this Want message
// if for catchup

// the peer must always send the proposal before sending parts, if they did
//  not, this node must disconnect from them.

// blockProp.Switch.StopPeerForError(p.peer, errors.New("received want part for unknown proposal"))

// if we have the parts, send them to the peer.

// p.SetHave(height, round, int(partIndex))

// for parts that we don't have, but they still want, store the wants.

// handleRecoveryPart is called when a peer sends a block part message. This is used
// to store the part and clear any wants for that part.
func (blockProp *Reactor) handleRecoveryPart(peer p2p.ID, part *proptypes.RecoveryPart) {
	_ = "STUB: not implemented"
	return
}

// the peer must always send the proposal before sending parts, if they did
// not this node must disconnect from them.

// blockProp.Switch.StopPeerForError(p.peer, errors.New("received recovery part for unknown proposal"))

// todo: add these defensive checks in a better way

// todo: we need to figure out a way to get the proof for a part that was
// sent during catchup.

// avoid blocking if a single peer is backed up. This means that they
// are sending us too many parts

// if the part was not added and there was no error, the part has already
// been seen, and therefore doesn't need to be cleared.

// only send original parts to the consensus reactor

// attempt to decode the remaining block parts. If they are decoded, then
// this node should send all the wanted parts that nodes have requested. cp
// == nil means that there was no compact block available and this was
// during catchup. todo: use the bool found in the state instead of checking
// for nil.

// broadcast haves for all parts since we've decoded the entire block.
// rely on the broadcast method to ensure that parts are only sent once.

// only send original parts to the consensus reactor

// clear all the wants if they exist

// clearWants checks the wantState to see if any peers want the given part, if
// so, it attempts to send them that part.
func (blockProp *Reactor) clearWants(part *proptypes.RecoveryPart, proof merkle.Proof) {
	_ = "STUB: not implemented"
	return
}
