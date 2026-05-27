package http

import (
	"context"
	"regexp"
	"time"

	"github.com/cometbft/cometbft/light/provider"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/types"
)

var (
	// This is very brittle, see: https://github.com/tendermint/tendermint/issues/4740
	regexpMissingHeight = regexp.MustCompile(`height \d+ is not available`)
	regexpTooHigh       = regexp.MustCompile(`height \d+ must be less than or equal to`)
	regexpTimedOut      = regexp.MustCompile(`Timeout exceeded`)

	maxRetryAttempts      = 5
	timeout          uint = 5 // sec.
)

// http provider uses an RPC client to obtain the necessary information.
type http struct {
	chainID string
	client  rpcclient.RemoteClient
}

// New creates a HTTP provider, which is using the rpchttp.HTTP client under
// the hood. If no scheme is provided in the remote URL, http will be used by
// default. The 5s timeout is used for all requests.
func New(chainID, remote string) (provider.Provider, error) {
	_ = "STUB: not implemented"
	// Ensure URL scheme is set (default HTTP) when not provided.
	return *new(provider.Provider), nil
}

// NewWithClient allows you to provide a custom client.
func NewWithClient(chainID string, client rpcclient.RemoteClient) provider.Provider {
	_ = "STUB: not implemented"
	return *new(provider.Provider)
}

// ChainID returns a chainID this provider was configured with.
func (p *http) ChainID() string { _ = "STUB: not implemented"; return "" }

func (p *http) String() string { _ = "STUB: not implemented"; return "" }

// LightBlock fetches a LightBlock at the given height and checks the
// chainID matches.
func (p *http) LightBlock(ctx context.Context, height int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReportEvidence calls `/broadcast_evidence` endpoint.
func (p *http) ReportEvidence(ctx context.Context, ev types.Evidence) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *http) validatorSet(ctx context.Context, height *int64) (*types.ValidatorSet, error) {
	_ = "STUB: not implemented"
	// Since the malicious node could report a massive number of pages, making us
	// spend a considerable time iterating, we restrict the number of pages here.
	// => 10000 validators max
	return nil, nil
}

// Validate response.

// if we have exceeded retry attempts then return no response error

// we wait and try again with exponential backoff

// context canceled or connection refused we return the error

func (p *http) signedHeader(ctx context.Context, height *int64) (*types.SignedHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// See https://github.com/cometbft/cometbft/issues/575
// If the node is starting at a non-zero height, but does not yet
// have any blocks, it can return an empty signed header without
// returning an error.
//nolint:staticcheck
// Technically this means that the provider still needs to
// catch up.

// we wait and try again with exponential backoff

// either context was canceled or connection refused.

func validateHeight(height int64) (*int64, error) { _ = "STUB: not implemented"; return nil, nil }

// exponential backoff (with jitter)
// 0.5s -> 2s -> 4.5s -> 8s -> 12.5 with 1s variation
func backoffTimeout(attempt uint16) time.Duration {
	_ = "STUB: not implemented"
	//nolint:gosec // G404: Use of weak random number generator
	return *new(time.Duration)
}
