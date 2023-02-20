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
	"github.com/berachain/stargazer/eth/common"
	coretypes "github.com/berachain/stargazer/eth/core/types"
)

// `GetStargazerHeader` returns the stargazer header at the given height.
func (p *plugin) GetStargazerBlockAtHeight(height int64) *coretypes.StargazerBlock {
	// Get the stargazer header at the given height.
	// header, ok := p.shg.GetStargazerHeader(p.ctx, height)
	// if !ok {
	// 	return nil, fmt.Errorf("stargazer header not found at height %d", height)
	// }
	// // Get the stargazer block at the given height.
	// block, ok := p.shg.GetStargazerBlock(p.ctx, height)
	// if !ok {
	// 	return nil, fmt.Errorf("stargazer block not found at height %d", height)
	// }
	// // Return the stargazer block.
	// return &StargazerBlock{
	// 	Header: header,
	// 	Block:  block,
	// }, nil
	return nil
}

func (p *plugin) GetStargazerBlockByHash(hash common.Hash) *coretypes.StargazerBlock {
	// // Get the stargazer header at the given height.
	// header, ok := p.shg.GetStargazerHeaderByHash(p.ctx, hash)
	// if !ok {
	// 	return nil
	// }
	// Return the stargazer block.
	return nil
}
