package pex

import (
	"time"

	"github.com/cometbft/cometbft/p2p"
)

// knownAddress tracks information about a known network address
// that is used to determine how viable an address is.
type knownAddress struct {
	Addr        *p2p.NetAddress `json:"addr"`
	Src         *p2p.NetAddress `json:"src"`
	Buckets     []int           `json:"buckets"`
	Attempts    int32           `json:"attempts"`
	BucketType  byte            `json:"bucket_type"`
	LastAttempt time.Time       `json:"last_attempt"`
	LastSuccess time.Time       `json:"last_success"`
	LastBanTime time.Time       `json:"last_ban_time"`
}

func newKnownAddress(addr *p2p.NetAddress, src *p2p.NetAddress) *knownAddress {
	_ = "STUB: not implemented"
	return nil
}

func (ka *knownAddress) ID() p2p.ID { _ = "STUB: not implemented"; return *new(p2p.ID) }

func (ka *knownAddress) isOld() bool { _ = "STUB: not implemented"; return false }

func (ka *knownAddress) isNew() bool { _ = "STUB: not implemented"; return false }

func (ka *knownAddress) markAttempt() { _ = "STUB: not implemented"; return }

func (ka *knownAddress) markGood() { _ = "STUB: not implemented"; return }

func (ka *knownAddress) ban(banTime time.Duration) { _ = "STUB: not implemented"; return }

func (ka *knownAddress) isBanned() bool { _ = "STUB: not implemented"; return false }

func (ka *knownAddress) addBucketRef(bucketIdx int) int { _ = "STUB: not implemented"; return 0 }

// TODO refactor to return error?
// log.Warn(Fmt("Bucket already exists in ka.Buckets: %v", ka))

func (ka *knownAddress) removeBucketRef(bucketIdx int) int { _ = "STUB: not implemented"; return 0 }

// TODO refactor to return error?
// log.Warn(Fmt("bucketIdx not found in ka.Buckets: %v", ka))

/*
An address is bad if the address in question is a New address, has not been tried in the last
minute, and meets one of the following criteria:

1) It claims to be from the future
2) It hasn't been seen in over a week
3) It has failed at least three times and never succeeded
4) It has failed ten times in the last week

All addresses that meet these criteria are assumed to be worthless and not
worth keeping hold of.
*/
func (ka *knownAddress) isBad() bool {
	_ = "STUB: not implemented"
	// Is Old --> good
	return false
}

// Has been attempted in the last minute --> good

// TODO: From the future?

// Too old?
// TODO: should be a timestamp of last seen, not just last attempt

// Never succeeded?

// Hasn't succeeded in too long?
