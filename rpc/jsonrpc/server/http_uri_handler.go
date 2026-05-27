package server

import (
	"net/http"
	"reflect"
	"regexp"

	"github.com/cometbft/cometbft/libs/log"
)

// HTTP + URI handler

var reInt = regexp.MustCompile(`^-?[0-9]+$`)

// convert from a function name to the http handler
func makeHTTPHandler(rpcFunc *RPCFunc, logger log.Logger) func(http.ResponseWriter, *http.Request) {
	_ = "STUB: not implemented"
	// Always return -1 as there's no ID here.
	return nil
}

// URIClientRequestID

// Exception for websocket endpoints

// All other endpoints

//nolint:prealloc

// Covert an http query to a list of properly typed values.
// To be properly decoded the arg must be a concrete type from CometBFT (if its an interface).
func httpParamsToArgs(rpcFunc *RPCFunc, r *http.Request) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	// skip types.Context
	return nil, nil
}

// set default for that type

// log.Notice("param to arg", "argType", argType, "name", name, "arg", arg)

func jsonStringToArg(rt reflect.Type, arg string) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func nonJSONStringToArg(rt reflect.Type, arg string) (reflect.Value, bool, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false, nil
}

// NOTE: rt.Kind() isn't a pointer.
func _nonJSONStringToArg(rt reflect.Type, arg string) (reflect.Value, bool, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false, nil
}

func getParam(r *http.Request, param string) string { _ = "STUB: not implemented"; return "" }
