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

	"cosmossdk.io/log"

	"github.com/ethereum/go-ethereum/consensus/misc"

	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/state"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
)

// Backend wraps all methods required for mining. Only full node is capable
// to offer all the functions here.
type Backend interface {
	// Blockchain returns the blockchain instance.
	Blockchain() core.Blockchain
}

// Miner defines the interface for a Polaris miner.
type Miner interface {
	// Prepare prepares the miner for a new block. This method is called before the first tx in
	// the block.
	Prepare(context.Context, uint64) *types.Header

	// ProcessTransaction processes the given transaction and returns the receipt after applying
	// the state transition. This method is called for each tx in the block.
	ProcessTransaction(context.Context, *types.Transaction) (*core.ExecutionResult, error)

	// Finalize is called after the last tx in the block.
	Finalize(context.Context) error
}

// miner implements the Miner interface.
type miner struct {
	backend   Backend
	processor *core.StateProcessor
	host      core.PolarisHostChain
	bp        core.BlockPlugin
	cp        core.ConfigurationPlugin
	gp        core.GasPlugin

	sp       core.StatePlugin
	logger   log.Logger
	vmConfig vm.Config
	statedb  vm.PolarisStateDB

	// TODO: historical plugin has no purpose here in the miner.
	// Should be handled async via channel
	hp core.HistoricalPlugin
}

// NewMiner creates a new Miner instance.
func NewMiner(backend Backend) Miner {
	host := backend.Blockchain().GetHost()

	m := &miner{
		host:    host,
		bp:      host.GetBlockPlugin(),
		cp:      host.GetConfigurationPlugin(),
		hp:      host.GetHistoricalPlugin(),
		gp:      host.GetGasPlugin(),
		sp:      host.GetStatePlugin(),
		backend: backend,
		logger:  log.NewNopLogger(), // todo: fix.
	}

	m.statedb = state.NewStateDB(m.sp)
	m.processor = core.NewStateProcessor(
		m.cp, m.gp, host.GetPrecompilePlugin(), m.statedb, &m.vmConfig,
	)

	return m
}

// Prepare prepares the blockchain for processing a new block at the given height.
func (m *miner) Prepare(ctx context.Context, number uint64) *types.Header {
	// Prepare the State, Block, Configuration, Gas, and Historical plugins for the block.
	m.sp.Prepare(ctx)
	m.bp.Prepare(ctx)
	m.cp.Prepare(ctx)
	m.gp.Prepare(ctx)

	if m.hp != nil {
		m.hp.Prepare(ctx)
	}

	coinbase, timestamp := m.host.GetBlockPlugin().GetNewBlockMetadata(number)
	chainCfg := m.cp.ChainConfig()

	// Build the new block header.
	parent := m.backend.Blockchain().CurrentFinalBlock()
	if number >= 1 && parent == nil {
		parent = m.backend.Blockchain().GetHeaderByNumber(number - 1)
	}

	// Polaris does not set Ethereum state root (Root), mix hash (MixDigest), extra data (Extra),
	// and block nonce (Nonce) on the new header.
	header := &types.Header{
		// Used in Polaris.
		ParentHash: parent.Hash(),
		Coinbase:   coinbase,
		Number:     new(big.Int).SetUint64(number),
		GasLimit:   m.gp.BlockGasLimit(),
		Time:       timestamp,
		BaseFee:    misc.CalcBaseFee(m.backend.Blockchain().Config(), parent),
	}

	m.logger.Info("preparing evm block", "seal_hash", header.Hash())

	// Prepare the State Processor, StateDB and the EVM for the block.
	m.processor.Prepare(
		vm.NewGethEVMWithPrecompiles(
			*m.backend.Blockchain().NewEVMBlockContext(header),
			vm.TxContext{}, m.statedb, chainCfg, m.vmConfig,
			m.backend.Blockchain().GetHost().GetPrecompilePlugin()),
		header,
	)

	return header
}

// ProcessTransaction processes the given transaction and returns the receipt after applying
// the state transition. This method is called for each tx in the block.
func (m *miner) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*core.ExecutionResult, error) {
	m.logger.Debug("processing evm transaction", "tx_hash", tx.Hash())

	// Reset the Gas and State plugins for the tx.
	m.gp.Reset(ctx) // TODO: may not need this.
	m.sp.Reset(ctx)

	return m.processor.ProcessTransaction(ctx, tx)
}

// Finalize is called after the last tx in the block.
func (m *miner) Finalize(ctx context.Context) error {
	block, receipts, logs, err := m.processor.Finalize(ctx)
	if err != nil {
		return err
	}

	return m.backend.Blockchain().InsertBlock(block, receipts, logs)
}
