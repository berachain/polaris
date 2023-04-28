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
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/node"

	"pkg.berachain.dev/polaris/eth/api"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/common/hexutil"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/log"
	"pkg.berachain.dev/polaris/eth/params"
	rpcapi "pkg.berachain.dev/polaris/eth/rpc/api"
	errorslib "pkg.berachain.dev/polaris/lib/errors"
	"pkg.berachain.dev/polaris/lib/utils"
)

// PolarisBackend represents the backend object for a Polaris chain. It extends the standard
// go-ethereum backend object.
type PolarisBackend interface {
	Backend
	rpcapi.NetBackend
}

// backend represents the backend for the JSON-RPC service.
type backend struct {
	chain      api.Chain
	rpcConfig  *Config
	nodeConfig *node.Config
	gpo        *gasprice.Oracle
	logger     log.Logger
}

// ==============================================================================
// Constructor
// ==============================================================================

// NewPolarisBackend returns a new `Backend` object.
func NewPolarisBackend(
	chain api.Chain,
	rpcConfig *Config,
	nodeConfig *node.Config,
) PolarisBackend {
	b := &backend{
		chain:     chain,
		rpcConfig: rpcConfig,
		logger:    log.Root(),
	}
	b.gpo = gasprice.NewOracle(b, rpcConfig.GPO)
	return b
}

// ==============================================================================
// General Ethereum API
// ==============================================================================

// SyncProgress returns the current progress of the sync algorithm.
func (b *backend) SyncProgress() ethereum.SyncProgress {
	// TODO: Consider implementing this in the future.
	b.logger.Warn("called eth.rpc.backend.SyncProgress", "sync_progress", "not implemented")
	return ethereum.SyncProgress{
		CurrentBlock: 0,
		HighestBlock: 0,
	}
}

// SuggestGasTipCap returns the recommended gas tip cap for a new transaction.
func (b *backend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	defer b.logger.Info("called eth.rpc.backend.SuggestGasTipCap", "suggested_tip_cap")
	return b.gpo.SuggestTipCap(ctx)
}

// FeeHistory returns the base fee and gas used history of the last N blocks.
func (b *backend) FeeHistory(ctx context.Context, blockCount int, lastBlock BlockNumber,
	rewardPercentiles []float64) (*big.Int, [][]*big.Int, []*big.Int, []float64, error) {
	b.logger.Info("called eth.rpc.backend.FeeHistory", "blockCount", blockCount,
		"lastBlock", lastBlock, "rewardPercentiles", rewardPercentiles)
	return b.gpo.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
}

// ChainDb is unused in Polaris.
func (b *backend) ChainDb() ethdb.Database { //nolint:stylecheck // conforms to interface.
	return ethdb.Database(nil)
}

// AccountManager is unused in Polaris.
func (b *backend) AccountManager() *accounts.Manager {
	return nil
}

// ExtRPCEnabled returns whether the RPC endpoints are exposed over external
// interfaces.
func (b *backend) ExtRPCEnabled() bool {
	return b.nodeConfig.ExtRPCEnabled()
}

// RPCGasCap returns the global gas cap for eth_call over rpc: this is
// if the user doesn't specify a cap.
func (b *backend) RPCGasCap() uint64 {
	return b.rpcConfig.RPCGasCap
}

// RPCEVMTimeout returns the global timeout for eth_call over rpc.
func (b *backend) RPCEVMTimeout() time.Duration {
	return b.rpcConfig.RPCEVMTimeout
}

// RPCTxFeeCap returns the global gas price cap for transactions over rpc.
func (b *backend) RPCTxFeeCap() float64 {
	return b.rpcConfig.RPCTxFeeCap
}

// UnprotectedAllowed returns whether unprotected transactions are alloweds.
// We will consider implementing these later, But our opinion is that
// there is no reason in 2023 not to use these.
func (b *backend) UnprotectedAllowed() bool {
	return false
}

// ==============================================================================
// Blockchain API
// ==============================================================================

// SetHead is used for state sync on ethereum, we leave state sync up to the host
// chain and thus it is not implemented in Polaris.
func (b *backend) SetHead(number uint64) {
	panic("not implemented")
}

