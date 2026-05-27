package schema

import "github.com/cometbft/cometbft/libs/trace"

// P2PTables returns the list of tables that are used for p2p tracing.
func P2PTables() []string { _ = "STUB: not implemented"; return nil }

const (
	// PeerUpdateTable is the name of the table that stores the p2p peer
	// updates.
	PeersTable = "peers"
)

// P2PPeerUpdate is an enum that represents the different types of p2p
// trace data.
type P2PPeerUpdate string

const (
	// PeerJoin is the action for when a peer is connected.
	PeerJoin P2PPeerUpdate = "connect"
	// PeerDisconnect is the action for when a peer is disconnected.
	PeerDisconnect P2PPeerUpdate = "disconnect"
)

// PeerUpdate describes schema for the "peer_update" table.
type PeerUpdate struct {
	PeerID string `json:"peer_id"`
	Action string `json:"action"`
	Reason string `json:"reason"`
}

// Table returns the table name for the PeerUpdate struct.
func (PeerUpdate) Table() string {
	_ = "STUB: not implemented"

	// WritePeerUpdate writes a tracing point for a peer update using the predetermined
	// schema for p2p tracing.
	return ""
}

func WritePeerUpdate(client trace.Tracer, peerID string, action P2PPeerUpdate, reason string) {
	_ = "STUB: not implemented"
	return
}

const (
	PendingBytesTable = "pending_bytes"
)

type PendingBytes struct {
	PeerID string       `json:"peer_id"`
	Bytes  map[byte]int `json:"bytes"`
}

func (PendingBytes) Table() string { _ = "STUB: not implemented"; return "" }

func WritePendingBytes(client trace.Tracer, peerID string, bytes map[byte]int) {
	_ = "STUB: not implemented"
	return
}

const (
	ReceivedBytesTable = "received_bytes"
)

type ReceivedBytes struct {
	PeerID  string `json:"peer_id"`
	Channel byte   `json:"channel"`
	Bytes   int    `json:"bytes"`
}

func (ReceivedBytes) Table() string { _ = "STUB: not implemented"; return "" }

func WriteReceivedBytes(client trace.Tracer, peerID string, channel byte, bytes int) {
	_ = "STUB: not implemented"
	return
}

const (
	QueueLimitTable = "queue_limit"
)

type QueueLimit struct {
	Channel  byte   `json:"channel"`
	Reactor  string `json:"reactor"`
	LimitHit bool   `json:"limit_hit"`
}

func (QueueLimit) Table() string { _ = "STUB: not implemented"; return "" }

func WriteQueueLimit(client trace.Tracer, channel byte, reactor string, limitHit bool) {
	_ = "STUB: not implemented"
	return
}
