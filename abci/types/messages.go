package types

import (
	"io"
	"math"

	"github.com/cosmos/gogoproto/proto"
)

const (
	maxMsgSize = math.MaxInt32 // 2GB
)

// WriteMessage writes a varint length-delimited protobuf message.
func WriteMessage(msg proto.Message, w io.Writer) error { _ = "STUB: not implemented"; return nil }

// ReadMessage reads a varint length-delimited protobuf message.
func ReadMessage(r io.Reader, msg proto.Message) error { _ = "STUB: not implemented"; return nil }

//----------------------------------------

func ToRequestEcho(message string) *Request { _ = "STUB: not implemented"; return nil }

func ToRequestFlush() *Request { _ = "STUB: not implemented"; return nil }

func ToRequestInfo(req *RequestInfo) *Request { _ = "STUB: not implemented"; return nil }

func ToRequestCheckTx(req *RequestCheckTx) *Request { _ = "STUB: not implemented"; return nil }

func ToRequestCommit() *Request { _ = "STUB: not implemented"; return nil }

func ToRequestQuery(req *RequestQuery) *Request { _ = "STUB: not implemented"; return nil }

func ToRequestInitChain(req *RequestInitChain) *Request { _ = "STUB: not implemented"; return nil }

func ToRequestListSnapshots(req *RequestListSnapshots) *Request {
	_ = "STUB: not implemented"
	return nil
}

func ToRequestOfferSnapshot(req *RequestOfferSnapshot) *Request {
	_ = "STUB: not implemented"
	return nil
}

func ToRequestLoadSnapshotChunk(req *RequestLoadSnapshotChunk) *Request {
	_ = "STUB: not implemented"
	return nil
}

func ToRequestApplySnapshotChunk(req *RequestApplySnapshotChunk) *Request {
	_ = "STUB: not implemented"
	return nil
}

func ToRequestPrepareProposal(req *RequestPrepareProposal) *Request {
	_ = "STUB: not implemented"
	return nil
}

func ToRequestProcessProposal(req *RequestProcessProposal) *Request {
	_ = "STUB: not implemented"
	return nil
}

func ToRequestExtendVote(req *RequestExtendVote) *Request { _ = "STUB: not implemented"; return nil }

func ToRequestVerifyVoteExtension(req *RequestVerifyVoteExtension) *Request {
	_ = "STUB: not implemented"
	return nil
}

func ToRequestFinalizeBlock(req *RequestFinalizeBlock) *Request {
	_ = "STUB: not implemented"
	return nil
}

func ToRequestQuerySequence(req *RequestQuerySequence) *Request {
	_ = "STUB: not implemented"
	return nil
}

//----------------------------------------

func ToResponseException(errStr string) *Response { _ = "STUB: not implemented"; return nil }

func ToResponseEcho(message string) *Response { _ = "STUB: not implemented"; return nil }

func ToResponseFlush() *Response { _ = "STUB: not implemented"; return nil }

func ToResponseInfo(res *ResponseInfo) *Response { _ = "STUB: not implemented"; return nil }

func ToResponseCheckTx(res *ResponseCheckTx) *Response { _ = "STUB: not implemented"; return nil }

func ToResponseCommit(res *ResponseCommit) *Response { _ = "STUB: not implemented"; return nil }

func ToResponseQuery(res *ResponseQuery) *Response { _ = "STUB: not implemented"; return nil }

func ToResponseInitChain(res *ResponseInitChain) *Response { _ = "STUB: not implemented"; return nil }

func ToResponseListSnapshots(res *ResponseListSnapshots) *Response {
	_ = "STUB: not implemented"
	return nil
}

func ToResponseOfferSnapshot(res *ResponseOfferSnapshot) *Response {
	_ = "STUB: not implemented"
	return nil
}

func ToResponseLoadSnapshotChunk(res *ResponseLoadSnapshotChunk) *Response {
	_ = "STUB: not implemented"
	return nil
}

func ToResponseApplySnapshotChunk(res *ResponseApplySnapshotChunk) *Response {
	_ = "STUB: not implemented"
	return nil
}

func ToResponsePrepareProposal(res *ResponsePrepareProposal) *Response {
	_ = "STUB: not implemented"
	return nil
}

func ToResponseProcessProposal(res *ResponseProcessProposal) *Response {
	_ = "STUB: not implemented"
	return nil
}

func ToResponseExtendVote(res *ResponseExtendVote) *Response { _ = "STUB: not implemented"; return nil }

func ToResponseVerifyVoteExtension(res *ResponseVerifyVoteExtension) *Response {
	_ = "STUB: not implemented"
	return nil
}

func ToResponseFinalizeBlock(res *ResponseFinalizeBlock) *Response {
	_ = "STUB: not implemented"
	return nil
}

func ToResponseQuerySequence(res *ResponseQuerySequence) *Response {
	_ = "STUB: not implemented"
	return nil
}
