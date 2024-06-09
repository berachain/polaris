// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
