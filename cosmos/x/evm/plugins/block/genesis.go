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
	"github.com/berachain/polaris/eth/core"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// InitGenesis stores the genesis block header in the KVStore under its own genesis key.
func (p *plugin) InitGenesis(ctx sdk.Context, ethGen *core.Genesis) error {
	p.Prepare(ctx)

	// Writing genesis block 0 to disk, available to query from any future IAVL height
	return p.StoreHeader(ethGen.ToBlock().Header())
}

// Export genesis modifies a pointer to a genesis state object and populates it.
func (p *plugin) ExportGenesis(ctx sdk.Context, ethGen *core.Genesis) {
	p.Prepare(ctx)

	header, err := p.getGenesisHeader()
	if err != nil {
		panic(err)
	}

	core.UnmarshalGenesisHeader(header, ethGen)
}

// getGenesisHeader returns the block header at height 0 and does a sanity check.
func (p *plugin) getGenesisHeader() (*ethtypes.Header, error) {
	header, err := p.GetHeaderByNumber(0)
	if err != nil {
		return nil, err
	}

	if err = header.SanityCheck(); err != nil {
		return nil, err
	}

	return header, nil
}
