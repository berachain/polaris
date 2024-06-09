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

package core

import (
	"errors"

	"github.com/berachain/polaris/eth/core/state"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// WriteGenesisBlock inserts the genesis block into the blockchain.
func (bc *blockchain) WriteGenesisBlock(block *ethtypes.Block) error {
	// Get the state with the latest finalize block context.
	sp := bc.spf.NewPluginWithMode(state.Genesis)
	state := state.NewStateDB(sp, bc.pp)

	// TODO: add more validation here.
	if block.NumberU64() != 0 {
		return errors.New("not the genesis block")
	}
	_, err := bc.WriteBlockAndSetHead(block, nil, nil, state, true)
	return err
}
