package clist

/*

The purpose of CList is to provide a goroutine-safe linked-list.
This list can be traversed concurrently by any number of goroutines.
However, removed CElements cannot be added back.
NOTE: Not all methods of container/list are (yet) implemented.
NOTE: Removed elements need to DetachPrev or DetachNext consistently
to ensure garbage collection of removed elements.

*/

import (
	"sync"

	cmtsync "github.com/cometbft/cometbft/libs/sync"
)

// MaxLength is the max allowed number of elements a linked list is
// allowed to contain.
// If more elements are pushed to the list it will panic.
const MaxLength = int(^uint(0) >> 1)

/*
CElement is an element of a linked-list
Traversal from a CElement is goroutine-safe.

We can't avoid using WaitGroups or for-loops given the documentation
spec without re-implementing the primitives that already exist in
golang/sync. Notice that WaitGroup allows many go-routines to be
simultaneously released, which is what we want. Mutex doesn't do
this. RWMutex does this, but it's clumsy to use in the way that a
WaitGroup would be used -- and we'd end up having two RWMutex's for
prev/next each, which is doubly confusing.

sync.Cond would be sort-of useful, but we don't need a write-lock in
the for-loop. Use sync.Cond when you need serial access to the
"condition". In our case our condition is if `next != nil || removed`,
and there's no reason to serialize that condition for goroutines
waiting on NextWait() (since it's just a read operation).
*/
type CElement struct {
	mtx        cmtsync.RWMutex
	prev       *CElement
	prevWg     *sync.WaitGroup
	prevWaitCh chan struct{}
	next       *CElement
	nextWg     *sync.WaitGroup
	nextWaitCh chan struct{}
	removed    bool

	Value interface{} // immutable
}

// Blocking implementation of Next().
// May return nil iff CElement was tail and got removed.
func (e *CElement) NextWait() *CElement { _ = "STUB: not implemented"; return nil }

// e.next doesn't necessarily exist here.
// That's why we need to continue a for-loop.

// Blocking implementation of Prev().
// May return nil iff CElement was head and got removed.
func (e *CElement) PrevWait() *CElement { _ = "STUB: not implemented"; return nil }

// PrevWaitChan can be used to wait until Prev becomes not nil. Once it does,
// channel will be closed.
func (e *CElement) PrevWaitChan() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// NextWaitChan can be used to wait until Next becomes not nil. Once it does,
// channel will be closed.
func (e *CElement) NextWaitChan() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// Nonblocking, may return nil if at the end.
func (e *CElement) Next() *CElement { _ = "STUB: not implemented"; return nil }

// Nonblocking, may return nil if at the end.
func (e *CElement) Prev() *CElement { _ = "STUB: not implemented"; return nil }

func (e *CElement) Removed() bool { _ = "STUB: not implemented"; return false }

func (e *CElement) DetachNext() { _ = "STUB: not implemented"; return }

func (e *CElement) DetachPrev() { _ = "STUB: not implemented"; return }

// NOTE: This function needs to be safe for
// concurrent goroutines waiting on nextWg.
func (e *CElement) SetNext(newNext *CElement) { _ = "STUB: not implemented"; return }

// See https://golang.org/pkg/sync/:
//
// If a WaitGroup is reused to wait for several independent sets of
// events, new Add calls must happen after all previous Wait calls have
// returned.
// WaitGroups are difficult to re-use.

// NOTE: This function needs to be safe for
// concurrent goroutines waiting on prevWg
func (e *CElement) SetPrev(newPrev *CElement) { _ = "STUB: not implemented"; return }

// WaitGroups are difficult to re-use.

func (e *CElement) SetRemoved() { _ = "STUB: not implemented"; return }

// This wakes up anyone waiting in either direction.

//--------------------------------------------------------------------------------

// CList represents a linked list.
// The zero value for CList is an empty list ready to use.
// Operations are goroutine-safe.
// Panics if length grows beyond the max.
type CList struct {
	mtx    cmtsync.RWMutex
	wg     *sync.WaitGroup
	waitCh chan struct{}
	head   *CElement // first element
	tail   *CElement // last element
	curLen int       // list length
	maxLen int       // max list length
}

func (l *CList) Init() *CList { _ = "STUB: not implemented"; return nil }

// Return CList with MaxLength. CList will panic if it goes beyond MaxLength.
func New() *CList { _ = "STUB: not implemented"; return nil }

// Return CList with given maxLength.
// Will panic if list exceeds given maxLength.
func newWithMax(maxLength int) *CList { _ = "STUB: not implemented"; return nil }

func (l *CList) Len() int { _ = "STUB: not implemented"; return 0 }

func (l *CList) Front() *CElement { _ = "STUB: not implemented"; return nil }

func (l *CList) FrontWait() *CElement {
	_ = "STUB: not implemented"
	// Loop until the head is non-nil else wait and try again
	return nil
}

// NOTE: If you think l.head exists here, think harder.

func (l *CList) Back() *CElement { _ = "STUB: not implemented"; return nil }

func (l *CList) BackWait() *CElement { _ = "STUB: not implemented"; return nil }

// l.tail doesn't necessarily exist here.
// That's why we need to continue a for-loop.

// WaitChan can be used to wait until Front or Back becomes not nil. Once it
// does, channel will be closed.
func (l *CList) WaitChan() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// Panics if list grows beyond its max length.
func (l *CList) PushBack(v interface{}) *CElement {
	_ = "STUB: not implemented"

	// Construct a new element
	return nil
}

// Release waiters on FrontWait/BackWait maybe

// Modify the tail

// We must init e first.
// This will make e accessible.
// Update the list.

// CONTRACT: Caller must call e.DetachPrev() and/or e.DetachNext() to avoid memory leaks.
// NOTE: As per the contract of CList, removed elements cannot be added back.
func (l *CList) Remove(e *CElement) interface{} { _ = "STUB: not implemented"; return nil }

// If we're removing the only item, make CList FrontWait/BackWait wait.

// WaitGroups are difficult to re-use.

// Update l.len

// Connect next/prev and set head/tail

// Set .Done() on e, otherwise waiters will wait forever.

func waitGroup1() (wg *sync.WaitGroup) { _ = "STUB: not implemented"; return nil }
