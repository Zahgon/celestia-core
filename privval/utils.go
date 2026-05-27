package privval

import (
	"github.com/cometbft/cometbft/libs/log"
)

// IsConnTimeout returns a boolean indicating whether the error is known to
// report that a connection timeout occurred. This detects both fundamental
// network timeouts, as well as ErrConnTimeout errors.
func IsConnTimeout(err error) bool { _ = "STUB: not implemented"; return false }

// NewSignerListener creates a new SignerListenerEndpoint using the corresponding listen address
func NewSignerListener(listenAddr string, logger log.Logger) (*SignerListenerEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: persist this key so external signer can actually authenticate us

// GetFreeLocalhostAddrPort returns a free localhost:port address
func GetFreeLocalhostAddrPort() string { _ = "STUB: not implemented"; return "" }
