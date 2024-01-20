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

package polar

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	pcore "github.com/berachain/polaris/eth/core"
	"github.com/berachain/polaris/eth/core/state"
	polarapi "github.com/berachain/polaris/eth/polar/api"
	"github.com/berachain/polaris/eth/version"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/bloombits"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/ethapi"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

// Backend represents the backend object for a Polaris chain. It extends the standard
// go-ethereum backend object.
type (
	APIBackend interface {
		ethapi.Backend
		polarapi.NetBackend
		polarapi.Web3Backend
		tracers.Backend
	}

	// SyncStatusProvider defines methods that allow the chain to have insight into the underlying
	// consensus engine of the host chain.
	SyncStatusProvider interface {
		// SyncProgress returns the current sync progress of the host chain.
		SyncProgress(ctx context.Context) (ethereum.SyncProgress, error)
		// IsListening returns whether or not the host chain is listening for new blocks.
		Listening(ctx context.Context) (bool, error)
		// PeerCount returns the current number of peers connected to the host chain.
		PeerCount(ctx context.Context) (uint64, error)
	}
)

// backend represents the backend for the JSON-RPC service.
type backend struct {
	polar               *Polaris
	cfg                 *Config
	extRPCEnabled       bool
	allowUnprotectedTxs bool
	hostChainVersion    string
	gpo                 *gasprice.Oracle
	logger              log.Logger
}

// ==============================================================================
// Constructor
// ==============================================================================

// NewAPIBackend returns a new `Backend` object.
func NewAPIBackend(
	polar *Polaris,
	extRPCEnabled bool,
	allowUnprotectedTxs bool,
	cfg *Config,
	hostChainVersion string,
) APIBackend {
	b := &backend{

		polar:               polar,
		cfg:                 cfg,
		extRPCEnabled:       extRPCEnabled,
		hostChainVersion:    hostChainVersion,
		allowUnprotectedTxs: allowUnprotectedTxs,
		logger:              log.Root(),
	}

	if cfg.GPO.Default == nil {
		panic("cfg.GPO.Default is nil")
	}
	b.gpo = gasprice.NewOracle(b, cfg.GPO)
	return b
}

// ==============================================================================
// General Ethereum API
// ==============================================================================

// ChainConfig returns the chain configuration.
func (b *backend) ChainConfig() *params.ChainConfig {
	b.logger.Debug("called eth.rpc.backend.ChainConfig")
	return b.polar.blockchain.Config()
}

// CurrentHeader returns the current header from the local chains.
func (b *backend) CurrentHeader() *ethtypes.Header {
	return b.polar.blockchain.CurrentHeader()
}

// CurrentBlock returns the current block from the local chain.
func (b *backend) CurrentBlock() *ethtypes.Header {
	return b.polar.blockchain.CurrentHeader()
}

// SuggestGasTipCap returns the recommended gas tip cap for a new transaction.
func (b *backend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	defer b.logger.Debug("called eth.rpc.backend.SuggestGasTipCap", "suggested_tip_cap")
	return b.gpo.SuggestTipCap(ctx)
}

// FeeHistory returns the base fee and gas used history of the last N blocks.
func (b *backend) FeeHistory(
	ctx context.Context,
	blockCount uint64,
	lastBlock rpc.BlockNumber,
	rewardPercentiles []float64,
) (*big.Int, [][]*big.Int, []*big.Int, []float64, error) {
	return b.gpo.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
}

// ChainDb is unused in Polaris.
func (b *backend) ChainDb() ethdb.Database { //nolint:stylecheck // conforms to interface.
	return ethdb.Database(nil)
}

// AccountManager is unused in Polaris.
func (b *backend) AccountManager() *accounts.Manager {
	return &accounts.Manager{}
}

// ExtRPCEnabled returns whether the RPC endpoints are exposed over external
// interfaces.
func (b *backend) ExtRPCEnabled() bool {
	return b.extRPCEnabled
}

// RPCGasCap returns the global gas cap for eth_call over rpc: this is
// if the user doesn't specify a cap.
func (b *backend) RPCGasCap() uint64 {
	return b.cfg.RPCGasCap
}

// RPCEVMTimeout returns the global timeout for eth_call over rpc.
func (b *backend) RPCEVMTimeout() time.Duration {
	return b.cfg.RPCEVMTimeout
}

// RPCTxFeeCap returns the global gas price cap for transactions over rpc.
func (b *backend) RPCTxFeeCap() float64 {
	return b.cfg.RPCTxFeeCap
}

// UnprotectedAllowed returns whether unprotected transactions are alloweds.
func (b *backend) UnprotectedAllowed() bool {
	return b.allowUnprotectedTxs
}

