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
