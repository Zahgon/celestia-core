package client

import (
	types "github.com/cometbft/cometbft/rpc/jsonrpc/types"
)

func unmarshalResponseBytes(
	responseBytes []byte,
	expectedID types.JSONRPCIntID,
	result interface{},
) (interface{}, error) {
	_ = "STUB: not implemented"

	// Read response.  If rpc/core/types is imported, the result will unmarshal
	// into the correct type.
	return nil, nil
}

// Unmarshal the RawMessage into the result.

// Separate the unmarshalling actions using different functions to improve readability and maintainability.
func unmarshalIndividualResponse(responseBytes []byte) (types.RPCResponse, error) {
	_ = "STUB: not implemented"
	return *new(types.RPCResponse), nil
}

func unmarshalMultipleResponses(responseBytes []byte) ([]types.RPCResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalResponseBytesArray(
	responseBytes []byte,
	expectedIDs []types.JSONRPCIntID,
	results []interface{},
) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to unmarshal as multiple responses

// if err == nil it could unmarshal in multiple responses

// No response error checking here as there may be a mixture of successful
// and unsuccessful responses.

// Intersect IDs from responses with expectedIDs.

// check if it's a single response that should be an error

// Here, an error means that even single response unmarshalling failed,
// so return the error.

func validateResponseIDs(ids, expectedIDs []types.JSONRPCIntID) error {
	_ = "STUB: not implemented"
	return nil
}

// From the JSON-RPC 2.0 spec:
// id: It MUST be the same as the value of the id member in the Request Object.
func validateAndVerifyID(res *types.RPCResponse, expectedID types.JSONRPCIntID) error {
	_ = "STUB: not implemented"
	return nil
}

// validateResponseID ensured res.ID has the right type

func validateResponseID(id interface{}) error { _ = "STUB: not implemented"; return nil }
