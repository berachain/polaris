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
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
)

// const testBaseFee = 69

//go:generate moq -out ./block_plugin.mock.go -pkg mock ../ BlockPlugin

func NewBlockPluginMock() *BlockPluginMock {
	bp := &BlockPluginMock{
		PrepareFunc: func(contextMoqParam context.Context) {},
		StoreHeaderFunc: func(header *types.Header) error {
			return nil
		},
		GetNewBlockMetadataFunc: func(n uint64) (common.Address, uint64) {
			return common.BytesToAddress([]byte{2}), uint64(3)
		},
	}
	bp.GetHeaderByNumberFunc = func(v uint64) (*types.Header, error) {
		if v == 0 { // handle genesis block
			return GenerateHeaderAtHeight(0), nil
		}
		for _, call := range bp.StoreHeaderCalls() {
			if call.Header.Number.Uint64() == v {
				return types.CopyHeader(call.Header), nil
			}
		}
		return nil, core.ErrHeaderNotFound
	}
	return bp
}

func GenerateHeaderAtHeight(height int64) *types.Header {
	return &types.Header{
		ParentHash:  common.Hash{0x01},
		UncleHash:   common.Hash{0x02},
		Coinbase:    common.Address{0x03},
		Root:        common.Hash{0x04},
		TxHash:      common.Hash{0x05},
		ReceiptHash: common.Hash{0x06},
		Number:      big.NewInt(height),
		BaseFee:     big.NewInt(69),
	}
}
