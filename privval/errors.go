package privval

import (
	"errors"
)

// EndpointTimeoutError occurs when endpoint times out.
type EndpointTimeoutError struct{}

// Implement the net.Error interface.
func (e EndpointTimeoutError) Error() string { _ = "STUB: not implemented"; return "" }
func (e EndpointTimeoutError) Timeout() bool { _ = "STUB: not implemented"; return false }
func (e EndpointTimeoutError) Temporary() bool {
	_ = "STUB: not implemented"

	// Socket errors.
	return false
}

var (
	ErrConnectionTimeout  = EndpointTimeoutError{}
	ErrNoConnection       = errors.New("endpoint is not connected")
	ErrReadTimeout        = errors.New("endpoint read timed out")
	ErrUnexpectedResponse = errors.New("empty response")
	ErrWriteTimeout       = errors.New("endpoint write timed out")
)

// RemoteSignerError allows (remote) validators to include meaningful error
// descriptions in their reply.
type RemoteSignerError struct {
	// TODO(ismail): create an enum of known errors
	Code        int
	Description string
}

func (e *RemoteSignerError) Error() string { _ = "STUB: not implemented"; return "" }
