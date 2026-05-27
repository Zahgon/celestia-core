package conn

import (
	"bufio"
	"net"
	"time"

	"github.com/cosmos/gogoproto/proto"

	"github.com/cometbft/cometbft/config"
	flow "github.com/cometbft/cometbft/libs/flowrate"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/protoio"
	"github.com/cometbft/cometbft/libs/service"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/libs/timer"
	tmp2p "github.com/cometbft/cometbft/proto/tendermint/p2p"
)

const (
	defaultMaxPacketMsgPayloadSize = 1024

	numBatchPacketMsgs = 10
	minReadBufferSize  = 1024
	minWriteBufferSize = 65536
	updateStats        = 2 * time.Second

	// some of these defaults are written in the user config
	// flushThrottle, sendRate, recvRate
	// TODO: remove values present in config
	defaultFlushThrottle = 100 * time.Millisecond

	defaultSendQueueCapacity   = 1
	defaultRecvBufferCapacity  = 4096
	defaultRecvMessageCapacity = 22020096      // 21MB
	defaultSendRate            = int64(512000) // 500KB/s
	defaultRecvRate            = int64(512000) // 500KB/s
	defaultSendTimeout         = 10 * time.Second
	defaultPingInterval        = 60 * time.Second
	defaultPongTimeout         = 45 * time.Second
)

type (
	receiveCbFunc func(chID byte, msgBytes []byte)
	errorCbFunc   func(interface{})
)

/*
Each peer has one `MConnection` (multiplex connection) instance.

__multiplex__ *noun* a system or signal involving simultaneous transmission of
several messages along a single channel of communication.

Each `MConnection` handles message transmission on multiple abstract communication
`Channel`s.  Each channel has a globally unique byte id.
The byte id and the relative priorities of each `Channel` are configured upon
initialization of the connection.

There are two methods for sending messages:

	func (m MConnection) Send(chID byte, msgBytes []byte) bool {}
	func (m MConnection) TrySend(chID byte, msgBytes []byte}) bool {}

`Send(chID, msgBytes)` is a blocking call that waits until `msg` is
successfully queued for the channel with the given id byte `chID`, or until the
request times out.  The message `msg` is serialized using Protobuf.

`TrySend(chID, msgBytes)` is a nonblocking call that returns false if the
channel's queue is full.

Inbound message bytes are handled with an onReceive callback function.
*/
type MConnection struct {
	service.BaseService

	conn          net.Conn
	bufConnReader *bufio.Reader
	bufConnWriter *bufio.Writer
	sendMonitor   *flow.Monitor
	recvMonitor   *flow.Monitor
	send          chan struct{}
	pong          chan struct{}
	channels      []*Channel
	channelsIdx   map[byte]*Channel
	onReceive     receiveCbFunc
	onError       errorCbFunc
	errored       uint32
	config        MConnConfig

	// Closing quitSendRoutine will cause the sendRoutine to eventually quit.
	// doneSendRoutine is closed when the sendRoutine actually quits.
	quitSendRoutine chan struct{}
	doneSendRoutine chan struct{}

	// Closing quitRecvRouting will cause the recvRouting to eventually quit.
	quitRecvRoutine chan struct{}

	// used to ensure FlushStop and OnStop
	// are safe to call concurrently.
	stopMtx cmtsync.Mutex

	flushTimer *timer.ThrottleTimer // flush writes as necessary but throttled.
	pingTimer  *time.Ticker         // send pings periodically

	// close conn if pong is not received in pongTimeout
	pongTimer     *time.Timer
	pongTimeoutCh chan bool // true - timeout, false - peer sent pong

	chStatsTimer *time.Ticker // update channel stats periodically

	created time.Time // time of creation

	_maxPacketMsgSize int
}

