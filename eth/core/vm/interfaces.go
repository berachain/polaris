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

package vm

import (
	"math/big"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/params"
	libtypes "pkg.berachain.dev/stargazer/lib/types"
)

type (
	// `StargazerEVM` defines an extension to the interface provided by Go-Ethereum to support
	// additional state transition functionalities.
	StargazerEVM interface {
		// `Create` creates a new contract using the given code and returns the contract address.
		Create(
			caller ContractRef, code []byte, gas uint64, value *big.Int,
		) (ret []byte, contractAddr common.Address, leftOverGas uint64, err error)
		// `Call` executes the given contract with the given input data and returns the output.
		Call(
			caller ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int,
		) (ret []byte, leftOverGas uint64, err error)
		// `SetTxContext` sets the transaction context for the new transaction.
		SetTxContext(txCtx TxContext)
		// `StateDB` returns the `StateDB` attached to this `StargazerEVM`.
		StateDB() StargazerStateDB
		// `ChainConfig` returns the `ChainConfig` attached to this `StargazerEVM`.
		ChainConfig() *params.ChainConfig
		UnderlyingEVM() *GethEVM
	}

	// `StargazerStateDB` defines an extension to the interface provided by Go-Ethereum to support
	// additional state transition functionalities.
	StargazerStateDB interface {
		GethStateDB
		libtypes.Finalizeable
		// `Reset` resets the context for the new transaction.
		libtypes.Resettable
		// `TransferBalance` transfers the balance from one account to another
		TransferBalance(common.Address, common.Address, *big.Int)
	}

	// `RegistrablePrecompile` is a type for the base precompile implementation, which only needs to
	// provide an Ethereum address of where its contract is found.
	RegistrablePrecompile = libtypes.Registrable[common.Address]
)
