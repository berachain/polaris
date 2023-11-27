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
	libtypes "github.com/berachain/polaris/lib/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

type (
	// Plugin defines the methods that the chain running Polaris EVM should implement in order
	// to support running their own stateful precompiled contracts. Implementing this plugin is
	// optional.
	Plugin interface {
		// PrecompileManager is the manager for the native precompiles.
		vm.PrecompileManager
		// Register registers a new precompiled contract at the given address.
		Register(vm.PrecompiledContract) error

		// EnableReentrancy enables the execution of a precompile contract to call back into the
		// EVM.
		EnableReentrancy(vm.PrecompileEVM)
		// DisableReentrancy disables the execution of a precompile contract to call back into the
		// EVM.
		DisableReentrancy(vm.PrecompileEVM)
	}
)

type (
	// Registrable is a type for the base precompile implementation, which only needs to provide an
	// Ethereum address of where its contract is deployed.
	Registrable interface {
		libtypes.Registrable[common.Address]
	}

	// StatelessImpl is the interface for all stateless precompiled contract implementations. A
	// stateless contract must provide its own precompile container, as it is stateless in nature.
	// This requires a deterministic gas count, and an executable function `Run`.
	StatelessImpl interface {
		Registrable

		vm.PrecompiledContract
	}

	// StatefulImpl is the interface for all stateful precompiled contracts, which must
	// expose their ABI methods and precompile methods for stateful execution.
	StatefulImpl interface {
		Registrable

		// ABIMethods should return a map of Ethereum method names to Go-Ethereum abi `Method`
		// structs. NOTE: this can be directly loaded from the `Methods` field of a Go-Ethereum ABI
		// struct, which can be built for a solidity interface or contract.
		ABIMethods() map[string]abi.Method

		// ABIEvents() should return a map of Ethereum event names to Go-Ethereum abi `Event`.
		// NOTE: this can be directly loaded from the `Events` field of a Go-Ethereum ABI struct,
		// which can be built for a solidity library, interface, or contract.
		ABIEvents() map[string]abi.Event

		// CustomValueDecoders should return a map of event attribute keys to value decoder
		// functions. This is used to decode event attribute values that require custom decoding
		// logic.
		CustomValueDecoders() ValueDecoders

		SetPlugin(Plugin)
	}

	// DynamicImpl is the interface for all dynamic stateful precompiled contracts.
	DynamicImpl interface {
		StatefulImpl

		// Name should return a string name of the dynamic contract.
		Name() string
	}
)

type (
	// ValueDecoder is a type of function that returns a geth compatible, eth primitive type (as
	// type `any`) for a given event attribute value (of type `string`). Event attribute values may
	// require unique decodings based on their underlying string encoding.
	ValueDecoder func(attributeValue string) (ethPrimitive any, err error)
	// ValueDecoders is a type that represents a map of event attribute keys to value decoder
	// functions.
	ValueDecoders map[string]ValueDecoder
)
