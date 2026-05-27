package privval

import (
	"time"

	"github.com/cometbft/cometbft/libs/trace"

	"github.com/cometbft/cometbft/crypto"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cometbft/cometbft/types"
)

// SignerClient implements PrivValidator.
// Handles remote validator connections that provide signing services
type SignerClient struct {
	endpoint *SignerListenerEndpoint
	chainID  string
	tracer   trace.Tracer
}

var _ types.PrivValidator = (*SignerClient)(nil)

// NewSignerClient returns an instance of SignerClient.
// it will start the endpoint (if not already started)
func NewSignerClient(endpoint *SignerListenerEndpoint, chainID string) (*SignerClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetTracer sets the tracer for the SignerClient
func (sc *SignerClient) SetTracer(tracer trace.Tracer) {
	_ = "STUB: not implemented"

	// Close closes the underlying connection
	return
}

func (sc *SignerClient) Close() error { _ = "STUB: not implemented"; return nil }

// IsConnected indicates with the signer is connected to a remote signing service
func (sc *SignerClient) IsConnected() bool { _ = "STUB: not implemented"; return false }

// WaitForConnection waits maxWait for a connection or returns a timeout error
func (sc *SignerClient) WaitForConnection(maxWait time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//--------------------------------------------------------
// Implement PrivValidator

// Ping sends a ping request to the remote signer
func (sc *SignerClient) Ping() error { _ = "STUB: not implemented"; return nil }

// GetPubKey retrieves a public key from a remote signer
// returns an error if client is not able to provide the key
func (sc *SignerClient) GetPubKey() (crypto.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey), nil
}

// SignVote requests a remote signer to sign a vote
func (sc *SignerClient) SignVote(chainID string, vote *cmtproto.Vote) error {
	_ = "STUB: not implemented"
	return nil
}

// SignProposal requests a remote signer to sign a proposal
func (sc *SignerClient) SignProposal(chainID string, proposal *cmtproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}

func (sc *SignerClient) SignRawBytes(chainID, uniqueID string, rawBytes []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
