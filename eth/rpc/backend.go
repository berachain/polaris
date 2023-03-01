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
	IgnorePrice:      gasprice.DefaultIgnorePrice,
}

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
	b.logger.Warn("called eth.rpc.backend.SyncProgress", "sync_progress", "not implemented")
	return ethereum.SyncProgress{
		CurrentBlock: 0,
		HighestBlock: 0,
	}
}

// `SuggestGasTipCap` returns the recommended gas tip cap for a new transaction.
func (b *backend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	defer b.logger.Info("called eth.rpc.backend.SuggestGasTipCap", "suggested_tip_cap")
	return b.gpo.SuggestTipCap(ctx)
}

// `FeeHistory` returns the base fee and gas used history of the last N blocks.
func (b *backend) FeeHistory(ctx context.Context, blockCount int, lastBlock BlockNumber,
	rewardPercentiles []float64) (*big.Int, [][]*big.Int, []*big.Int, []float64, error) {
	b.logger.Info("called eth.rpc.backend.FeeHistory", "blockCount", blockCount,
		"lastBlock", lastBlock, "rewardPercentiles", rewardPercentiles)
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
	block, err := b.stargazerBlockByNumber(number)
	if err != nil {
		b.logger.Error("eth.rpc.backend.HeaderByNumber", "number", number, "err", err)
		return nil, err
	}
	b.logger.Info("called eth.rpc.backend.HeaderByNumber", "header", block.Header)
	return block.Header, nil
}

// `HeaderByHash` returns the block header with the given hash.
func (b *backend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	block, err := b.stargazerBlockByHash(hash)
	if err != nil {
		b.logger.Error("eth.rpc.backend.HeaderByHash", "hash", hash, "err", err)
		return nil, errorslib.Wrapf(ErrBlockNotFound, "HeaderByHash [%s]", hash.String())
	}
	b.logger.Info("eth.rpc.backend.HeaderByHash", "header", block.Header)
	return block.EthBlock().Header(), nil
}

// `HeaderByNumberOrHash` returns the header identified by `number` or `hash`.
func (b *backend) HeaderByNumberOrHash(ctx context.Context,
	blockNrOrHash BlockNumberOrHash,
) (*types.Header, error) {
	block, err := b.stargazerBlockByNumberOrHash(blockNrOrHash)
	if err != nil {
		b.logger.Error("eth.rpc.backend.HeaderByNumberOrHash", "blockNrOrHash", blockNrOrHash, "err", err)
		return nil, err
	}
	b.logger.Info("eth.rpc.backend.HeaderByNumberOrHash", "header", block.Header)
	return block.Header, nil
}

// `CurrentHeader` returns the current header from the local chain.s.
func (b *backend) CurrentHeader() *types.Header {
	block, err := b.chain.CurrentBlock()
	if err != nil {
		b.logger.Error("eth.rpc.backend.CurrentHeader", "block", block, "err", err)
		return nil
	}
	b.logger.Info("called eth.rpc.backend.CurrentHeader", "header", block.Header)
	return block.Header
}

// `CurrentBlock` returns the current block from the local chain.
func (b *backend) CurrentBlock() *types.Block {
	block, err := b.chain.CurrentBlock()
	if err != nil {
		b.logger.Error("eth.rpc.backend.CurrentBlock", "block", block, "err", err)
		return nil
	}
	b.logger.Info("called eth.rpc.backend.CurrentBlock", "header", block.Header,
		"num_txs", len(block.GetTransactions()))
	return block.EthBlock()
}

// `BlockByNumber` returns the block identified by `number`.
func (b *backend) BlockByNumber(ctx context.Context, number BlockNumber) (*types.Block, error) {
	block, err := b.stargazerBlockByNumber(number)
	if err != nil {
		b.logger.Error("eth.rpc.backend.BlockByNumber", "number", number, "err", err)
		return nil, errorslib.Wrapf(err, "BlockByNumber [%d]", number)
	}
	b.logger.Info("called eth.rpc.backend.BlockByNumber", "header", block.Header,
		"num_txs", len(block.GetTransactions()))
	return block.EthBlock(), nil
}

// `BlockByHash` returns the block with the given `hash`.
func (b *backend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	block, err := b.stargazerBlockByHash(hash)
	b.logger.Info("BlockByHash", "hash", hash, "block", block)
	if err != nil {
		b.logger.Error("eth.rpc.backend.BlockByHash", "hash", hash, "err", err)
		return nil, errorslib.Wrapf(err, "BlockByHash [%s]", hash.String())
	}
	b.logger.Info("called eth.rpc.backend.BlockByHash", "header", block.Header,
		"num_txs", len(block.GetTransactions()))
	return block.EthBlock(), nil
}

