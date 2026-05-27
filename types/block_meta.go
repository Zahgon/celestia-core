package types

import (
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

// BlockMeta contains meta information.
type BlockMeta struct {
	BlockID   BlockID `json:"block_id"`
	BlockSize int     `json:"block_size"`
	Header    Header  `json:"header"`
	NumTxs    int     `json:"num_txs"`
}

// NewBlockMeta returns a new BlockMeta.
func NewBlockMeta(block *Block, blockParts *PartSet) *BlockMeta {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck

func (bm *BlockMeta) ToProto() *cmtproto.BlockMeta { _ = "STUB: not implemented"; return nil }

func BlockMetaFromProto(pb *cmtproto.BlockMeta) (*BlockMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func BlockMetaFromTrustedProto(pb *cmtproto.BlockMeta) (*BlockMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateBasic performs basic validation.
func (bm *BlockMeta) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
