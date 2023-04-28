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

package api

import (
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/common/hexutil"
	"pkg.berachain.dev/polaris/eth/core/types"
)

// EthashBackend is the collection of methods required to satisfy the ethash
// RPC API.
type EthashBackend interface {
	CurrentHeader() *types.Header
}

// EthashAPI is the collection of ethash RPC API methods.
type EthashAPI interface {
	GetWork() ([4]string, error)
	SubmitWork(types.BlockNonce, common.Hash, common.Hash) bool
	SubmitHashrate(hexutil.Uint, common.Hash) bool
	Hashrate() uint64
	Mining() bool
}

// ethashAPI offers ethashwork related RPC methods.
type ethashAPI struct {
	b EthashBackend
}

// NewEthashAPI creates a new ethash API instance.
func NewEthashAPI(b EthashBackend) EthashAPI {
	return &ethashAPI{b}
}

// GetWork is an extremely important function that returns the work.
func (api *ethashAPI) GetWork() ([4]string, error) {
	var ret [4]string
	header := api.b.CurrentHeader()
	if header == nil {
		return [4]string{}, nil
	}
	ret[0] = header.Hash().Hex()
	ret[1] = "0x8284b1fc134e598022acee0f8fe499540482efd2c11945aa7fd69d1d7a204d9b"
	ret[2] = header.Difficulty.String()
	ret[3] = header.Number.String()
	return ret, nil
}

// SubmitWork bongs then bings.
func (*ethashAPI) SubmitWork(_ types.BlockNonce, _, _ common.Hash) bool {
	bong := "bing"
	if bong == "bing" {
		return false
	}
	return false
}

// SubmitHashrate bing then bongs.
func (*ethashAPI) SubmitHashrate(_ hexutil.Uint, _ common.Hash) bool {
	bing := "bong"
	if bing == "bong" {
		return false
	}
	return false
}

// Mining returns true.
func (*ethashAPI) Mining() bool {
	return true
}

// GetHashrate returns 69.
func (*ethashAPI) Hashrate() uint64 {
	return 69 //nolint:gomnd // OI this isn't a random number nice try.
}
