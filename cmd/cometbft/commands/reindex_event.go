package commands

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	cmtcfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/state/indexer"
	"github.com/cometbft/cometbft/state/txindex"
)

const (
	reindexFailed = "event re-index failed: "
)

var (
	ErrHeightNotAvailable = errors.New("height is not available")
	ErrInvalidRequest     = errors.New("invalid request")
)

// ReIndexEventCmd constructs a command to re-index events in a block height interval.
var ReIndexEventCmd = &cobra.Command{
	Use:     "reindex-event",
	Aliases: []string{"reindex_event"},
	Short:   "reindex events to the event store backends",
	Long: `
reindex-event is an offline tooling to re-index block and tx events to the eventsinks,
you can run this command when the event store backend dropped/disconnected or you want to 
replace the backend. The default start-height is 0, meaning the tooling will start 
reindex from the base block height(inclusive); and the default end-height is 0, meaning 
the tooling will reindex until the latest block height(inclusive). User can omit
either or both arguments.

Note: This operation requires ABCI Responses. Do not set DiscardABCIResponses to true if you
want to use this command.
	`,
	Example: `
	cometbft reindex-event
	cometbft reindex-event --start-height 2
	cometbft reindex-event --end-height 10
	cometbft reindex-event --start-height 2 --end-height 10
	`,
	Run: func(cmd *cobra.Command, args []string) {
		bs, ss, err := loadStateAndBlockStore(config)
		if err != nil {
			fmt.Println(reindexFailed, err)
			return
		}

		state, err := ss.Load()
		if err != nil {
			fmt.Println(reindexFailed, err)
			return
		}

		if err := checkValidHeight(bs); err != nil {
			fmt.Println(reindexFailed, err)
			return
		}

		bi, ti, err := loadEventSinks(config, state.ChainID)
		if err != nil {
			fmt.Println(reindexFailed, err)
			return
		}

		riArgs := eventReIndexArgs{
			startHeight:  startHeight,
			endHeight:    endHeight,
			blockIndexer: bi,
			txIndexer:    ti,
			blockStore:   bs,
			stateStore:   ss,
		}
		if err := eventReIndex(cmd, riArgs); err != nil {
			panic(fmt.Errorf("%s: %w", reindexFailed, err))
		}

		fmt.Println("event re-index finished")
	},
}

var (
	startHeight int64
	endHeight   int64
)

func init() {
	ReIndexEventCmd.Flags().Int64Var(&startHeight, "start-height", 0, "the block height would like to start for re-index")
	ReIndexEventCmd.Flags().Int64Var(&endHeight, "end-height", 0, "the block height would like to finish for re-index")
}

func loadEventSinks(cfg *cmtcfg.Config, chainID string) (indexer.BlockIndexer, txindex.TxIndexer, error) {
	_ = "STUB: not implemented"
	return *new(indexer.BlockIndexer), *new(txindex.TxIndexer), nil
}

type eventReIndexArgs struct {
	startHeight  int64
	endHeight    int64
	blockIndexer indexer.BlockIndexer
	txIndexer    txindex.TxIndexer
	blockStore   state.BlockStore
	stateStore   state.Store
}

func eventReIndex(cmd *cobra.Command, args eventReIndexArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func checkValidHeight(bs state.BlockStore) error { _ = "STUB: not implemented"; return nil }
