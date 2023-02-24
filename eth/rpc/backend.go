// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

//nolint:gomnd // TODO: fix
package rpc

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/stargazer/eth/api"
	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/common/hexutil"
	"pkg.berachain.dev/stargazer/eth/core"
	"pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	"pkg.berachain.dev/stargazer/eth/log"
	"pkg.berachain.dev/stargazer/eth/params"
	rpcapi "pkg.berachain.dev/stargazer/eth/rpc/api"
	"pkg.berachain.dev/stargazer/eth/rpc/config"
	errorslib "pkg.berachain.dev/stargazer/lib/errors"
)

var DefaultGasPriceOracleConfig = gasprice.Config{
	Blocks:           20,
	Percentile:       60,
	MaxHeaderHistory: 256,
	MaxBlockHistory:  256,
	Default:          big.NewInt(1000000000),
	MaxPrice:         big.NewInt(1000000000000000000),
}

// Compile-time type assertion.
var _ StargazerBackend = (*backend)(nil)

type StargazerBackend interface {
	Backend
	rpcapi.NetBackend
}

// `backend` represents the backend for the JSON-RPC service.
type backend struct {
	chain     api.Chain
	rpcConfig *config.Server
	gpo       *gasprice.Oracle
	logger    log.Logger
}

// ==============================================================================
// Constructor
// ==============================================================================

// `NewStargazerBackend` returns a new `Backend` object.
func NewStargazerBackend(chain api.Chain, rpcConfig *config.Server) StargazerBackend {
	b := &backend{
		// accountManager: accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: true}),
		chain:     chain,
		rpcConfig: rpcConfig,
		logger:    log.Root(),
	}
	b.gpo = gasprice.NewOracle(b, DefaultGasPriceOracleConfig)
	return b
}

// ==============================================================================
// General Ethereum API
// ==============================================================================

// `SyncProgress` returns the current progress of the sync algorithm.
func (b *backend) SyncProgress() ethereum.SyncProgress {
	// Consider implementing this in the future.
	return ethereum.SyncProgress{
		CurrentBlock: 0,
		HighestBlock: 0,
	}
}

// `SuggestGasTipCap` returns the recommended gas tip cap for a new transaction.
func (b *backend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestTipCap(ctx)
}

// `FeeHistory` returns the base fee and gas used history of the last N blocks.
func (b *backend) FeeHistory(ctx context.Context, blockCount int, lastBlock BlockNumber,
	rewardPercentiles []float64) (*big.Int, [][]*big.Int, []*big.Int, []float64, error) {
	return b.gpo.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
}

// `ChainDb` is unused in Stargazer.
func (b *backend) ChainDb() ethdb.Database { //nolint:stylecheck // conforms to interface.
	return ethdb.Database(nil)
}

// `AccountManager` is unused in Stargazer.
func (b *backend) AccountManager() *accounts.Manager {
	return nil
}

// `ExtRPCEnabled` returns whether the RPC endpoints are exposed over external
// interfaces.
func (b *backend) ExtRPCEnabled() bool {
	return b.rpcConfig.Enabled
}

// `RPCGasCap` returns the global gas cap for eth_call over rpc: this is
// if the user doesn't specify a cap.
func (b *backend) RPCGasCap() uint64 {
	return b.rpcConfig.RPCGasCap
}

// `RPCEVMTimeout` returns the global timeout for eth_call over rpc.
func (b *backend) RPCEVMTimeout() time.Duration {
	return b.rpcConfig.RPCEVMTimeout
}

// `RPCTxFeeCap` returns the global gas price cap for transactions over rpc.
func (b *backend) RPCTxFeeCap() float64 {
	return b.rpcConfig.RPCTxFeeCap
}

// `UnprotectedAllowed` returns whether unprotected transactions are alloweds.
// We will consider implementing these later, But our opinion is that
// there is no reason in 2023 not to use these.
func (b *backend) UnprotectedAllowed() bool {
	return false
}

// ==============================================================================
// Blockchain API
// ==============================================================================

// `SetHead` is used for state sync on ethereum, we leave state sync up to the host
// chain and thus it is not implemented in Stargazer.
func (b *backend) SetHead(number uint64) {
	panic("not implemented")
}

// `HeaderByNumber` returns the block header at the given block number.
func (b *backend) HeaderByNumber(ctx context.Context, number BlockNumber) (*types.Header, error) {
	block, err := b.BlockByNumber(ctx, number)
	b.logger.Info("HeaderByNumber", "block", block)
	if err != nil {
		return nil, err
	}
	return block.Header(), nil
}

// `HeaderByHash` returns the block header with the given hash.
func (b *backend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	b.logger.Info("HeaderByHash", "hash", hash)
	block := b.chain.GetStargazerBlockByHash(hash)
	if block == nil {
		return nil, ErrBlockNotFound
	}
	return block.EthBlock().Header(), nil
}

