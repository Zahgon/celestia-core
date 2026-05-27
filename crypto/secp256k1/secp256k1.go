package secp256k1

import (
	"io"
	"math/big"

	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	//nolint:staticcheck
	"github.com/cometbft/cometbft/crypto"
	cmtjson "github.com/cometbft/cometbft/libs/json"
)

// -------------------------------------
const (
	PrivKeyName = "tendermint/PrivKeySecp256k1"
	PubKeyName  = "tendermint/PubKeySecp256k1"

	KeyType     = "secp256k1"
	PrivKeySize = 32
)

func init() {
	cmtjson.RegisterType(PubKey{}, PubKeyName)
	cmtjson.RegisterType(PrivKey{}, PrivKeyName)
}

var _ crypto.PrivKey = PrivKey{}

// PrivKey implements PrivKey.
type PrivKey []byte

// Bytes marshalls the private key using amino encoding.
func (privKey PrivKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// PubKey performs the point-scalar multiplication from the privKey on the
// generator point to get the pubkey.
func (privKey PrivKey) PubKey() crypto.PubKey {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey)
}

// Equals - you probably don't need to use this.
// Runs in constant time based on length of the keys.
func (privKey PrivKey) Equals(other crypto.PrivKey) bool { _ = "STUB: not implemented"; return false }

func (privKey PrivKey) Type() string {
	_ = "STUB: not implemented"

	// GenPrivKey generates a new ECDSA private key on curve secp256k1 private key.
	// It uses OS randomness to generate the private key.
	return ""
}

func GenPrivKey() PrivKey { _ = "STUB: not implemented"; return *new(PrivKey) }

// genPrivKey generates a new secp256k1 private key using the provided reader.
func genPrivKey(rand io.Reader) PrivKey { _ = "STUB: not implemented"; return *new(PrivKey) }

// break if we found a valid point (i.e. > 0 and < N == curverOrder)

var one = new(big.Int).SetInt64(1)

// GenPrivKeySecp256k1 hashes the secret with SHA2, and uses
// that 32 byte output to create the private key.
//
// It makes sure the private key is a valid field element by setting:
//
// c = sha256(secret)
// k = (c mod (n − 1)) + 1, where n = curve order.
//
// NOTE: secret should be the output of a KDF like bcrypt,
// if it's derived from user input.
func GenPrivKeySecp256k1(secret []byte) PrivKey { _ = "STUB: not implemented"; return *new(PrivKey) }

// to guarantee that we have a valid field element, we use the approach of:
// "Suite B Implementer’s Guide to FIPS 186-3", A.2.1
// https://apps.nsa.gov/iaarchive/library/ia-guidance/ia-solutions-for-classified/algorithm-guidance/suite-b-implementers-guide-to-fips-186-3-ecdsa.cfm
// see also https://github.com/golang/go/blob/0380c9ad38843d523d9c9804fe300cb7edd7cd3c/src/crypto/ecdsa/ecdsa.go#L89-L101

// copy feB over to fixed 32 byte privKey32 and pad (if necessary)

// Sign creates an ECDSA signature on curve Secp256k1, using SHA256 on the msg.
// The returned signature will be of the form R || S (in lower-S form).
func (privKey PrivKey) Sign(msg []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// remove the first byte which is compactSigRecoveryCode

//-------------------------------------

var _ crypto.PubKey = PubKey{}

// PubKeySize is comprised of 32 bytes for one field element
// (the x-coordinate), plus one byte for the parity of the y-coordinate.
const PubKeySize = 33

// PubKey implements crypto.PubKey.
// It is the compressed form of the pubkey. The first byte depends is a 0x02 byte
// if the y-coordinate is the lexicographically largest of the two associated with
// the x-coordinate. Otherwise the first byte is a 0x03.
// This prefix is followed with the x-coordinate.
type PubKey []byte

// Address returns a Bitcoin style addresses: RIPEMD160(SHA256(pubkey))
func (pubKey PubKey) Address() crypto.Address {
	_ = "STUB: not implemented"
	return *new(crypto.Address)
}

// does not error

// does not error

// Bytes returns the pubkey marshaled with amino encoding.
func (pubKey PubKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (pubKey PubKey) String() string { _ = "STUB: not implemented"; return "" }

func (pubKey PubKey) Equals(other crypto.PubKey) bool { _ = "STUB: not implemented"; return false }

func (pubKey PubKey) Type() string {
	_ = "STUB: not implemented"

	// VerifySignature verifies a signature of the form R || S.
	// It rejects signatures which are not in lower-S form.
	return ""
}

func (pubKey PubKey) VerifySignature(msg []byte, sigStr []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// parse the signature:

// Reject malleable signatures. libsecp256k1 does this check but decred doesn't.
// see: https://github.com/ethereum/go-ethereum/blob/f9401ae011ddf7f8d2d95020b7446c17f8d98dc1/crypto/signature_nocgo.go#L90-L93
// Serialize() would negate S value if it is over half order.
// Hence, if the signature is different after Serialize() if should be rejected.

// Read Signature struct from R || S. Caller needs to ensure
// that len(sigStr) == 64.
func signatureFromBytes(sigStr []byte) *ecdsa.Signature { _ = "STUB: not implemented"; return nil }
