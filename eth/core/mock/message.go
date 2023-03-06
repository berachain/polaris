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

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/types"
)

//go:generate moq -out ./message.mock.go -pkg mock ../ Message

func NewEmptyMessage() *MessageMock {
	mockedMessage := &MessageMock{
		AccessListFunc: func() types.AccessList {
			return nil
		},
		DataFunc: func() []byte {
			return nil
		},
		FromFunc: func() common.Address {
			return common.Address{}
		},
		GasFunc: func() uint64 {
			return 0
		},
		GasFeeCapFunc: func() *big.Int {
			return big.NewInt(0)
		},
		GasPriceFunc: func() *big.Int {
			return big.NewInt(0)
		},
		GasTipCapFunc: func() *big.Int {
			return big.NewInt(0)
		},
		IsFakeFunc: func() bool {
			return false
		},
		NonceFunc: func() uint64 {
			return 0
		},
		ToFunc: func() *common.Address {
			return nil
		},
		ValueFunc: func() *big.Int {
			return big.NewInt(0)
		},
	}
	return mockedMessage
}
