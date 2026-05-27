package p2p

import (
	"github.com/cometbft/cometbft/crypto"
)

// ID is a hex-encoded crypto.Address
type ID string

// IDByteLength is the length of a crypto.Address. Currently only 20.
// TODO: support other length addresses ?
const IDByteLength = crypto.AddressSize

//------------------------------------------------------------------------------
// Persistent peer ID
// TODO: encrypt on disk

// NodeKey is the persistent peer key.
// It contains the nodes private key for authentication.
type NodeKey struct {
	PrivKey crypto.PrivKey `json:"priv_key"` // our priv key
}

// ID returns the peer's canonical ID - the hash of its public key.
func (nodeKey *NodeKey) ID() ID { _ = "STUB: not implemented"; return *new(ID) }

// PubKey returns the peer's PubKey
func (nodeKey *NodeKey) PubKey() crypto.PubKey {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey)
}

// PubKeyToID returns the ID corresponding to the given PubKey.
// It's the hex-encoding of the pubKey.Address().
func PubKeyToID(pubKey crypto.PubKey) ID { _ = "STUB: not implemented"; return *new(ID) }

// LoadOrGenNodeKey attempts to load the NodeKey from the given filePath. If
// the file does not exist, it generates and saves a new NodeKey.
func LoadOrGenNodeKey(filePath string) (*NodeKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadNodeKey loads NodeKey located in filePath.
func LoadNodeKey(filePath string) (*NodeKey, error) { _ = "STUB: not implemented"; return nil, nil }

// SaveAs persists the NodeKey to filePath.
func (nodeKey *NodeKey) SaveAs(filePath string) error { _ = "STUB: not implemented"; return nil }

//------------------------------------------------------------------------------

// MakePoWTarget returns the big-endian encoding of 2^(targetBits - difficulty) - 1.
// It can be used as a Proof of Work target.
// NOTE: targetBits must be a multiple of 8 and difficulty must be less than targetBits.
func MakePoWTarget(difficulty, targetBits uint) []byte { _ = "STUB: not implemented"; return nil }