// ==============================================================================
// Blockchain API
// ==============================================================================

// SetHead is used for state sync on ethereum, we leave state sync up to the host
// chain and thus it is not implemented in Polaris.
func (b *backend) SetHead(_ uint64) {
	panic("not implemented")
}

func (b *backend) HeaderByNumber(
	_ context.Context,
	number rpc.BlockNumber,
) (*ethtypes.Header, error) {
	switch number {
	case rpc.PendingBlockNumber:
		// TODO: handle "miner" stuff, Pending block is only known by the miner
		block := b.polar.miner.PendingBlock()
		if block != nil {
			return block.Header(), nil
		}
		// To improve client compatibility we return the latest state if
		// pending is not available.
		return b.polar.blockchain.CurrentHeader(), nil
	case rpc.LatestBlockNumber:
		return b.polar.blockchain.CurrentHeader(), nil
	case rpc.FinalizedBlockNumber:
		block := b.polar.blockchain.CurrentFinalBlock()
		if block != nil {
			return block, nil
		}
		return nil, errors.New("finalized block not found")
	case rpc.SafeBlockNumber:
		block := b.polar.blockchain.CurrentSafeBlock()
		if block != nil {
			return block, nil
		}
		return nil, errors.New("safe block not found")
	case rpc.EarliestBlockNumber:
		return b.polar.blockchain.GetHeaderByNumber(0), nil
	default:
		return b.polar.blockchain.GetHeaderByNumber(uint64(number)), nil
	}
}

// HeaderByNumberOrHash returns the header identified by `number` or `hash`.
func (b *backend) HeaderByNumberOrHash(ctx context.Context,
	blockNrOrHash rpc.BlockNumberOrHash,
) (*ethtypes.Header, error) {
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.HeaderByNumber(ctx, blockNr)
	}
	if hash, ok := blockNrOrHash.Hash(); ok {
		return b.HeaderByHash(ctx, hash)
	}
	return nil, errors.New("invalid arguments; neither block nor hash specified")
}

// HeaderByHash returns the block header with the given hash.
func (b *backend) HeaderByHash(_ context.Context, hash common.Hash) (*ethtypes.Header, error) {
	return b.polar.blockchain.GetHeaderByHash(hash), nil
}

// BlockByNumber returns the block with the given `number`.
func (b *backend) BlockByNumber(
	_ context.Context,
	number rpc.BlockNumber,
) (*ethtypes.Block, error) {
	// Pending block is only known by the miner
	switch number {
	case rpc.PendingBlockNumber:
		block := b.polar.miner.PendingBlock()
		if block == nil {
			// To improve client compatibility we return the latest state if
			// pending is not available.
			header := b.polar.blockchain.CurrentBlock()
			return b.polar.blockchain.GetBlock(
				header.Hash(), header.Number.Uint64(),
			), nil
		}
		return block, nil
	// Otherwise resolve and return the block
	case rpc.LatestBlockNumber:
		header := b.polar.blockchain.CurrentBlock()
		return b.polar.blockchain.GetBlock(header.Hash(), header.Number.Uint64()), nil

	case rpc.FinalizedBlockNumber:
		header := b.polar.blockchain.CurrentFinalBlock()
		return b.polar.blockchain.GetBlock(header.Hash(), header.Number.Uint64()), nil

	case rpc.SafeBlockNumber:
		header := b.polar.blockchain.CurrentSafeBlock()
		return b.polar.blockchain.GetBlock(header.Hash(), header.Number.Uint64()), nil

	case rpc.EarliestBlockNumber:
		return b.polar.blockchain.GetBlockByNumber(0), nil
	}
	// safe to assume number > 0
	return b.polar.blockchain.GetBlockByNumber(uint64(number)), nil
}

// BlockByHash returns the block with the given `hash`.
func (b *backend) BlockByHash(_ context.Context, hash common.Hash) (*ethtypes.Block, error) {
	block := b.polar.blockchain.GetBlockByHash(hash)
	b.logger.Debug("BlockByHash", "hash", hash, "block", block)
	if block == nil {
		b.logger.Error("eth.rpc.backend.BlockByHash", "hash", hash, "nil", true)
		return nil, nil //nolint:nilnil // to match geth.
	}
	b.logger.Debug("called eth.rpc.backend.BlockByHash", "header", block.Header(),
		"num_txs", len(block.Transactions()))
	return block, nil
}

