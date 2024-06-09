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

package chain

import (
	"github.com/berachain/polaris/eth/core"
)

// WrappedBlockchain is a struct that wraps the core blockchain with additional
// application context.
type WrappedBlockchain struct {
	core.Blockchain           // chain is the core blockchain.
	app             txDecoder // App is the application context.
}

// New creates a new instance of WrappedBlockchain with the provided core blockchain
// and application context.
func New(chain core.Blockchain, app txDecoder) *WrappedBlockchain {
	return &WrappedBlockchain{Blockchain: chain, app: app}
}

func (wbc *WrappedBlockchain) SetBlockchain(chain core.Blockchain) {
	wbc.Blockchain = chain
}
