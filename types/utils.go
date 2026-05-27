package types

// Go lacks a simple and safe way to see if something is a typed nil.
// See:
//   - https://dave.cheney.net/2017/08/09/typed-nils-in-go-2
//   - https://groups.google.com/forum/#!topic/golang-nuts/wnH302gBa4I/discussion
//   - https://github.com/golang/go/issues/21538
func isTypedNil(o interface{}) bool { _ = "STUB: not implemented"; return false }

// Returns true if it has zero length.
func isEmpty(o interface{}) bool { _ = "STUB: not implemented"; return false }

// SaveBlockToFile saves a CometBFT block to a file in the given directory.
func SaveBlockToFile(dir string, filename string, block *Block) error {
	_ = "STUB: not implemented"
	return nil
}
