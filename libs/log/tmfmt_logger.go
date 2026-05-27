package log

import (
	"bytes"
	"io"
	"sync"

	kitlog "github.com/go-kit/log"
	"github.com/go-logfmt/logfmt"
)

type tmfmtEncoder struct {
	*logfmt.Encoder
	buf bytes.Buffer
}

func (l *tmfmtEncoder) Reset() { _ = "STUB: not implemented"; return }

var tmfmtEncoderPool = sync.Pool{
	New: func() interface{} {
		var enc tmfmtEncoder
		enc.Encoder = logfmt.NewEncoder(&enc.buf)
		return &enc
	},
}

type tmfmtLogger struct {
	w io.Writer
}

// NewTMFmtLogger returns a logger that encodes keyvals to the Writer in
// CometBFT custom format. Note complex types (structs, maps, slices)
// formatted as "%+v".
//
// Each log event produces no more than one call to w.Write.
// The passed Writer must be safe for concurrent use by multiple goroutines if
// the returned Logger will be used concurrently.
func NewTMFmtLogger(w io.Writer) kitlog.Logger {
	_ = "STUB: not implemented"
	return *new(kitlog.Logger)
}

func (l tmfmtLogger) Log(keyvals ...interface{}) error { _ = "STUB: not implemented"; return nil }

// indexes of keys to skip while encoding later

// Extract level

//nolint:gocritic

// and message

// and module (could be multiple keyvals; if such case last keyvalue wins)

// Print []byte as a hexadecimal string (uppercased)

// Realize stringers

// Form a custom CometBFT line
//
// Example:
//     D[2016-05-02|11:06:44.322]   Stopping AddrBook (ignoring: already stopped)
//
// Description:
//     D										- first character of the level, uppercase (ASCII only)
//     [2016-05-02|11:06:44.322]    - our time format (see https://golang.org/src/time/format.go)
//     Stopping ...					- message

//nolint:errcheck // no need to check error again

// Add newline to the end of the buffer

// The Logger interface requires implementations to be safe for concurrent
// use by multiple goroutines. For this implementation that means making
// only one call to l.w.Write() for each call to Log.
