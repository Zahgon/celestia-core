package consensus

import (
	"context"
	"hash/crc32"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/proxy"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/types"
)

var crc32c = crc32.MakeTable(crc32.Castagnoli)

// Functionality to replay blocks and messages on recovery from a crash.
// There are two general failure scenarios:
//
//  1. failure during consensus
//  2. failure while applying the block
//
// The former is handled by the WAL, the latter by the proxyApp Handshake on
// restart, which ultimately hands off the work to the WAL.

//-----------------------------------------
// 1. Recover from failure during consensus
// (by replaying messages from the WAL)
//-----------------------------------------

// Unmarshal and apply a single message to the consensus state as if it were
// received in receiveRoutine.  Lines that start with "#" are ignored.
// NOTE: receiveRoutine should not be running.
func (cs *State) readReplayMessage(msg *TimedWALMessage, newStepSub types.Subscription) error {
	_ = "STUB: not implemented"
	// Skip meta messages which exist for demarcating boundaries.
	return nil
}

// for logging

// these are playback checks

// Replay only those messages since the last block.  `timeoutRoutine` should
// run concurrently to read off tickChan.
func (cs *State) catchupReplay(csHeight int64) error {
	_ = "STUB: not implemented"

	// Set replayMode to true so we don't log signing errors.
	return nil
}

// Ensure that #ENDHEIGHT for this height doesn't exist.
// NOTE: This is just a sanity check. As far as we know things work fine
// without it, and Handshake could reuse State if it weren't for
// this check (since we can crash after writing #ENDHEIGHT).
//
// Ignore data corruption errors since this is a sanity check.

// Search for last height marker.
//
// Ignore data corruption errors in previous heights because we only care about last height

// NOTE: since the priv key is set when the msgs are received
// it will attempt to eg double sign but we can just ignore it
// since the votes will be replayed and we'll get to the next step

//--------------------------------------------------------------------------------

// Parses marker lines of the form:
// #ENDHEIGHT: 12345
/*
func makeHeightSearchFunc(height int64) auto.SearchFunc {
	return func(line string) (int, error) {
		line = strings.TrimRight(line, "\n")
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return -1, errors.New("line did not have 2 parts")
		}
		i, err := strconv.Atoi(parts[1])
		if err != nil {
			return -1, errors.New("failed to parse INFO: " + err.Error())
		}
		if height < i {
			return 1, nil
		} else if height == i {
			return 0, nil
		} else {
			return -1, nil
		}
	}
}*/

//---------------------------------------------------
// 2. Recover from failure while applying the block.
// (by handshaking with the app to figure out where
// we were last, and using the WAL to recover there.)
//---------------------------------------------------

type Handshaker struct {
	stateStore   sm.Store
	initialState sm.State
	store        sm.BlockStore
	eventBus     types.BlockEventPublisher
	genDoc       *types.GenesisDoc
	logger       log.Logger

	nBlocks int // number of blocks applied to the state
}

func NewHandshaker(stateStore sm.Store, state sm.State,
	store sm.BlockStore, genDoc *types.GenesisDoc) *Handshaker {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handshaker) SetLogger(l log.Logger) {
	_ = "STUB: not implemented"

	// SetEventBus - sets the event bus for publishing block related events.
	// If not called, it defaults to types.NopEventBus.
	return
}

func (h *Handshaker) SetEventBus(eventBus types.BlockEventPublisher) {
	_ = "STUB: not implemented"
	return

	// NBlocks returns the number of blocks applied to the state.
}

func (h *Handshaker) NBlocks() int {
	_ = "STUB: not implemented"

	// TODO: retry the handshake/replay if it fails ?
	return 0
}

func (h *Handshaker) Handshake(proxyApp proxy.AppConns) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// HandshakeWithContext is cancellable version of Handshake
func (h *Handshaker) HandshakeWithContext(ctx context.Context, proxyApp proxy.AppConns) (string, error) {
	_ = "STUB: not implemented"

	// Handshake is done via ABCI Info on the query conn.
	return "", nil
}

// Only set the version if there is no existing state.

// set app version if it's not set via genesis