func (b *backend) BlockByNumberOrHash(
	ctx context.Context,
	blockNrOrHash rpc.BlockNumberOrHash,
) (*ethtypes.Block, error) {
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.BlockByNumber(ctx, blockNr)
	}
	if hash, ok := blockNrOrHash.Hash(); ok {
		block := b.polar.blockchain.GetBlockByHash(hash)
		if block == nil {
			return nil, pcore.ErrBlockNotFound
		}
		// if blockNrOrHash.RequireCanonical &&
		// b.polar.blockchain.GetCanonicalHash(header.Number.Uint64()) != hash {
		// 	return nil, errors.New("hash is not currently canonical")
		// }
		// block := b.polar.blockchain.GetBlock(hash, header.Number.Uint64())
		// if block == nil {
		// 	return nil, errors.New("header found, but block body is missing")
		// }
		return block, nil
	}
	return nil, errors.New("invalid arguments; neither block nor hash specified")
}

func (b *backend) StateAndHeaderByNumber(
	ctx context.Context,
	number rpc.BlockNumber,
) (state.StateDB, *ethtypes.Header, error) {
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, nil, err
	}
	if header == nil {
		// to match Geth
		return nil, nil, pcore.ErrHeaderNotFound
	}
	b.logger.Debug("called eth.rpc.backend.StateAndHeaderByNumber", "header", header)

	// StateAtBlockNumber returns nil if the number is not found
	state, err := b.polar.blockchain.StateAtBlockNumber(header.Number.Uint64())
	if err != nil {
		b.logger.Error("eth.rpc.backend.StateAndHeaderByNumber", "number", number, "err", err)
		return nil, nil, err
	}

	return state, header, nil
}

func (b *backend) StateAndHeaderByNumberOrHash(
	ctx context.Context,
	blockNrOrHash rpc.BlockNumberOrHash,
) (state.StateDB, *ethtypes.Header, error) {
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.StateAndHeaderByNumber(ctx, blockNr)
	}

	if hash, ok := blockNrOrHash.Hash(); ok {
		header, err := b.HeaderByHash(ctx, hash)
		if err != nil {
			return nil, nil, err
		}
		if header == nil {
			// to match Geth
			return nil, nil, pcore.ErrBlockNotFound
		}
		// if blockNrOrHash.RequireCanonical &&
		// b.eth.blockchain.GetCanonicalHash(header.Number.Uint64()) != hash {
		// 	return nil, nil, errors.New("hash is not currently canonical")
		// }
		return b.StateAndHeaderByNumber(ctx, rpc.BlockNumber(header.Number.Int64()))
	}
	return nil, nil, errors.New("invalid arguments; neither block nor hash specified")
}

// StateAtBlock returns the state at a specific block.
func (b *backend) StateAtBlock(ctx context.Context, block *ethtypes.Block, reexec uint64,
	base state.StateDB, readOnly bool, preferDisk bool,
) (state.StateDB, tracers.StateReleaseFunc, error) {
	return b.polar.blockchain.StateAtBlock(ctx, block, reexec, base, readOnly, preferDisk)
}

// StateAtTransaction returns the state at a specific transaction.
func (b *backend) StateAtTransaction(
	ctx context.Context, block *ethtypes.Block,
	txIndex int, reexec uint64,
) (*core.Message, vm.BlockContext, state.StateDB, tracers.StateReleaseFunc, error,
) {
	return b.polar.blockchain.StateAtTransaction(ctx, block, txIndex, reexec)
}

// GetTransaction returns the transaction identified by `txHash`, along with
// information about the transaction.
func (b *backend) GetTransaction(
	_ context.Context,
	txHash common.Hash,
) (*ethtypes.Transaction, common.Hash, uint64, uint64, error) {
	b.logger.Debug("called eth.rpc.backend.GetTransaction", "tx_hash", txHash)
	txLookup := b.polar.blockchain.GetTransactionLookup(txHash)
	if txLookup == nil {
		return nil, common.Hash{}, 0, 0, nil
	}
	return txLookup.Tx, txLookup.BlockHash, txLookup.BlockNum, txLookup.TxIndex, nil
}

// PendingBlockAndReceipts returns the pending block (equivalent to current block in Polaris)
// and associated receipts.
func (b *backend) PendingBlockAndReceipts() (*ethtypes.Block, ethtypes.Receipts) {
	block, receipts := b.polar.miner.PendingBlockAndReceipts()
	// If the block is non-existent, return nil.
	// This is to maintain parity with the behavior of the geth backend.
	if block == nil {
		b.logger.Debug("called eth.rpc.backend.PendingBlockAndReceipts is nil", "block", block)
		return nil, nil
	}
	b.logger.Debug("called eth.rpc.backend.PendingBlockAndReceipts", "block", block,
		"num_receipts", len(receipts))
	return block, receipts
}

