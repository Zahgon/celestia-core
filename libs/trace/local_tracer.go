package trace

import (
	"os"
	"time"

	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/log"
)

const (
	PushBucketName = "TRACE_PUSH_BUCKET_NAME"
	PushRegion     = "TRACE_PUSH_REGION"
	PushAccessKey  = "TRACE_PUSH_ACCESS_KEY"
	PushKey        = "TRACE_PUSH_SECRET_KEY"
	PushDelay      = "TRACE_PUSH_DELAY"
)

// Event wraps some trace data with metadata that dictates the table and things
// like the chainID and nodeID.
type Event[T any] struct {
	ChainID   string    `json:"chain_id"`
	NodeID    string    `json:"node_id"`
	Table     string    `json:"table"`
	Timestamp time.Time `json:"timestamp"`
	Msg       T         `json:"msg"`
}

// NewEvent creates a new Event with the given chainID, nodeID, table, and msg.
// It adds the current time as the timestamp.
func NewEvent[T any](chainID, nodeID, table string, msg T) Event[T] {
	_ = "STUB: not implemented"
	return nil
}

// LocalTracer saves all of the events passed to the retuen channel to files
// based on their "type" (a string field in the event). Each type gets its own
// file. The internals are purposefully not *explicitly* thread safe to avoid the
// overhead of locking with each event save. Only pass events to the returned
// channel. Call CloseAll to close all open files.
type LocalTracer struct {
	chainID, nodeID string
	logger          log.Logger
	cfg             *config.Config
	s3Config        S3Config

	// fileMap maps tables to their open files files are threadsafe, but the map
	// is not. Therefore don't create new files after initialization to remain
	// threadsafe.
	fileMap map[string]*bufferedFile
	// canal is a channel for all events that are being written. It acts as an
	// extra buffer to avoid blocking the caller when writing to files.
	canal chan Event[Entry]
}

// NewLocalTracer creates a struct that will save all of the events passed to
// the retuen channel to files based on their "table" (a string field in the
// event). Each type gets its own file. The internal are purposefully not thread
// safe to avoid the overhead of locking with each event save. Only pass events
// to the returned channel. Call CloseAll to close all open files. Goroutine to
// save events is started in this function.
func NewLocalTracer(cfg *config.Config, logger log.Logger, chainID, nodeID string) (*LocalTracer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPushConfigFromEnv reads the required environment variables to push trace.
func GetPushConfigFromEnv() (S3Config, error) {
	_ = "STUB: not implemented"
	return *new(S3Config), nil
}

func (lt *LocalTracer) Write(e Entry) { _ = "STUB: not implemented"; return }

// ReadTable returns a file for the given table. If the table is not being
// collected, an error is returned. The caller should not close the file.
func (lt *LocalTracer) readTable(table string) (*os.File, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (lt *LocalTracer) IsCollecting(table string) bool { _ = "STUB: not implemented"; return false }

// getFile gets a file for the given type. This method is purposely
// not thread-safe to avoid the overhead of locking with each event save.
func (lt *LocalTracer) getFile(table string) (*bufferedFile, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// saveEventToFile marshals an Event into JSON and appends it to a file named after the event's Type.
func (lt *LocalTracer) saveEventToFile(event Event[Entry]) error {
	_ = "STUB: not implemented"
	return nil
}

// draincanal takes a variadic number of channels of Event pointers and drains them into files.
func (lt *LocalTracer) drainCanal() {
	_ = "STUB: not implemented"
	// purposefully do not lock, and rely on the channel to provide sync
	// actions, to avoid overhead of locking with each event save.
	return
}

// Stop optionally uploads and closes all open files.
func (lt *LocalTracer) Stop() { _ = "STUB: not implemented"; return }

// splitAndTrimEmpty slices s into all subslices separated by sep and returns a
// slice of the string s with all leading and trailing Unicode code points
// contained in cutset removed. If sep is empty, SplitAndTrim splits after each
// UTF-8 sequence. First part is equivalent to strings.SplitN with a count of
// -1.  also filter out empty strings, only return non-empty strings.
//
// NOTE: this is copy pasted from the config package to avoid a circular
// dependency. See the function of the same name for tests.
func splitAndTrimEmpty(s, sep, cutset string) []string { _ = "STUB: not implemented"; return nil }
