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

package miner

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core/txpool"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/state"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/log"
	errorslib "pkg.berachain.dev/polaris/lib/errors"
)

// Backend wraps all methods required for mining. Only full node is capable
// to offer all the functions here.
type Backend interface {
	// Blockchain returns the blockchain instance.
	Blockchain() core.Blockchain
	TxPool() *txpool.TxPool
	Host() core.PolarisHostChain
}

// Miner defines the interface for a Polaris miner.
type Miner interface {
	// Prepare prepares the miner for a new block. This method is called before the first tx in
	// the block.
	Prepare(context.Context, uint64) *types.Header

	// ProcessTransaction processes the given transaction and returns the receipt after applying
	// the state transition. This method is called for each tx in the block.
	ProcessTransaction(context.Context, *types.Transaction) (*types.Receipt, error)

	// Finalize is called after the last tx in the block.
	Finalize(context.Context) error

	// TODO: deprecate
	NextBaseFee() *big.Int
}

// miner implements the Miner interface.
// TODO: RENAME TO STATE PROCESSOR OR DEPRECATE.
type miner struct {
	backend   Backend
	chain     core.Blockchain
	processor *core.StateProcessor
	txPool    *txpool.TxPool
	bp        core.BlockPlugin
	cp        core.ConfigurationPlugin
	gp        core.GasPlugin
	pp        core.PrecompilePlugin
	sp        core.StatePlugin
	logger    log.Logger
	vmConfig  vm.Config
	statedb   state.PolarStateDB

	// TODO: historical plugin has no purpose here in the miner.
	// Should be handled async via channel
	hp core.HistoricalPlugin

	// workspace
	pendingHeader *types.Header
	gasPool       *core.GasPool
}

// NewMiner creates a new Miner instance.
func New(backend Backend) Miner {
	chain := backend.Blockchain()
	host := backend.Host()

	m := &miner{
		bp:      host.GetBlockPlugin(),
		cp:      host.GetConfigurationPlugin(),
		hp:      host.GetHistoricalPlugin(),
		gp:      host.GetGasPlugin(),
		sp:      host.GetStatePlugin(),
		txPool:  backend.TxPool(),
		chain:   chain,
		backend: backend,
		logger:  log.Root(), // todo: fix.
	}

	m.pp = host.GetPrecompilePlugin()
	if m.pp == nil {
		m.pp = precompile.NewDefaultPlugin()
	}

	return m
}

// TODO: deprecate and properly recalculate in prepare proposal, this is fine for now though.
func (m *miner) NextBaseFee() *big.Int {
	if m.pendingHeader == nil {
		return big.NewInt(0)
	}
	return eip1559.CalcBaseFee(m.cp.ChainConfig(), m.pendingHeader)
}

