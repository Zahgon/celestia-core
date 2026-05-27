package server

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/cometbft/cometbft/libs/log"
)

// HTTP + JSON handler

// jsonrpc calls grab the given method's function info and runs reflect.Call
func makeJSONRPCHandler(funcMap map[string]*RPCFunc, logger log.Logger) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// if it's an empty request (like from a browser), just display a list of
// functions

// first try to unmarshal the incoming request as an array of RPC requests

// next, try to unmarshal as a single request

// Set the default response cache to true unless
// 1. Any RPC request error.
// 2. Any RPC request doesn't allow to be cached.
// 3. Any RPC request has the height argument and the value is 0 (the default).

// A Notification is a Request object without an "id" member.
// The Server MUST NOT reply to a Notification, including those that are within a batch request.

func handleInvalidJSONRPCPaths(next http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Since the pattern "/" matches all paths not matched by other registered patterns,
//  we check whether the path is indeed "/", otherwise return a 404 error

func mapParamsToArgs(
	rpcFunc *RPCFunc,
	params map[string]json.RawMessage,
	argsOffset int,
) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use default for that type

func arrayParamsToArgs(
	rpcFunc *RPCFunc,
	params []json.RawMessage,
	argsOffset int,
) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// raw is unparsed json (from json.RawMessage) encoding either a map or an
// array.
//
// Example:
//
//	rpcFunc.args = [rpctypes.Context string]
//	rpcFunc.argNames = ["arg"]
func jsonParamsToArgs(rpcFunc *RPCFunc, raw []byte) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil,

		// TODO: Make more efficient, perhaps by checking the first character for '{' or '['?
		// First, try to get the map.
		nil
}

// Otherwise, try an array.

// Otherwise, bad format, we cannot parse

// writes a list of available rpc endpoints as an html page
func writeListOfEndpoints(w http.ResponseWriter, r *http.Request, funcMap map[string]*RPCFunc) {
	_ = "STUB: not implemented"
	return
}

//nolint:staticcheck

//nolint:staticcheck

//nolint: errcheck
