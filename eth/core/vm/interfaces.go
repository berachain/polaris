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
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/params"
	libtypes "github.com/berachain/stargazer/lib/types"

	"context"
	"math/big"

	"github.com/berachain/stargazer/eth/common"
)

type (
	// `StargazerEVM` defines an extension to the interface provided by Go-Ethereum to support additional
	// state transition functionalities.
	StargazerEVM interface {
		Reset(txCtx TxContext, sdb GethStateDB)
		Create(caller ContractRef, code []byte,
			gas uint64, value *big.Int,
		) (ret []byte, contractAddr common.Address, leftOverGas uint64, err error)
		Call(caller ContractRef, addr common.Address, input []byte,
			gas uint64, value *big.Int,
		) (ret []byte, leftOverGas uint64, err error)

		SetTxContext(txCtx TxContext)
		SetTracer(tracer EVMLogger)
		SetDebug(debug bool)
		StateDB() StargazerStateDB
		TxContext() TxContext
		Tracer() EVMLogger
		Context() BlockContext
		ChainConfig() *params.EthChainConfig
	}

	// `StargazerStateDB` defines an extension to the interface provided by Go-Ethereum to support
	// additional state transition functionalities.
	StargazerStateDB interface {
		GethStateDB
		libtypes.Finalizeable

		// `Reset` resets the context for the new transaction.
		Reset(context.Context)

		// `TransferBalance` transfers the balance from one account to another
		TransferBalance(common.Address, common.Address, *big.Int)

		// `BuildLogsAndClear` builds the logs for the tx with the given metadata. NOTE: must be
		// called after `Finalize`.
		BuildLogsAndClear(common.Hash, common.Hash, uint, uint) []*coretypes.Log
	}

	// `RegistrablePrecompile` is a type for the base precompile implementation, which only needs to
	// provide an Ethereum address of where its contract is found.
	RegistrablePrecompile = libtypes.Registrable[common.Address]
)
