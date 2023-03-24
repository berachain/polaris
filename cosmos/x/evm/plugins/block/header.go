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
	"context"
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

// SetQueryContextFn sets the query context func for the plugin.
func (p *plugin) SetQueryContextFn(gqc func(height int64, prove bool) (sdk.Context, error)) {
	p.getQueryContext = gqc
}

// GetHeaderByNumber returns the header at the given height. It verifies the height and determines
// whether to use the current context height or the plugin's query context for a historical height.
//
// GetHeaderByNumber implements core.BlockPlugin.
func (p *plugin) GetHeaderByNumber(ctx context.Context, height int64) (*coretypes.Header, error) {
	if p.getQueryContext == nil {
		return nil, errors.New("GetHeader: getQueryContext is nil")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentHeight := sdkCtx.BlockHeight()
	iavlHeight, err := p.resolveBlockHeight(height, currentHeight)
	if err != nil {
		return nil, errorslib.Wrapf(err, "GetHeader: invalid IAVL height")
	}

	var header *coretypes.Header
	if iavlHeight == currentHeight {
		header, err = p.getHeaderFromStore(sdkCtx)
	} else {
		header, err = p.GetHeaderByHistoricalNumber(iavlHeight)
	}

	if int64(header.Number.Uint64()) != height {
		panic("header number is not equal to the given iavl tree height")
	}

	return header, err
}

// GetHeaderByHistoricalNumber returns the header at the given height. It does not verify the
// height and always uses the query context for a historical height.
func (p *plugin) GetHeaderByHistoricalNumber(height int64) (*coretypes.Header, error) {
	sdkCtx, err := p.getQueryContext(height, false)
	if err != nil {
		return nil, errorslib.Wrap(err, "GetHeader: failed to use query context")
	}

	header, err := p.getHeaderFromStore(sdkCtx)
	if err != nil {
		return nil, errorslib.Wrap(err, "GetHeader: failed to get header from store")
	}

	if int64(header.Number.Uint64()) != height {
		panic("header number is not equal to the given iavl tree height")
	}

	return header, nil
}

// SetHeader saves a block to the store.
func (p *plugin) SetHeaderByNumber(ctx context.Context, _ int64, header *coretypes.Header) error {
	bz, err := coretypes.MarshalHeader(header)
	if err != nil {
		return errorslib.Wrap(err, "SetHeader: failed to marshal header")
	}
	sdk.UnwrapSDKContext(ctx).KVStore(p.storekey).Set([]byte{types.HeaderKey}, bz)
	return nil
}

func (p *plugin) getHeaderFromStore(ctx sdk.Context) (*coretypes.Header, error) {
	// Unmarshal the header from the context kv store.
	bz := ctx.KVStore(p.storekey).Get([]byte{types.HeaderKey})
	if bz == nil {
		return nil, errors.New("polaris header not found in kvstore")
	}
	header, err := coretypes.UnmarshalHeader(bz)
	if err != nil {
		return nil, errorslib.Wrap(err, "failed to unmarshal")
	}

	return header, nil
}

// resolveBlockHeight returns the IAVL height for the given block number and current height.
func (p *plugin) resolveBlockHeight(number int64, currentHeight int64) (int64, error) {
	var iavlHeight int64
	switch rpc.BlockNumber(number) { //nolint:nolintlint,exhaustive // covers all cases.
	case rpc.SafeBlockNumber, rpc.FinalizedBlockNumber:
		iavlHeight = currentHeight - 1
	case rpc.PendingBlockNumber, rpc.LatestBlockNumber:
		iavlHeight = currentHeight
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
