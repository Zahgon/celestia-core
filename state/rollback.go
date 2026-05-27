package state

// Rollback overwrites the current CometBFT state (height n) with the most
// recent previous state (height n - 1).
// Note that this function does not affect application state.
func Rollback(bs BlockStore, ss Store, removeBlock bool) (int64, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// NOTE: persistence of state and blocks don't happen atomically. Therefore it is possible that
// when the user stopped the node the state wasn't updated but the blockstore was. Discard the
// pending block before continuing.

// If the state store isn't one below nor equal to the blockstore height than this violates the
// invariant

// state store height is equal to blockstore height. We're good to proceed with rolling back state

// We also need to retrieve the latest block because the app hash and last
// results hash is only agreed upon in the following block.

// this can only happen if the validator set changed since the last block

// this can only happen if params changed from the last block

// build the new state from the old state and the prior block

// immutable fields

// persist the new state. This overrides the invalid one. NOTE: this will also
// persist the validator set and consensus params over the existing structures,
// but both should be the same

// If removeBlock is true then also remove the block associated with the previous state.
// This will mean both the last state and last block height is equal to n - 1
