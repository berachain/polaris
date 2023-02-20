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

//nolint:gomnd,nilnil // TODO: fix
package rpc

import (
	"context"
	"math/big"
	"time"

	"github.com/berachain/stargazer/eth/api"
	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/params"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
)

var _ Backend = (*backend)(nil)

type backend struct {
	chain api.Chain
}

func NewBackend(chain api.Chain) Backend {
	return &backend{
		chain: chain,
	}
}

// General Ethereum API.
func (b *backend) SyncProgress() ethereum.SyncProgress {
	// Implement your code here
	return ethereum.SyncProgress{
		CurrentBlock: 1000000,
		HighestBlock: 2000000,
	}
}

func (b *backend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	// Implement your code here
	return big.NewInt(1000000000), nil
}

func (b *backend) FeeHistory(ctx context.Context, blockCount int, lastBlock BlockNumber,
	rewardPercentiles []float64) (*big.Int, [][]*big.Int, []*big.Int, []float64, error) {
	// Implement your code here
	return big.NewInt(1000000000), nil, nil, nil, nil
}

func (b *backend) ChainDb() ethdb.Database { //nolint:stylecheck // conforms to interface.
	// Implement your code here
	return nil
}

func (b *backend) AccountManager() *accounts.Manager {
	// Implement your code here
	return &accounts.Manager{}
}

func (b *backend) ExtRPCEnabled() bool {
	// Implement your code here
	return true
}

func (b *backend) RPCGasCap() uint64 {
	// Implement your code here
	return 10000000
}

func (b *backend) RPCEVMTimeout() time.Duration {
	// Implement your code here
	return 30 * time.Second
}

func (b *backend) RPCTxFeeCap() float64 {
	// Implement your code here
	return 1.0
}

func (b *backend) UnprotectedAllowed() bool {
	// Implement your code here
	return false
}

// Blockchain API.
func (b *backend) SetHead(number uint64) {
	// Implement your code here
}

func (b *backend) HeaderByNumber(ctx context.Context,
	number BlockNumber) (*types.Header, error) {
	// Implement your code here
	return &types.Header{
		Number: big.NewInt(1000000),
	}, nil
	// return b.chain.GetBlockPlugin().GetStargazerHeaderAtHeight(number.Int64()).Header, nil
}

func (b *backend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	// Implement your code here
	return &types.Header{}, nil
}

func (b *backend) HeaderByNumberOrHash(ctx context.Context,
	blockNrOrHash BlockNumberOrHash) (*types.Header, error) {
	// Implement your code here
	return nil, nil
}

func (b *backend) CurrentHeader() *types.Header {
	// Implement your code here
	return &types.Header{}
}

func (b *backend) CurrentBlock() *types.Block {
	// Implement your code here
	return &types.Block{}
}

func (b *backend) BlockByNumber(ctx context.Context, number BlockNumber) (*types.Block, error) {
	// Implement your code here
	return &types.Block{}, nil
}

func (b *backend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	// Implement your code here
	return &types.Block{}, nil
}

func (b *backend) BlockByNumberOrHash(ctx context.Context,
	blockNrOrHash BlockNumberOrHash) (*types.Block, error) {
	// Implement your code here
	return &types.Block{}, nil
}

func (b *backend) StateAndHeaderByNumber(ctx context.Context,
	number BlockNumber) (*state.StateDB, *types.Header, error) {
	// Implement your code here
	return nil, nil, nil
}

func (b *backend) StateAndHeaderByNumberOrHash(ctx context.Context,
	blockNrOrHash BlockNumberOrHash) (*state.StateDB, *types.Header, error) {
	// Implement your code here
	return nil, nil, nil
}

func (b *backend) PendingBlockAndReceipts() (*types.Block, types.Receipts) {
	// Implement your code here
	return &types.Block{}, nil
}

func (b *backend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	// Implement your code here
	return nil, nil
}

func (b *backend) GetTd(ctx context.Context, hash common.Hash) *big.Int {
	// Implement your code here
	return new(big.Int)
}

func (b *backend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB,
	header *types.Header, vmConfig *vm.Config) (*vm.EVM, func() error, error) {
	// Implement your code here
	return nil, nil, nil
}

func (b *backend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	// Implement your code here
	return nil
}

func (b *backend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	// Implement your code here
	return nil
}

func (b *backend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	// Implement your code here
	return nil
}

// Transaction pool API.
func (b *backend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	// Implement your code here
	return nil
}

func (b *backend) GetTransaction(ctx context.Context,
	txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	// Implement your code here
	return nil, common.Hash{}, 0, 0, nil
}

func (b *backend) GetPoolTransactions() (types.Transactions, error) {
	// Implement your code here
	return nil, nil
}

func (b *backend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	// Implement your code here
	return nil
}

func (b *backend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	// Implement your code here
	return 0, nil
}

func (b *backend) Stats() (int, int) {
	pending := 0
	queued := 0
	// Implement your code here
	return pending, queued
}

func (b *backend) TxPoolContent() (map[common.Address]types.Transactions,
	map[common.Address]types.Transactions) {
	// Implement your code here
	return nil, nil
}

func (b *backend) TxPoolContentFrom(addr common.Address,
) (types.Transactions, types.Transactions) {
	// Implement your code here
	return nil, nil
}

func (b *backend) SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription {
	// Implement your code here
	return nil
}

func (b *backend) ChainConfig() *params.ChainConfig {
	// Implement your code here
	return b.chain.Host().GetConfigurationPlugin().ChainConfig()
}

func (b *backend) Engine() consensus.Engine {
	// Implement your code here
	return nil
}

// Filters Backend.
func (b *backend) GetBody(ctx context.Context, hash common.Hash,
	number BlockNumber) (*types.Body, error) {
	// Implement your code here
	return nil, nil
}

func (b *backend) GetLogs(ctx context.Context, blockHash common.Hash,
	number uint64) ([][]*types.Log, error) {
	// Implement your code here
	return nil, nil
}

func (b *backend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	// Implement your code here
	return nil
}

func (b *backend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	// Implement your code here
	return nil
}

func (b *backend) SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription {
	// Implement your code here
	return nil
}

func (b *backend) BloomStatus() (uint64, uint64) {
	// Implement your code here
	return 0, 0
}

func (b *backend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	// Implement your code here
}
