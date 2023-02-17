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

package core

import (
	"github.com/ethereum/go-ethereum/core"
)

type (
	// `ExecutionResult` is the result of executing a transaction.
	ExecutionResult = core.ExecutionResult
	// `Message` contains data used ype used to execute transactions.
	Message = core.Message
)

var (
	// `NewEVMTxContext` creates a new context for use in the EVM.
	NewEVMTxContext = core.NewEVMTxContext
)

var (
	// `ErrInsufficientFundsForTransfer` is the error returned when the account does not have enough funds to transfer.
	ErrInsufficientFundsForTransfer = core.ErrInsufficientFundsForTransfer
	// `ErrInsufficientFunds` is the error returned when the account does not have enough funds to execute the transaction.
	ErrInsufficientFunds = core.ErrInsufficientFunds
	// `ErrInsufficientBalanceForGas` is the error return when gas required to execute a transaction overflows.
	ErrGasUintOverflow = core.ErrGasUintOverflow
	// `ErrIntrinsicGas` is the error returned when the transaction does not have enough gas to cover the intrinsic cost.
	ErrIntrinsicGas = core.ErrIntrinsicGas
	// `IntrinsicGas` is the intrinsic gas of a transaction.
	EthIntrinsicGas = core.IntrinsicGas
)
