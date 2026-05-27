package client

import (
	"context"
	"net/http"

	types "github.com/cometbft/cometbft/rpc/jsonrpc/types"
)

const (
	// URIClientRequestID in a request ID used by URIClient
	URIClientRequestID = types.JSONRPCIntID(-1)
)

// URIClient is a JSON-RPC client, which sends POST form HTTP requests to the
// remote server.
//
// URIClient is safe for concurrent use by multiple goroutines.
type URIClient struct {
	address string
	client  *http.Client
}

var _ HTTPClient = (*URIClient)(nil)

// NewURI returns a new client.
// An error is returned on invalid remote.
// The function panics when remote is nil.
func NewURI(remote string) (*URIClient, error) { _ = "STUB: not implemented"; return nil, nil }

// Call issues a POST form HTTP request.
func (c *URIClient) Call(ctx context.Context, method string,
	params map[string]interface{}, result interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
