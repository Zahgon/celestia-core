package server

import (
	"net/http"
	"reflect"

	"github.com/cometbft/cometbft/libs/log"
)

// RegisterRPCFuncs adds a route for each function in the funcMap, as well as
// general jsonrpc and websocket handlers for all functions. "result" is the
// interface on which the result objects are registered, and is popualted with
// every RPCResponse
func RegisterRPCFuncs(mux *http.ServeMux, funcMap map[string]*RPCFunc, logger log.Logger) {
	_ = "STUB: not implemented"
	// HTTP endpoints
	return
}

// JSONRPC endpoints

type Option func(*RPCFunc)

// Cacheable enables returning a cache control header from RPC functions to
// which it is applied.
//
// `noCacheDefArgs` is a list of argument names that, if omitted or set to
// their defaults when calling the RPC function, will skip the response
// caching.
func Cacheable(noCacheDefArgs ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Ws enables WebSocket communication.
func Ws() Option { _ = "STUB: not implemented"; return *new(Option) }

// RPCFunc contains the introspected type information for a function
type RPCFunc struct {
	f              reflect.Value          // underlying rpc function
	args           []reflect.Type         // type of each function arg
	returns        []reflect.Type         // type of each return arg
	argNames       []string               // name of each argument
	cacheable      bool                   // enable cache control
	ws             bool                   // enable websocket communication
	noCacheDefArgs map[string]interface{} // a lookup table of args that, if not supplied or are set to default values, cause us to not cache
}

// NewRPCFunc wraps a function for introspection.
// f is the function, args are comma separated argument names
func NewRPCFunc(f interface{}, args string, options ...Option) *RPCFunc {
	_ = "STUB: not implemented"
	return nil
}

// NewWSRPCFunc wraps a function for introspection and use in the websockets.
func NewWSRPCFunc(f interface{}, args string, options ...Option) *RPCFunc {
	_ = "STUB: not implemented"
	return nil
}

// cacheableWithArgs returns whether or not a call to this function is cacheable,
// given the specified arguments.
func (f *RPCFunc) cacheableWithArgs(args []reflect.Value) bool {
	_ = "STUB: not implemented"
	return false
}

// Skip the context variable common to all RPC functions

// f.argNames does not include the context variable

// Argument with default value was not supplied

// Argument with default value is set to its zero value

func newRPCFunc(f interface{}, args string, options ...Option) *RPCFunc {
	_ = "STUB: not implemented"
	return nil
}

// return a function's argument types
func funcArgTypes(f interface{}) []reflect.Type { _ = "STUB: not implemented"; return nil }

// return a function's return types
func funcReturnTypes(f interface{}) []reflect.Type { _ = "STUB: not implemented"; return nil }

//-------------------------------------------------------------

// NOTE: assume returns is result struct and error. If error is not nil, return it
func unreflectResult(returns []reflect.Value) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the result is a registered interface,
// we need a pointer to it so we can marshal with type byte
