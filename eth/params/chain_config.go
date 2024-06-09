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

package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

const (
	artioChainID = 80085
)

var (
	zero = uint64(0)
)

var DefaultChainConfig = &params.ChainConfig{
	ChainID:                       big.NewInt(artioChainID),
	HomesteadBlock:                big.NewInt(0),
	DAOForkBlock:                  big.NewInt(0),
	DAOForkSupport:                true,
	EIP150Block:                   big.NewInt(0),
	EIP155Block:                   big.NewInt(0),
	EIP158Block:                   big.NewInt(0),
	ByzantiumBlock:                big.NewInt(0),
	ConstantinopleBlock:           big.NewInt(0),
	PetersburgBlock:               big.NewInt(0),
	IstanbulBlock:                 big.NewInt(0),
	MuirGlacierBlock:              big.NewInt(0),
	BerlinBlock:                   big.NewInt(0),
	LondonBlock:                   big.NewInt(0),
	ArrowGlacierBlock:             big.NewInt(0),
	GrayGlacierBlock:              big.NewInt(0),
	MergeNetsplitBlock:            big.NewInt(0),
	TerminalTotalDifficulty:       big.NewInt(0),
	TerminalTotalDifficultyPassed: true,
	ShanghaiTime:                  &zero,
}
