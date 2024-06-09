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
