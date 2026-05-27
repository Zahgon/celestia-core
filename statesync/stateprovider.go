package statesync

import (
	"context"

	"github.com/cometbft/cometbft/libs/log"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/light"
	lightprovider "github.com/cometbft/cometbft/light/provider"
	cmtstate "github.com/cometbft/cometbft/proto/tendermint/state"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/types"
)

//go:generate ../scripts/mockery_generate.sh StateProvider

// StateProvider is a provider of trusted state data for bootstrapping a node. This refers
// to the state.State object, not the state machine.
type StateProvider interface {
	// AppHash returns the app hash after the given height has been committed.
	AppHash(ctx context.Context, height uint64) ([]byte, error)
	// Commit returns the commit at the given height.
	Commit(ctx context.Context, height uint64) (*types.Commit, error)
	// State returns a state object at the given height.
	State(ctx context.Context, height uint64) (sm.State, error)
}

// lightClientStateProvider is a state provider using the light client.
type lightClientStateProvider struct {
	cmtsync.Mutex // light.Client is not concurrency-safe
	lc            *light.Client
	version       cmtstate.Version
	initialHeight int64
	providers     map[lightprovider.Provider]string
}

// NewLightClientStateProvider creates a new StateProvider using a light client and RPC clients.
func NewLightClientStateProvider(
	ctx context.Context,
	chainID string,
	version cmtstate.Version,
	initialHeight int64,
	servers []string,
	trustOptions light.TrustOptions,
	logger log.Logger,
) (StateProvider, error) {
	_ = "STUB: not implemented"
	return *new(StateProvider), nil
}

// We store the RPC addresses keyed by provider, so we can find the address of the primary
// provider used by the light client and use it to fetch consensus parameters.

// AppHash implements StateProvider.
func (s *lightClientStateProvider) AppHash(ctx context.Context, height uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// We have to fetch the next height, which contains the app hash for the previous height.
		nil
}

// We also try to fetch the blocks at height H and H+2, since we need these
// when building the state while restoring the snapshot. This avoids the race
// condition where we try to restore a snapshot before H+2 exists.
//
// FIXME This is a hack, since we can't add new methods to the interface without
// breaking it. We should instead have a Has(ctx, height) method which checks
// that the state provider has access to the necessary data for the height.
// We piggyback on AppHash() since it's called when adding snapshots to the pool.

// Commit implements StateProvider.
func (s *lightClientStateProvider) Commit(ctx context.Context, height uint64) (*types.Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// State implements StateProvider.
func (s *lightClientStateProvider) State(ctx context.Context, height uint64) (sm.State, error) {
	_ = "STUB: not implemented"
	return *new(sm.State), nil
}

// The snapshot height maps onto the state heights as follows:
//
// height: last block, i.e. the snapshotted height
// height+1: current block, i.e. the first block we'll process after the snapshot
// height+2: next block, i.e. the second block after the snapshot
//
// We need to fetch the NextValidators from height+2 because if the application changed
// the validator set at the snapshot height then this only takes effect at height+2.

// We'll also need to fetch consensus params via RPC, using light client verification.

// rpcClient sets up a new RPC client
func rpcClient(server string) (*rpchttp.HTTP, error) { _ = "STUB: not implemented"; return nil, nil }
