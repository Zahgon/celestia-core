package consensus

import (
	"os"

	cfg "github.com/cometbft/cometbft/config"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/types"
)

const (
	// event bus subscriber
	subscriber = "replay-file"
)

//--------------------------------------------------------
// replay messages interactively or all at once

// replay the wal file
func RunReplayFile(config cfg.BaseConfig, csConfig *cfg.ConsensusConfig, console bool) {
	_ = "STUB: not implemented"
	return
}

// Replay msgs in file or start the console
func (cs *State) ReplayFile(file string, console bool) error { _ = "STUB: not implemented"; return nil }

// ensure all new step events are regenerated as expected

// just open the file for reading, no need to use wal

// apply N msgs in a row

//------------------------------------------------
// playback manager

type playback struct {
	cs *State

	fp    *os.File
	dec   *WALDecoder
	count int // how many lines/msgs into the file are we

	// replays can be reset to beginning
	fileName     string   // so we can close/reopen the file
	genesisState sm.State // so the replay session knows where to restart from
}

func newPlayback(fileName string, fp *os.File, cs *State, genState sm.State) *playback {
	_ = "STUB: not implemented"
	return nil
}

// go back count steps by resetting the state and running (pb.count - count) steps
func (pb *playback) replayReset(count int, newStepSub types.Subscription) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *State) startForReplay() { _ = "STUB: not implemented"; return }

/* TODO:!
// since we replay tocks we just ignore ticks
	go func() {
		for {
			select {
			case <-cs.tickChan:
			case <-cs.Quit:
				return
			}
		}
	}()*/

// console function for parsing input and running commands
func (pb *playback) replayConsoleLoop() int { _ = "STUB: not implemented"; return 0 }

// "next" -> replay next message
// "next N" -> replay next N messages

// "back" -> go back one message
// "back N" -> go back N messages

// NOTE: "back" is not supported in the state machine design,
// so we restart and replay up to

// ensure all new step events are regenerated as expected

// "rs" -> print entire round state
// "rs short" -> print height/round/step
// "rs <field>" -> print another field of the round state

//--------------------------------------------------------------------------------

// convenience for replay mode
func newConsensusStateForReplay(config cfg.BaseConfig, csConfig *cfg.ConsensusConfig) *State {
	_ = "STUB: not implemented"
	return nil
}

// Get BlockStore

// Get State

// Create proxyAppConn connection (consensus, mempool, query)

// TODO pass a tracer from here
