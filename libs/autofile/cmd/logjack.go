package main

import (
	"fmt"
	"io"
	"os"

	auto "github.com/cometbft/cometbft/libs/autofile"
	cmtos "github.com/cometbft/cometbft/libs/os"
)

const (
	Version        = "0.0.1"
	readBufferSize = 1024 // 1KB at a time
)

// Parse command-line options
func parseFlags() (headPath string, chopSize int64, limitSize int64, version bool) {
	_ = "STUB: not implemented"
	return "", 0, 0, false
}

type fmtLogger struct{}

func (fmtLogger) Info(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }

func main() {
	// Stop upon receiving SIGTERM or CTRL-C.
	cmtos.TrapSignal(fmtLogger{}, func() {
		fmt.Println("logjack shutting down")
	})

	// Read options
	headPath, chopSize, limitSize, version := parseFlags()
	if version {
		fmt.Printf("logjack version %v\n", Version)
		return
	}

	// Open Group
	group, err := auto.OpenGroup(headPath, auto.GroupHeadSizeLimit(chopSize), auto.GroupTotalSizeLimit(limitSize))
	if err != nil {
		fmt.Printf("logjack couldn't create output file %v\n", headPath)
		os.Exit(1)
	}

	if err = group.Start(); err != nil {
		fmt.Printf("logjack couldn't start with file %v\n", headPath)
		os.Exit(1)
	}

	// Forever read from stdin and write to AutoFile.
	buf := make([]byte, readBufferSize)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			if err := group.Stop(); err != nil {
				fmt.Fprintf(os.Stderr, "logjack stopped with error %v\n", headPath)
				os.Exit(1)
			}
			if err == io.EOF {
				os.Exit(0)
			}
			fmt.Println("logjack errored")
			os.Exit(1)
		}
		_, err = group.Write(buf[:n])
		if err != nil {
			fmt.Fprintf(os.Stderr, "logjack failed write with error %v\n", headPath)
			os.Exit(1)
		}
		if err := group.FlushAndSync(); err != nil {
			fmt.Fprintf(os.Stderr, "logjack flushsync fail with error %v\n", headPath)
			os.Exit(1)
		}
	}
}

func parseBytesize(chopSize string) int64 {
	_ = "STUB: not implemented"
	// Handle suffix multiplier
	return 0
}

// Parse the numeric part
