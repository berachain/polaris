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
	"context"

	"github.com/berachain/stargazer/eth/core/types"
)

const testBaseFee = 69

//go:generate moq -out ./block_plugin.mock.go -pkg mock ../ BlockPlugin
func NewBlockPluginMock() *BlockPluginMock {
	// make and configure a mocked core.BlockPlugin
	mockedBlockPlugin := &BlockPluginMock{
		BaseFeeFunc: func() uint64 {
			return testBaseFee
		},
		GetStargazerHeaderAtHeightFunc: func(n int64) *types.StargazerHeader {
			panic("mock out the GetStargazerHeaderAtHeight method")
		},
		PrepareFunc: func(contextMoqParam context.Context) {
			// no-op
		},
	}
	return mockedBlockPlugin
}
