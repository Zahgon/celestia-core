package coregrpc

import (
	"context"
	"sync"

	"github.com/cometbft/cometbft/libs/pubsub"
	core "github.com/cometbft/cometbft/rpc/core"
	eventstypes "github.com/cometbft/cometbft/types"
)

type broadcastAPI struct {
	env *core.Environment
}

func (bapi *broadcastAPI) Ping(context.Context, *RequestPing) (*ResponsePing, error) {
	_ = "STUB: not implemented"
	// kvstore so we can check if the server is up
	return nil, nil
}

func (bapi *broadcastAPI) BroadcastTx(_ context.Context, req *RequestBroadcastTx) (*ResponseBroadcastTx, error) {
	_ = "STUB: not implemented"
	// NOTE: there's no way to get client's remote address
	// see https://stackoverflow.com/questions/33684570/session-and-remote-ip-address-in-grpc-go
	return nil, nil
}

type BlockAPI struct {
	env *core.Environment
	sync.Mutex
	heightListeners      map[chan SubscribeNewHeightsResponse]struct{}
	newBlockSubscription eventstypes.Subscription
	subscriptionID       string
	subscriptionQuery    pubsub.Query
}

func NewBlockAPI(env *core.Environment) *BlockAPI { _ = "STUB: not implemented"; return nil }

func (blockAPI *BlockAPI) StartNewBlockEventListener(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// this will happen when the context is done. we can stop here

// this will happen when the context is done. we can stop here

// RetryAttempts the number of retry times when the subscription is closed.
const RetryAttempts = 6

// SubscriptionCapacity the maximum number of pending blocks in the subscription.
const SubscriptionCapacity = 500

// subscribe creates the initial EventBus subscription if one does not
// already exist. Holding blockAPI.Lock keeps this write synchronized
// with retryNewBlocksSubscription (which also writes the field under
// the lock) and Stop (which reads it under the lock).
func (blockAPI *BlockAPI) subscribe(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (blockAPI *BlockAPI) retryNewBlocksSubscription(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (blockAPI *BlockAPI) broadcastToListeners(ctx context.Context, height int64, hash []byte) {
	_ = "STUB: not implemented"
	// Snapshot the current set of listeners under the lock so we do not
	// hold the lock during sends. A slow listener must not block the
	// broadcaster (see https://github.com/celestiaorg/celestia-core/issues/2967).
	return
}

// sendNonBlocking attempts to deliver event to ch without blocking.
// If ch is full, the event is dropped and a debug log is emitted: a
// slow subscriber must not block the broadcaster or any other
// subscriber. Callers that require lossless delivery should use
// BlockByHeight to back-fill any gaps. The recover guards against a
// concurrent close of ch and evicts the listener in that case
// (matching the prior behavior).
func (blockAPI *BlockAPI) sendNonBlocking(ch chan SubscribeNewHeightsResponse, event SubscribeNewHeightsResponse) {
	_ = "STUB: not implemented"
	return
}

func (blockAPI *BlockAPI) addHeightListener() chan SubscribeNewHeightsResponse {
	_ = "STUB: not implemented"
	return nil
}

func (blockAPI *BlockAPI) removeHeightListener(ch chan SubscribeNewHeightsResponse) {
	_ = "STUB: not implemented"
	return
}

// removeHeightListenerLocked removes ch from heightListeners. The caller
// must hold blockAPI.Lock().
func (blockAPI *BlockAPI) removeHeightListenerLocked(ch chan SubscribeNewHeightsResponse) {
	_ = "STUB: not implemented"
	return
}

// closeAllListenersLocked clears every registered height listener.
// The caller must hold blockAPI.Lock(); the function does not acquire
// the lock itself because doing so would deadlock against Stop, which
// already holds it (sync.Mutex is not reentrant).
func (blockAPI *BlockAPI) closeAllListenersLocked() { _ = "STUB: not implemented"; return }

// Stop cleans up the BlockAPI instance by closing all listeners
// and ensuring no further events are processed.
func (blockAPI *BlockAPI) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// close all height listeners

// stop the events subscription. We deliberately do not clear
// blockAPI.newBlockSubscription here: StartNewBlockEventListener reads
// the field without holding the lock, so a write would race with that
// goroutine. Unsubscribe is sufficient to drain the subscription; the
// goroutine exits via ctx.Done after the caller cancels the context.

func (blockAPI *BlockAPI) BlockByHash(req *BlockByHashRequest, stream BlockAPI_BlockByHashServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (blockAPI *BlockAPI) BlockByHeight(req *BlockByHeightRequest, stream BlockAPI_BlockByHeightServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (blockAPI *BlockAPI) Status(_ context.Context, _ *StatusRequest) (*StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (blockAPI *BlockAPI) Commit(_ context.Context, req *CommitRequest) (*CommitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (blockAPI *BlockAPI) ValidatorSet(_ context.Context, req *ValidatorSetRequest) (*ValidatorSetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (blockAPI *BlockAPI) SubscribeNewHeights(_ *SubscribeNewHeightsRequest, stream BlockAPI_SubscribeNewHeightsServer) error {
	_ = "STUB: not implemented"
	return nil
}

type BlobstreamAPI struct {
	env *core.Environment
}

func NewBlobstreamAPI(env *core.Environment) *BlobstreamAPI { _ = "STUB: not implemented"; return nil }

func (blobAPI *BlobstreamAPI) DataRootInclusionProof(_ context.Context, req *DataRootInclusionProofRequest) (*DataRootInclusionProofResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
