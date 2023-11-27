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
