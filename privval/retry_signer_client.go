package privval

import (
	"time"

	"github.com/cometbft/cometbft/crypto"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cometbft/cometbft/types"
)

// RetrySignerClient wraps SignerClient adding retry for each operation (except
// Ping) w/ a timeout.
type RetrySignerClient struct {
	next    *SignerClient
	retries int
	timeout time.Duration
}

// NewRetrySignerClient returns RetrySignerClient. If +retries+ is 0, the
// client will be retrying each operation indefinitely.
func NewRetrySignerClient(sc *SignerClient, retries int, timeout time.Duration) *RetrySignerClient {
	_ = "STUB: not implemented"
	return nil
}

var _ types.PrivValidator = (*RetrySignerClient)(nil)

func (sc *RetrySignerClient) Close() error { _ = "STUB: not implemented"; return nil }

func (sc *RetrySignerClient) IsConnected() bool { _ = "STUB: not implemented"; return false }

func (sc *RetrySignerClient) WaitForConnection(maxWait time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//--------------------------------------------------------
// Implement PrivValidator

func (sc *RetrySignerClient) Ping() error { _ = "STUB: not implemented"; return nil }

func (sc *RetrySignerClient) GetPubKey() (crypto.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey), nil
}

// If remote signer errors, we don't retry.

func (sc *RetrySignerClient) SignVote(chainID string, vote *cmtproto.Vote) error {
	_ = "STUB: not implemented"
	return nil
}

// If remote signer errors, we don't retry.

func (sc *RetrySignerClient) SignProposal(chainID string, proposal *cmtproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}

// If remote signer errors, we don't retry.

func (sc *RetrySignerClient) SignRawBytes(chainID, uniqueID string, rawBytes []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If remote signer errors, we don't retry.