// `HeaderByNumberOrHash` returns the header identified by `number` or `hash`.
func (b *backend) HeaderByNumberOrHash(ctx context.Context,
	blockNrOrHash BlockNumberOrHash,
) (*types.Header, error) {
	block, err := b.BlockByNumberOrHash(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	return block.Header(), nil
}

// `CurrentHeader` returns the current header from the local chain.s.
func (b *backend) CurrentHeader() *types.Header {
	header := b.chain.CurrentHeader()
	b.logger.Info("CurrentHeader", "header", header)
	if header == nil {
		return nil
	}
	return header.Header
}

// `CurrentBlock` returns the current block from the local chain.
func (b *backend) CurrentBlock() *types.Block {
	block := b.chain.CurrentBlock()
	b.logger.Info("CurrentHeader", "block", block)
	if block == nil {
		return nil
	}
	return block.EthBlock()
}

// `BlockByNumber` returns the block identified by `number`.
func (b *backend) BlockByNumber(ctx context.Context, number BlockNumber) (*types.Block, error) {
	block := b.stargazerBlockByNumber(number)
	b.logger.Info("BlockByNumber", "block", block)
	if block == nil {
		return nil, errorslib.Wrapf(ErrBlockNotFound, "number [%d]", number)
	}
	return block.EthBlock(), nil
}

// `BlockByHash` returns the block with the given hash.
func (b *backend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	block := b.chain.GetStargazerBlockByHash(hash)
	b.logger.Info("BlockByHash", "hash", hash, "block", block)
	if block == nil {
		return nil, errorslib.Wrapf(ErrBlockNotFound, "hash [%s]", hash.String())
	}
	return block.EthBlock(), nil
}

// `BlockByNumberOrHash` returns the block identified by `number` or `hash`.
func (b *backend) BlockByNumberOrHash(ctx context.Context,
	blockNrOrHash BlockNumberOrHash,
) (*types.Block, error) {
	block, err := b.stargazerBlockByNumberOrHash(blockNrOrHash)
	if err != nil {
		return nil, err
	}
	b.logger.Info("BlockByNumberOrHash", "block", block)
	return block.EthBlock(), nil
}

func (b *backend) StateAndHeaderByNumber(ctx context.Context,
	number BlockNumber) (*state.StateDB, *types.Header, error) {
	panic("StateAndHeaderByNumber not implemented")
	// return nil, nil, nil
}

func (b *backend) StateAndHeaderByNumberOrHash(ctx context.Context,
	blockNrOrHash BlockNumberOrHash) (*state.StateDB, *types.Header, error) {
	panic("StateAndHeaderByNumberOrHash not implemented")
	// return nil, nil, nil
}

// `PendingBlockAndReceipts` returns the current pending block and associated receipts.
func (b *backend) PendingBlockAndReceipts() (*types.Block, types.Receipts) {
	block := b.chain.CurrentBlock()
	return block.EthBlock(), block.GetReceipts()
}

// `GetReceipts` returns the receipts for the given block hash.
func (b *backend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	block := b.chain.GetStargazerBlockByHash(hash)
	if block != nil {
		return nil, errorslib.Wrapf(ErrBlockNotFound, "hash [%s]", hash.String())
	}
	return block.GetReceipts(), nil
}

// `GetTd` returns the total difficulty of a block in the canonical chain.
// This is hardcoded to 0, as it is only applicable in a PoW chain.
func (b *backend) GetTd(ctx context.Context, hash common.Hash) *big.Int {
	return new(big.Int)
}

func (b *backend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB,
	header *types.Header, vmConfig *vm.Config,
) (*vm.GethEVM, func() error, error) {
	// if vmConfig == nil {
	// 	vmConfig = new(vm.Config)
	// }
	// txContext := core.NewEVMTxContext(msg)
	panic("GetEVM not implemented")
	// _ = txContext
	// _ = vmConfig
	// TODO: finish
}

func (b *backend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	panic("SubscribeChainEvent not implemented")
}

func (b *backend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	b.logger.Info("SubscribeChainHeadEvent")
	return b.chain.SubscribeChainHeadEvent(ch)
}

func (b *backend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	panic("SubscribeChainSideEvent not implemented")
}

// ==============================================================================
// Transaction Pool API
// ==============================================================================

func (b *backend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	b.logger.Info("SendTx", signedTx.Hash())
	return b.chain.Host().GetTxPoolPlugin().SendTx(signedTx)
}

func (b *backend) GetTransaction(ctx context.Context,
	txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	// 1. Check the Mempool
	tx := b.GetPoolTransaction(txHash)
	if tx != nil {
		// todo get other info
		return tx, common.Hash{}, 0, 0, nil
	}
	// 2. Check the Historical Storage
	// tx := b.chain.Host().GetBlockPlugin().GetTransactionByHash(txHash)
	return nil, common.Hash{}, 0, 0, nil
}

func (b *backend) GetPoolTransactions() (types.Transactions, error) {
	return b.chain.Host().GetTxPoolPlugin().GetAllTransactions()
}

func (b *backend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	return b.chain.Host().GetTxPoolPlugin().GetTransaction(txHash)
}

func (b *backend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	// TODO: Implement your code here
	return 0, nil
}

func (b *backend) Stats() (int, int) {
	pending := 0
	queued := 0
	// TODO: Implement your code here
	return pending, queued
}

func (b *backend) TxPoolContent() (map[common.Address]types.Transactions,
	map[common.Address]types.Transactions) {
	// TODO: Implement your code here
	return nil, nil
}

func (b *backend) TxPoolContentFrom(addr common.Address,
) (types.Transactions, types.Transactions) {
	// TODO: Implement your code here
	return nil, nil
}

func (b *backend) SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription {
	// TODO: Implement your code here
	return nil
}

// `ChainConfig` returns the chain configuration.
func (b *backend) ChainConfig() *params.ChainConfig {
	return b.chain.Host().GetConfigurationPlugin().ChainConfig()
}

func (b *backend) Engine() consensus.Engine {
	panic("not implemented")
}

// `GetBody retrieves the block body corresponding to block by has or number.`.
func (b *backend) GetBody(ctx context.Context, hash common.Hash,
	number BlockNumber,
) (*types.Body, error) {
	if number < 0 || hash == (common.Hash{}) {
		return nil, errors.New("invalid arguments; expect hash and no special block numbers")
	}
	block, err := b.BlockByNumberOrHash(ctx, BlockNumberOrHash{BlockNumber: &number, BlockHash: &hash})
	if err != nil {
		return nil, err
	}
	return block.Body(), nil
}

// `GetLogs` returns the logs for the given block hash or number.
func (b *backend) GetLogs(ctx context.Context, blockHash common.Hash,
	number uint64,
) ([][]*types.Log, error) {
	bn := BlockNumber(number)
	block, err := b.stargazerBlockByNumberOrHash(BlockNumberOrHash{
		BlockNumber: &bn,
		BlockHash:   &blockHash,
	})
	if err != nil {
		return nil, err
	}
	receipts := block.GetReceipts()
	buf := make([][]*types.Log, len(receipts))
	for i, receipt := range receipts {
		buf[i] = receipt.Logs
	}
	return buf, nil
}

func (b *backend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	// TODO: Implement your code here
	return nil
}

func (b *backend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	// TODO: Implement your code here
	return nil
}

func (b *backend) SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription {
	// TODO: Implement your code here
	return nil
}

func (b *backend) BloomStatus() (uint64, uint64) {
	// TODO: Implement your code here
	return 0, 0
}

func (b *backend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	// TODO: Implement your code here
}

func (b *backend) Version() string {
	// TODO: Implement your code here
	return "1.0" // get from comet
}

func (b *backend) Listening() bool {
	// TODO: Implement your code here
	return true
}

func (b *backend) PeerCount() hexutil.Uint {
	// TODO: Implement your code here
	return 1
}

// ==============================================================================
// Stargazer Helpers
// ==============================================================================

// `stargazerBlockByNumberOrHash` returns the block identified by `number` or `hash`.
func (b *backend) stargazerBlockByNumberOrHash(blockNrOrHash BlockNumberOrHash) (*types.StargazerBlock, error) {
	// First we try to get by hash.
	if hash, ok := blockNrOrHash.Hash(); ok {
		block := b.chain.GetStargazerBlockByHash(hash)
		if block == nil {
			return nil, errorslib.Wrapf(ErrBlockNotFound, "hash [%s]", hash.String())
		}

		// If the has is found, we have the canonical chain.
		if block.Hash() == hash {
			return block, nil
		}
		if blockNrOrHash.RequireCanonical {
			return nil, errorslib.Wrapf(ErrHashNotCanonical, "hash [%s]", hash.String())
		}
		// If not we try to query by number as a backup.
	}

	// Then we try to get the block by number
	if blockNr, ok := blockNrOrHash.Number(); ok {
		block := b.stargazerBlockByNumber(blockNr)
		if block == nil {
			return nil, errorslib.Wrapf(ErrBlockNotFound, "number [%d]", blockNr)
		}
	}
	return nil, errors.New("invalid arguments; neither block nor hash specified")
}

// `stargazerBlockByNumber` returns the stargazer block identified by `number.
func (b *backend) stargazerBlockByNumber(number BlockNumber) *types.StargazerBlock {
	//nolint:exhaustive // finish later.
	switch number {
	// Pending and latest are the same in stargazer.
	case PendingBlockNumber:
	case LatestBlockNumber:
		return b.chain.CurrentBlock()
	case FinalizedBlockNumber:
	case SafeBlockNumber:
		return b.chain.FinalizedBlock()
	case EarliestBlockNumber:
		// no-op, since we are querying block 0, which is done below.
	}
	return b.chain.GetStargazerBlockByNumber(number.Int64())
}
