package mock

import (
	"context"

	"github.com/cometbft/cometbft/rpc/client"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
)

// StatusMock returns the result specified by the Call
type StatusMock struct {
	Call
}

var (
	_ client.StatusClient = (*StatusMock)(nil)
	_ client.StatusClient = (*StatusRecorder)(nil)
)

func (m *StatusMock) Status(context.Context) (*ctypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StatusRecorder can wrap another type (StatusMock, full client)
// and record the status calls
type StatusRecorder struct {
	Client client.StatusClient
	Calls  []Call
}

func NewStatusRecorder(client client.StatusClient) *StatusRecorder {
	_ = "STUB: not implemented"
	return nil
}

func (r *StatusRecorder) addCall(call Call) { _ = "STUB: not implemented"; return }

func (r *StatusRecorder) Status(ctx context.Context) (*ctypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
