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

	"github.com/berachain/stargazer/eth/api"
	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/params"
	"github.com/berachain/stargazer/eth/rpc/config"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"

	errorslib "github.com/berachain/stargazer/lib/errors"
)

// Compile-time type assertion.
var _ Backend = (*backend)(nil)

// `backend` represents the backend for the JSON-RPC service.
type backend struct {
	chain     api.Chain
	rpcConfig *config.Server
}

// ==============================================================================
// Constructor
// ==============================================================================

// `NewBackend` returns a new `Backend` object.
func NewBackend(chain api.Chain, rpcConfig *config.Server) Backend {
	return &backend{
		chain:     chain,
		rpcConfig: rpcConfig,
	}
}

// ==============================================================================
// General Ethereum API
// ==============================================================================

func (b *backend) SyncProgress() ethereum.SyncProgress {
	// TODO: Implement your code here
	return ethereum.SyncProgress{
		CurrentBlock: 1000000,
		HighestBlock: 2000000,
	}
}

func (b *backend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	// TODO: Implement your code here
	return big.NewInt(1000000000), nil
}

func (b *backend) FeeHistory(ctx context.Context, blockCount int, lastBlock BlockNumber,
	rewardPercentiles []float64) (*big.Int, [][]*big.Int, []*big.Int, []float64, error) {
	// TODO: Implement your code here
	return big.NewInt(1000000000), nil, nil, nil, nil
}

func (b *backend) ChainDb() ethdb.Database { //nolint:stylecheck // conforms to interface.
	// TODO: is this implementable? (I don't think we need it tho tbh)
	panic("not implemented")
}

// `AccountManager` stargazer does not have an account manager.
func (b *backend) AccountManager() *accounts.Manager {
	panic("not implemented")
}

