package trace

import (
	"bufio"
	"os"
	"sync"
	"sync/atomic"
)

// bufferedFile is a file that is being written to and read from. It is thread
// safe, however, when reading from the file, writes will be ignored.
type bufferedFile struct {
	// reading protects the file from being written to while it is being read
	// from. This is needed beyond in addition to the mutex so that writes can
	// be ignored while reading.
	reading atomic.Bool

	// mut protects the buffered writer.
	mut *sync.Mutex

	// file is the file that is being written to.
	file *os.File

	// writer is the buffered writer that is writing to the file.
	wr *bufio.Writer
}

// newbufferedFile creates a new buffered file that writes to the given file.
func newbufferedFile(file *os.File) *bufferedFile { _ = "STUB: not implemented"; return nil }

// Write writes the given bytes to the file. If the file is currently being read
// from, the write will be lost.
func (f *bufferedFile) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (f *bufferedFile) startReading() error { _ = "STUB: not implemented"; return nil }

func (f *bufferedFile) stopReading() error { _ = "STUB: not implemented"; return nil }

// File returns the underlying file with the seek point reset. The caller should
// not close the file. The caller must call the returned function when they are
// done reading from the file. This function resets the seek point to where it
// was being written to.
func (f *bufferedFile) File() (*os.File, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Close closes the file.
func (f *bufferedFile) Close() error {
	_ = "STUB: not implemented"
	// set reading to true to prevent writes while closing the file.
	return nil
}