// `BlockByNumberOrHash` returns the block identified by `number` or `hash`.
func (b *backend) BlockByNumberOrHash(ctx context.Context,
	blockNrOrHash BlockNumberOrHash,
) (*types.Block, error) {
	block, err := b.stargazerBlockByNumberOrHash(blockNrOrHash)
	if err != nil {
		b.logger.Error("eth.rpc.backend.BlockByNumberOrHash", "blockNrOrHash", blockNrOrHash, "err", err)
		return nil, err
	}
	b.logger.Info("called eth.rpc.backend.BlockByNumberOrHash", "header", block.Header,
		"num_txs", len(block.GetTransactions()))
	return block.EthBlock(), nil
}

func (b *backend) StateAndHeaderByNumber(
	ctx context.Context, number BlockNumber,
) (vm.GethStateDB, *types.Header, error) {
	state, err := b.chain.GetStateByNumber(number.Int64())
	if err != nil {
		b.logger.Error("eth.rpc.backend.StateAndHeaderByNumber", "number", number, "err", err)
		return nil, nil, err
	}
	block, err := b.stargazerBlockByNumber(number)
	if err != nil {
		b.logger.Error("eth.rpc.backend.StateAndHeaderByNumber", "number", number, "err", err)
		return nil, nil, err
	}
	b.logger.Info("called eth.rpc.backend.StateAndHeaderByNumber", "header", block.Header,
		"num_txs", len(block.GetTransactions()))
	return state, block.Header, nil
}

