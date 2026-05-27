package trace

import (
	"bufio"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const jsonL = ".jsonl"

func (lt *LocalTracer) getTableHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Parse the request to get the data

//nolint:errcheck

// Use the pump function to continuously read from the file and write to
// the response writer

// Set the content type to the writer's form data content type

// Copy the data from the reader to the response writer

// pump continuously reads from a bufio.Reader and writes to a multipart.Writer.
// It returns the reader end of the pipe and the writer for consumption by the
// server.
func pump(table string, br *bufio.Reader) (*io.PipeReader, *multipart.Writer) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lt *LocalTracer) servePullData() { _ = "STUB: not implemented"; return }

//nolint:gosec

// GetTable downloads a table from the server and saves it to the given directory. It uses a multipart
// response to download the file.
func GetTable(serverURL, table, dirPath string) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec

// End of multipart

// S3Config is a struct that holds the configuration for an S3 bucket.
type S3Config struct {
	BucketName string `json:"bucket_name"`
	Region     string `json:"region"`
	AccessKey  string `json:"access_key"`
	SecretKey  string `json:"secret_key"`
	// PushDelay is the time in seconds to wait before pushing the file to S3.
	// If this is 0, it defaults is used.
	PushDelay int64 `json:"push_delay"`
}

// readS3Config reads an S3Config from a file in the given directory.
func readS3Config(dir string) (S3Config, error) {
	_ = "STUB: not implemented"
	return *new(S3Config), nil
}

// PushS3 pushes a file to an S3 bucket using the given S3Config. It uses the
// chainID and the nodeID to organize the files in the bucket. The directory
// structure is chainID/nodeID/table.jsonl .
func PushS3(chainID, nodeID string, s3cfg S3Config, f *os.File) error {
	_ = "STUB: not implemented"
	return nil
}

func (lt *LocalTracer) pushLoop() { _ = "STUB: not implemented"; return }

func (lt *LocalTracer) PushAll() error { _ = "STUB: not implemented"; return nil }

//nolint:gosec

// S3Download downloads files that match some prefix from an S3 bucket to a
// local directory dst.
// fileNames is a list of traced jsonl file names to download. If it is empty, all traces are downloaded.
// fileNames should not have .jsonl suffix.
func S3Download(dst, prefix string, cfg S3Config, fileNames ...string) error {
	_ = "STUB: not implemented"
	// Ensure local directory structure exists
	return nil
}

// If no fileNames are specified, download all files

// Add .jsonl suffix to the fileNames

// Create the directories in the path

// Create a file to write the S3 Object contents to.

// Copy the contents of the S3 object to the local file
//
