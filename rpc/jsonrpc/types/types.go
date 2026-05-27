package types

import (
	"context"
	"encoding/json"
	"net/http"
)

// a wrapper to emulate a sum type: jsonrpcid = string | int
// TODO: refactor when Go 2.0 arrives https://github.com/golang/go/issues/19412
type jsonrpcid interface {
	isJSONRPCID()
}

// JSONRPCStringID a wrapper for JSON-RPC string IDs
type JSONRPCStringID string

func (JSONRPCStringID) isJSONRPCID() { _ = "STUB: not implemented"; return }
func (id JSONRPCStringID) String() string {
	_ = "STUB: not implemented"

	// JSONRPCIntID a wrapper for JSON-RPC integer IDs
	return ""
}

type JSONRPCIntID int

func (JSONRPCIntID) isJSONRPCID()      { _ = "STUB: not implemented"; return }
func (id JSONRPCIntID) String() string { _ = "STUB: not implemented"; return "" }

func idFromInterface(idInterface interface{}) (jsonrpcid, error) {
	_ = "STUB: not implemented"
	return *new(jsonrpcid), nil
}

// json.Unmarshal uses float64 for all numbers
// (https://golang.org/pkg/encoding/json/#Unmarshal),
// but the JSONRPC2.0 spec says the id SHOULD NOT contain
// decimals - so we truncate the decimals here.

//----------------------------------------
// REQUEST

type RPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      jsonrpcid       `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"` // must be map[string]interface{} or []interface{}
}

// UnmarshalJSON custom JSON unmarshalling due to jsonrpcid being string or int
func (req *RPCRequest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// must be map[string]interface{} or []interface{}

// notification

func NewRPCRequest(id jsonrpcid, method string, params json.RawMessage) RPCRequest {
	_ = "STUB: not implemented"
	return *new(RPCRequest)
}

func (req RPCRequest) String() string { _ = "STUB: not implemented"; return "" }

func MapToRequest(id jsonrpcid, method string, params map[string]interface{}) (RPCRequest, error) {
	_ = "STUB: not implemented"
	return *new(RPCRequest), nil
}

func ArrayToRequest(id jsonrpcid, method string, params []interface{}) (RPCRequest, error) {
	_ = "STUB: not implemented"
	return *new(RPCRequest), nil
}

//----------------------------------------
// RESPONSE

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

func (err RPCError) Error() string { _ = "STUB: not implemented"; return "" }

type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      jsonrpcid       `json:"id,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *RPCError       `json:"error,omitempty"`
}

// UnmarshalJSON custom JSON unmarshalling due to jsonrpcid being string or int
func (resp *RPCResponse) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRPCSuccessResponse(id jsonrpcid, res interface{}) RPCResponse {
	_ = "STUB: not implemented"
	return *new(RPCResponse)
}

func NewRPCErrorResponse(id jsonrpcid, code int, msg string, data string) RPCResponse {
	_ = "STUB: not implemented"
	return *new(RPCResponse)
}

func (resp RPCResponse) String() string { _ = "STUB: not implemented"; return "" }

// From the JSON-RPC 2.0 spec:
//
//	If there was an error in detecting the id in the Request object (e.g. Parse
//	error/Invalid Request), it MUST be Null.
func RPCParseError(err error) RPCResponse { _ = "STUB: not implemented"; return *new(RPCResponse) }

// From the JSON-RPC 2.0 spec:
//
//	If there was an error in detecting the id in the Request object (e.g. Parse
//	error/Invalid Request), it MUST be Null.
func RPCInvalidRequestError(id jsonrpcid, err error) RPCResponse {
	_ = "STUB: not implemented"
	return *new(RPCResponse)
}

func RPCMethodNotFoundError(id jsonrpcid) RPCResponse {
	_ = "STUB: not implemented"
	return *new(RPCResponse)
}

func RPCInvalidParamsError(id jsonrpcid, err error) RPCResponse {
	_ = "STUB: not implemented"
	return *new(RPCResponse)
}

func RPCInternalError(id jsonrpcid, err error) RPCResponse {
	_ = "STUB: not implemented"
	return *new(RPCResponse)
}

func RPCServerError(id jsonrpcid, err error) RPCResponse {
	_ = "STUB: not implemented"
	return *new(RPCResponse)
}

//----------------------------------------

// WSRPCConnection represents a websocket connection.
type WSRPCConnection interface {
	// GetRemoteAddr returns a remote address of the connection.
	GetRemoteAddr() string
	// WriteRPCResponse writes the response onto connection (BLOCKING).
	WriteRPCResponse(context.Context, RPCResponse) error
	// TryWriteRPCResponse tries to write the response onto connection (NON-BLOCKING).
	TryWriteRPCResponse(RPCResponse) bool
	// Context returns the connection's context.
	Context() context.Context
}

// Context is the first parameter for all functions. It carries a json-rpc
// request, http request and websocket connection.
//
// - JSONReq is non-nil when JSONRPC is called over websocket or HTTP.
// - WSConn is non-nil when we're connected via a websocket.
// - HTTPReq is non-nil when URI or JSONRPC is called over HTTP.
type Context struct {
	// json-rpc request
	JSONReq *RPCRequest
	// websocket connection
	WSConn WSRPCConnection
	// http request
	HTTPReq *http.Request
}

// RemoteAddr returns the remote address (usually a string "IP:port").
// If neither HTTPReq nor WSConn is set, an empty string is returned.
// HTTP:
//
//	http.Request#RemoteAddr
//
// WS:
//
//	result of GetRemoteAddr
func (ctx *Context) RemoteAddr() string { _ = "STUB: not implemented"; return "" }

// Context returns the request's context.
// The returned context is always non-nil; it defaults to the background context.
// HTTP:
//
//	The context is canceled when the client's connection closes, the request
//	is canceled (with HTTP/2), or when the ServeHTTP method returns.
//
// WS:
//
//	The context is canceled when the client's connections closes.
func (ctx *Context) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

//----------------------------------------
// SOCKETS

// Determine if its a unix or tcp socket.
// If tcp, must specify the port; `0.0.0.0` will return incorrectly as "unix" since there's no port
// TODO: deprecate
func SocketType(listenAddr string) string { _ = "STUB: not implemented"; return "" }
