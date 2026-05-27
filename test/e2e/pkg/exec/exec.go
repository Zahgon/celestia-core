package exec

import (
	"context"
)

// Command executes a shell command.
func Command(ctx context.Context, args ...string) error { _ = "STUB: not implemented"; return nil }

// CommandOutput executes a shell command and returns the command's output.
func CommandOutput(ctx context.Context, args ...string) ([]byte, error) {
	_ = "STUB: not implemented"
	//nolint: gosec
	// G204: Subprocess launched with a potential tainted input or cmd arguments
	return nil, nil
}

// CommandVerbose executes a shell command while displaying its output.
func CommandVerbose(ctx context.Context, args ...string) error {
	_ = "STUB: not implemented"
	//nolint: gosec
	// G204: Subprocess launched with a potential tainted input or cmd arguments
	return nil
}
