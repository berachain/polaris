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
	"time"

	"github.com/berachain/stargazer/eth/common"
	ethereumcorevm "github.com/ethereum/go-ethereum/core/vm"
)

//go:generate moq -out ./logger.mock.go -pkg mock ../ EVMLogger

func NewEVMLoggerMock() *EVMLoggerMock {
	mockedEVMLogger := &EVMLoggerMock{
		CaptureEndFunc: func(output []byte, gasUsed uint64, t time.Duration, err error) {
			// no-op
		},
		CaptureEnterFunc: func(typ ethereumcorevm.OpCode,
			from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
			// no-op
		},
		CaptureExitFunc: func(output []byte, gasUsed uint64, err error) {
			// no-op
		},
		CaptureFaultFunc: func(pc uint64,
			op ethereumcorevm.OpCode, gas uint64, cost uint64,
			scope *ethereumcorevm.ScopeContext, depth int, err error) {
			// no-op
		},
		CaptureStartFunc: func(env *ethereumcorevm.EVM,
			from common.Address, to common.Address, create bool, input []byte, gas uint64,
			value *big.Int) {
			// no-op
		},
		CaptureStateFunc: func(pc uint64,
			op ethereumcorevm.OpCode, gas uint64, cost uint64,
			scope *ethereumcorevm.ScopeContext, rData []byte, depth int, err error) {
			// no-op
		},
		CaptureTxEndFunc: func(restGas uint64) {
			// no-op
		},
		CaptureTxStartFunc: func(gasLimit uint64) {
			// no-op
		},
	}
	return mockedEVMLogger
}
