package core

import (
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
)

const (
	// maxQueryLength is the maximum length of a query string that will be
	// accepted. This is just a safety check to avoid outlandish queries.
	maxQueryLength = 512
)

// Subscribe for events via WebSocket.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Websocket/subscribe
func (env *Environment) Subscribe(ctx *rpctypes.Context, query string) (*ctypes.ResultSubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Capture the current ID, since it can change in the future.

// Unsubscribe from events via WebSocket.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Websocket/unsubscribe
func (env *Environment) Unsubscribe(ctx *rpctypes.Context, query string) (*ctypes.ResultUnsubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnsubscribeAll from all events via WebSocket.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Websocket/unsubscribe_all
func (env *Environment) UnsubscribeAll(ctx *rpctypes.Context) (*ctypes.ResultUnsubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
