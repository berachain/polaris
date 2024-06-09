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

package log

import (
	"github.com/berachain/polaris/cosmos/x/evm/plugins/state/events"
	"github.com/berachain/polaris/eth/core/precompile"
	"github.com/berachain/polaris/lib/registry"
	libtypes "github.com/berachain/polaris/lib/types"
	"github.com/berachain/polaris/lib/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// Factory is a `PrecompileLogFactory` that builds Ethereum logs from Cosmos events. All Ethereum
// events must be registered with the factory before it can build logs during state transitions.
type Factory struct {
	// events is a registry of precompile logs, indexed by the Cosmos event type.
	events libtypes.Registry[string, *precompileLog]
	// customValueDecoders is a map of Cosmos attribute keys to attribute value decoder
	// functions for custom events.
	customValueDecoders precompile.ValueDecoders
}

// NewFactory returns a `Factory` with the events and custom value decoders of the given
// precompiles registered.
func NewFactory(precompiles []precompile.Registrable) *Factory {
	f := &Factory{
		events:              registry.NewMap[string, *precompileLog](),
		customValueDecoders: make(precompile.ValueDecoders),
	}
	f.registerAllEvents(precompiles)
	return f
}

// Build builds an Ethereum log from a Cosmos event.
//
// Build implements `events.PrecompileLogFactory`.
func (f *Factory) Build(event *sdk.Event) (*ethtypes.Log, error) {
	// get the precompile log for the Cosmos event type
	pl := f.events.Get(event.Type)
	if pl == nil {
		return nil, events.ErrEthEventNotRegistered
	}

	var err error

	// validate the Cosmos event attributes
	if err = validateAttributes(pl, event); err != nil {
		return nil, err
	}

	// build the Ethereum log
	log := &ethtypes.Log{
		Address: pl.precompileAddr,
	}
	if log.Topics, err = f.makeTopics(pl, event); err != nil {
		return nil, err
	}
	if log.Data, err = f.makeData(pl, event); err != nil {
		return nil, err
	}

	return log, nil
}

// registerAllEvents registers all Ethereum events from the provided precompiles with the factory.
func (f *Factory) registerAllEvents(precompiles []precompile.Registrable) {
	for _, pc := range precompiles {
		if spc, ok := utils.GetAs[precompile.StatefulImpl](pc); ok {
			// register the ABI Event as a precompile log
			moduleEthAddr := spc.RegistryKey()
			for _, event := range spc.ABIEvents() {
				_ = f.events.Register(newPrecompileLog(moduleEthAddr, event))
			}

			// register the precompile's custom value decoders, if any are provided
			for attr, decoder := range spc.CustomValueDecoders() {
				f.customValueDecoders[attr] = decoder
			}
		}
	}
}
