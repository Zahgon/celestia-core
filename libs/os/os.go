package os

import (
	"os"
)

type logger interface {
	Info(msg string, keyvals ...interface{})
}

// TrapSignal catches the SIGTERM/SIGINT and executes cb function. After that it exits
// with code 0.
func TrapSignal(logger logger, cb func()) { _ = "STUB: not implemented"; return }

// Kill the running process by sending itself SIGTERM.
func Kill() error { _ = "STUB: not implemented"; return nil }

func Exit(s string) { _ = "STUB: not implemented"; return }

// EnsureDir ensures the given directory exists, creating it if necessary.
// Errors if the path already exists as a non-directory.
func EnsureDir(dir string, mode os.FileMode) error { _ = "STUB: not implemented"; return nil }

func FileExists(filePath string) bool { _ = "STUB: not implemented"; return false }

func ReadFile(filePath string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func MustReadFile(filePath string) []byte { _ = "STUB: not implemented"; return nil }

func WriteFile(filePath string, contents []byte, mode os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

func MustWriteFile(filePath string, contents []byte, mode os.FileMode) {
	_ = "STUB: not implemented"
	return
}

// CopyFile copies a file. It truncates the destination file if it exists.
func CopyFile(src, dst string) error { _ = "STUB: not implemented"; return nil }

// create new file, truncate if exists and apply same permissions as the original one
