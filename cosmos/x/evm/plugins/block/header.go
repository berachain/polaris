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
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	errorslib "pkg.berachain.dev/polaris/lib/errors"
)

// ===========================================================================
// Polaris Block Header Tracking
// ===========================================================================.

// SetQueryContextFn sets the query context func for the plugin.
func (p *plugin) SetQueryContextFn(gqc func(height int64, prove bool) (sdk.Context, error)) {
	p.getQueryContext = gqc
}

// GetHeaderByNumber returns the header at the given height, using the plugin's query context.
//
// GetHeaderByNumber implements core.BlockPlugin.
func (p *plugin) GetHeaderByNumber(number uint64) (*coretypes.Header, error) {
	if p.getQueryContext == nil {
		return nil, errors.New("GetHeader: getQueryContext is nil")
	}

	// TODO: ensure we aren't differing from geth / hiding errors here.
	// TODO: the GTE may be hiding a larger issue with the timing of the NewHead channel stuff.
	// Investigate and hopefully remove this GTE.
	if number > uint64(p.ctx.BlockHeight()) {
		number = uint64(p.ctx.BlockHeight())
	}

	ctx, err := p.getQueryContext(int64(number), false)
	if err != nil {
		return nil, errorslib.Wrap(err, "GetHeader: failed to use query context")
	}

	// Unmarshal the header from the context kv store.
	bz := ctx.KVStore(p.storekey).Get([]byte{types.HeaderKey})
	if bz == nil {
		return nil, errors.New("GetHeader: polaris header not found in kvstore")
	}

	header, err := coretypes.UnmarshalHeader(bz)
	if err != nil {
		return nil, errorslib.Wrap(err, "GetHeader: failed to unmarshal")
	}

	if header.Number.Uint64() != number {
		return nil, errorslib.Wrapf(err,
			"GetHeader: header number mismatch, got %d, expected %d",
			header.Number.Uint64(), number)
	}

	return header, nil
}

// SetHeader saves a block to the store.
func (p *plugin) StoreHeader(header *coretypes.Header) error {
	bz, err := coretypes.MarshalHeader(header)
	if err != nil {
		return errorslib.Wrap(err, "SetHeader: failed to marshal header")
	}
	p.ctx.KVStore(p.storekey).Set([]byte{types.HeaderKey}, bz)
	return nil
}
