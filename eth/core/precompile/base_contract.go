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

package precompile

import (
	"github.com/berachain/polaris/eth/accounts/abi"

	"github.com/ethereum/go-ethereum/common"
)

// ==============================================================================
// Precompile Injector
// ==============================================================================

// Injector is a precompile injector, that allows for precompiles to be injected with
// the Cosmos depinject framework.
type Injector struct {
	// precompiles stores the precompiles.
	precompiles []Registrable
}

func NewPrecompiles(precompiles ...Registrable) *Injector {
	return &Injector{
		precompiles: precompiles,
	}
}

// GetPrecompiles implements Precompiles.
func (pci *Injector) GetPrecompiles() []Registrable {
	return pci.precompiles
}

// AddPrecompile adds a new precompile to the injector.
func (pci *Injector) AddPrecompile(precompile Registrable) {
	pci.precompiles = append(pci.precompiles, precompile)
}

// ==============================================================================
// Base Precompile
// ==============================================================================

type BaseContract interface {
	StatefulImpl
	GetPlugin() Plugin
}

// baseContract is a base implementation of `StatefulImpl`.
type baseContract struct {
	// abi stores the ABI of the precompile.
	abi abi.ABI
	// address stores the address of the precompile.
	address common.Address
	// plugin stores the core precompile plugin.
	plugin Plugin
}

// NewBaseContract creates a new `BasePrecompile`.
func NewBaseContract(abiStr string, address common.Address) BaseContract {
	return &baseContract{
		abi:     abi.MustUnmarshalJSON(abiStr),
		address: address,
	}
}

// RegistryKey implements StatefulImpl.
func (c *baseContract) RegistryKey() common.Address {
	return c.address
}

// ABIMethods implements StatefulImpl.
func (c *baseContract) ABIMethods() map[string]abi.Method {
	return c.abi.Methods
}

// ABIEvents implements StatefulImpl.
func (c *baseContract) ABIEvents() map[string]abi.Event {
	return c.abi.Events
}

// CustomValueDecoders implements StatefulImpl.
func (c *baseContract) CustomValueDecoders() ValueDecoders {
	return nil
}

// SetPlugin implements BaseContract.
func (c *baseContract) SetPlugin(plugin Plugin) {
	c.plugin = plugin
}

// GetPlugin implements BaseContract.
func (c *baseContract) GetPlugin() Plugin {
	return c.plugin
}
