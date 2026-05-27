package core

import (
	rpc "github.com/cometbft/cometbft/rpc/jsonrpc/server"
)

// TODO: better system than "unsafe" prefix

type RoutesMap map[string]*rpc.RPCFunc

// Routes is a map of available routes.
func (env *Environment) GetRoutes() RoutesMap {
	_ = "STUB: not implemented"

	// subscribe/unsubscribe are reserved for websocket events.
	return *new(RoutesMap)
}

// info AP

// tx broadcast API

// abci API

// evidence API

// celestia-specific API

// AddUnsafeRoutes adds unsafe routes.
func (env *Environment) AddUnsafeRoutes(routes RoutesMap) {
	_ = "STUB: not implemented"
	// control API
	return
}