// MConnConfig is a MConnection configuration.
type MConnConfig struct {
	SendRate int64 `mapstructure:"send_rate"`
	RecvRate int64 `mapstructure:"recv_rate"`

	// Maximum payload size
	MaxPacketMsgPayloadSize int `mapstructure:"max_packet_msg_payload_size"`

	// Interval to flush writes (throttled)
	FlushThrottle time.Duration `mapstructure:"flush_throttle"`

	// Interval to send pings
	PingInterval time.Duration `mapstructure:"ping_interval"`

	// Maximum wait time for pongs
	PongTimeout time.Duration `mapstructure:"pong_timeout"`

	// Fuzz connection
	TestFuzz       bool                   `mapstructure:"test_fuzz"`
	TestFuzzConfig *config.FuzzConnConfig `mapstructure:"test_fuzz_config"`
}

// DefaultMConnConfig returns the default config.
func DefaultMConnConfig() MConnConfig { _ = "STUB: not implemented"; return *new(MConnConfig) }

// NewMConnection wraps net.Conn and creates multiplex connection
func NewMConnection(
	conn net.Conn,
	chDescs []*ChannelDescriptor,
	onReceive receiveCbFunc,
	onError errorCbFunc,
) *MConnection {
	_ = "STUB: not implemented"
	return nil
}

// NewMConnectionWithConfig wraps net.Conn and creates multiplex connection with a config
func NewMConnectionWithConfig(
	conn net.Conn,
	chDescs []*ChannelDescriptor,
	onReceive receiveCbFunc,
	onError errorCbFunc,
	config MConnConfig,
) *MConnection {
	_ = "STUB: not implemented"
	return nil
}

// Create channels

//nolint:prealloc

// maxPacketMsgSize() is a bit heavy, so call just once

func (c *MConnection) SetLogger(l log.Logger) { _ = "STUB: not implemented"; return }

// OnStart implements BaseService
func (c *MConnection) OnStart() error { _ = "STUB: not implemented"; return nil }

// stopServices stops the BaseService and timers and closes the quitSendRoutine.
// if the quitSendRoutine was already closed, it returns true, otherwise it returns false.
// It uses the stopMtx to ensure only one of FlushStop and OnStop can do this at a time.
func (c *MConnection) stopServices() (alreadyStopped bool) { _ = "STUB: not implemented"; return false }

// already quit

// already quit

// inform the recvRouting that we are shutting down

// FlushStop replicates the logic of OnStop.
// It additionally ensures that all successful
// .Send() calls will get flushed before closing
// the connection.
func (c *MConnection) FlushStop() { _ = "STUB: not implemented"; return }

// this block is unique to FlushStop

// wait until the sendRoutine exits
// so we dont race on calling sendSomePacketMsgs

// Send and flush all pending msgs.
// Since sendRoutine has exited, we can call this
// safely

// Now we can close the connection

// We can't close pong safely here because
// recvRoutine may write to it after we've stopped.
// Though it doesn't need to get closed at all,
// we close it @ recvRoutine.

// c.Stop()

// OnStop implements BaseService
func (c *MConnection) OnStop() { _ = "STUB: not implemented"; return }

// We can't close pong safely here because
// recvRoutine may write to it after we've stopped.
// Though it doesn't need to get closed at all,
// we close it @ recvRoutine.

func (c *MConnection) String() string { _ = "STUB: not implemented"; return "" }

func (c *MConnection) flush() {
	_ = "STUB: not implemented"
	// c.Logger.Debug("Flush", "conn", c)
	return
}

// Catch panics, usually caused by remote disconnects.
func (c *MConnection) _recover() { _ = "STUB: not implemented"; return }

func (c *MConnection) stopForError(r interface{}) { _ = "STUB: not implemented"; return }

