package light

import (
	"context"
	"time"

	"github.com/cometbft/cometbft/light/provider"
	"github.com/cometbft/cometbft/light/store"
)

// NewHTTPClient initiates an instance of a light client using HTTP addresses
// for both the primary provider and witnesses of the light client. A trusted
// header and hash must be passed to initialize the client.
//
// See all Option(s) for the additional configuration.
// See NewClient.
func NewHTTPClient(
	ctx context.Context,
	chainID string,
	trustOptions TrustOptions,
	primaryAddress string,
	witnessesAddresses []string,
	trustedStore store.Store,
	options ...Option) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewHTTPClientFromTrustedStore initiates an instance of a light client using
// HTTP addresses for both the primary provider and witnesses and uses a
// trusted store as the root of trust.
//
// See all Option(s) for the additional configuration.
// See NewClientFromTrustedStore.
func NewHTTPClientFromTrustedStore(
	chainID string,
	trustingPeriod time.Duration,
	primaryAddress string,
	witnessesAddresses []string,
	trustedStore store.Store,
	options ...Option) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func providersFromAddresses(addrs []string, chainID string) ([]provider.Provider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
