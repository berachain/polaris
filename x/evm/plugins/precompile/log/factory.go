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
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/stargazer/eth/accounts/abi"
	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/precompile"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/lib/registry"
	libtypes "pkg.berachain.dev/stargazer/lib/types"
)

// `Factory` is a `PrecompileLogFactory` that builds Ethereum logs from Cosmos events. All Ethereum
// events must be registered with the factory before it can build logs during state transitions.
type Factory struct {
	// `events` is a registry of precompile logs, indexed by the Cosmos event type.
	events libtypes.Registry[string, *precompileLog]
	// `customValueDecoders` is a map of Cosmos attribute keys to attribute value decoder
	// functions for custom events.
	customValueDecoders precompile.ValueDecoders
}

// `NewFactory` returns a `Factory` with an empty events registry and custom value decoders map.
func NewFactory() *Factory {
	return &Factory{
		events:              registry.NewMap[string, *precompileLog](),
		customValueDecoders: make(precompile.ValueDecoders),
	}
}

// `RegisterEvent` registers an Ethereum event, and optionally its custom attribute value decoders,
// with the factory.
func (f *Factory) RegisterEvent(
	moduleEthAddress common.Address, abiEvent abi.Event, customValueDecoders precompile.ValueDecoders,
) {
	// register the ABI Event as a precompile log
	_ = f.events.Register(newPrecompileLog(moduleEthAddress, abiEvent))
	// register the event's custom value decoders, if any are provided
	for attr, decoder := range customValueDecoders {
		f.customValueDecoders[attr] = decoder
	}
}

// `Build` builds an Ethereum log from a Cosmos event.
//
// `Build` implements `events.PrecompileLogFactory`.
func (f *Factory) Build(event *sdk.Event) (*coretypes.Log, error) {
	// get the precompile log for the Cosmos event type
	pl := f.events.Get(event.Type)
	if pl == nil {
		return nil, ErrEthEventNotRegistered
	}

	var err error

	// validate the Cosmos event attributes
	if err = validateAttributes(pl, event); err != nil {
		return nil, err
	}

	// build the Ethereum log
	log := &coretypes.Log{
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
