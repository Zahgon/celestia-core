// metricsdiff is a tool for generating a diff between two different files containing
// prometheus metrics. metricsdiff outputs which metrics have been added, removed,
// or have different sets of labels between the two files.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	dto "github.com/prometheus/client_model/go"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %[1]s <path1> <path2>

Generate the diff between the two files of Prometheus metrics.
The input should have the format output by a Prometheus HTTP endpoint.
The tool indicates which metrics have been added, removed, or use different
label sets from path1 to path2.

`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

// Diff contains the set of metrics that were modified between two files
// containing prometheus metrics output.
type Diff struct {
	Adds    []string
	Removes []string

	Changes []LabelDiff
}

// LabelDiff describes the label changes between two versions of the same metric.
type LabelDiff struct {
	Metric  string
	Adds    []string
	Removes []string
}

type parsedMetric struct {
	name   string
	labels []string
}

type metricsList []parsedMetric

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		log.Fatalf("Usage is '%s <path1> <path2>', got %d arguments",
			filepath.Base(os.Args[0]), flag.NArg())
	}
	fa, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("Open: %v", err)
	}
	defer fa.Close()
	fb, err := os.Open(flag.Arg(1))
	if err != nil {
		log.Fatalf("Open: %v", err)
	}
	defer fb.Close()
	md, err := DiffFromReaders(fa, fb)
	if err != nil {
		log.Fatalf("Generating diff: %v", err)
	}
	fmt.Print(md)
}

// DiffFromReaders parses the metrics present in the readers a and b and
// determines which metrics were added and removed in b.
func DiffFromReaders(a, b io.Reader) (Diff, error) {
	_ = "STUB: not implemented"
	return *new(Diff), nil
}

func toList(l map[string]*dto.MetricFamily) metricsList {
	_ = "STUB: not implemented"
	return *new(metricsList)
}

func labelsToStringList(ls []*dto.LabelPair) []string { _ = "STUB: not implemented"; return nil }

func listDiff(a, b []string) ([]string, []string) { _ = "STUB: not implemented"; return nil, nil }

func (m metricsList) Len() int           { _ = "STUB: not implemented"; return 0 }
func (m metricsList) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (m metricsList) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func (m Diff) String() string { _ = "STUB: not implemented"; return "" }
