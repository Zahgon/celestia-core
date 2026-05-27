package autofile

import (
	"bufio"
	"os"
	"sync"
	"time"

	"github.com/cometbft/cometbft/libs/service"
)

const (
	defaultGroupCheckDuration = 5000 * time.Millisecond
	defaultHeadSizeLimit      = 10 * 1024 * 1024       // 10MB
	defaultTotalSizeLimit     = 1 * 1024 * 1024 * 1024 // 1GB
	maxFilesToRemove          = 4                      // needs to be greater than 1
)

/*
You can open a Group to keep restrictions on an AutoFile, like
the maximum size of each chunk, and/or the total amount of bytes
stored in the group.

The first file to be written in the Group.Dir is the head file.

	Dir/
	- <HeadPath>

Once the Head file reaches the size limit, it will be rotated.

	Dir/
	- <HeadPath>.000   // First rolled file
	- <HeadPath>       // New head path, starts empty.
										 // The implicit index is 001.

As more files are written, the index numbers grow...

	Dir/
	- <HeadPath>.000   // First rolled file
	- <HeadPath>.001   // Second rolled file
	- ...
	- <HeadPath>       // New head path

The Group can also be used to binary-search for some line,
assuming that marker lines are written occasionally.
*/
type Group struct {
	service.BaseService

	ID                 string
	Head               *AutoFile // The head AutoFile to write to
	headBuf            *bufio.Writer
	Dir                string // Directory that contains .Head
	ticker             *time.Ticker
	mtx                sync.Mutex
	headSizeLimit      int64
	totalSizeLimit     int64
	groupCheckDuration time.Duration
	minIndex           int // Includes head
	maxIndex           int // Includes head, where Head will move to

	// close this when the processTicks routine is done.
	// this ensures we can cleanup the dir after calling Stop
	// and the routine won't be trying to access it anymore
	doneProcessTicks chan struct{}

	// TODO: When we start deleting files, we need to start tracking GroupReaders
	// and their dependencies.
}

