package types

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/crypto/merkle"
)

// ABCIResults wraps the deliver tx results to return a proof.
type ABCIResults []*abci.ExecTxResult

// NewResults strips non-deterministic fields from ExecTxResult responses
// and returns ABCIResults.
func NewResults(responses []*abci.ExecTxResult) ABCIResults {
	_ = "STUB: not implemented"
	return *new(ABCIResults)
}

// Hash returns a merkle hash of all results.
func (a ABCIResults) Hash() []byte { _ = "STUB: not implemented"; return nil }

// ProveResult returns a merkle proof of one result from the set
func (a ABCIResults) ProveResult(i int) merkle.Proof {
	_ = "STUB: not implemented"
	return *new(merkle.Proof)
}

func (a ABCIResults) toByteSlices() [][]byte { _ = "STUB: not implemented"; return nil }

// deterministicExecTxResult strips non-deterministic fields from
// ExecTxResult and returns another ExecTxResult.
func deterministicExecTxResult(response *abci.ExecTxResult) *abci.ExecTxResult {
	_ = "STUB: not implemented"
	return nil
}