// Queues a message to be sent to channel.
func (c *MConnection) Send(chID byte, msgBytes []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// Send message to channel.

// Wake up sendRoutine if necessary

// Queues a message to be sent to channel.
// Nonblocking, returns true if successful.
func (c *MConnection) TrySend(chID byte, msgBytes []byte) bool {
	_ = "STUB: not implemented"
	return false
}

//c.Logger.Debug("TrySend", "channel", chID, "conn", c, "msgBytes", log.NewLazySprintf("%X", msgBytes))

// Send message to channel.

// Wake up sendRoutine if necessary

// CanSend returns true if you can send more data onto the chID, false
// otherwise. Use only as a heuristic.
func (c *MConnection) CanSend(chID byte) bool { _ = "STUB: not implemented"; return false }

// sendRoutine polls for packets to send from channels.
func (c *MConnection) sendRoutine() { _ = "STUB: not implemented"; return }

// NOTE: flushTimer.Set() must be called every time
// something is written to .bufConnWriter.

// Send some PacketMsgs

// Keep sendRoutine awake.

// Cleanup

// Returns true if messages from channels were exhausted.
// Blocks in accordance to .sendMonitor throttling.
func (c *MConnection) sendSomePacketMsgs(w protoio.Writer) bool {
	_ = "STUB: not implemented"
	// Block until .sendMonitor says we can write.
	// Once we're ready we send more than we asked for,
	// but amortized it should even out.
	return false
}

// Now send some PacketMsgs.

// Returns true if messages from channels were exhausted.
func (c *MConnection) sendBatchPacketMsgs(w protoio.Writer, batchSize int) bool {
	_ = "STUB: not implemented"
	// Send a batch of PacketMsgs.
	return false
}

// nothing to send across any channel.

// selects a channel to gossip our next message on.
// TODO: Make "batchChannelToGossipOn", so we can do our proto marshaling overheads in parallel,
// and we can avoid re-checking for `isSendPending`.
// We can easily mock the recentlySent differences for the batch choosing.
func selectChannelToGossipOn(channels []*Channel) *Channel {
	_ = "STUB: not implemented"
	// Choose a channel to create a PacketMsg from.
	// The chosen channel will be the one whose recentlySent/priority is the least.
	return nil
}

// If nothing to send, skip this channel
// TODO: Skip continually looking for isSendPending on channels we've already skipped in this batch-send.

// Get ratio, and keep track of lowest ratio.
// TODO: RecentlySent right now is bytes. This should be refactored to num messages to fix
// gossip prioritization bugs.

// returns (num_bytes_written, error_occurred).
func (c *MConnection) sendPacketMsgOnChannel(w protoio.Writer, sendChannel *Channel) (int, bool) {
	_ = "STUB: not implemented"
	// Make & send a PacketMsg from this channel
	return 0, false
}

// TODO: Change this to only add flush signals at the start and end of the batch.

// recvRoutine reads PacketMsgs and reconstructs the message using the channels' "recving" buffer.
// After a whole message has been assembled, it's pushed to onReceive().
// Blocks depending on how the connection is throttled.
// Otherwise, it never blocks.
func (c *MConnection) recvRoutine() { _ = "STUB: not implemented"; return }

// Block until .recvMonitor says we can read.

// Peek into bufConnReader for debugging
/*
	if numBytes := c.bufConnReader.Buffered(); numBytes > 0 {
		bz, err := c.bufConnReader.Peek(cmtmath.MinInt(numBytes, 100))
		if err == nil {
			// return
		} else {
			c.Logger.Debug("Error peeking connection buffer", "err", err)
			// return nil
		}
		c.Logger.Info("Peek connection buffer", "numBytes", numBytes, "bz", bz)
	}
*/

// Read packet type

// stopServices was invoked and we are shutting down
// receiving is excpected to fail since we will close the connection

// Read more depending on packet type.

// TODO: prevent abuse, as they cause flush()'s.
// https://github.com/tendermint/tendermint/issues/1190

// never block

// never block

//c.Logger.Debug("Received bytes", "chID", channelID, "msgBytes", msgBytes)
// NOTE: This means the reactor.Receive runs in the same thread as the p2p recv routine

// Cleanup

//nolint:revive

// Drain

// not goroutine-safe
func (c *MConnection) stopPongTimer() { _ = "STUB: not implemented"; return }

// maxPacketMsgSize returns a maximum size of PacketMsg
func (c *MConnection) maxPacketMsgSize() int { _ = "STUB: not implemented"; return 0 }

type ConnectionStatus struct {
	Duration    time.Duration
	SendMonitor flow.Status
	RecvMonitor flow.Status
	Channels    []ChannelStatus
}

type ChannelStatus struct {
	ID                byte
	SendQueueCapacity int
	SendQueueSize     int
	Priority          int
	RecentlySent      int64
}

func (c *MConnection) Status() ConnectionStatus {
	_ = "STUB: not implemented"
	return *new(ConnectionStatus)
}

//-----------------------------------------------------------------------------

type ChannelDescriptor struct {
	ID                  byte
	Priority            int
	SendQueueCapacity   int
	RecvBufferCapacity  int
	RecvMessageCapacity int
	MessageType         proto.Message
}

func (chDesc ChannelDescriptor) FillDefaults() (filled ChannelDescriptor) {
	_ = "STUB: not implemented"
	return *new(ChannelDescriptor)
}

// TODO: lowercase.
// NOTE: not goroutine-safe.
type Channel struct {
	conn          *MConnection
	desc          ChannelDescriptor
	sendQueue     chan []byte
	sendQueueSize int32 // atomic.
	recving       []byte
	sending       []byte
	recentlySent  int64 // exponential moving average

	maxPacketMsgPayloadSize int

	Logger log.Logger
}

func newChannel(conn *MConnection, desc ChannelDescriptor) *Channel {
	_ = "STUB: not implemented"
	return nil
}

func (ch *Channel) SetLogger(l log.Logger) {
	_ = "STUB: not implemented"

	// Queues message to send to this channel.
	// Goroutine-safe
	// Times out (and returns false) after defaultSendTimeout
	return
}

func (ch *Channel) sendBytes(bytes []byte) bool { _ = "STUB: not implemented"; return false }

// Queues message to send to this channel.
// Nonblocking, returns true if successful.
// Goroutine-safe
func (ch *Channel) trySendBytes(bytes []byte) bool { _ = "STUB: not implemented"; return false }

// Goroutine-safe
func (ch *Channel) loadSendQueueSize() (size int) { _ = "STUB: not implemented"; return 0 }

// Goroutine-safe
// Use only as a heuristic.
func (ch *Channel) canSend() bool { _ = "STUB: not implemented"; return false }

// Returns true if any PacketMsgs are pending to be sent.
// Call before calling nextPacketMsg()
// Goroutine-safe
func (ch *Channel) isSendPending() bool { _ = "STUB: not implemented"; return false }

// Creates a new PacketMsg to send.
// Not goroutine-safe
func (ch *Channel) nextPacketMsg() tmp2p.PacketMsg {
	_ = "STUB: not implemented"
	return *new(tmp2p.PacketMsg)
}

// decrement sendQueueSize

// Writes next PacketMsg to w and updates c.recentlySent.
// Not goroutine-safe.
func (ch *Channel) writePacketMsgTo(w protoio.Writer) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Handles incoming PacketMsgs. It returns a message bytes if message is
// complete. NOTE message bytes may change on next call to recvPacketMsg.
// Not goroutine-safe
func (ch *Channel) recvPacketMsg(packet tmp2p.PacketMsg) ([]byte, error) {
	_ = "STUB: not implemented"
	//ch.Logger.Debug("Read PacketMsg", "conn", ch.conn, "packet", packet)
	return nil, nil
}

// Reset the receive buffer to the baseline capacity so that large
// message allocations are not retained for the lifetime of the
// connection. Without this, a peer can pin memory by sending a
// single large message and keeping the connection open.

// Call this periodically to update stats for throttling purposes.
// Not goroutine-safe
func (ch *Channel) updateStats() {
	_ = "STUB: not implemented"
	// Exponential decay of stats.
	// TODO: optimize.
	return
}

//----------------------------------------
// Packet

// mustWrapPacket takes a packet kind (oneof) and wraps it in a tmp2p.Packet message.
func mustWrapPacket(pb proto.Message) *tmp2p.Packet { _ = "STUB: not implemented"; return nil }

// already a packet