// HeaderByNumber returns the block header at the given block number.
func (b *backend) HeaderByNumber(_ context.Context, number BlockNumber) (*types.Header, error) {
	block, err := b.blockByNumber(number)
	if err != nil {
		b.logger.Error("eth.rpc.backend.HeaderByNumber", "number", number, "err", err)
		return nil, err
	}
	b.logger.Info("called eth.rpc.backend.HeaderByNumber", "header", block.Header())
	return block.Header(), nil
}

// HeaderByHash returns the block header with the given hash.
func (b *backend) HeaderByHash(_ context.Context, hash common.Hash) (*types.Header, error) {
	block, err := b.blockByHash(hash)
	if err != nil {
		b.logger.Error("eth.rpc.backend.HeaderByHash", "hash", hash, "err", err)
		return nil, errorslib.Wrapf(ErrBlockNotFound, "HeaderByHash [%s]", hash.String())
	}
	b.logger.Info("eth.rpc.backend.HeaderByHash", "header", block.Header())
	return block.Header(), nil
}

// HeaderByNumberOrHash returns the header identified by `number` or `hash`.
func (b *backend) HeaderByNumberOrHash(_ context.Context,
	blockNrOrHash BlockNumberOrHash,
) (*types.Header, error) {
	block, err := b.blockByNumberOrHash(blockNrOrHash)
	if err != nil {
		b.logger.Error("eth.rpc.backend.HeaderByNumberOrHash", "blockNrOrHash", blockNrOrHash, "err", err)
		return nil, err
	}
	b.logger.Info("eth.rpc.backend.HeaderByNumberOrHash", "header", block.Header())
	return block.Header(), nil
}

// CurrentHeader returns the current header from the local chains.
func (b *backend) CurrentHeader() *types.Header {
	block, err := b.chain.CurrentBlock()
	if err != nil {
		b.logger.Error("eth.rpc.backend.CurrentHeader", "block", block, "err", err)
		return nil
	}
	header := block.Header()
	b.logger.Info("called eth.rpc.backend.CurrentHeader", "header", header)
	return header
}

// CurrentBlock returns the current block from the local chain.
func (b *backend) CurrentBlock() *types.Header {
	block, err := b.chain.CurrentBlock()
	if err != nil {
		b.logger.Error("eth.rpc.backend.CurrentBlock", "block", block, "err", err)
		return nil
	}
	b.logger.Info("called eth.rpc.backend.CurrentBlock", "block", block)
	return block.Header()
}

// BlockByNumber returns the block identified by `number`.
func (b *backend) BlockByNumber(_ context.Context, number BlockNumber) (*types.Block, error) {
	block, err := b.blockByNumber(number)
	if err != nil {
		b.logger.Error("eth.rpc.backend.BlockByNumber", "number", number, "err", err)
		return nil, errorslib.Wrapf(err, "BlockByNumber [%d]", number)
	}
	b.logger.Info("called eth.rpc.backend.BlockByNumber", "header", block.Header(),
		"num_txs", len(block.Transactions()))
	return block, nil
}

// BlockByHash returns the block with the given `hash`.
func (b *backend) BlockByHash(_ context.Context, hash common.Hash) (*types.Block, error) {
	block, err := b.blockByHash(hash)
	b.logger.Info("BlockByHash", "hash", hash, "block", block)
	if err != nil {
		b.logger.Error("eth.rpc.backend.BlockByHash", "hash", hash, "err", err)
		return nil, errorslib.Wrapf(err, "BlockByHash [%s]", hash.String())
	}
	b.logger.Info("called eth.rpc.backend.BlockByHash", "header", block.Header(),
		"num_txs", len(block.Transactions()))
	return block, nil
}

// BlockByNumberOrHash returns the block identified by `number` or `hash`.
func (b *backend) BlockByNumberOrHash(_ context.Context,
	blockNrOrHash BlockNumberOrHash,
) (*types.Block, error) {
	block, err := b.blockByNumberOrHash(blockNrOrHash)
	if err != nil {
		b.logger.Error("eth.rpc.backend.BlockByNumberOrHash", "blockNrOrHash", blockNrOrHash, "err", err)
		return nil, err
	}
	b.logger.Info("called eth.rpc.backend.BlockByNumberOrHash", "header", block.Header(),
		"num_txs", len(block.Transactions()))
	return block, nil
}

