package conn

import (
	"crypto/cipher"
	"errors"
	"io"
	"net"
	"time"

	"golang.org/x/crypto/chacha20poly1305"

	"github.com/cometbft/cometbft/crypto"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

// 4 + 1024 == 1028 total frame size
const (
	dataLenSize      = 4
	dataMaxSize      = 1024
	totalFrameSize   = dataMaxSize + dataLenSize
	aeadSizeOverhead = 16 // overhead of poly 1305 authentication tag
	aeadKeySize      = chacha20poly1305.KeySize
	aeadNonceSize    = chacha20poly1305.NonceSize

	labelEphemeralLowerPublicKey = "EPHEMERAL_LOWER_PUBLIC_KEY"
	labelEphemeralUpperPublicKey = "EPHEMERAL_UPPER_PUBLIC_KEY"
	labelDHSecret                = "DH_SECRET"
	labelSecretConnectionMac     = "SECRET_CONNECTION_MAC"
)

var (
	ErrSmallOrderRemotePubKey = errors.New("detected low order point from remote peer")

	secretConnKeyAndChallengeGen = []byte("TENDERMINT_SECRET_CONNECTION_KEY_AND_CHALLENGE_GEN")
)

// SecretConnection implements net.Conn.
// It is an implementation of the STS protocol.
// See https://github.com/cometbft/cometbft/blob/0.1/docs/sts-final.pdf for
// details on the protocol.
//
// Consumers of the SecretConnection are responsible for authenticating
// the remote peer's pubkey against known information, like a nodeID.
// Otherwise they are vulnerable to MITM.
// (TODO(ismail): see also https://github.com/tendermint/tendermint/issues/3010)
type SecretConnection struct {

	// immutable
	recvAead cipher.AEAD
	sendAead cipher.AEAD

	remPubKey crypto.PubKey
	conn      io.ReadWriteCloser

	// net.Conn must be thread safe:
	// https://golang.org/pkg/net/#Conn.
	// Since we have internal mutable state,
	// we need mtxs. But recv and send states
	// are independent, so we can use two mtxs.
	// All .Read are covered by recvMtx,
	// all .Write are covered by sendMtx.
	recvMtx         cmtsync.Mutex
	recvBuffer      []byte
	recvNonce       *[aeadNonceSize]byte
	recvFrame       []byte
	recvSealedFrame []byte

	sendMtx         cmtsync.Mutex
	sendNonce       *[aeadNonceSize]byte
	sendFrame       []byte
	sendSealedFrame []byte
}

// MakeSecretConnection performs handshake and returns a new authenticated
// SecretConnection.
// Returns nil if there is an error in handshake.
// Caller should call conn.Close()
// See docs/sts-final.pdf for more information.
func MakeSecretConnection(conn io.ReadWriteCloser, locPrivKey crypto.PrivKey) (*SecretConnection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Generate ephemeral keys for perfect forward secrecy.

// Write local ephemeral pubkey and receive one too.
// NOTE: every 32-byte string is accepted as a Curve25519 public key (see
// DJB's Curve25519 paper: http://cr.yp.to/ecdh/curve25519-20060209.pdf)

// Sort by lexical order.

// Check if the local ephemeral public key was the least, lexicographically
// sorted.

// Compute common diffie hellman secret using X25519.

// Generate the secret used for receiving, sending, challenge via HKDF-SHA2
// on the transcript state (which itself also uses HKDF-SHA2 to derive a key
// from the dhSecret).

// Sign the challenge bytes for authentication.

// Share (in secret) each other's pubkey & challenge signature

// We've authorized.

// RemotePubKey returns authenticated remote pubkey
func (sc *SecretConnection) RemotePubKey() crypto.PubKey {
	_ = "STUB: not implemented"
	return *

	// Writes encrypted frames of `totalFrameSize + aeadSizeOverhead`.
	// CONTRACT: data smaller than dataMaxSize is written atomically.
	new(crypto.PubKey)
}

func (sc *SecretConnection) Write(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// encrypt the frame

// end encryption

// CONTRACT: data smaller than dataMaxSize is read atomically.
func (sc *SecretConnection) Read(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// read off and update the recvBuffer, if non-empty

// read off the conn

// decrypt the frame.
// reads and updates the sc.recvNonce

// end decryption

// copy checkLength worth into data,
// set recvBuffer to the rest.
// read the first four bytes

// Implements net.Conn
func (sc *SecretConnection) Close() error                  { _ = "STUB: not implemented"; return nil }
func (sc *SecretConnection) LocalAddr() net.Addr           { _ = "STUB: not implemented"; return *new(net.Addr) }
func (sc *SecretConnection) RemoteAddr() net.Addr          { _ = "STUB: not implemented"; return *new(net.Addr) }
func (sc *SecretConnection) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
func (sc *SecretConnection) SetReadDeadline(t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (sc *SecretConnection) SetWriteDeadline(t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func genEphKeys() (ephPub, ephPriv *[32]byte) {
	_ = "STUB: not implemented"

	// TODO: Probably not a problem but ask Tony: different from the rust implementation (uses x25519-dalek),
	// we do not "clamp" the private key scalar:
	// see: https://github.com/dalek-cryptography/x25519-dalek/blob/34676d336049df2bba763cc076a75e47ae1f170f/src/x25519.rs#L56-L74
	return nil, nil
}

func shareEphPubKey(conn io.ReadWriter, locEphPub *[32]byte) (remEphPub *[32]byte, err error) {
	_ = "STUB: not implemented"

	// Send our pubkey and receive theirs in tandem.
	return nil, nil
}

// abort

// abort

// If error:

// Otherwise:

func deriveSecrets(
	dhSecret *[32]byte,
	locIsLeast bool,
) (recvSecret, sendSecret *[aeadKeySize]byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get enough data for 2 aead keys, and a 32 byte challenge

// bytes 0 through aeadKeySize - 1 are one aead key.
// bytes aeadKeySize through 2*aeadKeySize -1 are another aead key.
// which key corresponds to sending and receiving key depends on whether
// the local key is less than the remote key.

// computeDHSecret computes a Diffie-Hellman shared secret key
// from our own local private key and the other's public key.
func computeDHSecret(remPubKey, locPrivKey *[32]byte) (*[32]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sort32(foo, bar *[32]byte) (lo, hi *[32]byte) { _ = "STUB: not implemented"; return nil, nil }

func signChallenge(challenge *[32]byte, locPrivKey crypto.PrivKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type authSigMessage struct {
	Key crypto.PubKey
	Sig []byte
}

func shareAuthSignature(sc io.ReadWriter, pubKey crypto.PubKey, signature []byte) (recvMsg authSigMessage, err error) {
	_ = "STUB: not implemented"

	// Send our info and receive theirs in tandem.
	return *new(authSigMessage), nil
}

// abort

// abort

// abort

// If error:

//--------------------------------------------------------------------------------

// Increment nonce little-endian by 1 with wraparound.
// Due to chacha20poly1305 expecting a 12 byte nonce we do not use the first four
// bytes. We only increment a 64 bit unsigned int in the remaining 8 bytes
// (little-endian in nonce[4:]).
func incrNonce(nonce *[aeadNonceSize]byte) { _ = "STUB: not implemented"; return }

// Terminates the session and makes sure the nonce would not re-used.
// See https://github.com/tendermint/tendermint/issues/3531
