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
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/rpc"
	errorslib "pkg.berachain.dev/polaris/lib/errors"
)

// ===========================================================================
// Polaris Block Header Tracking
// ===========================================================================.

// `SetQueryContextFn` sets the query context func for the plugin.
func (p *plugin) SetQueryContextFn(gqc func(height int64, prove bool) (sdk.Context, error)) {
	p.getQueryContext = gqc
}

// `GetHeaderByNumber` returns the header at the given height, using the plugin's query context.
//
// `GetHeaderByNumber` implements core.BlockPlugin.
func (p *plugin) GetHeaderByNumber(height int64) (*coretypes.Header, error) {
	if p.getQueryContext == nil {
		return nil, errors.New("GetHeader: getQueryContext is nil")
	}

	iavlHeight, err := p.getIAVLHeight(height)
	if err != nil {
		return nil, errors.New("GetHeader: invalid IAVL height")
	}

	ctx, err := p.getQueryContext(iavlHeight, false)
	if err != nil {
		return nil, errors.New("GetHeader: failed to use query context")
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

	if int64(header.Number.Uint64()) != iavlHeight {
		panic("header number is not equal to the given iavl tree height")
	}

	return header, nil
}

// `SetHeader` saves a block to the store.
func (p *plugin) SetHeaderByNumber(_ int64, header *coretypes.Header) error {
	bz, err := coretypes.MarshalHeader(header)
	if err != nil {
		return err
	}
	p.ctx.KVStore(p.storekey).Set([]byte{types.HeaderKey}, bz)
	return nil
}

// `getIAVLHeight` returns the IAVL height for the given block number.
func (p *plugin) getIAVLHeight(number int64) (int64, error) {
	var iavlHeight int64
	switch rpc.BlockNumber(number) { //nolint:nolintlint,exhaustive // covers all cases.
	case rpc.SafeBlockNumber, rpc.FinalizedBlockNumber:
		iavlHeight = p.ctx.BlockHeight() - 1
	case rpc.PendingBlockNumber, rpc.LatestBlockNumber:
		iavlHeight = p.ctx.BlockHeight()
	case rpc.EarliestBlockNumber:
		iavlHeight = 1
	default:
		iavlHeight = number
	}

	if iavlHeight < 1 {
		return 1, fmt.Errorf("invalid block number %d", number)
	}

	return iavlHeight, nil
}
