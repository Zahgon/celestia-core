package pex

/* Loading & Saving */

type addrBookJSON struct {
	Key   string          `json:"key"`
	Addrs []*knownAddress `json:"addrs"`
}

func (a *addrBook) saveToFile(filePath string) { _ = "STUB: not implemented"; return }

// Returns false if file does not exist.
// cmn.Panics if file is corrupt.
func (a *addrBook) loadFromFile(filePath string) bool {
	_ = "STUB: not implemented"
	// If doesn't exist, do nothing.
	return false
}

// Load addrBookJSON{}

// Restore all the fields...
// Restore the key

// Restore .bucketsNew & .bucketsOld