func (b *backend) StateAndHeaderByNumber(
	_ context.Context, number BlockNumber,
) (vm.GethStateDB, *types.Header, error) {
	state, err := b.chain.GetStateByNumber(number.Int64())
	if err != nil {
		b.logger.Error("eth.rpc.backend.StateAndHeaderByNumber", "number", number, "err", err)
		return nil, nil, err
	}
	block, err := b.blockByNumber(number)
	if err != nil {
		b.logger.Error("eth.rpc.backend.StateAndHeaderByNumber", "number", number, "err", err)
		return nil, nil, err
	}
	b.logger.Info("called eth.rpc.backend.StateAndHeaderByNumber", "header", block.Header(),
		"num_txs", len(block.Transactions()))
	return state, block.Header(), nil
}

func (b *backend) StateAndHeaderByNumberOrHash(
	_ context.Context, blockNrOrHash BlockNumberOrHash,
) (vm.GethStateDB, *types.Header, error) {
	var err error
	var number int64
	var hash common.Hash
	var block *types.Block
	if inputNum, ok := blockNrOrHash.Number(); ok {
		// Try to resolve by block number first.
		number = inputNum.Int64()
		block, err = b.blockByNumber(inputNum)
		if err != nil {
			b.logger.Error("eth.rpc.backend.StateAndHeaderByNumberOrHash", "number", inputNum,
				"err", err)
			return nil, nil, err
		}
	} else if hash, ok = blockNrOrHash.Hash(); ok {
		// Try to resolve by hash next.
		block, err = b.blockByHash(hash)
		if err != nil {
			b.logger.Error("eth.rpc.backend.StateAndHeaderByNumberOrHash", "hash", hash,
				"err", err)
			return nil, nil, err
		}
		number = block.Number().Int64()
	} else {
		return nil, nil, errors.New("invalid arguments; neither block nor hash specified")
	}

	// Now that we have a number, we can load up a statedb at the derived block number.
	state, err := b.chain.GetStateByNumber(number)
	if err != nil {
		b.logger.Error("eth.rpc.backend.StateAndHeaderByNumberOrHash", "number", number,
			"err", err)
		return nil, nil, err
	}
	b.logger.Info("called eth.rpc.backend.StateAndHeaderByNumberOrHash", "header", block.Header(),
		"num_txs", len(block.Transactions()))
	return state, block.Header(), nil
}

// PendingBlockAndReceipts returns the pending block (equivalent to current block in Polaris)
// and associated receipts.
func (b *backend) PendingBlockAndReceipts() (*types.Block, types.Receipts) {
	block, receipts, err := b.chain.CurrentBlockAndReceipts()
	if err != nil {
		b.logger.Error("eth.rpc.backend.PendingBlockAndReceipts", "err", err)
		return nil, nil
	}
	b.logger.Info("called eth.rpc.backend.PendingBlockAndReceipts", "block", block,
		"num_receipts", len(receipts))
	return block, receipts
}

// GetReceipts returns the receipts for the given block hash.
func (b *backend) GetReceipts(_ context.Context, bhash common.Hash) (types.Receipts, error) {
	receipts, err := b.chain.GetReceipts(bhash)
	if err != nil {
		b.logger.Error("eth.rpc.backend.GetReceipts", "block_hash", bhash, "err", err)
		return nil, err
	}
	b.logger.Info("called eth.rpc.backend.GetReceipts", "block_hash", bhash,
		"num_receipts", len(receipts))
	return receipts, nil
}

// GetTd returns the total difficulty of a block in the canonical chain.
// This is hardcoded to 69, as it is only applicable in a PoW chain.
func (b *backend) GetTd(_ context.Context, hash common.Hash) *big.Int {
	b.logger.Info("called eth.rpc.backend.GetTd", "hash", hash)
	return new(big.Int).SetInt64(69)
}