// GetReceipts returns the receipts for the given block hash.
func (b *backend) GetReceipts(_ context.Context, hash common.Hash) (ethtypes.Receipts, error) {
	return b.polar.blockchain.GetReceiptsByHash(hash), nil
}

// GetLogs returns the logs for the given block hash or number.
func (b *backend) GetLogs(
	_ context.Context, blockHash common.Hash, number uint64,
) ([][]*ethtypes.Log, error) {
	receipts := b.polar.blockchain.GetReceiptsByHash(blockHash)
	logs := make([][]*ethtypes.Log, len(receipts))
	for i, receipt := range receipts {
		logs[i] = receipt.Logs
	}
	b.logger.Debug("called eth.rpc.backend.GetLogs", "block_hash", blockHash, "number", number)
	return logs, nil
}

// GetTd returns the total difficulty of a block in the canonical chain.
// This is hardcoded to 69, as it is only applicable in a PoW chain.
func (b *backend) GetTd(_ context.Context, hash common.Hash) *big.Int {
	if header := b.polar.blockchain.GetHeaderByHash(hash); header != nil {
		return b.polar.blockchain.GetTd(hash, header.Number.Uint64())
	}
	return nil
}

// GetEVM returns a new EVM to be used for simulating a transaction, estimating gas etc.
func (b *backend) GetEVM(_ context.Context, msg *core.Message,
	state state.StateDB, header *ethtypes.Header, vmConfig *vm.Config,
	blockCtx *vm.BlockContext,
) *vm.EVM {
	if vmConfig == nil {
		vmConfig = b.polar.blockchain.GetVMConfig()
	}
	txContext := core.NewEVMTxContext(msg)
	var context vm.BlockContext
	if blockCtx != nil {
		context = *blockCtx
	} else {
		// TODO: we are hardcoding author to coinbase, this may be incorrect.
		// TODO: Suggestion -> implement Engine.Author() and allow host chain to decide.
		context = core.NewEVMBlockContext(header, b.polar.Blockchain(), &header.Coinbase)
	}
	return vm.NewEVM(context, txContext, state, b.polar.blockchain.Config(),
		*vmConfig)
}

// GetBlockContext returns a new block context to be used by a EVM.
func (b *backend) GetBlockContext(
	_ context.Context, header *ethtypes.Header,
) *vm.BlockContext {
	// TODO: we are hardcoding author to coinbase, this may be incorrect.
	// TODO: Suggestion -> implement Engine.Author() and allow host chain to decide.
	blockContext := core.NewEVMBlockContext(header, b.polar.Blockchain(), &header.Coinbase)
	return &blockContext
}

func (b *backend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	b.logger.Debug("called eth.rpc.backend.SubscribeChainEvent", "ch", ch)
	return b.polar.blockchain.SubscribeChainEvent(ch)
}

func (b *backend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	b.logger.Debug("called eth.rpc.backend.SubscribeChainHeadEvent", "ch", ch)
	return b.polar.blockchain.SubscribeChainHeadEvent(ch)
}

func (b *backend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	b.logger.Debug("called eth.rpc.backend.SubscribeChainSideEvent", "ch", ch)
	return b.polar.blockchain.SubscribeChainSideEvent(ch)
}

// ==============================================================================
// Transaction Pool API
// ==============================================================================

func (b *backend) SendTx(_ context.Context, signedTx *ethtypes.Transaction) error {
	return b.polar.txPool.Add([]*ethtypes.Transaction{signedTx}, true, false)[0]
}

func (b *backend) GetPoolTransactions() (ethtypes.Transactions, error) {
	b.logger.Debug("called eth.rpc.backend.GetPoolTransactions")
	pending := b.polar.txPool.Pending(false)
	var txs ethtypes.Transactions
	for _, batch := range pending {
		for _, lazy := range batch {
			if tx := lazy.Resolve(); tx != nil {
				txs = append(txs, tx)
			}
		}
	}
	return txs, nil
}

func (b *backend) GetPoolTransaction(hash common.Hash) *ethtypes.Transaction {
	b.logger.Debug("called eth.rpc.backend.GetPoolTransaction", "tx_hash", hash)
	return b.polar.txPool.Get(hash)
}

func (b *backend) GetPoolNonce(_ context.Context, addr common.Address) (uint64, error) {
	nonce := b.polar.txPool.Nonce(addr)
	b.logger.Debug("called eth.rpc.backend.GetPoolNonce", "addr", addr, "nonce", nonce)
	return nonce, nil
}

