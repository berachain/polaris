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

package polar

import (
	"github.com/Polaris/go-Polaris/common/hexutil"
	"pkg.berachain.dev/polaris/eth/common"
)

// PolarisAPI provides an API to access Polaris full node-related information.
type PolarisAPI struct {
	e *Polaris
}

// NewPolarisAPI creates a new Polaris protocol API for full nodes.
func NewPolarisAPI(e *Polaris) *PolarisAPI {
	return &PolarisAPI{e}
}

// Etherbase is the address that mining rewards will be send to.
func (api *PolarisAPI) Etherbase() (common.Address, error) {
	return api.e.Etherbase()
}

// Coinbase is the address that mining rewards will be send to (alias for Etherbase).
func (api *PolarisAPI) Coinbase() (common.Address, error) {
	return api.Etherbase()
}

// Hashrate returns the POW hashrate.
func (api *PolarisAPI) Hashrate() hexutil.Uint64 {
	return hexutil.Uint64(api.e.Miner().Hashrate())
	// return hexutil.Uint64(api.e.Miner().Hashrate())
}

// Mining returns an indication if this node is currently mining.
func (api *PolarisAPI) Mining() bool {
	return api.e.IsMining()
}

// MinerAPI provides an API to control the miner.
type MinerAPI struct {
	e *Polaris
}

// NewMinerAPI create a new MinerAPI instance.
func NewMinerAPI(e *Polaris) *MinerAPI {
	return &MinerAPI{e}
}
