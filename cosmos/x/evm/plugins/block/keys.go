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

package block

import (
	"github.com/berachain/polaris/cosmos/x/evm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// headerHashKeySize is the number of bytes in the header hash key: 1 (prefix) + 8 (block height).
const headerHashKeySize = 9

// headerHashKeyForHeight returns the key for the hash of the header at the given height.
func headerHashKeyForHeight(number int64) []byte {
	bz := make([]byte, headerHashKeySize)
	copy(bz, []byte{types.HeaderHashKeyPrefix})
	copy(bz[1:], sdk.Uint64ToBigEndian(uint64(number%prevHeaderHashes)))
	return bz
}
