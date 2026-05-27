package rand

import (
	mrand "math/rand"
	"time"

	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

const (
	strChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" // 62 characters
)

// Rand is a prng, that is seeded with OS randomness.
// The OS randomness is obtained from crypto/rand, however none of the provided
// methods are suitable for cryptographic usage.
// They all utilize math/rand's prng internally.
//
// All of the methods here are suitable for concurrent use.
// This is achieved by using a mutex lock on all of the provided methods.
type Rand struct {
	cmtsync.Mutex
	rand *mrand.Rand
}

var grand *Rand

func init() {
	grand = NewRand()
	grand.init()
}

func NewRand() *Rand { _ = "STUB: not implemented"; return nil }

func (r *Rand) init() {
	bz := cRandBytes(8)
	var seed uint64
	for i := 0; i < 8; i++ {
		seed |= uint64(bz[i])
		seed <<= 8
	}
	r.reset(int64(seed))
}

func (r *Rand) reset(seed int64) {
	_ = "STUB: not implemented"
	// G404: Use of weak random number generator (math/rand instead of crypto/rand)
	//
	//nolint:gosec
	return
}

//----------------------------------------
// Global functions

func Seed(seed int64) { _ = "STUB: not implemented"; return }

func Str(length int) string { _ = "STUB: not implemented"; return "" }

func Uint16() uint16 { _ = "STUB: not implemented"; return 0 }

func Uint32() uint32 { _ = "STUB: not implemented"; return 0 }

func Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

func Uint() uint { _ = "STUB: not implemented"; return 0 }

func Int16() int16 { _ = "STUB: not implemented"; return 0 }

func Int32() int32 { _ = "STUB: not implemented"; return 0 }

func Int64() int64 { _ = "STUB: not implemented"; return 0 }

func Int() int { _ = "STUB: not implemented"; return 0 }

func Int31() int32 { _ = "STUB: not implemented"; return 0 }

func Int31n(n int32) int32 { _ = "STUB: not implemented"; return 0 }

func Int63() int64 { _ = "STUB: not implemented"; return 0 }

func Int63n(n int64) int64 { _ = "STUB: not implemented"; return 0 }

func Bool() bool { _ = "STUB: not implemented"; return false }

func Float32() float32 { _ = "STUB: not implemented"; return 0 }

func Float64() float64 { _ = "STUB: not implemented"; return 0 }

func Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func Bytes(n int) []byte { _ = "STUB: not implemented"; return nil }

func Intn(n int) int { _ = "STUB: not implemented"; return 0 }

func Perm(n int) []int { _ = "STUB: not implemented"; return nil }

//----------------------------------------
// Rand methods

func (r *Rand) Seed(seed int64) { _ = "STUB: not implemented"; return }

// Str constructs a random alphanumeric string of given length.
func (r *Rand) Str(length int) string { _ = "STUB: not implemented"; return "" }

// rightmost 6 bits
// only 62 characters in strChars

func (r *Rand) Uint16() uint16 { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Uint32() uint32 { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Uint() uint { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Int16() int16 { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Int32() int32 { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Int64() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Int() int { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Int31() int32 { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Int31n(n int32) int32 { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Int63() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Int63n(n int64) int64 { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Float32() float32 { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Float64() float64 { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Bytes returns n random bytes generated from the internal
// prng.
func (r *Rand) Bytes(n int) []byte {
	_ = "STUB: not implemented"
	// cRandBytes isn't guaranteed to be fast so instead
	// use random bytes generated from the internal PRNG
	return nil
}

// Intn returns, as an int, a uniform pseudo-random number in the range [0, n).
// It panics if n <= 0.
func (r *Rand) Intn(n int) int { _ = "STUB: not implemented"; return 0 }

// Bool returns a uniformly random boolean
func (r *Rand) Bool() bool {
	_ = "STUB: not implemented"
	// See https://github.com/golang/go/issues/23804#issuecomment-365370418
	// for reasoning behind computing like this
	return false
}

// Perm returns a pseudo-random permutation of n integers in [0, n).
func (r *Rand) Perm(n int) []int { _ = "STUB: not implemented"; return nil }

// NOTE: This relies on the os's random number generator.
// For real security, we should salt that with some seed.
// See github.com/cometbft/cometbft/crypto for a more secure reader.
func cRandBytes(numBytes int) []byte { _ = "STUB: not implemented"; return nil }
