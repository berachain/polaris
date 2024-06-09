// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package mock

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

//go:generate moq -out ./state_plugin.mock.go -skip-ensure -pkg mock ../ selfDestructStatePlugin

var (
	a1 = common.HexToAddress("0x1")
	a3 = common.HexToAddress("0x3")
	a4 = common.HexToAddress("0x4")
)

func NewSelfDestructsStatePluginMock() *selfDestructStatePluginMock {
	return &selfDestructStatePluginMock{
		GetCodeHashFunc: func(address common.Address) common.Hash {
			if address == a1 || address == a3 || address == a4 {
				return common.Hash{0x1}
			}
			return common.Hash{}
		},
		GetBalanceFunc: func(address common.Address) *big.Int {
			return new(big.Int)
		},
		SubBalanceFunc: func(address common.Address, amount *big.Int) {
			// no-op
		},
	}
}