// Replay blocks up to the latest in the blockstore.

// TODO: (on restart) replay mempool

// ReplayBlocks replays all blocks since appBlockHeight and ensures the result
// matches the current state.
// Returns the final AppHash or an error.
func (h *Handshaker) ReplayBlocks(
	state sm.State,
	appHash []byte,
	appBlockHeight int64,
	proxyApp proxy.AppConns,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReplayBlocksWithContext is cancellable version of ReplayBlocks.
func (h *Handshaker) ReplayBlocksWithContext(
	ctx context.Context,
	state sm.State,
	appHash []byte,
	appBlockHeight int64,
	proxyApp proxy.AppConns,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If appBlockHeight == 0 it means that we are at genesis and hence should send InitChain.

// we only update state when we are in initial state
// If the app did not return an app hash, we keep the one set from the genesis doc in
// the state. We don't set appHash since we don't want the genesis doc app hash
// recorded in the genesis block. We should probably just remove GenesisDoc.AppHash.

// If the app returned validators or consensus params, update the state.

// If validator set is not set in genesis and still empty after InitChain, exit.

// update timeouts based on the InitChainSync response

// We update the last results hash with the empty hash, to conform with RFC-6962.

// First handle edge cases and constraints on the storeBlockHeight and storeBlockBase.

// the app has no state, and the block store is truncated above the initial height

// the app is too far behind truncated store (can be 1 behind since we replay the next)

// the app should never be ahead of the store (but this is under app's control)

// the state should never be ahead of the store (this is under CometBFT's control)

// store should be at most one ahead of the state (this is under CometBFT's control)

// Now either store is equal to state, or one ahead.
// For each, consider all cases of where the app could be, given app <= store
//nolint:staticcheck
// CometBFT ran Commit and saved the state.
// Either the app is asking for replay, or we're all synced up.

// the app is behind, so replay blocks, but no need to go through WAL (state is already synced to store)

// We're good!

// We saved the block in the store but haven't updated the state,
// so we'll need to replay a block using the WAL.

// the app is further behind than it should be, so replay blocks
// but leave the last block to go through the WAL

// We haven't run Commit (both the state and app are one block behind),
// so replayBlock with the real app.
// NOTE: We could instead use the cs.WAL on cs.Start,
// but we'd have to allow the WAL to replay a block that wrote it's #ENDHEIGHT

// We ran Commit, but didn't save the state, so replayBlock with mock app.

// NOTE: There is a rare edge case where a node has upgraded from
// v0.37 with endblock to v0.38 with finalize block and thus
// does not have the app hash saved from the previous height
// here we take the appHash provided from the Info handshake

func (h *Handshaker) replayBlocks(
	ctx context.Context,
	state sm.State,
	proxyApp proxy.AppConns,
	appBlockHeight,
	storeBlockHeight int64,
	mutateState bool) ([]byte, error) {
	_ = "STUB: not implemented"
	// App is further behind than it should be, so we need to replay blocks.
	// We replay all blocks from appBlockHeight+1.
	//
	// Note that we don't have an old version of the state,
	// so we by-pass state validation/mutation using sm.ExecCommitBlock.
	// This also means we won't be saving validator sets if they change during this period.
	// TODO: Load the historical information to fix this and just use state.ApplyBlock
	//
	// If mutateState == true, the final block is replayed with h.replayBlock()
	return nil, nil
}

// Extra check to ensure the app was not changed in a way it shouldn't have.

// sync the final block

// ApplyBlock on the proxyApp with the last block.
func (h *Handshaker) replayBlock(state sm.State, height int64, proxyApp proxy.AppConnConsensus) (sm.State, error) {
	_ = "STUB: not implemented"
	return *new(sm.State), nil
}

// Use stubs for both mempool and evidence pool since no transactions nor
// evidence are needed here - block already exists.

func assertAppHashEqualsOneFromBlock(appHash []byte, block *types.Block) {
	_ = "STUB: not implemented"
	return
}

func assertAppHashEqualsOneFromState(appHash []byte, state sm.State) {
	_ = "STUB: not implemented"
	return
}
