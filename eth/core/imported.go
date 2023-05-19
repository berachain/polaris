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
	// ChainContext provides information about the current blockchain to the EVM.
	ChainContext = core.ChainContext
	// ChainEvent contains information about the chain.
	ChainEvent = core.ChainEvent
	// ChainHeadEvent is posted when a new head block is added to the chain.
	ChainHeadEvent = core.ChainHeadEvent
	// ChainSideEvent is posted when a new side block is added to the chain.
	ChainSideEvent = core.ChainSideEvent
	// ExecutionResult is the result of executing a transaction.
	ExecutionResult = core.ExecutionResult
	// GasPool is a pool of gas that can be consumed by transactions.
	GasPool = core.GasPool
	// NewTxsEvent is posted when a batch of transactions enter the transaction pool.
	NewTxsEvent = core.NewTxsEvent
	// Message contains data used ype used to execute transactions.
	Message = core.Message
	// RemovedLogsEvent is posted pre-removal of a set of logs.
	RemovedLogsEvent = core.RemovedLogsEvent
)

var (
	// ApplyTransactionWithEVM applies a transaction to the current state of the blockchain.
	ApplyTransactionWithEVMWithResult = core.ApplyTransactionWithEVMWithResult
	// NewEVMTxContext creates a new context for use in the EVM.
	NewEVMTxContext = core.NewEVMTxContext
	// NewEVMBlockContext creates a new block context for a given header.
	NewEVMBlockContext = core.NewEVMBlockContext
	// GetHashFn returns a GetHashFunc.
	GetHashFn = core.GetHashFn
	// TransactionToMessage converts a transaction to a message.
	TransactionToMessage = core.TransactionToMessage

	Transfer    = core.Transfer
	CanTransfer = core.CanTransfer
)

var (
	// ErrInsufficientBalanceForGas is the error return when gas required to execute a transaction overflows.
	ErrGasUintOverflow = core.ErrGasUintOverflow
)
