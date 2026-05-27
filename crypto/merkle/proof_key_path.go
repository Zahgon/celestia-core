package merkle

/*

	For generalized Merkle proofs, each layer of the proof may require an
	optional key.  The key may be encoded either by URL-encoding or
	(upper-case) hex-encoding.
	TODO: In the future, more encodings may be supported, like base32 (e.g.
	/32:)

	For example, for a Cosmos-SDK application where the first two proof layers
	are ValueOps, and the third proof layer is an IAVLValueOp, the keys
	might look like:

	0: []byte("App")
	1: []byte("IBC")
	2: []byte{0x01, 0x02, 0x03}

	Assuming that we know that the first two layers are always ASCII texts, we
	probably want to use URLEncoding for those, whereas the third layer will
	require HEX encoding for efficient representation.

	kp := new(KeyPath)
	kp.AppendKey([]byte("App"), KeyEncodingURL)
	kp.AppendKey([]byte("IBC"), KeyEncodingURL)
	kp.AppendKey([]byte{0x01, 0x02, 0x03}, KeyEncodingURL)
	kp.String() // Should return "/App/IBC/x:010203"

	NOTE: Key paths must begin with a `/`.

	NOTE: All encodings *MUST* work compatibly, such that you can choose to use
	whatever encoding, and the decoded keys will always be the same.  In other
	words, it's just as good to encode all three keys using URL encoding or HEX
	encoding... it just wouldn't be optimal in terms of readability or space
	efficiency.

	NOTE: Punycode will never be supported here, because not all values can be
	decoded.  For example, no string decodes to the string "xn--blah" in
	Punycode.

*/

type keyEncoding int

const (
	KeyEncodingURL keyEncoding = iota
	KeyEncodingHex
	KeyEncodingMax // Number of known encodings. Used for testing
)

type Key struct {
	name []byte
	enc  keyEncoding
}

type KeyPath []Key

func (pth KeyPath) AppendKey(key []byte, enc keyEncoding) KeyPath {
	_ = "STUB: not implemented"
	return *new(KeyPath)
}

func (pth KeyPath) String() string { _ = "STUB: not implemented"; return "" }

// Decode a path to a list of keys. Path must begin with `/`.
// Each key must use a known encoding.
func KeyPathToKeys(path string) (keys [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO Test this with random bytes, I'm not sure that it works for arbitrary bytes...