// Prepare prepares the blockchain for processing a new block at the given height.
func (m *miner) Prepare(ctx context.Context, number uint64) *types.Header {
	// Prepare the State, Block, Configuration, Gas, and Historical plugins for the block.
	m.sp.Reset(ctx)
	m.bp.Prepare(ctx)
	m.cp.Prepare(ctx)
	m.gp.Prepare(ctx)

	// TODO: this shouldnt be in the miner.
	if m.hp != nil {
		m.hp.Prepare(ctx)
	}

	coinbase, timestamp := m.bp.GetNewBlockMetadata(number)
	chainConfig := m.cp.ChainConfig()

	// Build the new block m.pendingHeader.
	parent := m.chain.CurrentFinalBlock()
	if number >= 1 && parent == nil {
		parent = m.chain.GetHeaderByNumber(number - 1)
	}

	// Polaris does not set Ethereum state root (Root), mix hash (MixDigest), extra data (Extra),
	// and block nonce (Nonce) on the new m.pendingHeader.
	m.pendingHeader = &types.Header{
		// Used in Polaris.
		ParentHash: parent.Hash(),
		Coinbase:   coinbase,
		Number:     new(big.Int).SetUint64(number),
		GasLimit:   m.gp.BlockGasLimit(),
		Time:       timestamp,
		Difficulty: new(big.Int),
	}

	// TODO: Settable in PrepareProposal.
	// Set the extra field.
	if /*len(w.extra) != 0*/ true {
		m.pendingHeader.Extra = nil
	}

	// Set the randomness field from the beacon chain if it's available.
	// TODO: Settable in PrepareProposal.
	if /*genParams.random != (common.Hash{})*/ true {
		// m.pendingHeader.MixDigest = genParams.random
		m.pendingHeader.MixDigest = common.Hash{}
	}

	// Apply EIP-1559.
	// TODO: Move to PrepareProposal.
	if chainConfig.IsLondon(m.pendingHeader.Number) {
		m.pendingHeader.BaseFee = eip1559.CalcBaseFee(chainConfig, parent)
		// On switchover.
		// TODO: implement.
		// if !chainConfig.IsLondon(parent.Number) {
		// 	parentGasLimit := parent.GasLimit * chainConfig.ElasticityMultiplier()
		// 	m.pendingHeader.GasLimit = core.CalcGasLimit(parentGasLimit, bc.gp.BlockGasLimit())
		// }
	}

	// Apply EIP-4844, EIP-4788.
	// TODO: Move to PrepareProposal.
	if chainConfig.IsCancun(m.pendingHeader.Number, m.pendingHeader.Time) {
		var excessBlobGas uint64
		if chainConfig.IsCancun(parent.Number, parent.Time) {
			excessBlobGas = eip4844.CalcExcessBlobGas(*parent.ExcessBlobGas, *parent.BlobGasUsed)
		} else {
			// For the first post-fork block, both parent.data_gas_used and
			// parent.excess_data_gas are evaluated as 0
			excessBlobGas = eip4844.CalcExcessBlobGas(0, 0)
		}
		m.pendingHeader.BlobGasUsed = new(uint64)
		m.pendingHeader.ExcessBlobGas = &excessBlobGas
		m.pendingHeader.ParentBeaconRoot = &common.Hash{}
	}

	m.logger.Info("preparing evm block", "seal_hash", m.pendingHeader.Hash())

	// Create new statedb and processor every block to clear out journals and stuff.
	// DEPRECATED VIA 1 Block 1 Txn anyways, but works for now.
	m.statedb = state.NewStateDB(m.sp, m.pp)
	m.processor = core.NewStateProcessor(
		m.cp, m.pp, m.statedb, &m.vmConfig,
	)

	m.processor.Prepare(
		m.pendingHeader,
	)
	return m.pendingHeader
}

// ProcessTransaction processes the given transaction and returns the receipt after
// applying the state transition. This method is called for each tx in the block.
func (m *miner) ProcessTransaction(
	ctx context.Context, tx *types.Transaction,
) (*types.Receipt, error) {
	m.logger.Debug("processing evm transaction", "tx_hash", tx.Hash())

	// Reset the Gas and State plugins for the tx.
	m.gp.Reset(ctx) // TODO: may not need this.
	m.sp.Reset(ctx)

	// We set the gasPool = gasLimit - gasUsed.
	m.gasPool = new(core.GasPool).AddGas(m.pendingHeader.GasLimit - m.gp.BlockGasConsumed())

	// Header is out of sync with block plugin.
	// TODO: the miner will handle this systemically when properly done is PrepareProposal.
	if m.gp.BlockGasConsumed() != m.pendingHeader.GasUsed {
		panic("gas consumed mismatch")
	}

	receipt, err := m.processor.ProcessTransaction(ctx, m.chain, m.gasPool, tx)
	if err != nil {
		return nil, errorslib.Wrapf(
			err, "could not process transaction [%s]", tx.Hash().Hex(),
		)
	}

	// Consume the gas used by the state transition. In both the out of block gas as well as out of
	// gas on the plugin cases, the line below will consume the remaining gas for the block and
	// transaction respectively.
	if err = m.gp.ConsumeTxGas(receipt.GasUsed); err != nil {
		return nil, errorslib.Wrapf(
			err, "could not consume gas used [%s] %d", tx.Hash().Hex(), receipt.GasUsed,
		)
	}
	return receipt, nil
}

// Finalize is called after the last tx in the block.
func (m *miner) Finalize(ctx context.Context) error {
	block, receipts, logs, err := m.processor.Finalize(ctx)
	if err != nil {
		return err
	}

	return m.chain.InsertBlock(block, receipts, logs)
}
