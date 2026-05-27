package consensus

import (
	"io"
	"testing"
	"time"

	cfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/log"
)

// WALGenerateNBlocks generates a consensus WAL. It does this by spinning up a
// stripped down version of node (proxy app, event bus, consensus state) with a
// persistent kvstore application and special consensus wal instance
// (byteBufferWAL) and waits until numBlocks are created.
// If the node fails to produce given numBlocks, it returns an error.
func WALGenerateNBlocks(t *testing.T, wr io.Writer, numBlocks int, config *cfg.Config) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// COPY PASTE FROM node.go WITH A FEW MODIFICATIONS
// NOTE: we can't import node package because of circular dependency.
// NOTE: we don't do handshake so need to set state.Version.Consensus.App directly.

// END OF COPY PASTE

// set consensus wal to buffered WAL, which will write all incoming msgs to buffer

// see wal.go#103

// WALWithNBlocks returns a WAL content with numBlocks.
func WALWithNBlocks(t *testing.T, numBlocks int, config *cfg.Config) (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func randPort() int {
	_ = "STUB: not implemented"
	// returns between base and base + spread
	return 0
}

func makeAddrs() (string, string, string) { _ = "STUB: not implemented"; return "", "", "" }

// getConfig returns a config for test cases
func getConfig(t *testing.T) *cfg.Config { _ = "STUB: not implemented"; return nil }

// and we use random ports to run in parallel

// byteBufferWAL is a WAL which writes all msgs to a byte buffer. Writing stops
// when the heightToStop is reached. Client will be notified via
// signalWhenStopsTo channel.
type byteBufferWAL struct {
	enc               *WALEncoder
	stopped           bool
	heightToStop      int64
	signalWhenStopsTo chan<- struct{}

	logger log.Logger
}

// needed for determinism
var fixedTime, _ = time.Parse(time.RFC3339, "2017-01-02T15:04:05Z")

func newByteBufferWAL(logger log.Logger, enc *WALEncoder, nBlocks int64, signalStop chan<- struct{}) *byteBufferWAL {
	_ = "STUB: not implemented"
	return nil
}

// Save writes message to the internal buffer except when heightToStop is
// reached, in which case it will signal the caller via signalWhenStopsTo and
// skip writing.
func (w *byteBufferWAL) Write(m WALMessage) error { _ = "STUB: not implemented"; return nil }

func (w *byteBufferWAL) WriteSync(m WALMessage) error { _ = "STUB: not implemented"; return nil }

func (w *byteBufferWAL) FlushAndSync() error { _ = "STUB: not implemented"; return nil }

func (w *byteBufferWAL) SearchForEndHeight(
	int64,
	*WALSearchOptions,
) (rd io.ReadCloser, found bool, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), false, nil
}

func (w *byteBufferWAL) Start() error { _ = "STUB: not implemented"; return nil }
func (w *byteBufferWAL) Stop() error  { _ = "STUB: not implemented"; return nil }
func (w *byteBufferWAL) Wait()        { _ = "STUB: not implemented"; return }
