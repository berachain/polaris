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
	"pkg.berachain.dev/stargazer/eth/accounts/abi"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	libtypes "pkg.berachain.dev/stargazer/lib/types"
)

type (
	// `Plugin` defines the methods that the chain running Stargazer EVM should implement in order
	// to support running their own stateful precompiled contracts. Implementing this plugin is
	// optional.
	Plugin interface {
		// `Reset` sets the native precompile context before beginning a state transition.
		libtypes.Resettable
		// `Register` registers a new precompiled contract at the given address.
		Register(vm.PrecompileContainer) error
		// `PrecompileManager` is the manager for the native precompiles.
		vm.PrecompileManager
	}
)

type (
	// `StatelessImpl` is the interface for all stateless precompiled contract
	// implementations. A stateless contract must provide its own precompile container, as it is
	// stateless in nature. This requires a deterministic gas count, `RequiredGas`, and an
	// executable function `Run`.
	StatelessImpl interface {
		vm.RegistrablePrecompile

		vm.PrecompileContainer
	}

	// `StatefulImpl` is the interface for all stateful precompiled contracts, which must
	// expose their ABI methods and precompile methods for stateful execution.
	StatefulImpl interface {
		vm.RegistrablePrecompile

		// `ABIMethods` should return a map of Ethereum method names to Go-Ethereum abi `Method`
		// structs. NOTE: this can be directly loaded from the `Methods` field of a Go-Ethereum ABI
		// struct, which can be built for a solidity interface or contract.
		ABIMethods() map[string]abi.Method

		// `PrecompileMethods` should return all the stateful precompile's functions (and each of
		// their required gas).
		PrecompileMethods() Methods

		// `ABIEvents()` should return a map of Ethereum event names to Go-Ethereum abi `Event`.
		// NOTE: this can be directly loaded from the `Events` field of a Go-Ethereum ABI struct,
		// which can be built for a solidity library, interface, or contract.
		ABIEvents() map[string]abi.Event

		// `CustomValueDecoders` should return a map of event attribute keys to value decoder
		// functions. This is used to decode event attribute values that require custom decoding
		// logic.
		CustomValueDecoders() ValueDecoders
	}

	// `DynamicImpl` is the interface for all dynamic stateful precompiled contracts.
	DynamicImpl interface {
		StatefulImpl

		// `Name` should return a string name of the dynamic contract.
		Name() string
	}
)

type (
	// `ValueDecoder` is a type of function that returns a geth compatible, eth primitive type (as
	// type `any`) for a given event attribute value (of type `string`). Event attribute values may
	// require unique decodings based on their underlying string encoding.
	ValueDecoder func(attributeValue string) (ethPrimitive any, err error)
	// `ValueDecoders` is a type that represents a map of event attribute keys to value decoder
	// functions.
	ValueDecoders map[string]ValueDecoder
)
