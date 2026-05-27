package testsuite

import (
	"context"

	abcicli "github.com/cometbft/cometbft/abci/client"
	"github.com/cometbft/cometbft/abci/types"
)

func InitChain(ctx context.Context, client abcicli.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func Commit(ctx context.Context, client abcicli.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func FinalizeBlock(ctx context.Context, client abcicli.Client, txBytes [][]byte, codeExp []uint32, dataExp []byte, hashExp []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func PrepareProposal(ctx context.Context, client abcicli.Client, txBytes [][]byte, txExpected [][]byte, _ []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func ProcessProposal(ctx context.Context, client abcicli.Client, txBytes [][]byte, statusExp types.ResponseProcessProposal_ProposalStatus) error {
	_ = "STUB: not implemented"
	return nil
}

func CheckTx(ctx context.Context, client abcicli.Client, txBytes []byte, codeExp uint32, dataExp []byte) error {
	_ = "STUB: not implemented"
	return nil
}
