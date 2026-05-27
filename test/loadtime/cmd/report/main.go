package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	dbm "github.com/cometbft/cometbft-db"

	"github.com/cometbft/cometbft/store"
	"github.com/cometbft/cometbft/test/loadtime/report"
)

var (
	db     = flag.String("database-type", "pebbledb", "the type of database holding the blockstore")
	dir    = flag.String("data-dir", "", "path to the directory containing the CometBFT databases")
	csvOut = flag.String("csv", "", "dump the extracted latencies as raw csv for use in additional tooling")
)

func main() {
	flag.Parse()
	if *db == "" {
		log.Fatalf("must specify a database-type")
	}
	if *dir == "" {
		log.Fatalf("must specify a data-dir")
	}
	d := strings.TrimPrefix(*dir, "~/")
	if d != *dir {
		h, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		d = h + "/" + d
	}
	_, err := os.Stat(d)
	if err != nil {
		panic(err)
	}
	dbType := dbm.BackendType(*db)
	db, err := dbm.NewDB("blockstore", dbType, d)
	if err != nil {
		panic(err)
	}
	s := store.NewBlockStore(db)
	defer s.Close()
	rs, err := report.GenerateFromBlockStore(s)
	if err != nil {
		panic(err)
	}
	if *csvOut != "" {
		cf, err := os.Create(*csvOut)
		if err != nil {
			panic(err)
		}
		w := csv.NewWriter(cf)
		err = w.WriteAll(toCSVRecords(rs.List()))
		if err != nil {
			panic(err)
		}
		return
	}
	for _, r := range rs.List() {
		fmt.Printf(""+
			"Experiment ID: %s\n\n"+
			"\tConnections: %d\n"+
			"\tRate: %d\n"+
			"\tSize: %d\n\n"+
			"\tTotal Valid Tx: %d\n"+
			"\tTotal Negative Latencies: %d\n"+
			"\tMinimum Latency: %s\n"+
			"\tMaximum Latency: %s\n"+
			"\tAverage Latency: %s\n"+
			"\tStandard Deviation: %s\n\n", r.ID, r.Connections, r.Rate, r.Size, len(r.All), r.NegativeCount, r.Min, r.Max, r.Avg, r.StdDev)
	}
	fmt.Printf("Total Invalid Tx: %d\n", rs.ErrorCount())
}

func toCSVRecords(rs []report.Report) [][]string { _ = "STUB: not implemented"; return nil }
