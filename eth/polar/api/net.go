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

package polarapi

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// NetBackend is the collection of methods required to satisfy the net
// RPC API.
type NetBackend interface {
	NetAPI
}

// NetAPI is the collection of net RPC API methods.
type NetAPI interface {
	PeerCount() hexutil.Uint
	Listening() bool
	Version() string
}

// netAPI offers network related RPC methods.
type netAPI struct {
	b NetBackend
}

// NewNetAPI creates a new net API instance.
func NewNetAPI(b NetBackend) NetAPI {
	return &netAPI{b}
}

// Listening returns an indication if the node is listening for network connections.
func (api *netAPI) Listening() bool {
	return api.b.Listening()
}

// PeerCount returns the number of connected peers.
func (api *netAPI) PeerCount() hexutil.Uint {
	return api.b.PeerCount()
}

// Version returns the current ethereum protocol version.
func (api *netAPI) Version() string {
	return api.b.Version()
}