// GetEVM returns a new EVM to be used for simulating a transaction, estimating gas etc.
func (b *backend) GetEVM(ctx context.Context, msg *core.Message, state vm.GethStateDB,
	header *types.Header, vmConfig *vm.Config,
) (*vm.GethEVM, func() error, error) {
	if vmConfig == nil {
		b.logger.Info("eth.rpc.backend.GetEVM", "vmConfig", "nil")
		vmConfig = new(vm.Config)
	}
	if msg == nil {
		b.logger.Error("eth.rpc.backend.GetEVM", "msg", "nil")
		return nil, nil, errors.New("msg is nil")
	}
	txContext := core.NewEVMTxContext(msg)
	b.logger.Info("called eth.rpc.backend.GetEVM", "header", header, "txContext", txContext, "vmConfig", vmConfig)
	gethEVM := b.chain.GetEVM(ctx, txContext, utils.MustGetAs[vm.PolarisStateDB](state), header, vmConfig)
	return gethEVM, state.Error, nil
}

func (b *backend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	b.logger.Info("called eth.rpc.backend.SubscribeChainEvent", "ch", ch)
	return b.chain.SubscribeChainEvent(ch)
}

func (b *backend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	b.logger.Info("called eth.rpc.backend.SubscribeChainHeadEvent", "ch", ch)
	return b.chain.SubscribeChainHeadEvent(ch)
}

func (b *backend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	b.logger.Info("called eth.rpc.backend.SubscribeChainSideEvent", "ch", ch)
	return b.chain.SubscribeChainSideEvent(ch)
}

// ==============================================================================
// Transaction Pool API
// ==============================================================================

func (b *backend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.chain.SendTx(ctx, signedTx)
}

func (b *backend) GetTransaction(
	_ context.Context, txHash common.Hash,
) (*types.Transaction, common.Hash, uint64, uint64, error) {
	b.logger.Info("called eth.rpc.backend.GetTransaction", "tx_hash", txHash)
	return b.chain.GetTransaction(txHash)
}

func (b *backend) GetPoolTransactions() (types.Transactions, error) {
	b.logger.Info("called eth.rpc.backend.GetPoolTransactions")
	return b.chain.GetPoolTransactions()
}

func (b *backend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	b.logger.Info("called eth.rpc.backend.GetPoolTransaction", "tx_hash", txHash)
	return b.chain.GetPoolTransaction(txHash)
}

func (b *backend) GetPoolNonce(_ context.Context, addr common.Address) (uint64, error) {
	nonce, err := b.chain.GetPoolNonce(addr)
	b.logger.Info("called eth.rpc.backend.GetPoolNonce", "addr", addr, "nonce", nonce)
	return nonce, err
}

func (b *backend) Stats() (int, int) {
	pending, queued := b.chain.GetPoolStats()
	b.logger.Info("called eth.rpc.backend.Stats", "pending", pending, "queued", queued)
	return pending, queued
}

func (b *backend) TxPoolContent() (
	map[common.Address]types.Transactions, map[common.Address]types.Transactions,
) {
	pending, queued := b.chain.GetPoolContent()
	b.logger.Info("called eth.rpc.backend.TxPoolContent", "pending", len(pending), "queued", len(queued))
	return pending, queued
}

func (b *backend) TxPoolContentFrom(addr common.Address) (
	types.Transactions, types.Transactions,
) {
	pending, queued := b.chain.GetPoolContentFrom(addr)
	b.logger.Info("called eth.rpc.backend.TxPoolContentFrom", "addr", addr, "pending", len(pending), "queued", len(queued))
	return pending, queued
}

func (b *backend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return b.chain.SubscribeNewTxsEvent(ch)
}

// ChainConfig returns the chain configuration.
func (b *backend) ChainConfig() *params.ChainConfig {
	b.logger.Info("called eth.rpc.backend.ChainConfig")
	return b.chain.ChainConfig()
}

func (b *backend) Engine() consensus.Engine {
	panic("not implemented")
}

