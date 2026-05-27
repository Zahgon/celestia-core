// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Modified from original GoGo Protobuf to not buffer the reader.

package protoio

import (
	"io"

	"github.com/cosmos/gogoproto/proto"
)

// NewDelimitedReader reads varint-delimited Protobuf messages from a reader.
// Unlike the gogoproto NewDelimitedReader, this does not buffer the reader,
// which may cause poor performance but is necessary when only reading single
// messages (e.g. in the p2p package). It also returns the number of bytes
// read, which is necessary for the p2p package.
func NewDelimitedReader(r io.Reader, maxSize int) ReadCloser {
	_ = "STUB: not implemented"
	return *new(ReadCloser)
}

type varintReader struct {
	r io.Reader
	// ReadUvarint needs an io.ByteReader, and we also need to keep track of the
	// number of bytes read, so we use our own byteReader. This can't be
	// buffered, so the caller should pass a buffered io.Reader to avoid poor
	// performance.
	byteReader *byteReader
	buf        []byte
	maxSize    int
	closer     io.Closer
}

func (r *varintReader) ReadMsg(msg proto.Message) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Make sure length doesn't overflow the native int size (e.g. 32-bit),
// and that the returned sum of n+length doesn't overflow either.

func (r *varintReader) Close() error { _ = "STUB: not implemented"; return nil }

func UnmarshalDelimited(data []byte, msg proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}
