package types

import (
	"encoding/json"

	"github.com/cosmos/gogoproto/jsonpb"
)

const (
	CodeTypeOK uint32 = 0
)

// IsOK returns true if Code is OK.
func (r ResponseCheckTx) IsOK() bool { _ = "STUB: not implemented"; return false }

// IsErr returns true if Code is something other than OK.
func (r ResponseCheckTx) IsErr() bool { _ = "STUB: not implemented"; return false }

// IsOK returns true if Code is OK.
func (r ExecTxResult) IsOK() bool { _ = "STUB: not implemented"; return false }

// IsErr returns true if Code is something other than OK.
func (r ExecTxResult) IsErr() bool { _ = "STUB: not implemented"; return false }

// IsOK returns true if Code is OK.
func (r ResponseQuery) IsOK() bool { _ = "STUB: not implemented"; return false }

// IsErr returns true if Code is something other than OK.
func (r ResponseQuery) IsErr() bool { _ = "STUB: not implemented"; return false }

// IsAccepted returns true if Code is ACCEPT
func (r ResponseProcessProposal) IsAccepted() bool { _ = "STUB: not implemented"; return false }

// IsStatusUnknown returns true if Code is UNKNOWN
func (r ResponseProcessProposal) IsStatusUnknown() bool { _ = "STUB: not implemented"; return false }

func (r ResponseVerifyVoteExtension) IsAccepted() bool { _ = "STUB: not implemented"; return false }

// IsStatusUnknown returns true if Code is Unknown
func (r ResponseVerifyVoteExtension) IsStatusUnknown() bool {
	_ = "STUB: not implemented"
	return false
}

//---------------------------------------------------------------------------
// override JSON marshaling so we emit defaults (ie. disable omitempty)

var (
	jsonpbMarshaller = jsonpb.Marshaler{
		EnumsAsInts:  true,
		EmitDefaults: true,
	}
	jsonpbUnmarshaller = jsonpb.Unmarshaler{}
)

func (r *ResponseCheckTx) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *ResponseCheckTx) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (r *ExecTxResult) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *ExecTxResult) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (r *ResponseQuery) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *ResponseQuery) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (r *ResponseCommit) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *ResponseCommit) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (r *EventAttribute) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON was modified in the Celestia fork to be backwards compatible
// with the EventAttribute from CometBFT v0.34.x. CometBFT v0.38.x uses the type
// string for keys and values. CometBFT v0.34.x used the type bytes for keys and
// values. CometBFT v0.34.x event attributes that were marshaled to JSON
// previously encoded the keys and values as base64 strings so this method
// attempts to base64 decode the keys and values if they look like base64
// encoded data.
func (r *EventAttribute) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func maybeBase64Decode(input string) string { _ = "STUB: not implemented"; return "" }

// input is not a base64 encoded string so return the input

// isLikelyBase64Encoded returns true if input is likely a base64 encoded string.
func isLikelyBase64Encoded(input string) bool { _ = "STUB: not implemented"; return false }

// Only decode if the result is printable ASCII/UTF-8 text

// Allow printable ASCII characters, spaces, and common unicode

// Reject high-value bytes that are likely binary garbage

// Additional heuristic: old format base64 was typically longer
// and had padding or specific characteristics

// Short strings that happen to be valid base64 are probably just normal strings

// Some compile time assertions to ensure we don't
// have accidental runtime surprises later on.

// jsonEncodingRoundTripper ensures that asserted
// interfaces implement both MarshalJSON and UnmarshalJSON
type jsonRoundTripper interface {
	json.Marshaler
	json.Unmarshaler
}

var _ jsonRoundTripper = (*ResponseCommit)(nil)
var _ jsonRoundTripper = (*ResponseQuery)(nil)
var _ jsonRoundTripper = (*ExecTxResult)(nil)
var _ jsonRoundTripper = (*ResponseCheckTx)(nil)

var _ jsonRoundTripper = (*EventAttribute)(nil)

// deterministicExecTxResult constructs a copy of response that omits
// non-deterministic fields. The input response is not modified.
func deterministicExecTxResult(response *ExecTxResult) *ExecTxResult {
	_ = "STUB: not implemented"
	return nil
}

// MarshalTxResults encodes the the TxResults as a list of byte
// slices. It strips off the non-deterministic pieces of the TxResults
// so that the resulting data can be used for hash comparisons and used
// in Merkle proofs.
func MarshalTxResults(r []*ExecTxResult) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// -----------------------------------------------
// construct Result data