func (b *backend) Stats() (int, int) {
	pending, queued := b.polar.txPool.Stats()
	b.logger.Debug("called eth.rpc.backend.Stats", "pending", pending, "queued", queued)
	return pending, queued
}

func (b *backend) TxPoolContent() (
	map[common.Address][]*ethtypes.Transaction,
	map[common.Address][]*ethtypes.Transaction,
) {
	pending, queued := b.polar.txPool.Content()
	b.logger.Debug(
		"called eth.rpc.backend.TxPoolContent", "pending", len(pending), "queued", len(queued))
	return pending, queued
}

func (b *backend) TxPoolContentFrom(addr common.Address) (
	[]*ethtypes.Transaction,
	[]*ethtypes.Transaction,
) {
	pending, queued := b.polar.txPool.ContentFrom(addr)
	b.logger.Debug("called eth.rpc.backend.TxPoolContentFrom",
		"addr", addr, "pending", len(pending), "queued", len(queued))
	return pending, queued
}

func (b *backend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return b.polar.txPool.SubscribeTransactions(ch, true)
}

func (b *backend) Engine() consensus.Engine {
	return b.polar.blockchain.Engine()
}

// GetBody retrieves the block body corresponding to block by has or number..
func (b *backend) GetBody(
	ctx context.Context,
	hash common.Hash,
	number rpc.BlockNumber,
) (*ethtypes.Body, error) {
	if number < 0 || hash == (common.Hash{}) {
		b.logger.Error("eth.rpc.backend.GetBody", "number", number, "hash", hash)
		return nil, errors.New("invalid arguments; expect hash and no special block numbers")
	}
	block, err := b.BlockByNumberOrHash(
		ctx, rpc.BlockNumberOrHash{BlockNumber: &number, BlockHash: &hash})
	if block == nil || err != nil {
		b.logger.Error("eth.rpc.backend.GetBody", "number", number, "hash", hash, "err", err)
		return nil, nil //nolint:nilnil // to match geth.
	}
	b.logger.Debug("called eth.rpc.backend.GetBody", "hash", hash, "number", number)
	return block.Body(), nil
}

func (b *backend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	b.logger.Debug("called eth.rpc.backend.SubscribeRemovedLogsEvent", "ch", ch)
	return b.polar.blockchain.SubscribeRemovedLogsEvent(ch)
}

func (b *backend) SubscribeLogsEvent(ch chan<- []*ethtypes.Log) event.Subscription {
	return b.polar.blockchain.SubscribeLogsEvent(ch)
}

func (b *backend) SubscribePendingLogsEvent(ch chan<- []*ethtypes.Log) event.Subscription {
	return b.polar.miner.SubscribePendingLogs(ch)
}

func (b *backend) BloomStatus() (uint64, uint64) {
	// TODO: Implement your code here
	return 0, 0
}

func (b *backend) ServiceFilter(_ context.Context, _ *bloombits.MatcherSession) {
	// TODO: Implement your code here
}

// Version returns the current chain protocol version.
// For education:
// https://medium.com/@pedrouid/chainid-vs-networkid-how-do-they-differ-on-ethereum-eec2ed41635b
func (b *backend) Version() string {
	chainID := b.ChainConfig().ChainID
	if chainID == nil {
		b.logger.Error("eth.rpc.backend.Version", "ChainID is nil")
		return "-1"
	}
	return chainID.String()
}

// SyncProgress returns the current progress of the sync algorithm.
func (b *backend) SyncProgress() ethereum.SyncProgress {
	sp, err := b.polar.syncStatus.SyncProgress(context.Background())
	if err != nil {
		b.logger.Error("eth.rpc.backend.SyncProgress", "err", err)
		return ethereum.SyncProgress{}
	}
	return sp
}

// Listening returns whether the node is listening for connections.
func (b *backend) Listening() bool {
	listening, err := b.polar.syncStatus.Listening(context.Background())
	if err != nil {
		b.logger.Error("eth.rpc.backend.Listening", "err", err)
		return false
	}
	return listening
}

// PeerCount returns the number of connected peers.
func (b *backend) PeerCount() hexutil.Uint {
	peerCount, err := b.polar.syncStatus.PeerCount(context.Background())
	if err != nil {
		b.logger.Error("eth.rpc.backend.PeerCount", "err", err)
		return hexutil.Uint(0)
	}
	return hexutil.Uint(peerCount)
}

// ClientVersion returns the current client version.
func (b *backend) ClientVersion() string {
	return fmt.Sprintf(
		"%s:%s", b.hostChainVersion, strings.ToLower(version.ClientName("polaris-geth")),
	)
}
