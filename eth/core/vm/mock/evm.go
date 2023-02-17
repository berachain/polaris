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

package mock

import (
	"math/big"

	stargazercorevm "github.com/berachain/stargazer/eth/core/vm"
	"github.com/ethereum/go-ethereum/common"
	ethereumcorevm "github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

//go:generate moq -out ./evm.mock.go -pkg mock ../ StargazerEVM

func NewStargazerEVM() *StargazerEVMMock {
	mockedStargazerEVM := &StargazerEVMMock{
		CallFunc: func(caller ethereumcorevm.ContractRef, addr common.Address,
			input []byte, gas uint64, value *big.Int) ([]byte, uint64, error) {
			return []byte{}, 0, nil
		},
		ChainConfigFunc: func() *params.ChainConfig {
			return &params.ChainConfig{
				LondonBlock:    big.NewInt(0),
				HomesteadBlock: big.NewInt(0),
			}
		},
		ContextFunc: func() ethereumcorevm.BlockContext {
			return stargazercorevm.BlockContext{
				CanTransfer: func(db stargazercorevm.GethStateDB, addr common.Address, amount *big.Int) bool {
					return true
				},
				BlockNumber: big.NewInt(1), // default to block == 1 to pass all forks,
			}
		},
		CreateFunc: func(caller ethereumcorevm.ContractRef, code []byte,
			gas uint64, value *big.Int) ([]byte, common.Address, uint64, error) {
			return []byte{}, common.Address{}, 0, nil
		},
		ResetFunc: func(txCtx ethereumcorevm.TxContext, sdb ethereumcorevm.StateDB) {
			panic("mock out the Reset method")
		},
		SetDebugFunc: func(debug bool) {
			// no-op
		},
		SetTracerFunc: func(tracer ethereumcorevm.EVMLogger) {
			// no-op
		},
		SetTxContextFunc: func(txCtx ethereumcorevm.TxContext) {
			// no-op
		},
		StateDBFunc: func() stargazercorevm.StargazerStateDB {
			return NewEmptyStateDB()
		},
		TracerFunc: func() ethereumcorevm.EVMLogger {
			return &EVMLoggerMock{}
		},
		TxContextFunc: func() ethereumcorevm.TxContext {
			panic("mock out the TxContext method")
		},
	}
	return mockedStargazerEVM
}