// OpenGroup creates a new Group with head at headPath. It returns an error if
// it fails to open head file.
func OpenGroup(headPath string, groupOptions ...func(*Group)) (*Group, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GroupCheckDuration allows you to overwrite default groupCheckDuration.
func GroupCheckDuration(duration time.Duration) func(*Group) { _ = "STUB: not implemented"; return nil }

// GroupHeadSizeLimit allows you to overwrite default head size limit - 10MB.
func GroupHeadSizeLimit(limit int64) func(*Group) { _ = "STUB: not implemented"; return nil }

// GroupTotalSizeLimit allows you to overwrite default total size limit of the group - 1GB.
func GroupTotalSizeLimit(limit int64) func(*Group) { _ = "STUB: not implemented"; return nil }

// OnStart implements service.Service by starting the goroutine that checks file
// and group limits.
func (g *Group) OnStart() error { _ = "STUB: not implemented"; return nil }

// OnStop implements service.Service by stopping the goroutine described above.
// NOTE: g.Head must be closed separately using Close.
func (g *Group) OnStop() { _ = "STUB: not implemented"; return }

// Wait blocks until all internal goroutines are finished. Supposed to be
// called after Stop.
func (g *Group) Wait() {
	_ = "STUB: not implemented"
	// wait for processTicks routine to finish
	return
}

// Close closes the head file. The group must be stopped by this moment.
func (g *Group) Close() { _ = "STUB: not implemented"; return }

// HeadSizeLimit returns the current head size limit.
func (g *Group) HeadSizeLimit() int64 { _ = "STUB: not implemented"; return 0 }

// TotalSizeLimit returns total size limit of the group.
func (g *Group) TotalSizeLimit() int64 { _ = "STUB: not implemented"; return 0 }

// MaxIndex returns index of the last file in the group.
func (g *Group) MaxIndex() int { _ = "STUB: not implemented"; return 0 }

// MinIndex returns index of the first file in the group.
func (g *Group) MinIndex() int { _ = "STUB: not implemented"; return 0 }

// Write writes the contents of p into the current head of the group. It
// returns the number of bytes written. If nn < len(p), it also returns an
// error explaining why the write is short.
// NOTE: Writes are buffered so they don't write synchronously
// TODO: Make it halt if space is unavailable
func (g *Group) Write(p []byte) (nn int, err error) { _ = "STUB: not implemented"; return 0, nil }

// WriteLine writes line into the current head of the group. It also appends "\n".
// NOTE: Writes are buffered so they don't write synchronously
// TODO: Make it halt if space is unavailable
func (g *Group) WriteLine(line string) error { _ = "STUB: not implemented"; return nil }

// Buffered returns the size of the currently buffered data.
func (g *Group) Buffered() int { _ = "STUB: not implemented"; return 0 }

// FlushAndSync writes any buffered data to the underlying file and commits the
// current content of the file to stable storage (fsync).
func (g *Group) FlushAndSync() error { _ = "STUB: not implemented"; return nil }

func (g *Group) processTicks() { _ = "STUB: not implemented"; return }

// NOTE: this function is called manually in tests.
func (g *Group) checkHeadSizeLimit() { _ = "STUB: not implemented"; return }

func (g *Group) checkTotalSizeLimit() { _ = "STUB: not implemented"; return }

// Special degenerate case, just do nothing.

// RotateFile causes group to close the current head and assign it some index.
// Note it does not create a new head.
func (g *Group) RotateFile() { _ = "STUB: not implemented"; return }

// NewReader returns a new group reader.
// CONTRACT: Caller must close the returned GroupReader.
func (g *Group) NewReader(index int) (*GroupReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GroupInfo holds information about the group.
type GroupInfo struct {
	MinIndex  int   // index of the first file in the group, including head
	MaxIndex  int   // index of the last file in the group, including head
	TotalSize int64 // total size of the group
	HeadSize  int64 // size of the head
}

// Returns info after scanning all files in g.Head's dir.
func (g *Group) ReadGroupInfo() GroupInfo { _ = "STUB: not implemented"; return *new(GroupInfo) }

// Index includes the head.
// CONTRACT: caller should have called g.mtx.Lock
func (g *Group) readGroupInfo() GroupInfo { _ = "STUB: not implemented"; return *new(GroupInfo) }

//nolint:staticcheck

// For each file in the directory, filter by pattern

// Matches

// Now account for the head.

// If there were no numbered files,
// then the head is index 0.

// Otherwise, the head file is 1 greater

func filePathForIndex(headPath string, index int, maxIndex int) string {
	_ = "STUB: not implemented"
	return ""
}

//--------------------------------------------------------------------------------

// GroupReader provides an interface for reading from a Group.
type GroupReader struct {
	*Group
	mtx       sync.Mutex
	curIndex  int
	curFile   *os.File
	curReader *bufio.Reader
	curLine   []byte
}

func newGroupReader(g *Group) *GroupReader { _ = "STUB: not implemented"; return nil }

// Close closes the GroupReader by closing the cursor file.
func (gr *GroupReader) Close() error { _ = "STUB: not implemented"; return nil }

// Read implements io.Reader, reading bytes from the current Reader
// incrementing index until enough bytes are read.
func (gr *GroupReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Open file if not open yet

// Iterate over files until enough bytes are read

// Open the next file

// empty file

// IF index > gr.Group.maxIndex, returns io.EOF
// CONTRACT: caller should hold gr.mtx
func (gr *GroupReader) openFile(index int) error {
	_ = "STUB: not implemented"
	// Lock on Group to ensure that head doesn't move in the meanwhile.
	return nil
}

//nolint:staticcheck

//nolint:staticcheck

// Update gr.cur*

// TODO return error?

// CurIndex returns cursor's file index.
func (gr *GroupReader) CurIndex() int { _ = "STUB: not implemented"; return 0 }

// SetIndex sets the cursor's file index to index by opening a file at this
// position.
func (gr *GroupReader) SetIndex(index int) error { _ = "STUB: not implemented"; return nil }