// `ExtRPCEnabled` returns whether the RPC endpoints are exposed over external
// interfaces.
func (b *backend) ExtRPCEnabled() bool {
	return b.rpcConfig.Address != "" || b.rpcConfig.WSAddress != ""
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

// `UnprotectedAllowed` returns whether unprotected transactions are alloweds
// For now, we have unprotected transactions fully disabled in stargazer.
func (b *backend) UnprotectedAllowed() bool {
	return false
}

// ==============================================================================
// Blockchain API
// ==============================================================================

// `SetHead` is used for state sync on ethereum, we leave state sync up to the chain
// implementing and thus it is not implemented in Stargazer.
func (b *backend) SetHead(number uint64) {
	panic("not implemented")
}

// `HeaderByNumber` returns the block header at the given block number.
func (b *backend) HeaderByNumber(ctx context.Context, number BlockNumber) (*types.Header, error) {
	block, err := b.BlockByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	return block.Header(), nil
}

// `HeaderByHash` returns the block header with the given hash.
func (b *backend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	block, err := b.BlockByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	return block.Header(), nil
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

// `CurrentHeader` returns the current header from the local chain.
func (b *backend) CurrentHeader() *types.Header {
	return b.chain.CurrentHeader().Header
}

// `CurrentBlock` returns the current block from the local chain.
func (b *backend) CurrentBlock() *types.Block {
	return b.chain.CurrentBlock().EthBlock()
}

// `BlockByNumber` returns the block identified by `number`.
func (b *backend) BlockByNumber(ctx context.Context, number BlockNumber) (*types.Block, error) {
	block := b.stargazerBlockByNumber(ctx, number)
	if block == nil {
		return nil, errorslib.Wrapf(ErrBlockNotFound, "number [%d]", number)
	}
	return block.EthBlock(), nil
}

// `BlockByHash` returns the block with the given hash.
func (b *backend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	block := b.stargazerBlockByHash(ctx, hash)
	if block == nil {
		return nil, errorslib.Wrapf(ErrBlockNotFound, "hash [%s]", hash.String())
	}
	return block.EthBlock(), nil
}

// `BlockByNumberOrHash` returns the block identified by `number` or `hash`.
func (b *backend) BlockByNumberOrHash(ctx context.Context,
	blockNrOrHash BlockNumberOrHash,
) (*types.Block, error) {
	block, err := b.stargazerBlockByNumberOrHash(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	return block.EthBlock(), nil
}

func (b *backend) StateAndHeaderByNumber(ctx context.Context,
	number BlockNumber) (*state.StateDB, *types.Header, error) {
	// TODO: Implement your code here
	return nil, nil, nil
}

func (b *backend) StateAndHeaderByNumberOrHash(ctx context.Context,
	blockNrOrHash BlockNumberOrHash) (*state.StateDB, *types.Header, error) {
	// TODO: Implement your code here
	return nil, nil, nil
}

// `PendingBlockAndReceipts` returns the current pending block and associated receipts.
func (b *backend) PendingBlockAndReceipts() (*types.Block, types.Receipts) {
	block := b.chain.CurrentBlock()
	return block.EthBlock(), block.GetReceipts()
}

// `GetReceipts` returns the receipts for the given block hash.
func (b *backend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	block := b.stargazerBlockByHash(ctx, hash)
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
) (*vm.EVM, func() error, error) {
	if vmConfig == nil {
		vmConfig = new(vm.Config)
	}
	txContext := core.NewEVMTxContext(msg)
	_ = txContext
	_ = vmConfig
	// TODO: finish
	return nil, nil, nil
}

func (b *backend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	// TODO: Implement your code here
	return nil
}

func (b *backend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	// TODO: Implement your code here
	return nil
}

func (b *backend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	// TODO: Implement your code here
	return nil
}

// ==============================================================================
// Transaction Pool API
// ==============================================================================

func (b *backend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	// TODO: Implement your code here
	return nil
}

func (b *backend) GetTransaction(ctx context.Context,
	txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	// TODO: Implement your code here
	// 1. Check the Mempool
	// 2. Check the Historical Storage
	return nil, common.Hash{}, 0, 0, nil
}

func (b *backend) GetPoolTransactions() (types.Transactions, error) {
	// TODO: Implement your code here
	return nil, nil
}

func (b *backend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	// TODO: Implement your code here
	return nil
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
	block, err := b.BlockByNumberOrHash(ctx, rpc.BlockNumberOrHash{BlockNumber: &number, BlockHash: &hash})
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
	block, err := b.stargazerBlockByNumberOrHash(ctx, BlockNumberOrHash{
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

// ==============================================================================
// Stargazer Helpers
// ==============================================================================

// TODO: consider actually using the context?

// `stargazerBlockByNumberOrHash` returns the block identified by `number` or `hash`.
func (b *backend) stargazerBlockByNumberOrHash(ctx context.Context,
	blockNrOrHash BlockNumberOrHash,
) (*types.StargazerBlock, error) {
	// First we try to get the block by number
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.stargazerBlockByNumber(ctx, blockNr), nil
	}
	if hash, ok := blockNrOrHash.Hash(); ok {
		_ = hash
		block := b.stargazerBlockByHash(ctx, hash)
		if block == nil {
			return nil, errorslib.Wrapf(ErrBlockNotFound, "hash [%s]", hash.String())
		}
		// if blockNrOrHash.RequireCanonical && b.eth.blockchain.GetCanonicalHash(header.Number.Uint64()) != hash {
		// 	return nil, errors.New("hash is not currently canonical")
		// }
		// block := b.eth.blockchain.GetBlock(hash, header.Number.Uint64())
		// if block == nil {
		// 	return nil, errors.New("header found, but block body is missing")
		// }
		// return block, nil
	}
	return nil, errors.New("invalid arguments; neither block nor hash specified")
}

// `stargazerBlockByNumber` returns the stargazer block identified by `number.
func (b *backend) stargazerBlockByNumber(_ context.Context, number BlockNumber) *types.StargazerBlock {
	//nolint:exhaustive // finish later.
	switch number {
	// Pending and Latest are the same since no Pow?
	case PendingBlockNumber:
	case LatestBlockNumber:
		// We just read the current processing block off the canonical
		// state processor.
		return b.chain.CurrentBlock()
	case FinalizedBlockNumber:
		// current block minus 1
		// return b.chain.FinalizedBlock().EthBlock()
	case SafeBlockNumber:
		// current block minus 1
		// return /b.chain.FinalizedBlock().EthBlock()
	case EarliestBlockNumber:
		// on mainnet, this doesn't even exist?
	}
	return b.chain.Host().GetBlockPlugin().GetStargazerBlockAtHeight(number.Int64())
}

func (b *backend) stargazerBlockByHash(_ context.Context, hash common.Hash) *types.StargazerBlock {
	return b.chain.Host().GetBlockPlugin().GetStargazerBlockByHash(hash)
}
