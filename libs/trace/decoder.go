package trace

import (
	"os"
)

// DecodeFile reads a file and decodes it into a slice of events via
// scanning. The table parameter is used to determine the type of the events.
// The file should be a jsonl file. The generic here are passed to the event
// type.
func DecodeFile[T any](f *os.File) ([]Event[T], error) { _ = "STUB: not implemented"; return nil, nil }
