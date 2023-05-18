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
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/state/journal"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/params"
	"pkg.berachain.dev/polaris/lib/snapshot"
	libtypes "pkg.berachain.dev/polaris/lib/types"
)

// stateDB is a struct that holds the plugins and controller to manage Ethereum state.
type stateDB struct {
	// Plugin is injected by the chain running the Polaris EVM.
	Plugin

	// Journals built internally and required for the stateDB.
	journal.LogsI
	RefundJournal
	AccessListJournal
	SuicidesJournal
	TransientStorageJournal

	// ctrl is used to manage snapshots and reverts across plugins and journals.
	ctrl libtypes.Controller[string, libtypes.Controllable[string]]
}

// NewStateDB returns a vm.PolarisStateDB with the given StatePlugin.
func NewStateDB(sp Plugin) vm.PolarisStateDB {
	return newStateDBWithJournals(sp, journal.NewLogs())
}

func newStateDBWithJournals(sp Plugin, lj journal.LogsI) vm.PolarisStateDB {
	// Build the journals required for the stateDB
	rj := journal.NewRefund()
	aj := journal.NewAccesslist()
	sj := journal.NewSuicides(sp)
	tj := journal.NewTransientStorage()

	// Build the controller and register the plugins and journals
	ctrl := snapshot.NewController[string, libtypes.Controllable[string]]()
	_ = ctrl.Register(sp)
	_ = ctrl.Register(lj)
	_ = ctrl.Register(rj)
	_ = ctrl.Register(aj)
	_ = ctrl.Register(sj)
	_ = ctrl.Register(tj)

	return &stateDB{
		Plugin:                  sp,
		LogsI:                   lj,
		RefundJournal:           rj,
		AccessListJournal:       aj,
		TransientStorageJournal: tj,
		SuicidesJournal:         sj,
		ctrl:                    ctrl,
	}
}

// =============================================================================
// Snapshot
// =============================================================================

// Snapshot implements vm.PolarisStateDB.
func (sdb *stateDB) Snapshot() int {
	return sdb.ctrl.Snapshot()
}

// RevertToSnapshot implements vm.PolarisStateDB.
func (sdb *stateDB) RevertToSnapshot(id int) {
	sdb.ctrl.RevertToSnapshot(id)
}

// =============================================================================
// Commit state
// =============================================================================

// Finalise deletes the suicided accounts and finalizes all plugins, preparing the statedb for the
// next transaction.
func (sdb *stateDB) Finalise(bool) {
	sdb.DeleteAccounts(sdb.GetSuicides())
	sdb.ctrl.Finalize()
}

func (sdb *stateDB) Commit(deleteEmptyObjects bool) (common.Hash, error) {
	sdb.Finalise(deleteEmptyObjects)
	return common.Hash{}, nil
}

// =============================================================================
// Prepare
// =============================================================================

// Implementation taken directly from the vm.PolarisStateDB in Go-Ethereum.
//
// Prepare implements vm.PolarisStateDB.
func (sdb *stateDB) Prepare(rules params.Rules, sender, coinbase common.Address,
	dest *common.Address, precompiles []common.Address, txAccesses coretypes.AccessList) {
	if rules.IsBerlin {
		// Clear out any leftover from previous executions
		sdb.AccessListJournal = journal.NewAccesslist()

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
}

// =============================================================================
// PreImage
// =============================================================================

// AddPreimage implements the the vm.PolarisStateDB interface, but currently
// performs a no-op since the EnablePreimageRecording flag is disabled.
func (sdb *stateDB) AddPreimage(hash common.Hash, preimage []byte) {}

// AddPreimage implements the the `StateDB“ interface, but currently
// performs a no-op since the EnablePreimageRecording flag is disabled.
func (sdb *stateDB) Preimages() map[common.Hash][]byte {
	return nil
}

// =============================================================================
// Code Size
// =============================================================================

// GetCodeSize implements the vm.PolarisStateDB interface by returning the size of the
// code associated with the given account.
func (sdb *stateDB) GetCodeSize(addr common.Address) int {
	return len(sdb.GetCode(addr))
}

// =============================================================================
// Other
// =============================================================================

func (sdb *stateDB) Copy() StateDBI {
	return newStateDBWithJournals(sdb.Plugin.Clone(), sdb.LogsI.Clone())
}

func (sdb *stateDB) DumpToCollector(_ DumpCollector, _ *DumpConfig) []byte {
	return nil
}

func (sdb *stateDB) Dump(_ *DumpConfig) []byte {
	return nil
}

func (sdb *stateDB) RawDump(_ *DumpConfig) Dump {
	return Dump{}
}

func (sdb *stateDB) IteratorDump(_ *DumpConfig) IteratorDump {
	return IteratorDump{}
}

func (sdb *stateDB) Database() Database {
	return nil
}

func (sdb *stateDB) StartPrefetcher(_ string) {}

func (sdb *stateDB) StopPrefetcher() {}

func (sdb *stateDB) IntermediateRoot(_ bool) common.Hash {
	return common.Hash{}
}

func (sdb *stateDB) StorageTrie(_ common.Address) (Trie, error) {
	return nil, nil
}

func (sdb *stateDB) Error() error {
	return nil
}

func (sdb *stateDB) GetStorageProof(_ common.Address, _ common.Hash) ([][]byte, error) {
	return nil, nil
}

func (sdb *stateDB) GetProof(_ common.Address) ([][]byte, error) {
	return nil, nil
}

func (sdb *stateDB) GetOrNewStateObject(_ common.Address) *StateObject {
	return nil
}
