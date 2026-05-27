package debug

// zipDir zips all the contents found in src, including both files and
// directories, into a destination file dest. It returns an error upon failure.
// It assumes src is a directory.
func zipDir(src, dest string) error { _ = "STUB: not implemented"; return nil }

// Each execution of this utility on a CometBFT process will result in a
// unique file.

// Handle cases where the content to be zipped is a file or a directory,
// where a directory must have a '/' suffix.

// copyFile copies a file from src to dest and returns an error upon failure. The
// copied file retains the source file's permissions.
func copyFile(src, dest string) error { _ = "STUB: not implemented"; return nil }

// writeStateToFile pretty JSON encodes an object and writes it to file composed
// of dir and filename. It returns an error upon failure to encode or write to
// file.
func writeStateJSONToFile(state interface{}, dir, filename string) error {
	_ = "STUB: not implemented"
	return nil
}
