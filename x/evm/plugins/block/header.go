package block

import (
	"errors"
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"pkg.berachain.dev/stargazer/eth/common"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/rpc"
)

var SGHeaderKey = []byte("header")

// `SetQueryContextFn` sets the query context func for the plugin.
func (p *plugin) SetQueryContextFn(gqc func(height int64, prove bool) (sdk.Context, error)) {
	p.getQueryContext = gqc
}

// `ProcessHeader` takes in the header and process it using the `ctx` and stores it in the context store.
func (p *plugin) ProcessHeader(ctx sdk.Context, header *coretypes.StargazerHeader) error {
	header = p.PrepareHeader(ctx, header)
	bz, err := header.MarshalBinary()
	if err != nil {
		return err
	}
	ctx.KVStore(p.storekey).Set(SGHeaderKey, bz)
	return nil
}

// `GetStargazerHeaderByNumber` returns the stargazer header for the given block number.
func (p *plugin) GetStargazerHeaderByNumber(number int64) (*coretypes.StargazerHeader, error) {
	if p.getQueryContext == nil {
		return nil, errors.New("query context func not set")
	}

	// Handle the special cases of the block number.
	var iavlHeight int64
	switch rpc.BlockNumber(number) {
	case rpc.SafeBlockNumber:
	case rpc.FinalizedBlockNumber:
		iavlHeight = p.ctx.BlockHeight() - 1
	case rpc.PendingBlockNumber:
	case rpc.LatestBlockNumber:
		iavlHeight = p.ctx.BlockHeight()
	case rpc.EarliestBlockNumber:
		iavlHeight = 0
	default:
		iavlHeight = number
	}
	fmt.Println("iavlHeight", iavlHeight)
	// Get the query context for the given block number.
	ctx, err := p.getQueryContext(iavlHeight, false)
	if err != nil {
		return nil, err
	}

	// Get the stargazer header from the query context.
	var header coretypes.StargazerHeader
	bz := ctx.KVStore(p.storekey).Get(SGHeaderKey)
	if bz == nil {
		return nil, errors.New("stargazer header not found")
	}
	if err := header.UnmarshalBinary(bz); err != nil {
		return nil, err
	}
	return &header, nil
}

// `ProcessHeader` takes in a `coretypes.StargazerHeader` and returns a `coretypes.StargazerHeader` with the
// Fields set to the correct values.
func (p *plugin) PrepareHeader(ctx sdk.Context, header *coretypes.StargazerHeader) *coretypes.StargazerHeader {
	cometHeader := ctx.BlockHeader()

	// We retrieve the `TxHash` from the `DataHash` field of the `sdk.Context` opposed to deriving it
	// from solely the ethereum transaction information.
	txHash := coretypes.EmptyRootHash
	if len(cometHeader.DataHash) == 0 {
		txHash = common.BytesToHash(cometHeader.DataHash)
	}

	parentHash := common.Hash{}
	if ctx.BlockHeight() > 1 {
		header, err := p.GetStargazerHeaderByNumber(ctx.BlockHeight() - 1)
		if err != nil || header == nil {
			panic("failed to get parent stargazer header")
		}
		parentHash = header.Hash()
	}

	return coretypes.NewStargazerHeader(
		&coretypes.Header{
			// `ParentHash` is set to the hash of the previous block.
			ParentHash: parentHash,
			// `UncleHash` is set empty as CometBFT does not have uncles.
			UncleHash: coretypes.EmptyUncleHash,
			// TODO: Use staking keeper to get the operator address.
			Coinbase: common.BytesToAddress(cometHeader.ProposerAddress),
			// `Root` is set to the hash of the state after the transactions are applied.
			Root: common.BytesToHash(cometHeader.AppHash),
			// `TxHash` is set to the hash of the transactions in the block. We take the
			// `DataHash` from the `sdk.Context` opposed to using DeriveSha on the StargazerBlock,
			// in order to include non-evm transactions block hash.
			TxHash: txHash,
			// We simply map the cosmos "BlockHeight" to the ethereum "BlockNumber".
			Number: big.NewInt(cometHeader.Height),
			// `GasLimit` is set to the block gas limit.
			GasLimit: blockGasLimitFromCosmosContext(p.ctx),
			// `Time` is set to the block timestamp.
			Time: uint64(cometHeader.Time.UTC().Unix()),
			// `BaseFee` is set to the block base fee.
			BaseFee: big.NewInt(int64(p.BaseFee())),
			// `ReceiptHash` set to empty. It is filled during `Finalize` in the StateProcessor.
			ReceiptHash: common.Hash{},
			// `Bloom` is set to empty. It is filled during `Finalize` in the StateProcessor.
			Bloom: coretypes.Bloom{},
			// `GasUsed` is set to 0. It is filled during `Finalize` in the StateProcessor.
			GasUsed: 0,
			// `Difficulty` is set to 0 as it is only used in PoW consensus.
			Difficulty: big.NewInt(0),
			// `MixDigest` is set empty as it is only used in PoW consensus.
			MixDigest: common.Hash{},
			// `Nonce` is set empty as it is only used in PoW consensus.
			Nonce: coretypes.BlockNonce{},
			// `Extra` is unused in Stargazer.
			Extra: []byte(nil),
		},
		blockHashFromCosmosContext(ctx),
	)
}
