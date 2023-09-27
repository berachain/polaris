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
	"fmt"
	"math/big"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
)

// InitGenesis takes in a pointer to a genesis state object and populates the KV store.
func (p *plugin) InitGenesis(ctx sdk.Context, ethGen *core.Genesis) error {
	p.Reset(ctx)

	// Sort the addresses so that they are in a consistent order.
	sortedAddresses := make(core.SortableAddresses, 0, len(ethGen.Alloc))
	for address := range ethGen.Alloc {
		sortedAddresses = append(sortedAddresses, address)
	}
	sort.Sort(sortedAddresses)

	// Iterate over the sorted genesis accounts and set nonces, balances, codes, and storage.
	for _, address := range sortedAddresses {
		account := ethGen.Alloc[address]

		// initialize the account on the auth keeper
		// NOTE: The auth module's init genesis runs before the evm module's init genesis.
		if p.Exist(address) {
			// if the account exists on the auth keeper, ensure the nonce is the same.
			if p.GetNonce(address) != account.Nonce {
				return fmt.Errorf(
					"account nonce mismatch for (%s) between auth (%d) and evm (%d) genesis state",
					address.Hex(), p.GetNonce(address), account.Nonce,
				)
			}
		} else if account.Nonce != 0 {
			p.SetNonce(address, account.Nonce)
		}

		// initialize the account data on the state plugin
		if account.Balance != nil {
			p.SetBalance(address, account.Balance)
		}
		if account.Code != nil {
			if ethGen.Config.IsEIP155(big.NewInt(0)) && account.Nonce == 0 {
				// NOTE: EIP 161 was released at the same block as EIP 155.
				return fmt.Errorf(
					"EIP-161 requires an account with code (%s) to have nonce of at least 1, given (0)",
					address.Hex(),
				)
			}
			p.SetCode(address, account.Code)
		} else {
			// initialize the code hash to be empty by default
			p.cms.GetKVStore(p.storeKey).Set(CodeHashKeyFor(address), emptyCodeHashBytes)
		}
		if account.Storage != nil {
			p.SetStorage(address, account.Storage)
		}
	}

	p.Finalize()
	return nil
}

// Export genesis modifies a pointer to a genesis state object and populates it.
func (p *plugin) ExportGenesis(ctx sdk.Context, ethGen *core.Genesis) {
	p.Reset(ctx)
	ethGen.Alloc = make(core.GenesisAlloc)

	// Iterate Balances and set the genesis accounts.
	p.IterateBalances(func(address common.Address, balance *big.Int) bool {
		account, ok := ethGen.Alloc[address]
		if !ok {
			account = core.GenesisAccount{}
		}
		account.Code = p.GetCode(address)
		if account.Code != nil {
			account.Storage = make(map[common.Hash]common.Hash)
		}
		account.Balance = p.GetBalance(address)
		account.Nonce = p.GetNonce(address)
		ethGen.Alloc[address] = account
		return false
	})

	// Iterate Storage and set the genesis accounts.
	p.IterateState(func(address common.Address, key common.Hash, value common.Hash) bool {
		account, ok := ethGen.Alloc[address]
		if !ok {
			account = core.GenesisAccount{}
		}
		if account.Storage == nil {
			account.Storage = make(map[common.Hash]common.Hash)
		}
		account.Storage[key] = value

		account.Code = p.GetCode(address)
		account.Balance = p.GetBalance(address)
		account.Nonce = p.GetNonce(address)
		ethGen.Alloc[address] = account

		return false
	})

	// Iterate Code and set the genesis accounts.
	p.IterateCode(func(address common.Address, codeHash common.Hash) bool {
		account, ok := ethGen.Alloc[address]
		if !ok {
			account = core.GenesisAccount{}
		}
		account.Code = p.GetCode(address)
		account.Nonce = p.GetNonce(address)
		if account.Balance == nil {
			account.Balance = big.NewInt(0)
		}
		ethGen.Alloc[address] = account

		return false
	})
}