func (b *backend) StateAndHeaderByNumberOrHash(
	ctx context.Context, blockNrOrHash BlockNumberOrHash,
) (vm.GethStateDB, *types.Header, error) {
	var err error
	var number int64
	var hash common.Hash
	var block *types.StargazerBlock
	if inputNum, ok := blockNrOrHash.Number(); ok {
		// Try to resolve by block number first.
		number = inputNum.Int64()
		block, err = b.stargazerBlockByNumber(inputNum)
		if err != nil {
			b.logger.Error("eth.rpc.backend.StateAndHeaderByNumberOrHash", "number", inputNum,
				"err", err)
			return nil, nil, err
		}
	} else if hash, ok = blockNrOrHash.Hash(); ok {
		// Try to resolve by hash next.
		block, err = b.stargazerBlockByHash(hash)
		if err != nil {
			b.logger.Error("eth.rpc.backend.StateAndHeaderByNumberOrHash", "hash", hash,
				"err", err)
			return nil, nil, err
		}
		number = block.Number.Int64()
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
	b.logger.Info("called eth.rpc.backend.StateAndHeaderByNumberOrHash", "header", block.Header,
		"num_txs", len(block.GetTransactions()))
	return state, block.Header, nil
}

// `PendingBlockAndReceipts` returns the current pending block and associated receipts.
func (b *backend) PendingBlockAndReceipts() (*types.Block, types.Receipts) {
	block, err := b.chain.CurrentBlock()
	if err != nil {
		b.logger.Error("eth.rpc.backend.PendingBlockAndReceipts", "err", err)
		return nil, nil
	}
	b.logger.Info("called eth.rpc.backend.PendingBlockAndReceipts", "header", block.Header,
		"num_receipts", len(block.GetReceipts()))
	return block.EthBlock(), block.GetReceipts()
}

// `GetReceipts` returns the receipts for the given block hash.
func (b *backend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	block, err := b.stargazerBlockByHash(hash)
	if err != nil {
		b.logger.Error("eth.rpc.backend.GetReceipts", "hash", hash, "err", err)
		return nil, err
	}
	b.logger.Info("called eth.rpc.backend.GetReceipts", "header", block.Header,
		"num_receipts", len(block.GetReceipts()))
	return block.GetReceipts(), nil
}

// `GetTd` returns the total difficulty of a block in the canonical chain.
// This is hardcoded to 69, as it is only applicable in a PoW chain.
func (b *backend) GetTd(ctx context.Context, hash common.Hash) *big.Int {
	b.logger.Info("called eth.rpc.backend.GetTd", "hash", hash)
	return new(big.Int).SetInt64(69)
}

// `GetEVM` returns a new EVM to be used for simulating a transaction, estimating gas etc.
func (b *backend) GetEVM(ctx context.Context, msg core.Message, state vm.GethStateDB,
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
	return b.chain.GetEVM(ctx, txContext, state, header, vmConfig), state.Error, nil
}

func (b *backend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	b.logger.Info("called eth.rpc.backend.SubscribeChainEvent", "ch", ch)
	panic("SubscribeChainEvent not implemented")
}

func (b *backend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	b.logger.Info("called eth.rpc.backend.SubscribeChainHeadEvent", "ch", ch)
	return b.chain.SubscribeChainHeadEvent(ch)
}

func (b *backend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	b.logger.Info("called eth.rpc.backend.SubscribeChainSideEvent", "ch", ch)
	panic("SubscribeChainSideEvent not implemented")
}

// ==============================================================================
// Transaction Pool API
// ==============================================================================

func (b *backend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.chain.Host().GetTxPoolPlugin().SendTx(signedTx)
}

func (b *backend) GetTransaction(
	ctx context.Context, txHash common.Hash,
) (*types.Transaction, common.Hash, uint64, uint64, error) {
	b.logger.Info("called eth.rpc.backend.GetTransaction", "tx_hash", txHash)
	return b.chain.GetTransaction(txHash)
}

func (b *backend) GetPoolTransactions() (types.Transactions, error) {
	b.logger.Info("called eth.rpc.backend.GetPoolTransactions")
	return b.chain.Host().GetTxPoolPlugin().GetAllTransactions()
}

func (b *backend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	b.logger.Info("called eth.rpc.backend.GetPoolTransaction", "tx_hash", txHash)
	return b.chain.Host().GetTxPoolPlugin().GetTransaction(txHash)
}

func (b *backend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	nonce, err := b.chain.Host().GetTxPoolPlugin().GetNonce(addr)
	defer b.logger.Info("called eth.rpc.backend.GetPoolNonce", "addr", addr, "nonce", nonce)
	return nonce, err
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
	b.logger.Info("called eth.rpc.backend.ChainConfig")
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
		b.logger.Error("eth.rpc.backend.GetBody", "number", number, "hash", hash)
		return nil, errors.New("invalid arguments; expect hash and no special block numbers")
	}
	block, err := b.stargazerBlockByNumberOrHash(BlockNumberOrHash{BlockNumber: &number, BlockHash: &hash})
	if err != nil {
		b.logger.Error("eth.rpc.backend.GetBody", "number", number, "hash", hash)
		return nil, err
	}
	b.logger.Info("called eth.rpc.backend.GetBody", "hash", hash, "number", number)
	return block.EthBlock().Body(), nil
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
		b.logger.Error("eth.rpc.backend.GetLogs", "number", number, "hash", blockHash)
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
func (b *backend) stargazerBlockByNumberOrHash(
	blockNrOrHash BlockNumberOrHash,
) (*types.StargazerBlock, error) {
	// First we try to get by hash.
	if hash, ok := blockNrOrHash.Hash(); ok {
		block, err := b.chain.GetStargazerBlockByHash(hash)
		if err != nil {
			return nil, errorslib.Wrapf(ErrBlockNotFound,
				"stargazerBlockByNumberOrHash: hash [%s]", hash.String())
		}

		// If the has is found, we have the canonical chain.
		if block.Hash() == hash {
			return block, nil
		}
		if blockNrOrHash.RequireCanonical {
			return nil, errorslib.Wrapf(ErrHashNotCanonical,
				"stargazerBlockByNumberOrHash: hash [%s]", hash.String())
		}
		// If not we try to query by number as a backup.
	}

	// Then we try to get the block by number
	if blockNr, ok := blockNrOrHash.Number(); ok {
		block, err := b.stargazerBlockByNumber(blockNr)
		if err != nil {
			return nil, errorslib.Wrapf(ErrBlockNotFound,
				"stargazerBlockByNumberOrHash: number [%d]", blockNr)
		}
		return block, nil
	}
	return nil, errors.New("invalid arguments; neither block nor hash specified")
}

// `stargazerBlockByHash` returns the stargazer block identified by `hash`.
func (b *backend) stargazerBlockByHash(hash common.Hash) (*types.StargazerBlock, error) {
	return b.chain.GetStargazerBlockByHash(hash)
}

// `stargazerBlockByNumber` returns the stargazer block identified by `number.
func (b *backend) stargazerBlockByNumber(number BlockNumber) (*types.StargazerBlock, error) {
	switch number { //nolint:nolintlint,exhaustive // golangci-lint bug?
	case SafeBlockNumber, FinalizedBlockNumber:
		return b.chain.FinalizedBlock()
	case PendingBlockNumber, LatestBlockNumber:
		return b.chain.CurrentBlock()
	default:
		// CONTRACT: GetStargazerBlockByNumber receives number >=0
		return b.chain.GetStargazerBlockByNumber(number.Int64())
	}
}
