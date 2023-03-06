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

package params

import "github.com/ethereum/go-ethereum/params"

type (
	ChainConfig = params.ChainConfig
	Rules       = params.Rules
)

var (
	// `RefundQuotient` is the refund quotient parameter.
	RefundQuotient = params.RefundQuotient
	// `RefundQuotientEIP3529` is the refund quotient parameter for EIP-3529.
	RefundQuotientEIP3529 = params.RefundQuotientEIP3529
	// `TxAccessListAddressGas` is the cost of an address for a transaction with an access list.
	TxAccessListAddressGas = params.TxAccessListAddressGas
	// `TxAccessListAddressGasEIP2930` is the cost of an address for a transaction with an access
	// list.
	TxAccessListStorageKeyGas = params.TxAccessListStorageKeyGas
	// `TxDataNonZeroGasFrontier` is the cost of a non-zero byte of data for a transaction.
	TxDataNonZeroGasFrontier = params.TxDataNonZeroGasFrontier
	// `TxDataNonZeroGasEIP2028` is the cost of a non-zero byte of data for a transaction.
	TxDataNonZeroGasEIP2028 = params.TxDataNonZeroGasEIP2028
	// `TxDataZeroGas` is the cost of a zero byte of data or code for a transaction.
	TxDataZeroGas = params.TxDataZeroGas
	// `TxGasContractCreation` is the amount of gas that is refunded for a contract creation
	// transaction.
	TxGasContractCreation = params.TxGasContractCreation
	// `TxGas` is the amount of gas that is refunded for a transaction.
	TxGas = params.TxGas
	// `MainnetChainConfig` is the chain parameters to run a node on the main network.
	MainnetChainConfig = params.MainnetChainConfig
	// `IdentityBaseGas` is the base gas required to execute the identity pre-compiled contract.
	IdentityBaseGas = params.IdentityBaseGas
)
