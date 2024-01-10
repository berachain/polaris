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

package state

import (
	"context"

	"github.com/berachain/polaris/eth/core/precompile"
	"github.com/berachain/polaris/eth/core/state/journal"
	"github.com/berachain/polaris/lib/snapshot"
	libtypes "github.com/berachain/polaris/lib/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

// For mocks.
type PrecompilePlugin interface {
	precompile.Plugin
}

// stateDB is a struct that holds the plugins and controller to manage Ethereum state.
type stateDB struct {
	// Plugin is injected by the chain running the Polaris EVM.
	Plugin
	pp precompile.Plugin

	// Journals built internally and required for the stateDB.
	journal.Log
	journal.Refund
	journal.Accesslist
	journal.SelfDestructs
	journal.TransientStorage

	// ctrl is used to manage snapshots and reverts across plugins and journals.
	ctrl libtypes.Controller[string, libtypes.Controllable[string]]

	// rules is used to store the rules for the chain.
	rules *params.Rules
}

type (
	// StateDB is an alias for StateDBI.
	StateDB = state.StateDBI //nolint:revive // to match geth naming.

	// PolarStateDB is a Polaris StateDB that has a context.
	PolarStateDB interface {
		StateDB
		GetContext() context.Context
	}
)

// NewStateDB returns a vm.PolarStateDB with the given StatePlugin and new journals.
func NewStateDB(sp Plugin, pp precompile.Plugin) PolarStateDB {
	return newStateDBWithJournals(
		sp, pp, journal.NewLogs(), journal.NewRefund(), journal.NewAccesslist(),
		journal.NewSelfDestructs(sp), journal.NewTransientStorage(),
	)
}

// newStateDBWithJournals returns a vm.PolarStateDB with the given StatePlugin and journals.
func newStateDBWithJournals(
	sp Plugin, pp precompile.Plugin, lj journal.Log, rj journal.Refund, aj journal.Accesslist,
	sj journal.SelfDestructs, tj journal.TransientStorage,
) *stateDB {
	if sp == nil {
		panic("StatePlugin is nil in newStateDBWithJournals")
	}

	// Build the controller and register the plugins and journals
	ctrl := snapshot.NewController[string, libtypes.Controllable[string]]()
	_ = ctrl.Register(sp)
	_ = ctrl.Register(lj)
	_ = ctrl.Register(rj)
	_ = ctrl.Register(aj)
	_ = ctrl.Register(sj)
	_ = ctrl.Register(tj)

	return &stateDB{
		Plugin:           sp,
		pp:               pp,
		Log:              lj,
		Refund:           rj,
		Accesslist:       aj,
		SelfDestructs:    sj,
		TransientStorage: tj,
		ctrl:             ctrl,
	}
}

// =============================================================================
// Plugin
// =============================================================================

// GetPlugin returns the plugin from statedb.
func (sdb *stateDB) GetPlugin() Plugin {
	return sdb.Plugin
}

func (sdb *stateDB) GetPrecompileManager() any {
	return sdb.pp
}

// =============================================================================
// Snapshot
// =============================================================================

// Snapshot implements vm.PolarStateDB.
func (sdb *stateDB) Snapshot() int {
	return sdb.ctrl.Snapshot()
}

// RevertToSnapshot implements vm.PolarStateDB.
func (sdb *stateDB) RevertToSnapshot(id int) {
	sdb.ctrl.RevertToSnapshot(id)
}

// =============================================================================
// Commit state
// =============================================================================

// Finalise deletes the SelfDestructd accounts and finalizes all plugins, preparing
// the statedb for the next transaction.
func (sdb *stateDB) Finalise(bool) {
	sdb.DeleteAccounts(sdb.GetSelfDestructs())
	sdb.ctrl.Finalize()
}

// Intermediate root is called in lieu of Finalise pre-Byzantium. They are
// equivalent at the moment in Polaris as we do not leverage the state root.
func (sdb *stateDB) IntermediateRoot(bool) common.Hash {
	sdb.Finalise(true)
	return common.Hash{}
}

// Commit implements vm.PolarStateDB.
func (sdb *stateDB) Commit(_ uint64, _ bool) (common.Hash, error) {
	if err := sdb.Error(); err != nil {
		return common.Hash{}, err
	}
	return common.Hash{}, nil
}

// =============================================================================
// Prepare
// =============================================================================

// Implementation taken directly from the vm.PolarStateDB in Go-Ethereum.
//
// Prepare implements vm.PolarStateDB.
func (sdb *stateDB) Prepare(rules params.Rules, sender, coinbase common.Address,
	dest *common.Address, precompiles []common.Address, txAccesses ethtypes.AccessList,
) {
	copyRules := rules
	sdb.rules = &copyRules

	if rules.IsBerlin {
		// Clear out any leftover from previous executions
		sdb.Accesslist = journal.NewAccesslist()

		sdb.AddAddressToAccessList(sender)
		if dest != nil {
			sdb.AddAddressToAccessList(*dest)
			// If it's a create-tx, the destination will be added inside evm.create
		}
		for _, addr := range precompiles {
			sdb.AddAddressToAccessList(addr)
		}
		for _, el := range txAccesses {
			sdb.AddAddressToAccessList(el.Address)
			for _, key := range el.StorageKeys {
				sdb.AddSlotToAccessList(el.Address, key)
			}
		}
		if rules.IsShanghai { // EIP-3651: warm coinbase
			sdb.AddAddressToAccessList(coinbase)
		}
	}
	// Reset TransientStorage for the new transaction
	sdb.TransientStorage = journal.NewTransientStorage()
}

// =============================================================================
// PreImage
// =============================================================================

// AddPreimage implements the the vm.PolarStateDB interface, but currently
// performs a no-op since the EnablePreimageRecording flag is disabled.
func (sdb *stateDB) AddPreimage(_ common.Hash, _ []byte) {}

// AddPreimage implements the the `StateDB“ interface, but currently
// performs a no-op since the EnablePreimageRecording flag is disabled.
func (sdb *stateDB) Preimages() map[common.Hash][]byte {
	return nil
}

// =============================================================================
// Code
// =============================================================================

// GetCodeSize implements the vm.PolarStateDB interface by returning the size of the
// code associated with the given account.
func (sdb *stateDB) GetCode(addr common.Address) []byte {
	// We return a single byte for client compatibility w/precompiles.
	if sdb.pp != nil {
		if _, ok := sdb.pp.Get(addr, sdb.rules); ok {
			return []byte{0x01}
		}
	}
	return sdb.Plugin.GetCode(addr)
}

// GetCodeSize implements the vm.PolarStateDB interface by returning the size of the
// code associated with the given account.
func (sdb *stateDB) GetCodeSize(addr common.Address) int {
	return len(sdb.GetCode(addr))
}

// =============================================================================
// Other
// =============================================================================

// Copy returns a new statedb with cloned plugin and journals.
func (sdb *stateDB) Copy() StateDB {
	return newStateDBWithJournals(
		sdb.Plugin.Clone(), sdb.pp, sdb.Log.Clone(), sdb.Refund.Clone(),
		sdb.Accesslist.Clone(), sdb.SelfDestructs.Clone(), sdb.TransientStorage.Clone(),
	)
}

func (sdb *stateDB) DumpToCollector(_ state.DumpCollector, _ *state.DumpConfig) []byte {
	return nil
}

func (sdb *stateDB) Dump(_ *state.DumpConfig) []byte {
	return nil
}

func (sdb *stateDB) RawDump(_ *state.DumpConfig) state.Dump {
	return state.Dump{}
}

func (sdb *stateDB) Database() state.Database {
	return nil
}

func (sdb *stateDB) StartPrefetcher(_ string) {}

func (sdb *stateDB) StopPrefetcher() {}

func (sdb *stateDB) StorageTrie(_ common.Address) (state.Trie, error) {
	return nil, nil
}

func (sdb *stateDB) GetStorageProof(_ common.Address, _ common.Hash) ([][]byte, error) {
	return nil, nil
}

func (sdb *stateDB) GetProof(_ common.Address) ([][]byte, error) {
	return nil, nil
}

func (sdb *stateDB) GetOrNewStateObject(_ common.Address) *state.StateObject {
	return nil
}

func (sdb *stateDB) GetStorageRoot(_ common.Address) common.Hash {
	return common.Hash{}
}
