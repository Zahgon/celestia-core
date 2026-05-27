// Commons for HTTP handling
package server

import (
	"bufio"
	"net"
	"net/http"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	types "github.com/cometbft/cometbft/rpc/jsonrpc/types"
)

// Config is a RPC server configuration.
type Config struct {
	// see netutil.LimitListener
	MaxOpenConnections int
	// mirrors http.Server#ReadTimeout
	ReadTimeout time.Duration
	// mirrors http.Server#WriteTimeout
	WriteTimeout time.Duration
	// MaxBodyBytes controls the maximum number of bytes the
	// server will read parsing the request body.
	MaxBodyBytes int64
	// mirrors http.Server#MaxHeaderBytes
	MaxHeaderBytes int
	// maximum number of requests in a batch request
	MaxRequestBatchSize int
}

// DefaultConfig returns a default configuration.
func DefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

// unlimited

// 1MB
// same as the net/http default
// default to max 10 requests per batch

// Serve creates a http.Server and calls Serve with the given listener. It
// wraps handler with RecoverAndLogHandler and a handler, which limits the max
// body size to config.MaxBodyBytes.
//
// NOTE: This function blocks - you may want to call it in a go-routine.
func Serve(listener net.Listener, handler http.Handler, logger log.Logger, config *Config) error {
	_ = "STUB: not implemented"
	return nil
}

// ServeTLS creates a http.Server and calls ServeTLS with the given listener,
// certFile and keyFile. It wraps handler with RecoverAndLogHandler and a
// handler, which limits the max body size to config.MaxBodyBytes.
//
// NOTE: This function blocks - you may want to call it in a go-routine.
func ServeTLS(
	listener net.Listener,
	handler http.Handler,
	certFile, keyFile string,
	logger log.Logger,
	config *Config,
) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteRPCResponseHTTPError marshals res as JSON (with indent) and writes it
// to w.
//
// source: https://www.jsonrpc.org/historical/json-rpc-over-http.html
func WriteRPCResponseHTTPError(
	w http.ResponseWriter,
	httpCode int,
	res types.RPCResponse,
) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteRPCResponseHTTP marshals res as JSON (with indent) and writes it to w.
func WriteRPCResponseHTTP(w http.ResponseWriter, res ...types.RPCResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteCacheableRPCResponseHTTP marshals res as JSON (with indent) and writes
// it to w. Adds cache-control to the response header and sets the expiry to
// one day.
func WriteCacheableRPCResponseHTTP(w http.ResponseWriter, res ...types.RPCResponse) error {
	_ = "STUB: not implemented"
	return nil
}

type httpHeader struct {
	name  string
	value string
}

func writeRPCResponseHTTP(w http.ResponseWriter, headers []httpHeader, res ...types.RPCResponse) error {
	_ = "STUB: not implemented"
	return nil
}

//-----------------------------------------------------------------------------

// RecoverAndLogHandler wraps an HTTP handler, adding error logging.
// If the inner function panics, the outer function recovers, logs, sends an
// HTTP 500 error response.
func RecoverAndLogHandler(handler http.Handler, logger log.Logger) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Wrap the ResponseWriter to remember the status

// Handle any panics in the panic handler below. Does not use the logger, since we want
// to avoid any further panics. However, we try to return a 500, since it otherwise
// defaults to 200 and there is no other way to terminate the connection. If that
// should panic for whatever reason then the Go HTTP server will handle it and
// terminate the connection - panicing is the de-facto and only way to get the Go HTTP
// server to terminate the request and close the connection/stream:
// https://github.com/golang/go/issues/17790#issuecomment-258481416

// Send a 500 error if a panic happens during a handler.
// Without this, Chrome & Firefox were retrying aborted ajax requests,
// at least to my localhost.

// If RPCResponse

// Panics can contain anything, attempt to normalize it as an error.

// Finally, log.

// Remember the status for logging
type responseWriterWrapper struct {
	Status int
	http.ResponseWriter
}

func (w *responseWriterWrapper) WriteHeader(status int) { _ = "STUB: not implemented"; return }

// implements http.Hijacker
func (w *responseWriterWrapper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

type defaultHandler struct {
	h http.Handler
}

func (h defaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return

	// Listen starts a new net.Listener on the given address.
	// It returns an error if the address is invalid or the call to Listen() fails.
}

func Listen(addr string, maxOpenConnections int) (listener net.Listener, err error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

// Middleware

// PreChecksHandler is a middleware function that checks the size of batch requests and returns an error
// if it exceeds the maximum configured size. It also checks if the request body is not greater than the
// configured maximum request body bytes limit.
func PreChecksHandler(next http.Handler, config *Config) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// ensure that the current request body bytes is not greater than the configured maximum request body bytes

// if maxBatchSize is 0 then don't constraint the limit of requests per batch
// It cannot be negative because the config.toml validation requires it to be
// greater than or equal to 0

// if no err it means multiple requests, check if the number of request exceeds
// the maximum batch size configured

// if the number of requests in batch exceed the maximum configured then return an error

// ensure the request body can be read again by other handlers

// next handler
