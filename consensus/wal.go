package consensus

import (
	"io"
	"time"

	auto "github.com/cometbft/cometbft/libs/autofile"
	cmtjson "github.com/cometbft/cometbft/libs/json"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
)

const (
	// time.Time + max consensus msg size
	maxMsgSizeBytes = maxMsgSize + 24

	// how often the WAL should be sync'd during period sync'ing
	walDefaultFlushInterval = 2 * time.Second
)

//--------------------------------------------------------
// types and functions for savings consensus messages

// TimedWALMessage wraps WALMessage and adds Time for debugging purposes.
type TimedWALMessage struct {
	Time time.Time  `json:"time"`
	Msg  WALMessage `json:"msg"`
}

// EndHeightMessage marks the end of the given height inside WAL.
// @internal used by scripts/wal2json util.
type EndHeightMessage struct {
	Height int64 `json:"height"`
}

type WALMessage interface{}

func init() {
	cmtjson.RegisterType(msgInfo{}, "tendermint/wal/MsgInfo")
	cmtjson.RegisterType(timeoutInfo{}, "tendermint/wal/TimeoutInfo")
	cmtjson.RegisterType(EndHeightMessage{}, "tendermint/wal/EndHeightMessage")
}

//--------------------------------------------------------
// Simple write-ahead logger

// WAL is an interface for any write-ahead logger.
type WAL interface {
	Write(WALMessage) error
	WriteSync(WALMessage) error
	FlushAndSync() error

	SearchForEndHeight(height int64, options *WALSearchOptions) (rd io.ReadCloser, found bool, err error)

	// service methods
	Start() error
	Stop() error
	Wait()
}

// Write ahead logger writes msgs to disk before they are processed.
// Can be used for crash-recovery and deterministic replay.
// TODO: currently the wal is overwritten during replay catchup, give it a mode
// so it's either reading or appending - must read to end to start appending
// again.
type BaseWAL struct {
	service.BaseService

	group *auto.Group

	enc *WALEncoder

	flushTicker   *time.Ticker
	flushInterval time.Duration
}

var _ WAL = &BaseWAL{}

// NewWAL returns a new write-ahead logger based on `baseWAL`, which implements
// WAL. It's flushed and synced to disk every 2s and once when stopped.
func NewWAL(walFile string, groupOptions ...func(*auto.Group)) (*BaseWAL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetFlushInterval allows us to override the periodic flush interval for the WAL.
func (wal *BaseWAL) SetFlushInterval(i time.Duration) { _ = "STUB: not implemented"; return }

func (wal *BaseWAL) Group() *auto.Group { _ = "STUB: not implemented"; return nil }

func (wal *BaseWAL) SetLogger(l log.Logger) { _ = "STUB: not implemented"; return }

//nolint:staticcheck

func (wal *BaseWAL) OnStart() error { _ = "STUB: not implemented"; return nil }

func (wal *BaseWAL) processFlushTicks() { _ = "STUB: not implemented"; return }

// FlushAndSync flushes and fsync's the underlying group's data to disk.
// See auto#FlushAndSync
func (wal *BaseWAL) FlushAndSync() error { _ = "STUB: not implemented"; return nil }

// Stop the underlying autofile group.
// Use Wait() to ensure it's finished shutting down
// before cleaning up files.
func (wal *BaseWAL) OnStop() { _ = "STUB: not implemented"; return }

// Wait for the underlying autofile group to finish shutting down
// so it's safe to cleanup files.
func (wal *BaseWAL) Wait() {
	_ = "STUB: not implemented"

	// Write is called in newStep and for each receive on the
	// peerMsgQueue and the timeoutTicker.
	// NOTE: does not call fsync()
	return
}

func (wal *BaseWAL) Write(msg WALMessage) error { _ = "STUB: not implemented"; return nil }

// WriteSync is called when we receive a msg from ourselves
// so that we write to disk before sending signed messages.
// NOTE: calls fsync()
func (wal *BaseWAL) WriteSync(msg WALMessage) error { _ = "STUB: not implemented"; return nil }

// WALSearchOptions are optional arguments to SearchForEndHeight.
type WALSearchOptions struct {
	// IgnoreDataCorruptionErrors set to true will result in skipping data corruption errors.
	IgnoreDataCorruptionErrors bool
}

// SearchForEndHeight searches for the EndHeightMessage with the given height
// and returns an auto.GroupReader, whenever it was found or not and an error.
// Group reader will be nil if found equals false.
//
// CONTRACT: caller must close group reader.
func (wal *BaseWAL) SearchForEndHeight(
	height int64,
	options *WALSearchOptions,
) (rd io.ReadCloser, found bool, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), false, nil
}

// NOTE: starting from the last file in the group because we're usually
// searching for the last height. See replay.go

// OPTIMISATION: no need to look for height in older files if we've seen h < height

// check next file

// do nothing

// found

// A WALEncoder writes custom-encoded WAL messages to an output stream.
//
// Format: 4 bytes CRC sum + 4 bytes length + arbitrary-length value
type WALEncoder struct {
	wr io.Writer
}

// NewWALEncoder returns a new encoder that writes to wr.
func NewWALEncoder(wr io.Writer) *WALEncoder { _ = "STUB: not implemented"; return nil }

// Encode writes the custom encoding of v to the stream. It returns an error if
// the encoded size of v is greater than 1MB. Any error encountered
// during the write is also returned.
func (enc *WALEncoder) Encode(v *TimedWALMessage) error { _ = "STUB: not implemented"; return nil }

// IsDataCorruptionError returns true if data has been corrupted inside WAL.
func IsDataCorruptionError(err error) bool { _ = "STUB: not implemented"; return false }

// DataCorruptionError is an error that occures if data on disk was corrupted.
type DataCorruptionError struct {
	cause error
}

func (e DataCorruptionError) Error() string { _ = "STUB: not implemented"; return "" }

func (e DataCorruptionError) Cause() error {
	_ = "STUB: not implemented"

	// A WALDecoder reads and decodes custom-encoded WAL messages from an input
	// stream. See WALEncoder for the format used.
	//
	// It will also compare the checksums and make sure data size is equal to the
	// length from the header. If that is not the case, error will be returned.
	return nil
}

type WALDecoder struct {
	rd io.Reader
}

// NewWALDecoder returns a new decoder that reads from rd.
func NewWALDecoder(rd io.Reader) *WALDecoder { _ = "STUB: not implemented"; return nil }

// Decode reads the next custom-encoded value from its reader and returns it.
func (dec *WALDecoder) Decode() (*TimedWALMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check checksum before decoding data

type nilWAL struct{}

var _ WAL = nilWAL{}

func (nilWAL) Write(WALMessage) error     { _ = "STUB: not implemented"; return nil }
func (nilWAL) WriteSync(WALMessage) error { _ = "STUB: not implemented"; return nil }
func (nilWAL) FlushAndSync() error        { _ = "STUB: not implemented"; return nil }
func (nilWAL) SearchForEndHeight(int64, *WALSearchOptions) (rd io.ReadCloser, found bool, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), false, nil
}

func (nilWAL) Start() error { _ = "STUB: not implemented"; return nil }
func (nilWAL) Stop() error  { _ = "STUB: not implemented"; return nil }
func (nilWAL) Wait()        { _ = "STUB: not implemented"; return }