// GetBody retrieves the block body corresponding to block by has or number..
func (b *backend) GetBody(_ context.Context, hash common.Hash,
	number BlockNumber,
) (*types.Body, error) {
	if number < 0 || hash == (common.Hash{}) {
		b.logger.Error("eth.rpc.backend.GetBody", "number", number, "hash", hash)
		return nil, errors.New("invalid arguments; expect hash and no special block numbers")
	}
	block, err := b.blockByNumberOrHash(BlockNumberOrHash{BlockNumber: &number, BlockHash: &hash})
	if err != nil {
		b.logger.Error("eth.rpc.backend.GetBody", "number", number, "hash", hash)
		return nil, err
	}
	b.logger.Info("called eth.rpc.backend.GetBody", "hash", hash, "number", number)
	return block.Body(), nil
}

// GetLogs returns the logs for the given block hash or number.
func (b *backend) GetLogs(
	_ context.Context, blockHash common.Hash, number uint64,
) ([][]*types.Log, error) {
	receipts, err := b.chain.GetReceipts(blockHash)
	if err != nil {
		b.logger.Error("eth.rpc.backend.GetLogs", "block_hash", blockHash, "err", err)
		return nil, err
	}
	logs := make([][]*types.Log, len(receipts))
	for i, receipt := range receipts {
		logs[i] = receipt.Logs
	}
	b.logger.Info("called eth.rpc.backend.GetBody", "block_hash", blockHash, "number", number)
	return logs, nil
}

func (b *backend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	b.logger.Info("called eth.rpc.backend.SubscribeRemovedLogsEvent", "ch", ch)
	return b.chain.SubscribeRemovedLogsEvent(ch)
}

func (b *backend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	b.logger.Info("called eth.rpc.backend.SubscribeLogsEvent", "ch", ch)
	return b.chain.SubscribeLogsEvent(ch)
}

func (b *backend) SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.chain.SubscribePendingLogsEvent(ch)
}

func (b *backend) BloomStatus() (uint64, uint64) {
	// TODO: Implement your code here
	return params.BloomBitsBlocks, 0
}

func (b *backend) ServiceFilter(_ context.Context, session *bloombits.MatcherSession) {
	// TODO: Implement your code here
}

// Version returns the current chain protocol version.
// For education:
// https://medium.com/@pedrouid/chainid-vs-networkid-how-do-they-differ-on-ethereum-eec2ed41635b
func (b *backend) Version() string {
	return b.ChainConfig().ChainID.String()
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
// Polaris Helpers
// ==============================================================================

// blockByNumberOrHash returns the block identified by `number` or `hash`.
func (b *backend) blockByNumberOrHash(
	blockNrOrHash BlockNumberOrHash,
) (*types.Block, error) {
	// First we try to get by hash.
	if hash, ok := blockNrOrHash.Hash(); ok {
		block, err := b.chain.GetBlockByHash(hash)
		if err != nil {
			return nil, errorslib.Wrapf(ErrBlockNotFound,
				"blockByNumberOrHash: hash [%s]", hash.String())
		}

		// If the has is found, we have the canonical chain.
		if block.Hash() == hash {
			return block, nil
		}
		if blockNrOrHash.RequireCanonical {
			return nil, errorslib.Wrapf(ErrHashNotCanonical,
				"blockByNumberOrHash: hash [%s]", hash.String())
		}
		// If not we try to query by number as a backup.
	}

	// Then we try to get the block by number
	if blockNr, ok := blockNrOrHash.Number(); ok {
		block, err := b.blockByNumber(blockNr)
		if err != nil {
			return nil, errorslib.Wrapf(ErrBlockNotFound,
				"blockByNumberOrHash: number [%d]", blockNr)
		}
		return block, nil
	}
	return nil, errors.New("invalid arguments; neither block nor hash specified")
}

// blockByHash returns the polaris block identified by `hash`.
func (b *backend) blockByHash(hash common.Hash) (*types.Block, error) {
	return b.chain.GetBlockByHash(hash)
}

// blockByNumber returns the polaris block identified by `number.
func (b *backend) blockByNumber(number BlockNumber) (*types.Block, error) {
	switch number { //nolint:nolintlint,exhaustive // golangci-lint bug?
	case SafeBlockNumber, FinalizedBlockNumber:
		return b.chain.FinalizedBlock()
	case PendingBlockNumber, LatestBlockNumber:
		return b.chain.CurrentBlock()
	default:
		// CONTRACT: GetBlockByNumber receives number >=0
		return b.chain.GetBlockByNumber(number.Int64())
	}
}
